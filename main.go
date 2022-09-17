package main

import (
	"machine"
	"time"
)

func main() {
	// Turn on the LED.
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led.High()
	for {
		// Turn on and off all of the GPIOs on the board.
		machine.GP0.Configure(machine.PinConfig{Mode: machine.PinOutput})
		machine.GP0.High()
		time.Sleep(time.Second * 1)
		machine.GP0.Low()
		machine.GP1.Configure(machine.PinConfig{Mode: machine.PinOutput})
		machine.GP1.High()
		time.Sleep(time.Second * 1)
		machine.GP1.Low()
		machine.GP2.Configure(machine.PinConfig{Mode: machine.PinOutput})
		machine.GP2.High()
		time.Sleep(time.Second * 1)
		machine.GP2.Low()
		machine.GP3.Configure(machine.PinConfig{Mode: machine.PinOutput})
		machine.GP3.High()
		time.Sleep(time.Second * 1)
		machine.GP3.Low()
		machine.GP4.Configure(machine.PinConfig{Mode: machine.PinOutput})
		machine.GP4.High()
		time.Sleep(time.Second * 1)
		machine.GP4.Low()
		machine.GP5.Configure(machine.PinConfig{Mode: machine.PinOutput})
		machine.GP5.High()
		time.Sleep(time.Second * 1)
		machine.GP5.Low()
		machine.GP6.Configure(machine.PinConfig{Mode: machine.PinOutput})
		machine.GP6.High()
		time.Sleep(time.Second * 1)
		machine.GP6.Low()
		machine.GP7.Configure(machine.PinConfig{Mode: machine.PinOutput})
		machine.GP7.High()
		time.Sleep(time.Second * 1)
		machine.GP7.Low()
		machine.GP8.Configure(machine.PinConfig{Mode: machine.PinOutput})
		machine.GP8.High()
		time.Sleep(time.Second * 1)
		machine.GP8.Low()
		machine.GP9.Configure(machine.PinConfig{Mode: machine.PinOutput})
		machine.GP9.High()
		time.Sleep(time.Second * 1)
		machine.GP9.Low()
		machine.GP10.Configure(machine.PinConfig{Mode: machine.PinOutput})
		machine.GP10.High()
		time.Sleep(time.Second * 1)
		machine.GP10.Low()
		machine.GP11.Configure(machine.PinConfig{Mode: machine.PinOutput})
		machine.GP11.High()
		time.Sleep(time.Second * 1)
		machine.GP11.Low()
		machine.GP12.Configure(machine.PinConfig{Mode: machine.PinOutput})
		machine.GP12.High()
		time.Sleep(time.Second * 1)
		machine.GP12.Low()
		machine.GP13.Configure(machine.PinConfig{Mode: machine.PinOutput})
		machine.GP13.High()
		time.Sleep(time.Second * 1)
		machine.GP13.Low()
		machine.GP14.Configure(machine.PinConfig{Mode: machine.PinOutput})
		machine.GP14.High()
		time.Sleep(time.Second * 1)
		machine.GP14.Low()
		machine.GP15.Configure(machine.PinConfig{Mode: machine.PinOutput})
		machine.GP15.High()
		time.Sleep(time.Second * 1)
		machine.GP15.Low()
		machine.GP16.Configure(machine.PinConfig{Mode: machine.PinOutput})
		machine.GP16.High()
		time.Sleep(time.Second * 1)
		machine.GP16.Low()
		machine.GP17.Configure(machine.PinConfig{Mode: machine.PinOutput})
		machine.GP17.High()
		time.Sleep(time.Second * 1)
		machine.GP17.Low()
		machine.GP18.Configure(machine.PinConfig{Mode: machine.PinOutput})
		machine.GP18.High()
		time.Sleep(time.Second * 1)
		machine.GP18.Low()
		machine.GP19.Configure(machine.PinConfig{Mode: machine.PinOutput})
		machine.GP19.High()
		time.Sleep(time.Second * 1)
		machine.GP19.Low()
		machine.GP20.Configure(machine.PinConfig{Mode: machine.PinOutput})
		machine.GP20.High()
		time.Sleep(time.Second * 1)
		machine.GP20.Low()
		machine.GP21.Configure(machine.PinConfig{Mode: machine.PinOutput})
		machine.GP21.High()
		time.Sleep(time.Second * 1)
		machine.GP21.Low()
		machine.GP22.Configure(machine.PinConfig{Mode: machine.PinOutput})
		machine.GP22.High()
		time.Sleep(time.Second * 1)
		machine.GP22.Low()
		machine.GP26.Configure(machine.PinConfig{Mode: machine.PinOutput})
		machine.GP26.High()
		time.Sleep(time.Second * 1)
		machine.GP26.Low()
		machine.GP27.Configure(machine.PinConfig{Mode: machine.PinOutput})
		machine.GP27.High()
		time.Sleep(time.Second * 1)
		machine.GP27.Low()
		machine.GP28.Configure(machine.PinConfig{Mode: machine.PinOutput})
		machine.GP28.High()
		time.Sleep(time.Second * 1)
		machine.GP28.Low()
	}
}
