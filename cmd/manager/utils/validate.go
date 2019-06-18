package utils

// aciv1 "github.com/gaurav-dalvi/snat-operator/pkg/apis/aci/v1"

const (
	MAX_PORT = 65535
	MIN_PORT = 1
)

// Validator defines validator struct
type Validator struct {
	Validated    bool
	ErrorMessage string
}

// // Validate validates SnatSubnet Custom Resource
// func (v *Validator) ValidateSnatSubnet(cr *aciv1.SnatSubnet) {
// 	v.Validated = true
// 	reservedPorts := GetReservedPortRanges()
// 	for _, item := range cr.Spec.Snatipsubnets {
// 		_, _, err := net.ParseCIDR(item)
// 		if err != nil {
// 			v.ErrorMessage = v.ErrorMessage + "Invalid subnet\n"
// 			v.Validated = false
// 		}
// 	}

// 	if cr.Spec.Pernodeports > MAX_PORT {
// 		v.ErrorMessage = v.ErrorMessage + "Invalid number of ports per node\n"
// 		v.Validated = false
// 	}

// 	for _, port_range := range cr.Spec.Snatports {
// 		if port_range.Start < MIN_PORT || port_range.Start > MAX_PORT || port_range.End < MIN_PORT || port_range.End > MAX_PORT {
// 			v.ErrorMessage = v.ErrorMessage + "Invalid port number in the range\n"
// 			v.Validated = false
// 		}
// 		if port_range.Start > port_range.End {
// 			v.ErrorMessage = v.ErrorMessage + "Start can not be bigger thant End in port_range\n"
// 			v.Validated = false
// 		}
// 	}

// 	// To check if any port from the given range falls into reserved port range or not
// 	for _, port_range := range cr.Spec.Snatports {
// 		for _, rPort := range reservedPorts {
// 			if inBetween(rPort.Start, port_range.Start, port_range.End) || inBetween(rPort.End, port_range.Start, port_range.End) {
// 				v.ErrorMessage = v.ErrorMessage + "Port can not be from reserved port ranges\n"
// 				v.Validated = false
// 			}
// 		}
// 	}

// }

// // Validate validates SnatIP Custom Resource
// func (v *Validator) ValidateSnatIP(cr *aciv1.SnatIP) {
// 	v.Validated = true

// 	for _, item := range cr.Spec.Snatipsubnets {
// 		_, _, err := net.ParseCIDR(item)
// 		if err != nil {
// 			v.ErrorMessage = v.ErrorMessage + "Invalid subnet\n"
// 			v.Validated = false
// 		}
// 	}

// 	is_found := false
// 	for _, item := range VALID_RESOURCE_TYPES {
// 		if item == cr.Spec.Resourcetype {
// 			is_found = true
// 			break
// 		}
// 	}
// 	if !is_found {
// 		v.ErrorMessage = v.ErrorMessage + "Invalid resourcetype\n"
// 		v.Validated = false
// 	}
// }

func inBetween(num, min, max int) bool {
	if num >= min && num <= max {
		return true
	}
	return false
}
