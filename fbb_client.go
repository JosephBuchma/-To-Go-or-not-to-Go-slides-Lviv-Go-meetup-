package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

// GetFBB OMIT
func GetFBB(n int) (body string, err error) {
	r, err := http.Get(fmt.Sprintf("http://localhost:8080/%d", n)) // HL
	if err != nil {
		return
	}
	defer r.Body.Close()
	bbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	return string(bbody), err
}

// GetFBB OMIT

// FBBHTTPConcurrent OMIT
func FBBHTTPConcurrent(n int) {
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 1; i <= n; i++ {
		go func(n int) {
			defer wg.Done()
			if r, e := GetFBB(n); e != nil {
				fmt.Println("Error: %s", e)
			} else {
				fmt.Println(r)
			}
		}(i)
	}
	wg.Wait()
}

// FBBHTTPConcurrent OMIT

// FBBHTTPConcurrentSorted OMIT
func FBBHTTPConcurrentSorted(n int) {
	returns := make([]chan string, n)
	for i := 1; i <= n; i++ {
		ret := make(chan string)
		returns[i-1] = ret
		go func(n int) {
			if r, e := GetFBB(n); e != nil {
				fmt.Println("Error: %s", e)
				close(ret)
			} else {
				ret <- r
			}
		}(i)
	}
	for _, r := range returns {
		fmt.Println(<-r)
	}
}

// FBBHTTPConcurrentSorted OMIT

// FBBNUMGEN OMIT
func NumGen(n int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

// FBBNUMGEN OMIT

// FBBHTTPConcurrentFanInFanOut OMIT
func FBBHTTPConcurrentFanInFanOut(fanOut <-chan int) {
	fanIn := make(chan string)
	for i := 0; i < 100; i++ {
		go func() {
			for n := range fanOut {
				if r, e := GetFBB(n); e != nil {
					fmt.Println("Error: %s", e)
				} else {
					fanIn <- r
				}
			}
		}()
	}
	for resp := range fanIn {
		fmt.Println(resp)
	}
}

// FBBHTTPConcurrentFanInFanOut OMIT

// FBBNUMGENCANCEL OMIT
func NumGenCancel(n int, done <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			select {
			case ch <- i:
			case <-done:
				return
			}
		}
		close(ch)
	}()
	return ch
}

// FBBNUMGENCANCEL OMIT

// FBBHTTPConcurrentFanInFanOutCancel OMIT
func FBBHTTPConcurrentFanInFanOutCancel(n int) {
	done := make(chan struct{})
	defer close(done)
	fanOut := NumGenCancel(n, done)
	fanIn := make(chan string)
	for i := 0; i < 100; i++ {
		go func() {
			for n := range fanOut {
				if r, e := GetFBB(n); e != nil {
					fmt.Println("Error: %s", e)
					return
				} else {
					fanIn <- r
				}
			}
		}()
	}
	for resp := range fanIn {
		fmt.Println(resp)
	}
}

// FBBHTTPConcurrentFanInFanOutCancel OMIT

// FBBHTTP OMIT
func FBBHTTP(n int) {
	for i := 1; i <= n; i++ {
		if r, e := GetFBB(i); e != nil {
			fmt.Println("Error: %s", e)
		} else {
			fmt.Println(r)
		}
	}
}

// FBBHTTP OMIT

func main() {
	// FBBHTTPM OMIT
	//FBBHTTP(30)
	// FBBHTTPM OMIT
	// FBBHTTPConcurrentM OMIT
	//FBBHTTPConcurrent(30)
	// FBBHTTPConcurrentM OMIT
	// FBBHTTPConcurrentSortedM OMIT
	//FBBHTTPConcurrentSorted(100)
	// FBBHTTPConcurrentSortedM OMIT

	// FBBHTTPConcurrentFanInFanOutM OMIT
	//FBBHTTPConcurrentFanInFanOut(NumGen(100000))
	// FBBHTTPConcurrentFanInFanOutM OMIT

	// FBBHTTPConcurrentFanInFanOutCancelM OMIT
	//FBBHTTPConcurrentFanInFanOutCancel(1000000)
	// FBBHTTPConcurrentFanInFanOutCancelM OMIT

	// FBBHTTPConcurrentSortedCancelM OMIT
	//FBBHTTPConcurrentSortedCancel(100)
	// FBBHTTPConcurrentSortedCancelM OMIT
}
