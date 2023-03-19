package lightlog

import (
	"fmt"
	"net"
	"path/filepath"
	"runtime"
)

// GetIPAddresses This function returns the IP addresses of the current machine.
func GetIPAddresses() (ipv4 string, ipv6 string) {
	adds, _ := net.InterfaceAddrs()
	for _, addr := range adds {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ip4 := ipnet.IP.To4(); ip4 != nil {
				ipv4 = ip4.String()
			} else if ip6 := ipnet.IP.To16(); ip6 != nil {
				ipv6 = ip6.String()
			}
		}
	}
	return ipv4, ipv6
}

// GetLocation returns the file name and line number of the caller of the method that called
func GetLocation() string {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "???"
		line = 0
	} else {
		file = filepath.Base(file)
	}
	location := fmt.Sprintf("%s:%d", file, line)
	return location
}
