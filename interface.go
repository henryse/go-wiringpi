package wiringpi

import (
	"os"
	"fmt"
)

var (
	board2pin = []int{
		-1,
		-1,
		-1,
		8,
		-1,
		9,
		-1,
		7,
		15,
		-1,
		16,
		0,
		1,
		2,
		-1,
		-1,
		4,
		-1,
		5,
		12,
		-1,
		13,
		6,
		14,
		10,
		-1,
		11,
	}
	gpio2pin = []int{
		8,
		9,
		-1,
		-1,
		7,
		-1,
		-1,
		11,
		10,
		13,
		12,
		14,
		-1,
		-1,
		15,
		16,
		-1,
		0,
		1,
		-1,
		-1,
		2,
		3,
		4,
		5,
		6,
		-1,
		-1,
		17,
		18,
		19,
		20,
	}

	gpioModes = []string  {"IN", "OUT", "ALT5", "ALT4", "ALT0", "ALT1", "ALT2", "ALT3"}
)

const (
	PIN_GPIO_0  = 0
	PIN_GPIO_1  = 1
	PIN_GPIO_2  = 2
	PIN_GPIO_3  = 3
	PIN_GPIO_4  = 4
	PIN_GPIO_5  = 5
	PIN_GPIO_6  = 6
	PIN_GPIO_7  = 7
	PIN_SDA     = 8
	PIN_SCL     = 9
	PIN_CE0     = 10
	PIN_CE1     = 11
	PIN_MOSI    = 12
	PIN_MOSO    = 13
	PIN_SCLK    = 14
	PIN_TXD     = 15
	PIN_RXD     = 16
	PIN_GPIO_8  = 17
	PIN_GPIO_9  = 18
	PIN_GPIO_10 = 19
	PIN_GPIO_11 = 20

	MODE_IN  = 0
	MODE_OUT = 1
	MODE_ALT5 = 2
	MODE_ALT4 = 3
	MODE_ALT0 = 4
	MODE_ALT1 = 5
	MODE_ALT2 = 6
	MODE_ALT3 = 7
)

//use RPi.GPIO's BOARD numbering
func BoardToPin(pin int) int {
	if pin < 1 || pin >= len(board2pin) {
		panic(fmt.Sprintf("Invalid board pin number: %d", pin))
	}
	return board2pin[pin]
}

func GpioToPin(pin int) int {
	if pin < 0 || pin >= len(gpio2pin) {
		panic(fmt.Sprintf("Invalid bcm gpio number: %d", pin))
	}
	return gpio2pin[pin]
}

func PinToGpio(pin int) int {
	return internalPinToGpio(pin)
}

func Setup() error {
	return internalSetup()
}

func PinMode(pin int, mode int) {
	internalPinMode(pin, mode)
}

func DigitalWrite(pin int, mode int) {
	internalDigitalWrite(pin, mode)
}

func DigitalRead(pin int) int {
	return internalDigitalRead(pin)
}

func DigitalReadStr(pin int) string {
	if internalDigitalRead(pin) == LOW {
		return "LOW"
	}
	return "HIGH"
}

func GetMode(pin int) int {
	return internalGetMode(pin)
}

func GetModeStr(pin int) string {
	var mode = internalGetMode(pin)

	if mode > len(gpioModes) {
		return "INVALID"
	}

	return gpioModes[GetMode(pin)]
}

func Delay(ms int) {
	internalDelay(ms)
}

func DelayMicroseconds(microSec int) {
	internalDelayMicroseconds(microSec)
}

func WiringISR(pin int, mode int) chan int {
	return internalWiringISR(pin, mode)
}

func IsRaspberryPi() bool{
	if _, err := os.Stat("/opt/vc/include/bcm_host.h"); os.IsNotExist(err) {
		return false
	}

	return true
}