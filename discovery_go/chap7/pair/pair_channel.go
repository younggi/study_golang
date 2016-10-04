package main

import (
	"fmt"
	"sync"
)

// Request is type with ID and Response channel
type Request struct {
	Num  int
	Resp chan Response
}

// Response is type with ID and Worker ID
type Response struct {
	Num      int
	WorkerID int
}

// PlusOneService returns Responses to requested channel
func PlusOneService(reqs <-chan Request, workerID int) {
	for req := range reqs {
		go func(req Request) {
			defer close(req.Resp)
			req.Resp <- Response{req.Num + 1, workerID}
		}(req)
	}
}

// Request channel x 1
// Response channel x 5
func main() {
	reqs := make(chan Request)
	defer close(reqs)
	for i := 0; i < 3; i++ {
		go PlusOneService(reqs, i)
	}
	var wg sync.WaitGroup
	for i := 3; i < 53; i += 10 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			resps := make(chan Response)
			reqs <- Request{i, resps}
			// multi responses
			for resp := range resps {
				fmt.Println(i, "=>", resp)
			}
			// or only single one
			//fmt.Println(i, "=>", <-resp)
		}(i)
	}
	wg.Wait()
}
