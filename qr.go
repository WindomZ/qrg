package qrg

import (
	"errors"
	"io"
	"io/ioutil"

	"github.com/WindomZ/qrw"
)

// Mode generate QR Code mode
type Mode int

const (
	// ModeTerminal print on the terminal
	ModeTerminal = iota
	// ModeText output a file with text
	ModeText
	// ModePNG output a png file
	ModePNG
	// ModeJPEG output a jpeg/jpg file
	ModeJPEG
)

// DefaultQRLevel the default level of QR Code
var DefaultQRLevel = qrw.L

var (
	// ErrNil nil error
	ErrNil = errors.New("input/output should not be nil")
	// ErrEmpty empty input error
	ErrEmpty = errors.New("input should not be empty")
)

// QRGenerate generate QR Code
func QRGenerate(r io.Reader, w io.Writer, m Mode, cs ...Config) error {
	if r == nil || w == nil {
		return ErrNil
	}

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	} else if len(data) == 0 {
		return ErrEmpty
	}

	level := DefaultQRLevel
	invert := false

	if len(cs) != 0 {
		c := cs[0]
		switch c.Level {
		case 1:
			level = qrw.M
		case 2:
			level = qrw.Q
		case 3:
			level = qrw.H
		default:
			level = qrw.L
		}
		invert = c.Invert
	}

	switch m {
	case ModeText:
		wrt := qrw.NewBlockWriter(level, w)
		if invert {
			wrt.Invert()
		}
		return wrt.QR(string(data))
	case ModePNG:
		wrt := qrw.NewPNGWriter(level, w)
		return wrt.QR(string(data))
	case ModeJPEG:
		wrt := qrw.NewJPEGWriter(level, w)
		return wrt.QR(string(data))
	default:
		wrt := qrw.NewBashWriter(level, w)
		if invert {
			wrt.Invert()
		}
		return wrt.QR(string(data))
	}
}
