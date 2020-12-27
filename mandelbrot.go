package main

import (
	"image"
	"image/png"
	"image/color"
	"math/cmplx"
	"os"
	"sync"
)

// number of iterations
// 0 if doesn't terminate within iteration limit
func mandel_point(c complex128) int {
	const threshold = 4096 //16^3 for 32-bit rgba vals`
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

// returs a 2d array iters that ranges form iters[0][0] to iters[|x0|+|x1|+1][|y0|+|y1|+1]
// e.g scale=100, x0=-2500, y0=-1000, x1=y1=1000 will give a picture of the range (-2.5,-1) to (1,1)
func mandel_range(scale, x0, y0, x1, y1 int) [][]int {
	iters := make([][]int, x1-x0+1)
	var wg sync.WaitGroup

	for x := x0; x <= x1; x++ {
		wg.Add(1)
		go func(x int) {
			iters[x-x0] = make([]int, y1-y0+1)
			for y := y0; y <= y1; y++ {
				iters[x-x0][y-y0] = mandel_point(complex(float64(x)/float64(scale), float64(y)/float64(scale))) 
			}
			wg.Done()
		}(x)
	}
	wg.Wait()
	return iters
}

// the dimensions of the image are (-25, -10)*scale to (10, 10)*scale
// normally it'd be (-2.5, 1)*scale but that would invalidate odd scale values and make the image too small
//const scale = 100

type MandelbrotImg struct{
	scale int
	x0, y0, x1, y1 int
	iters [][]int
}

func (m *MandelbrotImg) ColorModel() color.Model {
	return color.RGBAModel
}

func (m *MandelbrotImg) Bounds() image.Rectangle {
	return image.Rect(m.x0, m.y0, m.x1, m.y1)
//	return image.Rect(-25*scale, -10*scale, 10*scale, 10*scale) //x0, y0, w, h
}

func (m *MandelbrotImg) At(x, y int) color.Color {
//	iters := mandel_point(complex(float64(x)/scale/10.0, float64(y)/scale/10.0))
	iters := m.iters[x-m.x0][y-m.y0]
	if iters < 0 {
		return color.RGBA{0, 0, 0, 255}
	}
	return color.RGBA{uint8(iters%16)*16+15, uint8((iters/16)%16)*16+15, uint8((iters/256)%16)*16+15, 255}
}

func main() {
	f, _ := os.Create("mbrot.png")
	m := MandelbrotImg{
		scale: 1000,
		x0: -2500,
		y0: -1000,
		x1: 1000,
		y1: 1000,
		iters: nil,
	}
	m.iters = mandel_range(m.scale, m.x0, m.y0, m.x1, m.y1)
	png.Encode(f, &m)
	f.Close()
}
