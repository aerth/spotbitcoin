package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"io/ioutil"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

func drawpng(s string) (image.Image, error) {

	dest := image.NewRGBA(image.Rect(0, 0, 550, 22))

	draw.Draw(dest, dest.Bounds(), &image.Uniform{color.Black}, image.ZP, draw.Src)

	// grab font
	fontBytes, err := ioutil.ReadFile("TerminusTTF-4.40.1.ttf")
	if err != nil {
		return nil, err
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return nil, err
	}

	// font options
	opts := &truetype.Options{}
	opts.DPI = 96
	opts.Size = 14
	opts.Hinting = font.HintingNone

	// write white text on the (already) black background
	d := font.Drawer{}
	d.Dst = dest
	d.Src = image.White
	d.Face = truetype.NewFace(f, opts)
	d.Dot = freetype.Pt(10, 15)
	d.DrawString(s)
	return dest, nil
}

func WritePNG(img image.Image, w io.Writer) {
	encoder := png.Encoder{}
	encoder.CompressionLevel = png.DefaultCompression
	encoder.Encode(w, img)
}
