package https

// - Create a test suite for a function or package you have written, and include benchmarks to measure its performance.
import (
	"context"
	"fmt"
	"testing"
	"time"
)

// Unit test for HttpHandler
func exceptedResStatusToBe(reqDat Req, excpStatus int, results chan Res, t *testing.T) Res {
	go HttpHandler(context.Background(), reqDat, results)
	res := <-results
	if res.Status != excpStatus {
		t.Errorf("Expected status to be %d, but got %d", excpStatus, res.Status)
	}
	return res

}

func TestHttpHandler(t *testing.T) {
	// Test code goes here

	results := make(chan Res)

	test := []struct {
		reqDat Req
		wants  int
	}{
		{Req{Target: "https://jsonplaceholder.typicode.com/todos/1", Verb: "GET", Dat: nil}, 200},
		{Req{Target: "https://jsonplaceholder.typicode.com/todos/2", Verb: "GET", Dat: nil}, 200},
		{Req{Target: "https://jsonplaceholder.typicode.com/todos/3", Verb: "GET", Dat: nil}, 200},
		{Req{Target: "https://jsonplaceholder.typicode.com/todos/4", Verb: "GET", Dat: nil}, 200},
		{Req{Target: "https://jsonplaceholder.typicode.com/todos/5", Verb: "GET", Dat: nil}, 200},
	}

	for _, tt := range test {
		testname := fmt.Sprintf("%s,%s", tt.reqDat.Target, tt.reqDat.Verb)
		t.Run(testname, func(t *testing.T) {
			exceptedResStatusToBe(tt.reqDat, tt.wants, results, t)
		})
	}
}

// benchmarking for HttpHandler
func BenchmarkTestHttpHandlerWithRateLimiter(b *testing.B) {
	limiter := time.Tick(100 * time.Millisecond) // set limiter = 200 mills

	results := make(chan Res)
	reqDat := Req{Target: "https://jsonplaceholder.typicode.com/todos/1", Verb: "GET", Dat: nil}
	for i := 0; i < b.N; i++ {
		<-limiter
		go HttpHandler(context.Background(), reqDat, results)
	}
}

func BenchmarkTestHttpHandlerWithBurstyLimiter(b *testing.B) {
	buffers := 20
	burstyLimiter := make(chan time.Time, buffers)

	// Fill up the channel to represent allowed bursting.
	for i := 0; i < buffers; i++ {
		burstyLimiter <- time.Now()
	}

	// using goroutine to tick every 200 mills to fill up the channal if it can.
	go func() {
		for t := range time.Tick(100 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	results := make(chan Res)
	reqDat := Req{Target: "https://jsonplaceholder.typicode.com/todos/1", Verb: "GET", Dat: nil}

	for req := 0; req < b.N; req++ {
		<-burstyLimiter // limit reqs followed bursty limiter
		go HttpHandler(context.Background(), reqDat, results)
	}
}
