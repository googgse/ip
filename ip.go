package ip

import (
	// import built-in packages
	"net"
)

// 内部IPv4地址
func InternalV4() string {
	addrs, err := net.InterfaceAddrs()
	// control flow
	if err != nil {
		// return value
		return ""
	}
	for _, address := range addrs {
		// control flow
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			// control flow
			if ipnet.IP.To4() != nil {
				ip := ipnet.IP.String()
				// return value
				return ip
			}
		}
	}
	// return value
	return ""
}

// 内部IPv6地址
func InternalV6() string {
	addrs, err := net.InterfaceAddrs()
	// control flow
	if err != nil {
		// return value
		return ""
	}
	// control flow
	for _, address := range addrs {
		ipnet, ok := address.(*net.IPNet)
		// control flow
		if ok && !ipnet.IP.IsLoopback() {
			// control flow
			if ipnet.IP.To16() != nil {
				ip := ipnet.IP.String()
				// return value
				return ip
			}
		}
	}
	// return value
	return ""
}
