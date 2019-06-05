package controller

import (
	"github.com/gaurav-dalvi/snat-operator/pkg/controller/snatip"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, snatip.Add)
}
