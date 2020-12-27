package main

import (
	"image"
	"image/png"
	"image/color"
	"math/cmplx"
	"os"
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

type MandelbrotImg struct{}

func (m MandelbrotImg) ColorModel() color.Model {
	return color.RGBAModel
}

func (m MandelbrotImg) Bounds() image.Rectangle {
	return image.Rect(-2500, -1000, 1000, 1000) //x0, y0, w, h
}

func (m MandelbrotImg) At(x, y int) color.Color {
	iters := mandel_point(complex(float64(x)/1000.0, float64(y)/1000.0))
	if iters < 0 {
		return color.RGBA{0, 0, 0, 255}
	}
	return color.RGBA{uint8(iters%10)*25, uint8((iters/10)%10)*25, uint8((iters/100)%10)*25, 255}
}

func main() {
	f, _ := os.Create("mbrot.png")
	m := MandelbrotImg{}
	png.Encode(f, m)
	f.Close()
}
