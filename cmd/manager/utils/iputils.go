package utils

import (
	"bytes"
	"fmt"
	"net"
	"sort"
	// snattypes "github.com/gaurav-dalvi/snat-operator/pkg/apis/aci/v1"
)

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// Given CIDR, this function returns list of IP addresses in that CIDR
// It does omit network and host reserved IP address from that.
func GetIPsFromCIDR(cidr string) []string {

	var output []string

	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		UtilLog.Error(err, "Invalid CIDR, skiiping this one", cidr)
	}
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		output = append(output, ip.String())
	}
	return output[1 : len(output)-1]
}

// Given two slices of IP addresses, finds the difference of two sets
func SetDifferenceIPs(inputIps []string, usedIps []string) []string {

	var output []string

	return output
}

// Draw free IP address from the given list
func DrawIP(freeIps []string) {
	sort.Strings(freeIps)
}

// Given generic list of CIDRs for subnets
// return sorted array(based on IP address) of IP addresses.
func ExpandCIDRs(currCIDRs []string) []string {

	var expandedIPs []string
	for _, item := range currCIDRs {
		ips := GetIPsFromCIDR(item)
		expandedIPs = append(expandedIPs, ips...)
	}

	// Sort list of IPs
	expandedIPs = sortIps(expandedIPs)
	UtilLog.Info("Inside ExpandCIDRs", "currCIDRs:", expandedIPs)

	return expandedIPs
}

// This function sorts IP by parsing them to net.IP struct. String sort does not work
// eg: 10.0.0.9 should come before 10.0.0.10
func sortIps(ips []string) []string {
	realIPs := make([]net.IP, 0, len(ips))
	for _, ip := range ips {
		realIPs = append(realIPs, net.ParseIP(ip))
	}

	sort.Slice(realIPs, func(i, j int) bool { return bytes.Compare(realIPs[i], realIPs[j]) < 0 })

	outputIps := make([]string, 0, len(realIPs))
	for _, ip := range realIPs {
		outputIps = append(outputIps, fmt.Sprintf("%s", ip))
	}
	return outputIps
}

// // This function will be repaced depeding upon design choice.
// func GetReservedPortRanges() []snattypes.PortRange {
// 	reservedPorts := []snattypes.PortRange{
// 		snattypes.PortRange{Start: 1, End: 1},
// 		snattypes.PortRange{Start: 20, End: 22},
// 		snattypes.PortRange{Start: 29, End: 29},
// 		snattypes.PortRange{Start: 37, End: 37},
// 		snattypes.PortRange{Start: 42, End: 43},
// 		snattypes.PortRange{Start: 108, End: 110},
// 		snattypes.PortRange{Start: 443, End: 445},
// 		snattypes.PortRange{Start: 1080, End: 1080},
// 	}
// 	return reservedPorts
// }
