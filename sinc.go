package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime/trace"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	var state = make(map[int]int)

	var mutex = &sync.Mutex{}

	var readOp uint64
	var writeOp uint64

	for r := 0; r < 2; r++ {
		go func() {
			total := 0
			for {

				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOp, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 1; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOp, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOp)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOp)
	fmt.Println("writeOps:", writeOpsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
