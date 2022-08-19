package main

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/uc8151"
	"tinygo.org/x/tinydraw"
)

var (
	display  uc8151.Device
	led      machine.Pin
	a_btn    machine.Pin
	b_btn    machine.Pin
	c_btn    machine.Pin
	up_btn   machine.Pin
	down_btn machine.Pin
)

var (
	black       = color.RGBA{1, 1, 1, 255}
	white       = color.RGBA{0, 0, 0, 255}
	w     int16 = 296
	h     int16 = 128
)

func main() {
	a_btn = machine.BUTTON_A
	a_btn.Configure(machine.PinConfig{Mode: machine.PinInput})
	b_btn = machine.BUTTON_B
	b_btn.Configure(machine.PinConfig{Mode: machine.PinInput})
	c_btn = machine.BUTTON_C
	c_btn.Configure(machine.PinConfig{Mode: machine.PinInput})
	up_btn = machine.BUTTON_UP
	up_btn.Configure(machine.PinConfig{Mode: machine.PinInput})
	down_btn = machine.BUTTON_DOWN
	down_btn.Configure(machine.PinConfig{Mode: machine.PinInput})

	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 12000000,
		SCK:       machine.EPD_SCK_PIN,
		SDO:       machine.EPD_SDO_PIN,
	})

	display = uc8151.New(machine.SPI0, machine.EPD_CS_PIN, machine.EPD_DC_PIN, machine.EPD_RESET_PIN, machine.EPD_BUSY_PIN)
	display.Configure(uc8151.Config{
		Rotation: uc8151.ROTATION_270,
		Speed:    uc8151.MEDIUM,
		Blocking: true,
	})

	display.ClearBuffer()
	display.Display()
	display.WaitUntilIdle()

	tinydraw.FilledRectangle(&display, 0, 0, 100, 100, black)

	display.Display()
	display.WaitUntilIdle()
}
