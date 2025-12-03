package main

import (
	"fmt"
	"sync"
)

// NOTE:
// race condition problem in go language -> when multiple goroutines are accessing the same variable at the same time and changing the value of the variable.

// solution -> use mutex(mutual exclusion) to lock the variable and allow only one goroutine to access the variable at a time.

// --------------------------------------------------------------->

// type post struct{
// 	views int
// }

// func (p *post) inc(wg *sync.WaitGroup){
// 	defer func(){
// 		wg.Done()
// 	}()

// 	p.views += 1
// }

// func main() {

// 	var wg sync.WaitGroup

// 	myPost := post{
// 		views: 0,
// 	}

// 	for i := 0; i < 100; i++ {
// 		wg.Add(1)
// 		go myPost.inc(&wg)
// 	}

// 	wg.Wait()

// 	fmt.Println(myPost.views)
// }

// --------------------------------------------------------------->

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()

	p.mu.Lock()
	p.views += 1
}

func main() {

	var wg sync.WaitGroup

	myPost := post{
		views: 0,
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}

	wg.Wait()

	fmt.Println(myPost.views)
}
