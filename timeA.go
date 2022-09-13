package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	signalTerm()
}

func signalTerm() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	timer := time.After(10 * time.Second)
	select {
	case <-timer:
		fmt.Println("time is over")
		return
	case s := <-sig:
		fmt.Println("Stopped by signal: ", s)
		return
	}

}
