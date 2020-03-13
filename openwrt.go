package openwrt

import (
	"log"
	"net"
	"os/exec"
	"strings"
)

// Version export
const Version = "0.0.1"

// OpenWRT type
type OpenWRT struct {
}

// NewOpenWRT export
func NewOpenWRT() *OpenWRT {
	return &OpenWRT{}
}

// Service export
func (id *OpenWRT) Service(service string, command string) (string, error) {
	cmd := exec.Command("service", service, command)
	err := cmd.Run()
	if err != nil {
		log.Println("service failed to exec", service, command, err)
		return "", err
	}
	return "", nil
}

// UCI export
func (id *OpenWRT) UCI(command string, arg string) (string, error) {
	cmd := exec.Command("uci", command, arg)
	err := cmd.Run()
	if err != nil {
		log.Println("uci failed to exec", command, arg, err)
		return "", err
	}
	return "", nil
}

// SetDNS export
func (id *OpenWRT) SetDNS(servers []string) bool {
	serverString := strings.Join(servers, " ")
	_, err := id.UCI("set", "network.wan.dns=\""+serverString+"\"")
	if err != nil {
		log.Println("setDNS failed to uci set")
		return false
	}
	_, err = id.UCI("commit", "network")
	if err != nil {
		log.Println("setDNS failed to uci commit")
		return false
	}
	_, err = id.Service("network", "reload")
	if err != nil {
		log.Println("setDNS failed to network reload")
		return false
	}
	return true
}

// GetWANIPV4 export
func (id *OpenWRT) GetWANIPV4() string {
	var result = ""
	ifcs, _ := net.Interfaces()
	for _, ifc := range ifcs {
		addrs, _ := ifc.Addrs()
		for _, addr := range addrs {
			addrStr := addr.String()
			if strings.Index(addrStr, ".") > -1 && strings.HasPrefix(ifc.Name, "en") && ifc.Flags&net.FlagLoopback == 0 && ifc.Flags&net.FlagUp != 0 {
				// eliminate non-WAN routable blocks
				if !strings.HasPrefix(addrStr, "192.168.") && !strings.HasPrefix(addrStr, "172.16.") && !strings.HasPrefix(addrStr, "10.") && !strings.HasPrefix(addrStr, "127.") {
					mask := strings.Index(addrStr, "/")
					if mask != -1 {
						return addrStr[:mask]
					}
					return addrStr
				}
			}
		}
	}

	return result
}
