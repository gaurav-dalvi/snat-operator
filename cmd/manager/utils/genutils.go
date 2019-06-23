package utils

import (
	"sort"

	snattypes "github.com/gaurav-dalvi/snat-operator/pkg/apis/aci/v1"

	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var UtilLog = logf.Log.WithName("Utils:")

// StartSorter sorts PortRanges based on Start field.
type StartSorter []snattypes.PortRange

func (a StartSorter) Len() int           { return len(a) }
func (a StartSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a StartSorter) Less(i, j int) bool { return a[i].Start < a[j].Start }

// Given generic list of start and end of each port range,
// return sorted array(based on start of the range) of portranges based on number of per node
func ExpandPortRanges(currPortRange []snattypes.PortRange, step int) []snattypes.PortRange {

	UtilLog.Info("Inside ExpandPortRanges", "currPortRange:", currPortRange, "Step:", step)
	expandedPortRange := []snattypes.PortRange{}
	for _, item := range currPortRange {
		temp := item.Start
		for temp < item.End-1 {
			expandedPortRange = append(expandedPortRange, snattypes.PortRange{Start: temp, End: temp + step - 1})
			temp = temp + step
		}
	}

	// Sort based on `Start` field
	sort.Sort(StartSorter(expandedPortRange))

	return expandedPortRange
}

func Contains(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}

func Remove(list []string, s string) []string {
	for i, v := range list {
		if v == s {
			list = append(list[:i], list[i+1:]...)
		}
	}
	return list
}
