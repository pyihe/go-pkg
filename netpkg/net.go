package netpkg

import (
	"net"
	"strconv"
	"strings"
	"sync"
)

var (
	once    sync.Once
	localIP string
)

// GetLocalIP 获取本地内网IP
func GetLocalIP() string {
	once.Do(func() {
		ips, _ := internetIP()
		if len(ips) > 0 {
			localIP = ips[0]
		} else {
			localIP = "127.0.0.1"
		}
	})
	return localIP
}

// GetInternalIP 获取Internal IP
func GetInternalIP() string {
	inters, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, inter := range inters {
		if inter.Flags&net.FlagUp != 0 && !strings.HasPrefix(inter.Name, "lo") {
			addrs, err := inter.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					return ipnet.IP.String()
				}
			}
		}
	}
	return ""
}

func internetIP() ([]string, error) {
	var (
		ips []string
		err error
	)
	ifaces, err := net.Interfaces()
	if err != nil {
		return ips, err
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		if strings.HasPrefix(iface.Name, "docker") || strings.HasPrefix(iface.Name, "w-") {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			return ips, err
		}

		for _, ad := range addrs {
			var ip net.IP
			switch t := ad.(type) {
			case *net.IPNet:
				ip = t.IP
			case *net.IPAddr:
				ip = t.IP
			}

			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue
			}

			ipStr := ip.String()
			if isIntranet(ipStr) {
				ips = append(ips, ipStr)
			}
		}
	}
	return ips, nil
}

func isIntranet(s string) bool {
	if strings.HasPrefix(s, "10.") || strings.HasPrefix(s, "192.168.") {
		return true
	}

	if strings.HasPrefix(s, "172.") {
		// 172.16.0.0-172.31.255.255
		arr := strings.Split(s, ".")
		if len(arr) != 4 {
			return false
		}

		second, err := strconv.ParseInt(arr[1], 10, 64)
		if err != nil {
			return false
		}

		if second >= 16 && second <= 31 {
			return true
		}
	}

	return false
}
