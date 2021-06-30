package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ssd1306"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			println("an error really happened")
		}
	}()
	machine.I2C0.Configure(machine.I2CConfig{})
	device := ssd1306.NewI2C(machine.I2C0)
	device.Configure(ssd1306.Config{
		Address: 0x3C,
	})

	device.SetPixel(20, 20, color.RGBA{255, 255, 255, 255})

	device.Display()

	led := machine.D13
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledSwitch := true

	for {
		led.Set(ledSwitch)
		ledSwitch = !ledSwitch
		delay(250)
	}

}

func delay(t uint32) {
	time.Sleep(time.Duration(1000000 * t))
}
