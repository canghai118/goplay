package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"image/gif"
	"image"
	"bytes"
	"image/png"
)

func main() {
	flag.Parse()

	gifFile := flag.Args()[0]

	fmt.Print(gifFile)

	file, err := os.Open(gifFile)

	if err != nil {
		panic(err)
	}

	gifImg, err := gif.DecodeAll(file)

	revertGif(gifImg)

	writeGif(gifImg, "revert.gif")
}

func writeGif(g *gif.GIF, fileName string) error {
	var buf bytes.Buffer

	if err := gif.EncodeAll(&buf, g); err != nil {
		return err
	}

	if err := ioutil.WriteFile(fileName, buf.Bytes(), 0x755); err != nil {
		return err
	}

	return nil
}

func writeImg(img image.Image, fileName string) error {
	var buf bytes.Buffer

	if err := png.Encode(&buf, img); err != nil {
		return err
	}

	if err := ioutil.WriteFile(fileName, buf.Bytes(), 0x755); err != nil {
		return err
	}

	return nil
}

func revertGif(g *gif.GIF) {
	l := len(g.Image)
	for i := 0; i < l/2; i ++ {
		g.Image[i], g.Image[l-i-1] = g.Image[l-i-1], g.Image[i]
	}

	l = len(g.Delay)

	for i := 0; i < l/2; i ++ {
		g.Delay[i], g.Delay[l-i-1] = g.Delay[l-i-1], g.Delay[i]
	}
}

func revert(p *[]interface{}) {
	arr := *p
	l := len(arr)
	for i := 0; i < l/2; i ++ {
		t := arr[l-i-1]
		arr[l-i-1] = arr[i]
		arr[i] = t
	}
}

func doSth(arr []interface{}) {

}
