package pl2303relay

import (
	"fmt"
	"io"
	"time"

	"github.com/tarm/serial"
)

const pause = 500 * time.Millisecond

// PL2303Relay represents a set of USB-connected relays.
type PL2303Relay struct {
	f io.ReadWriter
	// Value is the last value successfully written to the LEDs,
	// or nil if no value has yet been written.
	Value *byte
}

// New created a new PL2303Relay or returns an error if it is unable to
// open the device.
func New(device string) (*PL2303Relay, error) {
	config := &serial.Config{
		Name:        device,
		Baud:        9600,
		ReadTimeout: pause,
	}
	file, err := serial.OpenPort(config)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s: %w", device, err)
	}
	return &PL2303Relay{f: file}, nil
}

// Init initialises the relay board and returns the number of relays available.
// This call can only be made once after the board gets power. If a second
// calls made then 0 will be returned, and the relays will be in an undefined
// state.
func (p *PL2303Relay) Init() (count int, err error) {
	_, err = p.f.Write([]byte{0x50})
	if err != nil {
		err = fmt.Errorf("failed to write initialisation sequence: %w", err)
		return
	}
	buf := make([]byte, 1)
	count, err = p.f.Read(buf)
	if err == io.EOF {
		err = nil
		return
	}
	if err != nil || count == 0 {
		err = fmt.Errorf("failed to read board type: %w", err)
		return
	}
	switch buf[0] {
	case 0xad:
		count = 2
	case 0xab:
		count = 4
	case 0xac:
		count = 8
	default:
		return 0, fmt.Errorf("unexpected response %02x", buf[0])
	}
	_, err = p.f.Write([]byte{0x51})
	if err != nil {
		err = fmt.Errorf("failed to complete initialisation: %w", err)
	}
	time.Sleep(pause)
	return
}

// Update turns the relays on or off in a pattern described by b.
func (p *PL2303Relay) Update(b byte) error {
	_, err := p.f.Write([]byte{b})
	if err == nil {
		p.Value = &b
	}
	p.Value = &b
	return err
}
