package main

import (
	"bytes"
	"encoding/json"
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

	// Load up config data
	type Config struct {
		Images     int
		Paragraphs int
	}

	// Open our jsonFile
	jsonFile, err := os.Open("sitegenconfig.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	//defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var cfg *Config
	cfg = new(Config)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &cfg)

	var imagecount = cfg.Images
	//var paragraphs = cfg.Paragraphs

	// start HTML generation
	print(imagecount)

	var imagehtml = "<div class=\"row\">\n"

	for i := 1; i <= imagecount; i++ {
		imagehtml += "<div class=\"col-sm-2\"><img src=\"" + makeImage(i) + "\" /></div>\n"
	}

	imagehtml += "</div>\n</div>"

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

func makeImage(i int) string {

	//flag.Parse()
	rand.Seed(time.Now().UTC().UnixNano())

	var filename = fmt.Sprintf("image-%d.jpg", i)

	out, err := os.Create("./" + filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// declared outside the loop so values are consistent with x and y
	var randx = getRandom()
	var randy = getRandom()

	// generate some QR code look a like image
	imgRect := image.Rect(0, 0, randx, randy)
	img := image.NewGray(imgRect)
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)
	for y := 0; y < randy; y += 10 {
		for x := 0; x < randx; x += 10 {
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

	return filename
}

// generates a psuedorandom number from 100-300
func getRandom() int {

	rand.Seed(time.Now().UnixNano())
	min := 100
	max := 300
	return (rand.Intn(max-min+1) + min)
}
