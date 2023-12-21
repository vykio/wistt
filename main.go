package main

import (
	"bytes"
	"github.com/skanehira/clipboard-image"
	"log"
	"os"
	"wistt/cmd"
	"wistt/image"
)

func main() {
	input, output, err := cmd.GetIO()
	if err != nil {
		log.Fatal(err)
	}

	b, err := image.GenerateBuffer(image.Image{Input: input, Output: output})
	if err != nil {
		log.Fatal(err)
	}

	reader := bytes.NewReader(b)
	if err := clipboard.CopyToClipboard(reader); err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
