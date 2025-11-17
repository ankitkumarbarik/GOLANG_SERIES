package main

import (
	"fmt"
	"sync"
)

func work(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("doing task", id)
}

func main() {
	fmt.Println("Hello")
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go work(i, &wg)
	}

	wg.Wait()

}

/*
1️⃣ The loop in main() runs 5 times

That means:
wg.Add(1) runs 5 times
Every time Add(1) runs → counter increases by 1
Counter starts at 0
After 5 Add(1) calls → counter becomes 5

This means:
👉 "There are 5 goroutines that we need to wait for."

2️⃣ go work(i, &wg) also runs 5 times

So 5 workers (goroutines) are started in the background
running parallel to the main function.

3️⃣ Every worker, when it finishes, calls:
wg.Done()

wg.Done() means:
👉 “My task is finished. Reduce the counter by 1.”

So the counter goes:
5 → 4 → 3 → 2 → 1 → 0

4️⃣ wg.Wait() checks the counter

If counter is not 0, it waits
If counter becomes 0, Wait() returns → program continues → exits
*/
