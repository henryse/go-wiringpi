// +build !linux, !arm

package wiringpi

import (
	"time"
	"log"
)

//noinspection ALL
const (
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

	PWM_MODE_MS  = 0
	PWM_MODE_BAL = 1

	INT_EDGE_SETUP   = 0
	INT_EDGE_FALLING = 1
	INT_EDGE_RISING  = 2
	INT_EDGE_BOTH    = 3

	// Used for emulator:
	gpio_pin_count = 26
)

var(
	gpio_list [gpio_pin_count]int
	gpio_mode_list [gpio_pin_count]int
)

func PinToGpio(pin int) int {
	return 0
}

func Setup() error {
	log.Println("Warning: Running in emulation mode")

	for i:= 0; i < gpio_pin_count; i++ {
		gpio_list[i] = LOW
		gpio_mode_list[i] = MODE_IN
	}

	return nil
}

func PinMode(pin int, mode int) {
	if pin < gpio_pin_count {
		gpio_mode_list[pin] = mode
	}
}

func DigitalWrite(pin int, mode int) {
	if pin < gpio_pin_count {
		gpio_list[pin] = mode
	}
}

func DigitalRead(pin int) int {
	if pin < gpio_pin_count {
		return gpio_list[pin]
	}

	return LOW
}

func DigitalReadStr(pin int) string {
	if DigitalRead(pin) == LOW {
		return "LOW"
	}
	return "HIGH"
}

func GetMode(pin int) int {
	if pin < gpio_pin_count {
		return gpio_mode_list[pin]
	}
}

func GetModeStr(pin int) string {
	var mode = GetMode(pin)

	if mode > len(gpioModes) {
		return "INVALID"
	}

	return gpioModes[GetMode(pin)]
}

func Delay(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func DelayMicroseconds(microSec int) {
	time.Sleep(time.Duration(microSec) * time.Microsecond)
}

func WiringISR(pin int, mode int) chan int {
	if pin < gpio_pin_count {
		gpio_list[pin] = mode
	}
	return nil
}