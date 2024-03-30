package main

import (
	"fmt"
	"sync/atomic"

	"math/rand"
	"sync"

	"github.com/textures1245/practical-examples/examples/concurrently"
	"github.com/textures1245/practical-examples/examples/resource"
)

type task struct {
	opt int
}

func (t *task) run() {
	switch t.opt {
	case 1:
		resourceTask()
	case 2:
		concurrentlyTask()
	default:
		fmt.Println("Invalid option")
	}

}

func main() {
	t := task{opt: 2}
	t.run()
}

func resourceTask() {
	// -  Write a program that reads a file and prints its contents to the console.
	file := resource.FileRead{FilePath: "/tmp/dat", Len: 0}
	file.Reader()

	// - Implement a function that takes a slice of integers and returns the sum of all even numbers.
	evenNum := resource.EvenNum{Nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}
	sum := evenNum.SumEven()
	println(sum)

	// - Create a program that simulates a simple bank account with deposit and withdrawal functions, ensuring thread safety using mutexes or channels.
	Ba := resource.BackAccount{Balance: 0}
	Ba.Deposit(500)
	Ba.Withdraw(400)
	fmt.Println(Ba.Balance)
}

func concurrentlyTask() {

	// - Implement a producer-consumer pattern using goroutines and channels, where one goroutine produces data, and another consumes it.
	service1 := concurrently.ProduceService{C: 0}
	const numJobs = 10
	jobs1 := make(chan func() int32, numJobs)
	res1 := make(chan int32, numJobs)

	// create workers
	for w := 1; w <= 3; w++ {
		go concurrently.Worker(w, jobs1, res1)
	}

	for j := 1; j <= numJobs; j++ {
		if rand.Intn(2) == 1 {
			jobs1 <- service1.Produce
		} else {
			jobs1 <- service1.Consume
		}
	}
	close(jobs1)

	for a := 1; a <= numJobs; a++ {
		fmt.Println(<-res1)
	}

	fmt.Printf("Final Value: %d\n", atomic.LoadInt32(&service1.C))

	// - Write a program that concurrently fetches data from multiple URLs and combines the results.
	urls := [5]string{}
	for i := 0; i < 5; i++ {
		urls[i] = "https://www.dochord.com/chord_charts/"
	}

	// - Create a program that simulates a simple web crawler using goroutines and channels, where each goroutine fetches and processes a URL.
	service2 := concurrently.FetchService{Urls: urls[:]}
	var res2 [5]string
	var wg sync.WaitGroup

	for i, url := range service2.Urls {
		wg.Add(1)
		index := i
		go func(url string) {
			defer wg.Done()
			res2[index] = service2.Fetch(url, 1)
		}(url)
	}

	wg.Wait()
	fmt.Println(res2)

	var fetcher = concurrently.FakeFetcher{
		"https://golang.org/": &concurrently.FakeResult{
			Body: "The Go Programming Language",
			Urls: []string{
				"https://golang.org/pkg/",
				"https://golang.org/cmd/",
			},
		},
		"https://golang.org/pkg/": &concurrently.FakeResult{
			Body: "Packages",
			Urls: []string{
				"https://golang.org/",
				"https://golang.org/cmd/",
				"https://golang.org/pkg/fmt/",
				"https://golang.org/pkg/os/",
			},
		},
		"https://golang.org/pkg/fmt/": &concurrently.FakeResult{
			Body: "Package fmt",
			Urls: []string{
				"https://golang.org/",
				"https://golang.org/pkg/",
			},
		},
		"https://golang.org/pkg/os/": &concurrently.FakeResult{
			Body: "Package os",
			Urls: []string{
				"https://golang.org/",
				"https://golang.org/pkg/",
			},
		},
	}

	service3 := concurrently.CrawService{}
	res3 := make(chan string)

	for i := 1; i < 5; i++ {
		index := i
		go func() {
			res := service3.Crawl("https://golang.org/", index, fetcher)
			res3 <- res

		}()
	}
	fmt.Println(<-res3)
}
