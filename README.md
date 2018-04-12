# QR Code Generator - CLI

> Generate QR Code on the terminal.

## Usage
print help:
```bash
qrg -h

Usage of qrg:
  -i    reverse QR Code display color
  -l int
        set QR Code quality level: [0~3]
```

from text:
```bash
echo 'hello world!' | qrg
```

from file:
```bash
cat 'qr.go' | qrg
```

## Install
```bash
go get -u github.com/WindomZ/qrg/...
```
