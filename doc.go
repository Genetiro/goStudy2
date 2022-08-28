// Package errors implements functions to demonstrate godoc features
//
// The Squre returns circle square
//
// Squre() float64
//
// So, you are allowed to call this function and get some float64
package doc

import "math"

func Squre(r float64) (res float64, err error) {

	return math.Pi * r * r, nil

}
