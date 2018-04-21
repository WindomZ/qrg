# qrg - QR Code Generator

> Generate QR Code on the terminal.

[![Build Status](https://travis-ci.org/WindomZ/qrg.svg?branch=master)](https://travis-ci.org/WindomZ/qrg)
[![Coverage Status](https://coveralls.io/repos/github/WindomZ/qrg/badge.svg?branch=master)](https://coveralls.io/github/WindomZ/qrg?branch=master)

## Usage
help:
```bash
qrg -h

Usage of qrg:
  -i    reverse QR Code display color
  -l int
        set QR Code quality level: [0~3]
  -o string
        output QR Code file: [.png|.jpeg|.text|.?]
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

## License
[MIT](https://github.com/WindomZ/qrg/blob/master/LICENSE)
