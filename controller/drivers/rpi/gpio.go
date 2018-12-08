package rpi

import (
	"fmt"
	"sort"

	"github.com/kidoman/embd"

	"github.com/reef-pi/reef-pi/controller/types/driver"
)

var (
	validGPIOPins = map[int]bool{
		2:  true,
		3:  true,
		4:  true,
		5:  true,
		6:  true,
		7:  true,
		8:  true,
		9:  true,
		10: true,
		11: true,
		12: true,
		13: true,
		14: true,
		15: true,
		16: true,
		17: true,
		18: true,
		19: true,
		20: true,
		21: true,
		22: true,
		23: true,
		24: true,
		25: true,
		26: true,
		27: true,
	}
)

type rpiPin struct {
	pin       int
	name      string
	lastState bool

	digitalPin embd.DigitalPin
}

func (p *rpiPin) Name() string {
	return p.name
}

func (p *rpiPin) Close() error {
	return p.digitalPin.Close()
}

func (p *rpiPin) Read() (bool, error) {
	err := p.digitalPin.SetDirection(embd.In)
	if err != nil {
		return false, fmt.Errorf("can't read input from channel %d: %v", p.pin, err)
	}

	v, err := p.digitalPin.Read()
	if err != nil {
		return false, err
	}
	return v == 1, nil
}

func (p *rpiPin) Write(state bool) error {
	err := p.digitalPin.SetDirection(embd.Out)
	if err != nil {
		return fmt.Errorf("can't set output on channel %d: %v", p.pin, err)
	}
	value := 0
	if state {
		value = 1
	}
	p.lastState = state
	return p.digitalPin.Write(value)
}

func (p *rpiPin) LastState() bool {
	return p.lastState
}

func (r *rpiDriver) InputPins() []driver.InputPin {
	var pins []driver.InputPin
	for _, pin := range r.pins {
		pins = append(pins, pin)
	}
	sort.Slice(pins, func(i, j int) bool { return pins[i].Name() < pins[j].Name() })
	return pins
}

func (r *rpiDriver) GetInputPin(name string) (driver.InputPin, error) {
	pin, ok := r.pins[name]
	if !ok {
		return nil, fmt.Errorf("pin %s unknown", name)
	}
	return pin, nil
}

func (r *rpiDriver) OutputPins() []driver.OutputPin {
	var pins []driver.OutputPin
	for _, pin := range r.pins {
		pins = append(pins, pin)
	}
	sort.Slice(pins, func(i, j int) bool { return pins[i].Name() < pins[j].Name() })
	return pins
}

func (r *rpiDriver) GetOutputPin(name string) (driver.OutputPin, error) {
	pin, ok := r.pins[name]
	if !ok {
		return nil, fmt.Errorf("pin %s unknown", name)
	}
	return pin, nil
}
