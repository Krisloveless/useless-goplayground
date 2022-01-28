package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"io/ioutil"
	"os"
	"strings"

	"github.com/golang/freetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func main() {
	bstream, err := os.ReadFile("x.gif")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
	breader := bytes.NewReader(bstream)
	pgif, err := gif.DecodeAll(breader)
	if err != nil {
		panic(err)
	}
	// 57 pgif.Image
	//frames := len(pgif.Image)
	//height, width := pgif.Image[0].Rect.Max.Y-pgif.Image[0].Rect.Min.Y, pgif.Image[0].Rect.Max.X-pgif.Image[0].Rect.Min.X
	pointX, pointY := x, y
	fmt.Print(pointX, pointY)
	//DrawDots(pgif, pointX, pointY)
	DrawBoard(pgif, x, y)
	// for _, d := range pgif.Delay {
	// 	fmt.Println(d)
	// }
	// step(1, pgif)
	f, err := os.OpenFile("x_out.gif", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = gif.EncodeAll(f, pgif)
	if err != nil {
		panic(err)
	}

}

var x = []int{20, 58, 100, 144, 191, 237, 283, 326, 365, 400, 400, 365, 326, 283, 237, 191, 144, 100, 58, 20, 20, 58, 100, 144, 191, 237, 283, 326, 365, 400, 400, 365, 326, 283, 237, 191, 144, 100, 58, 20, 20, 58, 100, 144, 191, 237, 283, 326, 365, 400, 400, 365, 326, 283, 237, 191, 144, 100, 58, 20}
var y = []int{426, 467, 497, 516, 524, 521, 506, 479, 443, 397, 397, 443, 479, 506, 521, 524, 516, 497, 467, 426, 426, 467, 497, 516, 524, 521, 506, 479, 443, 397, 397, 443, 479, 506, 521, 524, 516, 497, 467, 426, 426, 467, 497, 516, 524, 521, 506, 479, 443, 397, 397, 443, 479, 506, 521, 524, 516, 497, 467, 426}

func DrawDots(pgif *gif.GIF, x, y []int) {
	for i := 0; i < len(pgif.Delay); i++ {
		// sp is where src starts to crop, the size and position is from r
		draw.Draw(pgif.Image[i], image.Rect(x[i], y[i], x[i]+34, y[i]+34), &image.Uniform{image.Black}, image.Point{0, 0}, draw.Over)
	}
}

func DrawBoard(pgif *gif.GIF, x, y []int) {
	message := []string{"this is a very very long line"}
	for i, img := range pgif.Image {
		addLabel(img, x[i], y[i], message[0], color.RGBA{0, 0, 0, 255})
		//addLabel2(img, x[i], y[i], message, color.RGBA{0, 0, 0, 255})
	}
}

func addLabel(img draw.Image, x, y int, label string, col color.Color) {
	words := strings.Split(label, " ")
	m := len(words)
	dy := 10 + (76 / 2) - ((m * 13) / 2)

	for _, word := range words {
		n := len(word)
		dx := 4 + (100 / 2) - ((n * 7) / 2)
		point := fixed.Point26_6{
			fixed.Int26_6((dx + x) * 64),
			fixed.Int26_6((dy + y) * 64),
		}
		face := basicfont.Face7x13
		d := &font.Drawer{
			Dst:  img,
			Src:  image.NewUniform(col),
			Face: face,
			Dot:  point,
		}
		d.DrawString(word)
		dy += 13
	}
}

func addLabel2(img draw.Image, x, y int, label []string, col color.Color) {
	c := freetype.NewContext()
	size := float64(34)
	fontBytes, err := ioutil.ReadFile("luximr.ttf")
	if err != nil {
		panic(err)
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		panic(err)
	}
	c.SetFont(f)
	c.SetFontSize(size)
	c.SetDst(img)
	c.SetSrc(&image.Uniform{col})
	pt := freetype.Pt(x, y)
	for _, s := range label {
		_, err := c.DrawString(s, pt)
		if err != nil {
			panic(err)
		}
		pt.Y += c.PointToFixed(size * 1.5)
	}
}

func step(value int, pgif *gif.GIF) {
	total := len(pgif.Image)
	var arrayP []*image.Paletted
	var arrayD []int
	var arrayDis []byte
	if value == 0 {
		value = 1
	}
	for i := 0; i < total; {
		arrayP = append(arrayP, pgif.Image[i])
		arrayD = append(arrayD, 0)
		arrayDis = append(arrayDis, pgif.Disposal[i])
		i += value
	}
	pgif.Image = arrayP
	pgif.Delay = arrayD
	pgif.Disposal = arrayDis
}
