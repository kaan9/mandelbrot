package main

import (
	"fmt"
	"math/cmplx"
)

// number of iterations
// 0 if doesn't terminate within iteration limit
func mandel_point(c complex128) int {
	const threshold = 1000
	var z complex128
	var i int
	for z, i = 0+0i, 0; cmplx.Abs(z) < 2*2 && i < threshold; i++ {
		z = z*z+c
	}
	if i >= threshold {
		return -1
	}
	return i
}


func main() {
	for y := -30; y <= 30; y++ {
		for x := -30; x <= 30; x++ {
			z := complex(float64(x)/30.0, float64(y)/30.0)
			iters := mandel_point(z)
			if iters == -1 {
				fmt.Print(0)
			} else {
				fmt.Print(iters%10)
			}
		}
		fmt.Println()
	}
	c1, c2, c3 := 0.0+0.0i, 0.5+0.5i, -1+0.5i
	fmt.Println(mandel_point(c1), mandel_point(c2), mandel_point(c3))
}
