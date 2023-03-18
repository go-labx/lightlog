package lightlog

import "net"

func GetIPAddresses() (ipv4 string, ipv6 string) {
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
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
