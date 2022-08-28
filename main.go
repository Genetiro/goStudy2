package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func Squre(r float64) (res float64, err error) {

	defer func() {
		if v := recover(); v != nil {
			fmt.Println("recovered2", v)
			err = v.(error)
		}
	}()

	defer func() {
		if v := recover(); v != nil {
			fmt.Println("recovered", v)
			panic(v)
		}
	}()
	//var r float64

	if r < 0 {
		// panic("r must be > 0")
		panic(errors.New("r must be > 0"))

	}
	return math.Pi * r * r, nil

}

func main() {
	var r float64
	fmt.Println("enter radius")
	fmt.Fscan(os.Stdin, &r)
	res, err := Squre(r)
	fmt.Println(res, err)

}
