// Copyright 2018 The Operator-SDK Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package build

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/operator-framework/operator-sdk/internal/pkg/scaffold"
	"github.com/operator-framework/operator-sdk/internal/pkg/scaffold/input"
	"github.com/operator-framework/operator-sdk/internal/util/projutil"
	"github.com/operator-framework/operator-sdk/internal/util/yamlutil"
	"github.com/operator-framework/operator-sdk/pkg/test"

	"github.com/ghodss/yaml"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	namespacedManBuild string
	testLocationBuild  string
	enableTests        bool
	imageBuildArgs     string
	imageBuilder       string
)

func NewCmd() *cobra.Command {
	buildCmd := &cobra.Command{
		Use:   "build <image>",
		Short: "Compiles code and builds artifacts",
		Long: `The operator-sdk build command compiles the code, builds the executables,
and generates Kubernetes manifests.

<image> is the container image to be built, e.g. "quay.io/example/operator:v0.0.1".
This image will be automatically set in the deployment manifests.

After build completes, the image would be built locally in docker. Then it needs to
be pushed to remote registry.
For example:
	$ operator-sdk build quay.io/example/operator:v0.0.1
	$ docker push quay.io/example/operator:v0.0.1
`,
		RunE: buildFunc,
	}
	buildCmd.Flags().BoolVar(&enableTests, "enable-tests", false, "Enable in-cluster testing by adding test binary to the image")
	buildCmd.Flags().StringVar(&testLocationBuild, "test-location", "./test/e2e", "Location of tests")
	buildCmd.Flags().StringVar(&namespacedManBuild, "namespaced-manifest", "deploy/operator.yaml", "Path of namespaced resources manifest for tests")
	buildCmd.Flags().StringVar(&imageBuildArgs, "image-build-args", "", "Extra image build arguments as one string such as \"--build-arg https_proxy=$https_proxy\"")
	buildCmd.Flags().StringVar(&imageBuilder, "image-builder", "docker", "Tool to build OCI images. One of: [docker, buildah]")
	return buildCmd
}

/*
 * verifyDeploymentImages checks image names of pod 0 in deployments found in the provided yaml file.
 * This is done because e2e tests require a namespaced manifest file to configure a namespace with
 * required resources. This function is intended to identify if a user used a different image name
 * for their operator in the provided yaml, which would result in the testing of the wrong operator
 * image. As it is possible for a namespaced yaml to have multiple deployments (such as the vault
 * operator, which depends on the etcd-operator), this is just a warning, not a fatal error.
 */
func verifyDeploymentImage(yamlFile []byte, imageName string) error {
	warningMessages := ""
	scanner := yamlutil.NewYAMLScanner(yamlFile)
	for scanner.Scan() {
		yamlSpec := scanner.Bytes()

		yamlMap := make(map[string]interface{})
		err := yaml.Unmarshal(yamlSpec, &yamlMap)
		if err != nil {
			return fmt.Errorf("could not unmarshal YAML namespaced spec: (%v)", err)
		}
		kind, ok := yamlMap["kind"].(string)
		if !ok {
			return fmt.Errorf("yaml manifest file contains a 'kind' field that is not a string")
		}
		if kind == "Deployment" {
			// this is ugly and hacky; we should probably make this cleaner
			nestedMap, ok := yamlMap["spec"].(map[string]interface{})
			if !ok {
				continue
			}
			nestedMap, ok = nestedMap["template"].(map[string]interface{})
			if !ok {
				continue
			}
			nestedMap, ok = nestedMap["spec"].(map[string]interface{})
			if !ok {
				continue
			}
			containersArray, ok := nestedMap["containers"].([]interface{})
			if !ok {
				continue
			}
			for _, item := range containersArray {
				image, ok := item.(map[string]interface{})["image"].(string)
				if !ok {
					continue
				}
				if image != imageName {
					warningMessages = fmt.Sprintf("%s\nWARNING: Namespace manifest contains a deployment with image %v, which does not match the name of the image being built: %v", warningMessages, image, imageName)
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to verify deployment image: (%v)", err)
	}
	if warningMessages == "" {
		return nil
	}
	return errors.New(warningMessages)
}

func verifyTestManifest(image string) error {
	namespacedBytes, err := ioutil.ReadFile(namespacedManBuild)
	if err != nil {
		return fmt.Errorf("could not read namespaced manifest: (%v)", err)
	}

	err = verifyDeploymentImage(namespacedBytes, image)
	// the error from verifyDeploymentImage is just a warning, not fatal error
	if err != nil {
		log.Warn(err)
	}
	return nil
}

func createBuildCommand(imageBuilder, context, dockerFile, image string, imageBuildArgs ...string) (*exec.Cmd, error) {
	var args []string
	switch imageBuilder {
	case "docker":
		args = append(args, "build", "-f", dockerFile, "-t", image)
	case "buildah":
		args = append(args, "bud", "--format=docker", "-f", dockerFile, "-t", image)
	default:
		return nil, fmt.Errorf("%s is not supported image builder", imageBuilder)
	}

	for _, bargs := range imageBuildArgs {
		if bargs != "" {
			splitArgs := strings.Fields(bargs)
			args = append(args, splitArgs...)
		}
	}

	args = append(args, context)

	return exec.Command(imageBuilder, args...), nil
}

func buildFunc(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("command %s requires exactly one argument", cmd.CommandPath())
	}

	projutil.MustInProjectRoot()
	goBuildEnv := append(os.Environ(), "GOOS=linux", "GOARCH=amd64", "CGO_ENABLED=0")
	goTrimFlags := []string{"-gcflags", "all=-trimpath=${GOPATH}", "-asmflags", "all=-trimpath=${GOPATH}"}
	absProjectPath := projutil.MustGetwd()
	projectName := filepath.Base(absProjectPath)

	// Don't need to build Go code if a non-Go Operator.
	if projutil.IsOperatorGo() {
		opts := projutil.GoCmdOptions{
			BinName:     filepath.Join(absProjectPath, scaffold.BuildBinDir, projectName),
			PackagePath: filepath.Join(projutil.CheckAndGetProjectGoPkg(), scaffold.ManagerDir),
			Args:        goTrimFlags,
			Env:         goBuildEnv,
			GoMod:       projutil.IsDepManagerGoMod(),
		}
		if err := projutil.GoBuild(opts); err != nil {
			return fmt.Errorf("failed to build operator binary: (%v)", err)
		}
	}

	image := args[0]
	baseImageName := image
	if enableTests {
		baseImageName += "-intermediate"
	}

	log.Infof("Building OCI image %s", baseImageName)

	buildCmd, err := createBuildCommand(imageBuilder, ".", "build/Dockerfile", baseImageName, imageBuildArgs)
	if err != nil {
		return err
	}

	if err := projutil.ExecCmd(buildCmd); err != nil {
		if enableTests {
			return fmt.Errorf("failed to output intermediate image %s: (%v)", image, err)
		}
		return fmt.Errorf("failed to output build image %s: (%v)", image, err)
	}

	if enableTests {
		if projutil.IsOperatorGo() {
			opts := projutil.GoTestOptions{
				GoCmdOptions: projutil.GoCmdOptions{
					BinName:     filepath.Join(absProjectPath, scaffold.BuildBinDir, projectName+"-test"),
					PackagePath: testLocationBuild + "/...",
					Args:        append(goTrimFlags, "-c"),
					Env:         goBuildEnv,
					GoMod:       projutil.IsDepManagerGoMod(),
				},
			}
			if err := projutil.GoTest(opts); err != nil {
				return fmt.Errorf("failed to build test binary: (%v)", err)
			}
		}

		// if a user is using an older sdk repo as their library, make sure they have required build files
		testDockerfile := filepath.Join(scaffold.BuildTestDir, scaffold.DockerfileFile)
		_, err := os.Stat(testDockerfile)
		if err != nil && os.IsNotExist(err) {

			log.Info("Generating build manifests for test-framework.")

			cfg := &input.Config{
				Repo:           projutil.CheckAndGetProjectGoPkg(),
				AbsProjectPath: absProjectPath,
				ProjectName:    projectName,
			}

			s := &scaffold.Scaffold{}
			switch t := projutil.GetOperatorType(); t {
			case projutil.OperatorTypeGo:
				err = s.Execute(cfg,
					&scaffold.TestFrameworkDockerfile{},
					&scaffold.GoTestScript{},
					&scaffold.TestPod{Image: image, TestNamespaceEnv: test.TestNamespaceEnv},
				)
			case projutil.OperatorTypeAnsible:
				return fmt.Errorf("test scaffolding for Ansible Operators is not implemented")
			case projutil.OperatorTypeHelm:
				return fmt.Errorf("test scaffolding for Helm Operators is not implemented")
			default:
				return projutil.ErrUnknownOperatorType{}
			}

			if err != nil {
				return fmt.Errorf("test framework manifest scaffold failed: (%v)", err)
			}
		}

		log.Infof("Building test OCI image %s", image)

		testImageBuildArgs := fmt.Sprintf("--build-arg NAMESPACEDMAN=%s --build-arg BASEIMAGE=%s", namespacedManBuild, baseImageName)
		testBuildCmd, err := createBuildCommand(imageBuilder, ".", testDockerfile, image, imageBuildArgs, testImageBuildArgs)
		if err != nil {
			return err
		}

		if err := projutil.ExecCmd(testBuildCmd); err != nil {
			return fmt.Errorf("failed to output test image %s: (%v)", image, err)
		}
		// Check image name of deployments in namespaced manifest
		if err := verifyTestManifest(image); err != nil {
			return nil
		}
	}

	log.Info("Operator build complete.")
	return nil
}
