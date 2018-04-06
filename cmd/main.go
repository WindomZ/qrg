package main

import (
	"flag"
	"os"

	"github.com/WindomZ/qrg"
)

const (
	usageLevel  = `set QR Code quality level: [0~3]`
	usageInvert = `reverse QR Code display color`
)

func main() {
	c := qrg.Config{}

	flag.IntVar(&c.Level, "l", 0, usageLevel)
	flag.BoolVar(&c.Invert, "i", false, usageInvert)

	flag.Parse()

	qrg.QRGenerate(os.Stdin, os.Stdout, qrg.ModeTerminal, c)
}
