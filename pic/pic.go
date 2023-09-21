package pic

import (
	"image"
	"os"

	"github.com/mattn/go-sixel"
)

// Show displays a picture defined by the function f
// encoded as a sixtel image.
//
// f should return a slice of length dy,
// each element of which is a slice of dx
// 8-bit unsigned int. The integers are
// interpreted as bluescale values,
// where the value 0 means full blue,
// and the value 255 means full white.
func Show(f func(dx, dy int) [][]uint8) {
	const (
		dx = 256
		dy = 256
	)
	data := f(dx, dy)
	m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] = v
			m.Pix[i+1] = v
			m.Pix[i+2] = 255
			m.Pix[i+3] = 255
		}
	}
	sixel.NewEncoder(os.Stdout).Encode(m)
}
