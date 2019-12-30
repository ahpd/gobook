// Lissajous generates GIF animations of random Lissajous figures.

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xFF, 0x00, 0xff}}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		cycles := readIntParam(r, "cycles", 5)
		size := readIntParam(r, "size", 100)
		nframes := readIntParam(r, "nframes", 64)
		delay := readIntParam(r, "delay", 8)
		lissajous(w, cycles, size, nframes, delay)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func readIntParam(r *http.Request, paramName string, defaultValue int) int {
	paramValue := defaultValue
	val := r.FormValue(paramName)
	if val != "" {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			log.Print(err)
		} else {
			paramValue = intVal
		}
	}
	return paramValue
}

func lissajous(out io.Writer, cycles int, size int, nframes int, delay int) {

	const (
		res = 0.001
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		dim := 2*size + 1
		rect := image.Rect(0, 0, dim, dim)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(
				size+int(x*float64(size)+0.5),
				size+int(y*float64(size)+0.5),
				blackIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)

}
