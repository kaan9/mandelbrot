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
	const threshold = 32765 //32^3 for 32-bit rgba vals`
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

// the dimensions of the image are (-25, -10)*scale to (10, 10)*scale
// normally it'd be (-2.5, 1)*scale but that would invalidate odd scale values and make the image too small
const scale = 100

type MandelbrotImg struct{}

func (m MandelbrotImg) ColorModel() color.Model {
	return color.RGBAModel
}

func (m MandelbrotImg) Bounds() image.Rectangle {
	return image.Rect(-25*scale, -10*scale, 10*scale, 10*scale) //x0, y0, w, h
}

func (m MandelbrotImg) At(x, y int) color.Color {
	iters := mandel_point(complex(float64(x)/scale/10.0, float64(y)/scale/10.0))
	if iters < 0 {
		return color.RGBA{0, 0, 0, 255}
	}
	return color.RGBA{uint8(iters%32)*8+7, uint8((iters/32)%32)*8+7, uint8((iters/32/32)%(32*32))*8+7, 255}
}

func main() {
	f, _ := os.Create("mbrot.png")
	m := MandelbrotImg{}
	png.Encode(f, m)
	f.Close()
}
