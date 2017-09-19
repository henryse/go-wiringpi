// +build linux, arm

package wiringpi

/*
#cgo LDFLAGS: -lwiringPi

#include <wiringPi.h>
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

static void native_pin_mode(int p, int m) {
    pinMode(p,m);
}

static void native_digital_write(int p, int m) {
    digitalWrite(p,m);
}

static int native_digital_read(int p) {
    return digitalRead(p);
}

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
	"errors"
	"sync"
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

func PinToGpio(pin int) int {
	return int(C.wpiPinToGpio(C.int(pin)))
}

func Setup() error {
	if -1 == int(C.wiringPiSetup()) {
		return errors.New("wiringPiSetup failed to call")
	}
	return nil
}

func PinMode(pin int, mode int) {
	C.native_pin_mode(C.int(pin), C.int(mode))
}

func DigitalWrite(pin int, mode int) {
	C.native_digital_write(C.int(pin), C.int(mode))
}

func DigitalRead(pin int) int {
	return int(C.native_digital_read(C.int(pin)))
}

func DigitalReadStr(pin int) string {
	if DigitalRead(pin) == LOW {
		return "LOW"
	}
	return "HIGH"
}

func GetMode(pin int) int {
	return int(C.getAlt(C.int(pin)))
}

func GetModeStr(pin int) string {
	var mode = GetMode(pin)

	if mode > len(gpioModes) {
		return "INVALID"
	}

	return gpioModes[GetMode(pin)]
}

func Delay(ms int) {
	C.delay(C.uint(ms))
}

func DelayMicroseconds(microSec int) {
	C.delayMicroseconds(C.uint(microSec))
}

func WiringISR(pin int, mode int) chan int {
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