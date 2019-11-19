package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func main() {

	var imagehtml string

	for i := 1; i <= 5; i++ {
		makeImage(i)
	}

	input, err := ioutil.ReadFile("template.html")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output := bytes.Replace(input, []byte("{CONTENT}"), []byte(imagehtml), -1)

	if err = ioutil.WriteFile("index.html", output, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func makeImage(i int) {
	flag.Parse()
	rand.Seed(time.Now().UTC().UnixNano())

	out, err := os.Create("./output" + string(i) + ".jpg")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// generate some QR code look a like image

	imgRect := image.Rect(0, 0, 300, 300)
	img := image.NewGray(imgRect)
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)
	for y := 0; y < 300; y += 10 {
		for x := 0; x < 300; x += 10 {
			fill := &image.Uniform{color.Black}
			if rand.Intn(10)%2 == 0 {
				fill = &image.Uniform{color.White}
			}
			draw.Draw(img, image.Rect(x, y, x+10, y+10), fill, image.ZP, draw.Src)
		}
	}

	var opt jpeg.Options

	opt.Quality = 80
	// ok, write out the data into the new JPEG file

	err = jpeg.Encode(out, img, &opt) // put quality to 80%
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
