package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	http.HandleFunc("/", handleLissajous)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handleLissajous(w http.ResponseWriter, r *http.Request) {
	path := fmt.Sprint("%s", r.URL)
	split := strings.Split(path, "=")
	if len(split) == 2 {
		cycles, error := strconv.Atoi(split[1])
		if error != nil {
			fmt.Fprintf(w, "%v\n", error)
			return
		}
		lissajous(w, cycles)
	}
	lissajous(w, 5)
}

func lissajous(out io.Writer, num_of_cyles int) {
	cycles := float64(num_of_cyles)
	const (
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), 
												size+int(y*size+0.5), 
												blackIndex)
			
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}