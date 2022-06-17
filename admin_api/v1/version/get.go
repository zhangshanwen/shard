package version

import (
	"fmt"
	"net"

	"github.com/zhangshanwen/shard/initialize/service"
)

var buildTime string
var git string

func getIPs() (ips []string) {

	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("fail to get net interface addrs: %v", err)
		return ips
	}

	for _, address := range interfaceAddr {
		ipNet, isValidIpNet := address.(*net.IPNet)
		if isValidIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	return ips
}
func Get(c *service.AdminContext) (r service.Res) {
	r.Data = map[string]interface{}{
		"build_time": buildTime,
		"git":        git,
		"ip":         getIPs(),
	}
	return
}
