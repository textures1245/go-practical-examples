package main

import (
	"fmt"
	"sync/atomic"
	"time"

	"math/rand"
	"sync"

	"github.com/textures1245/practical-examples/examples/concurrently"
	"github.com/textures1245/practical-examples/examples/error"
	"github.com/textures1245/practical-examples/examples/generic"
	"github.com/textures1245/practical-examples/examples/https"
	"github.com/textures1245/practical-examples/examples/logger"
	"github.com/textures1245/practical-examples/examples/resource"
	"github.com/textures1245/practical-examples/examples/timer"
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
	case 3:
		genericTask()
	case 4:
		errorTask()
	case 5:
		timeTask()
	case 6:
		debugTask()
	default:
		fmt.Println("Invalid option")
	}

}

func main() {
	t := task{opt: 6}
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

func genericTask() {
	// -\ Implement a generic function and structs and achieved assign by the following
	sum1 := generic.Generic[int]{}
	fmt.Println(sum1.Sum([]int{1, 2, 3, 4, 5}))
	fmt.Println(sum1.Pos.GetElemsSumHistory())

	sum2 := generic.Generic[string]{}
	fmt.Println(sum2.Sum([]string{"a", "b", "c", "d", "e"}).(string))
	fmt.Println(sum2.Pos.GetElemsSumHistory())
	fmt.Println(sum2.Sum([]string{"a", "b", "c", "d", "e", "f"}).(string))
	fmt.Println(sum2.Pos.GetElemsSumHistory())

	// //- Write a program that simulates a client-server interaction, where the client sends a multiply requests and waits for a response with a timeout.
	// urls := []string{"https://json3placeholder.typicode.com/todos/1", "https://jsonplac3eholder.typicode.com/todos/2", "https://jsonplac3eholder.typicode.com/todos/3"}
	// task3 := generic.Client{Results: make(chan string, len(urls))}

	// task3.OnSendReqs(urls)

}

func errorTask() {

	// - Write a function that reads data from a file and handles various errors that may occur during the process.
	// f := error.File{}
	// f.Read("test.txt")

	// - Implement a program that uses pointers to modify the values of variables passed to a function.
	// lst := error.List[int]{}
	// lst.ChangeTo([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	// lst.ChangeTo([]int{1, 2, 3, 4, 5})

	// fmt.Println(lst.GetElems())
	// lst.ChangeToHistory(0)
	// fmt.Println(lst.GetElems())

	// - Create a program that uses goroutines and channels to perform a parallel calculation, such as finding the sum of elements in a large array.

	counter := error.Counter{C: 0}

	var sample = []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	arrNum := error.Divider(5, sample)

	fmt.Println(arrNum)

	jobs1 := make(chan func() int32, len(arrNum))
	res1 := make(chan int32, len(arrNum))

	// create workers
	for w := 1; w <= 3; w++ {
		go concurrently.Worker(w, jobs1, res1)
	}

	for _, nums := range arrNum {

		// store nums to local variable for each new iteration to prevent race condition
		// when working as concurrently
		s := nums
		jobs1 <- func() int32 {
			return counter.Sum(s)
		}
	}
	close(jobs1)

	for a := 1; a <= len(arrNum); a++ {

		fmt.Println(<-res1)
	}

	fmt.Printf("Final Value: %d\n", atomic.LoadInt32(&counter.C))

}

func timeTask() {
	//- create ticker handler, every tick will call the func from arg then return the result to the channel
	t := timer.OnTick[int]{}
	ticker := time.NewTicker(500 * time.Millisecond)
	num := 0
	done := make(chan bool, 1)
	results := make(chan int)

	incre := func(n *int) {
		*n++
	}

	go t.TimestampTicker(ticker, done, results, func() int {
		incre(&num)
		return num
	})

	time.Sleep(4 * time.Second)
	t.StopTicker(done)
	for n := range results {
		fmt.Println(n)

	}

	// - Implement a simple HTTP server that serves static files and handles different HTTP methods (GET, POST, etc.).
	s := https.Server{}

	reqs := []https.Req{
		{Target: "https://jsonplaceholder.typicode.com/todos/1", Verb: "GET", Dat: nil},
		{Target: "https://jsonplaceholder.typicode.com/todos/2", Verb: "GET", Dat: nil},
	}
	var wg sync.WaitGroup

	res := make(chan https.Res, len(reqs))

	for _, req := range reqs {
		wg.Add(1)
		go func(req https.Req, wg *sync.WaitGroup) {
			defer wg.Done()
			s.HandleReq(req, res)
		}(req, &wg)
	}
	wg.Wait()

	close(res) // close channel for no incoming reqs

	for r := range res {
		fmt.Println(r)
	}
}

func debugTask() {
	normal_d := logger.NewConsoleLogger()

	normal_d.Db.Error("Error")
	normal_d.Db.Info("Info")
	normal_d.Db.Debug("Debug")
	normal_d.Db.Warn("Warn")

	level := logger.Level("FETAL")
	d_with_level := logger.NewConsoleLogger(level)

	d_with_level.Db.Error("Error")
	d_with_level.Db.Info("Info")
	d_with_level.Db.Debug("Debug")
	d_with_level.Db.Warn("Warn")
}
