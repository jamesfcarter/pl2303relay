[![PkgGoDev](https://pkg.go.dev/badge/github.com/jamesfcarter/pl2303relay)](https://pkg.go.dev/github.com/jamesfcarter/pl2303relay)

# pl2303relay

A Go driver for common USB-attached relay boards using the PL2303
USB to RS232 chip.

The package comes with a simple command-line tool to control relays.
For example, to turn on both relays on a two relay board:
```
go run github.com/jamesfcarter/pl2303relay/cmd/relay /dev/ttyUSB0 3
```

