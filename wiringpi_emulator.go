// +build !linux, !arm

package wiringpi

import (
	"time"
)

//noinspection ALL
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

	WPI_MODE_PINS          = 0
	WPI_MODE_GPIO          = 1
	WPI_MODE_GPIO_SYS      = 2
	WPI_MODE_PIFACE        = 3
	WPI_MODE_UNINITIALISED = 4

	INPUT      = 0
	OUTPUT     = 1
	PWM_OUTPUT = 2
	GPIO_CLOCK = 3

	LOW  = 0
	HIGH = 1

	PUD_OFF  = 1
	PUD_DOWN = 2
	PUD_UP   = 3

	// PWM

	PWM_MODE_MS  = 0
	PWM_MODE_BAL = 1

	INT_EDGE_SETUP   = 0
	INT_EDGE_FALLING = 1
	INT_EDGE_RISING  = 2
	INT_EDGE_BOTH    = 3
)

func PinToGpio(pin int) int {
	// TODO: Need code
	return 0
}

func WiringPiSetup() error {
	// TODO: Need code
	return nil
}

func PinMode(pin int, mode int) {
	// TODO: Need code
}

func DigitalWrite(pin int, mode int) {
	// TODO: Need code
}

func DigitalRead(pin int) int {
	// TODO: Need code
	return 0
}

func Delay(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func DelayMicroseconds(microSec int) {
	time.Sleep(time.Duration(microSec) * time.Microsecond)
}

func WiringPiISR(pin int, mode int) chan int {
	// TODO: Need code
	return nil
}

func init() {
	// TODO: Need code
}

func IsRaspberryPiEmulator() bool{
	return true
}