package main

import (
	"flag"
	"os"

	"github.com/WindomZ/qrg"
)

func main() {
	c := qrg.Config{}

	flag.IntVar(&c.Level, "l", 0, `set QR Code quality level: [0~3]`)
	flag.BoolVar(&c.Invert, "i", false, `reverse QR Code display color`)

	flag.Parse()

	qrg.QRGenerate(os.Stdin, os.Stdout, qrg.ModeTerminal, c)
}
