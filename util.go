package lightlog

import (
	"fmt"
	"net"
	"path/filepath"
	"runtime"
)

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

func GetLocation() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	} else {
		file = filepath.Base(file)
	}
	location := fmt.Sprintf("%s:%d", file, line)

	return location
}
