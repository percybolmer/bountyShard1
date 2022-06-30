package benchmarker

import (
	"errors"
	"net/http"
	"time"
)

// Benchmarker is used to perform benchmarkin
type Benchmarker struct {
	reqs int
	max  int
}

// BenchMetric contains information about the request
type BenchMetric struct {
	Duration int64 // in nano
}

func NewBenchmarker(requestsToSend int, maxConcurrent int) *Benchmarker {
	return &Benchmarker{
		reqs: requestsToSend,
		max:  maxConcurrent,
	}
}

type GenerateRequestFunc func() *http.Request

// Response is the response from the request
type Response struct {
	Message string
	Err     error
	Metric  BenchMetric
}

// Dispatcher is used to dispatch the amount of requests
// once done it will close the requestchannel and trigger the consumer to exit
func (bm *Benchmarker) Dispatcher(reqChan chan *http.Request, generateRequest GenerateRequestFunc) {
	defer close(reqChan)
	for i := 0; i < bm.reqs; i++ {
		reqChan <- generateRequest()
	}
}

// Worker Pool handles the work load to avoid over concurreny
// Will only send bm.max amount of rquests at the same timeS
func (bm *Benchmarker) WorkerPool(reqChan chan *http.Request, respChan chan Response) {
	t := &http.Transport{}
	for i := 0; i < bm.max; i++ {
		go bm.Worker(t, reqChan, respChan)
	}
}

// Worker performs the actual work and measures execution time of the request
func (bm *Benchmarker) Worker(t *http.Transport, reqChan chan *http.Request, respChan chan Response) {
	for req := range reqChan {
		// Measure time
		start := time.Now()
		resp, err := t.RoundTrip(req)

		duration := time.Since(start)
		if resp.StatusCode != http.StatusOK {
			err = errors.New(resp.Status)
		}
		response := Response{
			Err: err,
			Metric: BenchMetric{
				Duration: duration.Nanoseconds(),
			},
		}
		// Send request
		respChan <- response
	}
}

// Consumer listens on the Responses
// Will cancel when respChan cancels
// Consumer is responsible for absobing all the requests
// And concatenating data into a result
// returns the following params
// requestsProcessed - amount of responses ingested
// totalDuration -- total duartion in nanoseconds
// failures -- amount of requests that failed
func (bm *Benchmarker) Consumer(respChan chan Response) (requestsProcessed int64, totalDuration int64, failures int64) {

	for requestsProcessed < int64(bm.reqs) {
		select {
		case r, ok := <-respChan:
			if ok {
				if r.Err != nil {
					failures++
				}
				totalDuration = totalDuration + r.Metric.Duration
				requestsProcessed++
			}
		}
	}
	return
}
