package benchmarker

import (
	"log"
	"net/http"
)

var (
	reqs int
	max  int
)

type Response struct {
	*http.Response
	err error
}

// Dispatcher
func Dispatcher(reqChan chan *http.Request, generateRequest func() *http.Request) {
	defer close(reqChan)
	for i := 0; i < reqs; i++ {
		reqChan <- generateRequest()
	}
}

// Worker Pool
func WorkerPool(reqChan chan *http.Request, respChan chan Response) {
	t := &http.Transport{}
	for i := 0; i < max; i++ {
		go Worker(t, reqChan, respChan)
	}
}

// Worker
func Worker(t *http.Transport, reqChan chan *http.Request, respChan chan Response) {
	for req := range reqChan {
		resp, err := t.RoundTrip(req)
		r := Response{resp, err}
		respChan <- r
	}
}

// Consumer
func Consumer(respChan chan Response) (int64, int64) {
	var (
		conns int64
		size  int64
	)
	for conns < int64(reqs) {
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
