package main

import (
	f "fmt"
	"strings"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	parts := make([]string, len(ip))
	for i, b := range ip {
		parts[i] = f.Sprintf("%d", b)
	}
	return strings.Join(parts, ".")
}
func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		f.Printf("%v: %v\n", name, ip)
	}
}
