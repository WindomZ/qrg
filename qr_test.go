package qrg

import (
	"bytes"
	"os"
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestQRGenerate(t *testing.T) {
	data := []byte("hello world!")

	assert.Error(t, QRGenerate(nil, os.Stdout, ModeTerminal))
	assert.Error(t, QRGenerate(bytes.NewReader(data), nil, ModeTerminal))
	assert.Error(t, QRGenerate(bytes.NewReader([]byte{}), os.Stdout, ModeTerminal))

	assert.NoError(t,
		QRGenerate(bytes.NewReader(data), os.Stdout, ModeTerminal,
			Config{Level: 0, Invert: true}))
	assert.NoError(t,
		QRGenerate(bytes.NewReader(data), os.Stdout, ModeTerminal,
			Config{Level: 1, Invert: false}))
	assert.NoError(t,
		QRGenerate(bytes.NewReader(data), os.Stdout, ModeTerminal,
			Config{Level: 2, Invert: false}))
	assert.NoError(t,
		QRGenerate(bytes.NewReader(data), os.Stdout, ModeTerminal,
			Config{Level: 3, Invert: true}))

	if f, err := os.OpenFile("./test.txt",
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666); assert.NoError(t, err) {
		QRGenerate(bytes.NewReader(data), f, ModeText)
		f.Close()
	}
	if f, err := os.OpenFile("./test.txt",
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666); assert.NoError(t, err) {
		QRGenerate(bytes.NewReader(data), f, ModeText, Config{Level: 0, Invert: true})
		f.Close()
	}
	if f, err := os.OpenFile("./test.png",
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666); assert.NoError(t, err) {
		QRGenerate(bytes.NewReader(data), f, ModePNG)
		f.Close()
	}
	if f, err := os.OpenFile("./test.jpeg",
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666); assert.NoError(t, err) {
		QRGenerate(bytes.NewReader(data), f, ModeJPEG)
		f.Close()
	}
}
