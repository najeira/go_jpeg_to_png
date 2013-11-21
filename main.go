package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"strings"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("no files")
		return
	}
	filenames := args[1:]
	for _, filename := range filenames {
		err := convert(filename)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(filename)
		}
	}
}

func read(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return jpeg.Decode(file)
}

func convert(filename string) error {
	ext := path.Ext(filename)
	pngFilename := strings.TrimSuffix(filename, ext) + ".png"

	jpeg, err := read(filename)
	if err != nil {
		return err
	}

	file, err := os.Create(pngFilename)
	if err != nil {
		return err
	}

	err = png.Encode(file, jpeg)
	if err != nil {
		os.Remove(pngFilename)
		return err
	}
	return nil
}
