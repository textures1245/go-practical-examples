package concurrently

import (
	"bufio"
	"fmt"
	"net/http"
	"sync/atomic"
)

// - Implement a producer-consumer pattern using goroutines and channels, where one goroutine produces data, and another consumes it.
type ProduceService struct {
	C int32
}

func Worker(id int, job <-chan func() int32, res chan<- int32) {
	for j := range job {

		doJob := j
		go func() {
			result := doJob()
			res <- result
		}()
	}
}

func (p *ProduceService) Produce() int32 {
	atomic.AddInt32(&p.C, 1)
	return p.C
}

func (p *ProduceService) Consume() int32 {
	atomic.AddInt32(&p.C, -1)
	return p.C
}

// - Write a program that concurrently fetches data from multiple URLs and combines the results.
type FetchService struct {
	Urls []string
	C    int32
}

func (f *FetchService) Fetch(url string, lineLimiter int) string {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var context string
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < lineLimiter; i++ {
		context = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return context

}

// - Create a program that simulates a simple web crawler using goroutines and channels, where each goroutine fetches and processes a URL.

type CrawService struct {
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (c CrawService) Crawl(url string, depth int, fetcher Fetcher) string {
	// TODO: Don't fetch the same URL twice.
	// TODO: Fetch URLs in parallel.
	// This implementation doesn't do either:
	if depth <= 0 {
		return ""
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	res := fmt.Sprintf("found: %s %q\n", url, body)
	for _, u := range urls {
		c.Crawl(u, depth-1, fetcher)
	}
	return res
}

// fakeFetcher is Fetcher that returns canned results.
type FakeFetcher map[string]*FakeResult

type FakeResult struct {
	Body string
	Urls []string
}

func (f FakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.Body, res.Urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
