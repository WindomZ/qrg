package qrg

import (
	"errors"
	"io"
	"io/ioutil"

	"github.com/WindomZ/qrw"
)

type Mode int

const (
	ModeTerminal = iota
	ModeText
	ModeHTTP
)

var DefaultQRLevel = qrw.L

var (
	ErrNil   = errors.New("input/output should not be nil")
	ErrEmpty = errors.New("input should not be empty")
)

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
	case ModeTerminal:
		wrt := qrw.NewBashWriter(level, w)
		if invert {
			wrt.Invert()
		}
		return wrt.QR(string(data))
	case ModeText:
		wrt := qrw.NewBlockWriter(level, w)
		if invert {
			wrt.Invert()
		}
		return wrt.QR(string(data))
	}
	return nil
}