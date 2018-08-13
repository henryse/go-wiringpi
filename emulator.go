// +build !linux, !arm

// **********************************************************************
//    Copyright (c) 2017 Henry Seurer
//
//   Permission is hereby granted, free of charge, to any person
//    obtaining a copy of this software and associated documentation
//    files (the "Software"), to deal in the Software without
//    restriction, including without limitation the rights to use,
//    copy, modify, merge, publish, distribute, sublicense, and/or sell
//    copies of the Software, and to permit persons to whom the
//    Software is furnished to do so, subject to the following
//    conditions:
//
//   The above copyright notice and this permission notice shall be
//   included in all copies or substantial portions of the Software.
//
//    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
//    EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
//    OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
//    NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
//    HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
//    WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
//    FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
//    OTHER DEALINGS IN THE SOFTWARE.
//
// **********************************************************************

package wiringpi

import (
	"fmt"
	"log"
	"time"
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

var (
	gpio_list      [gpio_pin_count]int
	gpio_mode_list [gpio_pin_count]int
)

func internalPinToGpio(_ int) int {
	return 0
}

func internalSetup() int {
	log.Println("Warning: Running in emulation mode")

	for i := 0; i < gpio_pin_count; i++ {
		gpio_list[i] = LOW
		gpio_mode_list[i] = MODE_IN
	}

	return 0
}

func internalSetupGpio() int {
	return 0
}

func internalSetupPhys() int {
	return 0
}

func internalSetupSys() int {
	return 0
}

func internalPinMode(pin int, mode int) {
	if pin < gpio_pin_count {
		gpio_mode_list[pin] = mode
	}
}

func internalPullUpDnControl(pin int, pud int) {
	switch pud {
	case PUD_OFF:
		fmt.Println("PUD_OFF")
	case PUD_UP:
		fmt.Println("PUD_UP")
	case PUD_DOWN:
		fmt.Println("PUD_DOWN")
	default:
		fmt.Printf("Error invliad pud: %d\n", pud)
	}

	if pin < gpio_pin_count {
		gpio_mode_list[pin] = pud
	}
}

func internalPwmWrite(pin int, value int) {
	if pin < gpio_pin_count {
		gpio_mode_list[pin] = value
	}
}

func internalDigitalWrite(pin int, mode int) {
	if pin < gpio_pin_count {
		gpio_list[pin] = mode
	}
}

func internalDigitalRead(pin int) int {
	if pin < gpio_pin_count {
		return gpio_list[pin]
	}

	return LOW
}

func internalGetMode(pin int) int {
	if pin < gpio_pin_count {
		return gpio_mode_list[pin]
	}

	return MODE_IN
}

func internalDelay(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func internalDelayMicroseconds(microSec int) {
	time.Sleep(time.Duration(microSec) * time.Microsecond)
}

func internalWiringISR(pin int, mode int) chan int {
	if pin < gpio_pin_count {
		gpio_list[pin] = mode
	}
	return nil
}

func internalSetupI2C(_ int) int {
	return -1
}

// Simple device read. Some devices present data when you read them
// without having to do any register transactions.
//
func internalI2CRead(_ int) int {
	return 0
}
