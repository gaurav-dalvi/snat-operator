package utils

import (
	"log"
	"net"
	"sort"
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
// It does not omit network and host reserved IP address from that.
func GetIPsFromCIDR(cidr string) []string {

	var output []string

	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		log.Fatal("Invalid CIDR, skiiping this one", cidr, err)
	}
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		output = append(output, ip.String())
	}
	return output
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
