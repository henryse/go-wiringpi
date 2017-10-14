// +build linux, arm

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

/*
#cgo LDFLAGS: -lwiringPi

#include <wiringPi.h>
#include <wiringPiI2C.h>
#include <stdio.h>
#include <stdlib.h>
#define nil ((void*)0)

#define GEN_INTERRUPTER(PIN) static void interrupt_handler_##PIN() { \
	context ctxt;   \
	ctxt.pin = PIN;  \
	ctxt.ret = PIN;  \
	callback_func(goCallback, &ctxt); \
}

typedef struct context context;
struct context {
	int pin;
	int ret;
};

static void(*callback_func)(void (*f)(void*), void*);

extern void goCallback(void *);

GEN_INTERRUPTER(0)
GEN_INTERRUPTER(1)
GEN_INTERRUPTER(2)
GEN_INTERRUPTER(3)
GEN_INTERRUPTER(4)
GEN_INTERRUPTER(5)
GEN_INTERRUPTER(6)
GEN_INTERRUPTER(7)
GEN_INTERRUPTER(8)
GEN_INTERRUPTER(9)
GEN_INTERRUPTER(10)
GEN_INTERRUPTER(11)
GEN_INTERRUPTER(12)
GEN_INTERRUPTER(13)
GEN_INTERRUPTER(14)
GEN_INTERRUPTER(15)
GEN_INTERRUPTER(16)
GEN_INTERRUPTER(17)
GEN_INTERRUPTER(18)
GEN_INTERRUPTER(19)
GEN_INTERRUPTER(20)

static int native_wiring_isr(int pin, int mode) {
	switch(pin) {
		case 0: return wiringPiISR(pin, mode, &interrupt_handler_0);
		case 1: return wiringPiISR(pin, mode, &interrupt_handler_1);
		case 2: return wiringPiISR(pin, mode, &interrupt_handler_2);
		case 3: return wiringPiISR(pin, mode, &interrupt_handler_3);
		case 4: return wiringPiISR(pin, mode, &interrupt_handler_4);
		case 5: return wiringPiISR(pin, mode, &interrupt_handler_5);
		case 6: return wiringPiISR(pin, mode, &interrupt_handler_6);
		case 7: return wiringPiISR(pin, mode, &interrupt_handler_7);
		case 8: return wiringPiISR(pin, mode, &interrupt_handler_8);
		case 9: return wiringPiISR(pin, mode, &interrupt_handler_9);
		case 10: return wiringPiISR(pin, mode, &interrupt_handler_10);
		case 11: return wiringPiISR(pin, mode, &interrupt_handler_11);
		case 12: return wiringPiISR(pin, mode, &interrupt_handler_12);
		case 13: return wiringPiISR(pin, mode, &interrupt_handler_13);
		case 14: return wiringPiISR(pin, mode, &interrupt_handler_14);
		case 15: return wiringPiISR(pin, mode, &interrupt_handler_15);
		case 16: return wiringPiISR(pin, mode, &interrupt_handler_16);
		case 17: return wiringPiISR(pin, mode, &interrupt_handler_17);
		case 18: return wiringPiISR(pin, mode, &interrupt_handler_18);
		case 19: return wiringPiISR(pin, mode, &interrupt_handler_19);
		case 20: return wiringPiISR(pin, mode, &interrupt_handler_20);
	}
	return -1;
}

static void init(void *p) {
	callback_func = p;
}
*/
import "C"
import "unsafe"

import (
	"github.com/henryse/go-callback"
	"sync"
	"os"
)

const (
	WPI_MODE_PINS          = C.WPI_MODE_PINS
	WPI_MODE_GPIO          = C.WPI_MODE_GPIO
	WPI_MODE_GPIO_SYS      = C.WPI_MODE_GPIO_SYS
	WPI_MODE_PIFACE        = C.WPI_MODE_PIFACE
	WPI_MODE_UNINITIALISED = C.WPI_MODE_UNINITIALISED

	INPUT      = C.INPUT
	OUTPUT     = C.OUTPUT
	PWM_OUTPUT = C.PWM_OUTPUT
	GPIO_CLOCK = C.GPIO_CLOCK

	LOW  = C.LOW
	HIGH = C.HIGH

	PUD_OFF  = C.PUD_OFF
	PUD_DOWN = C.PUD_DOWN
	PUD_UP   = C.PUD_UP

	PWM_MODE_MS  = C.PWM_MODE_MS
	PWM_MODE_BAL = C.PWM_MODE_BAL

	INT_EDGE_SETUP   = C.INT_EDGE_SETUP
	INT_EDGE_FALLING = C.INT_EDGE_FALLING
	INT_EDGE_RISING  = C.INT_EDGE_RISING
	INT_EDGE_BOTH    = C.INT_EDGE_BOTH
)

var mutex = &sync.Mutex{}

func internalPinToGpio(pin int) int {
	return int(C.wpiPinToGpio(C.int(pin)))
}

func internalSetup() error {
	return int(C.wiringPiSetup())
}

func internalSetupGpio() int {
	return int(C.wiringPiSetupGpio())
}

func internalSetupPhys() int {
	return int(C.wiringPiSetupPhys())
}

func internalSetupSys() int {
	return int(C.wiringPiSetupSys())
}

func internalPinMode(pin int, mode int) {
	C.pinMode(C.int(pin), C.int(mode))
}

func internalPullUpDnControl(pin int, pud int){
	C.pullUpDnControl(C.int(pin), C.int(pud))
}

func internalPwmWrite (pin int, value int) {
	C.pwmWrite(C.int(pin), C.int(value))
}

func internalDigitalWrite(pin int, mode int) {
	C.digitalWrite(C.int(pin), C.int(mode))
}

func internalDigitalRead(pin int) int {
	return int(C.digitalRead(C.int(pin)))
}

func internalGetMode(pin int) int {
	return int(C.getAlt(C.int(pin)))
}

func internalDelay(ms int) {
	C.delay(C.uint(ms))
}

func internalDelayMicroseconds(microSec int) {
	C.delayMicroseconds(C.uint(microSec))
}

func internalWiringISR(pin int, mode int) chan int {
	mutex.Lock()
	defer mutex.Unlock()
	if interrupt_channels[pin] == nil {
		interrupt_channels[pin] = make(chan int)
	}
	C.native_wiring_isr(C.int(pin), C.int(mode))
	return interrupt_channels[pin]
}

func init() {
	C.init(callback.Func)
}

var interrupt_channels = [64]chan int{}

//export goCallback
func goCallback(arg unsafe.Pointer) {
	ctxt := (*C.context)(arg)
	interrupt_channels[int(ctxt.pin)] <- int(ctxt.ret)
}

func internalGetPiRevision() int {
	inFile, err := os.Open("/proc/cpuinfo")
	if err != nil {
		fmt.Println(err.Error() + `: ` + path)
		return
	} else {
		defer inFile.Close()
	}

	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // the line
	}

}

// This initialises the I2C system with your given device identifier.
// The ID is the I2C number of the device and you can use the i2cdetect
// program to find this out. wiringPiI2CSetup() will work out which
// revision Raspberry Pi you have and open the appropriate device in /dev.
//
// The return value is the standard Linux filehandle, or -1 if any
// error – in which case, you can consult errno as usual.
//
func internalSetupI2C(devId int) int {
	return int(C.wiringPiI2CSetup(C.int(defId)))
}

// Simple device read. Some devices present data when you read them
// without having to do any register transactions.
//
func internalI2CRead(fd int) int {
	return int(C.wiringPiI2CRead(C.int(fd)))
}


