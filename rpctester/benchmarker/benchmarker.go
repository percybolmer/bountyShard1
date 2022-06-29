package benchmarker

import (
	"log"
	"net/http"
)

type Benchmarker struct {
	reqs int
	max  int
}

func NewBenchmarker(requestsToSend int, maxConcurrent int) *Benchmarker {
	return &Benchmarker{
		reqs: requestsToSend,
		max:  maxConcurrent,
	}
}

type Response struct {
	*http.Response
	err error
}

// Dispatcher
func (bm *Benchmarker) Dispatcher(reqChan chan *http.Request, generateRequest func() *http.Request) {
	defer close(reqChan)
	for i := 0; i < bm.reqs; i++ {
		reqChan <- generateRequest()
	}
}

// Worker Pool
func (bm *Benchmarker) WorkerPool(reqChan chan *http.Request, respChan chan Response) {
	t := &http.Transport{}
	for i := 0; i < bm.max; i++ {
		go bm.Worker(t, reqChan, respChan)
	}
}

// Worker
func (bm *Benchmarker) Worker(t *http.Transport, reqChan chan *http.Request, respChan chan Response) {
	for req := range reqChan {
		resp, err := t.RoundTrip(req)
		r := Response{resp, err}
		respChan <- r
	}
}

// Consumer
func (bm *Benchmarker) Consumer(respChan chan Response) (int64, int64) {
	var (
		conns int64
		size  int64
	)
	for conns < int64(bm.reqs) {
		select {
		case r, ok := <-respChan:
			if ok {
				if r.err != nil {
					log.Println(r.err)
				} else {
					size += r.ContentLength
					if err := r.Body.Close(); err != nil {
						log.Println(r.err)
					}
				}
				conns++
			}
		}
	}
	return conns, size
}
