package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/WindomZ/qrg"
)

func main() {
	c := qrg.Config{}

	flag.IntVar(&c.Level, "l", 0, `set QR Code quality level: [0~3]`)
	flag.BoolVar(&c.Invert, "i", false, `reverse QR Code display color`)

	var outputFile string
	flag.StringVar(&outputFile, "o", "", "output QR Code file: [.png|.jpeg|.text|.?]")
	flag.Parse()

	if outputFile == "" {
		qrg.QRGenerate(os.Stdin, os.Stdout, qrg.ModeTerminal, c)
	} else {
		f, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		defer f.Close()

		switch strings.ToLower(path.Ext(outputFile)) {
		case ".png":
			qrg.QRGenerate(os.Stdin, f, qrg.ModePNG, c)
		case ".jpeg", ".jpg":
			qrg.QRGenerate(os.Stdin, f, qrg.ModeJPEG, c)
		default:
			qrg.QRGenerate(os.Stdin, f, qrg.ModeText, c)
		}
	}
}
