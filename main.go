package main

import (
	"machine"
)

func main() {
	// Turn on the LED.
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led.High()
}
