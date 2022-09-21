package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

type Count struct {
	mutex sync.Mutex
	a     map[string]int
}

func (b *Count) Increment(key string) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.a[key]++
}

func (b *Count) GetValue(key string) int {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.a[key]
}
func main() {
	key := "count"
	b := Count{a: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go b.Increment(key)
	}
	time.Sleep(3 * time.Second)
	fmt.Println(b.GetValue(key))
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
}
