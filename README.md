# qrg - QR Code Generator

> Generate QR Code on the terminal.

[![Build Status](https://travis-ci.org/WindomZ/qrg.svg?branch=master)](https://travis-ci.org/WindomZ/qrg)
[![Coverage Status](https://coveralls.io/repos/github/WindomZ/qrg/badge.svg?branch=master)](https://coveralls.io/github/WindomZ/qrg?branch=master)

## Usage
help:
```bash
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
# or
echo 'hello world!' | qrg -i -l 3
# or
echo 'hello world!' | qrg -o hello.txt
# or
echo 'hello world!' | qrg -l 2 -o hello.png
```

from file:
```bash
cat LICENSE | qrg -l 1
# or
cat LICENSE | qrg -o LICENSE.png
```

## Install
If you have a Golang environment:
```bash
go get -u github.com/WindomZ/qrg/...
```

Or download the latest binary [release](https://github.com/WindomZ/qrg/releases)

## Changelog
See [CHANGELOG.md](https://github.com/WindomZ/qrg/blob/master/CHANGELOG.md#readme)

## Contributing
Welcome to pull requests, report bugs, suggest ideas and discuss, 
i would love to hear what you think on [issues page](https://github.com/WindomZ/qrg/issues).

If you like it then you can put a :star: on it.

## License
[MIT](https://github.com/WindomZ/qrg/blob/master/LICENSE)
