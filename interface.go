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
#include <time.h>

unsigned long long as_nanoseconds(struct timespec* ts) {
    return ts->tv_sec * (unsigned long long)1000000000L + ts->tv_nsec;
}

unsigned long long monotonic_time() {
    struct timespec last_t;
    clock_gettime(CLOCK_MONOTONIC, &last_t);
    return as_nanoseconds(&last_t);
}
*/
import "C"
import (
	"fmt"
	"os"
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

	gpioModes = []string{"IN", "OUT", "ALT5", "ALT4", "ALT0", "ALT1", "ALT2", "ALT3"}
)

//noinspection GoUnusedConst
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

	MODE_IN   = 0
	MODE_OUT  = 1
	MODE_ALT5 = 2
	MODE_ALT4 = 3
	MODE_ALT0 = 4
	MODE_ALT1 = 5
	MODE_ALT2 = 6
	MODE_ALT3 = 7
)

//use RPi.GPIO's BOARD numbering
//noinspection GoUnusedExportedFunction
func BoardToPin(pin int) int {
	if pin < 1 || pin >= len(board2pin) {
		panic(fmt.Sprintf("Invalid board pin number: %d", pin))
	}
	return board2pin[pin]
}

//noinspection GoUnusedExportedFunction
func GpioToPin(pin int) int {
	if pin < 0 || pin >= len(gpio2pin) {
		panic(fmt.Sprintf("Invalid bcm gpio number: %d", pin))
	}
	return gpio2pin[pin]
}

//noinspection GoUnusedExportedFunction
func PinToGpio(pin int) int {
	return internalPinToGpio(pin)
}

// This initialises wiringPi and assumes that the calling program is going to be using the wiringPi pin numbering scheme.
// This is a simplified numbering scheme which provides a mapping from virtual pin numbers 0 through 16 to the real
// underlying Broadcom GPIO pin numbers. See the pins page for a table which maps the wiringPi pin number to the
// Broadcom GPIO pin number to the physical location on the edge connector.
//
// This function needs to be called with root privileges.
//
//noinspection GoUnusedExportedFunction
func Setup() int {
	return internalSetup()
}

//This is identical to above, however it allows the calling programs to use the Broadcom GPIO pin numbers
// directly with no re-mapping.
//
// As above, this function needs to be called with root privileges, and note that some pins are different
// from revision 1 to revision 2 boards.
//
//noinspection GoUnusedExportedFunction
func SetupGpio() int {
	return internalSetupGpio()
}

// Identical to above, however it allows the calling programs to use the physical pin numbers on the P1 connector only.
//
// As above, this function needs to be called with root privileges.
//
//noinspection GoUnusedExportedFunction
func SetupPhys() int {
	return internalSetupPhys()
}

// This initialises wiringPi but uses the /sys/class/gpio interface rather than accessing the hardware directly.
// This can be called as a non-root user provided the GPIO pins have been exported before-hand using the gpio program.
// Pin numbering in this mode is the native Broadcom GPIO numbers – the same as wiringPiSetupGpio() above, so be
// aware of the differences between Rev 1 and Rev 2 boards.
//
// Note: In this mode you can only use the pins which have been exported via the /sys/class/gpio interface
// before you run your program. You can do this in a separate shell-script, or by using the system() function
// from inside your program to call the gpio program.
//
//Also note that some functions have no effect when using this mode as they’re not currently possible to action unless called with root privileges. (although you can use system() to call gpio to set/change modes if needed)
//
//noinspection GoUnusedExportedFunction
func SetupSys() int {
	return internalSetupSys()
}

// This sets the mode of a pin to either INPUT, OUTPUT, PWM_OUTPUT or GPIO_CLOCK. Note that only wiringPi pin 1
// (BCM_GPIO 18) supports PWM output and only wiringPi pin 7 (BCM_GPIO 4) supports CLOCK output modes.
//
// This function has no effect when in Sys mode. If you need to change the pin mode, then you can do it with the
// gpio program in a script before you start your program.
//
//noinspection GoUnusedExportedFunction
func PinMode(pin int, mode int) {
	internalPinMode(pin, mode)
}

// This sets the pull-up or pull-down resistor mode on the given pin, which should be set as an input. Unlike the
// Arduino, the BCM2835 has both pull-up an down internal resistors. The parameter pud should be; PUD_OFF, (no pull up/down), PUD_DOWN (pull to ground) or PUD_UP (pull to 3.3v) The internal pull up/down resistors have a value of approximately 50KΩ on the Raspberry Pi.
//
// This function has no effect on the Raspberry Pi’s GPIO pins when in Sys mode. If you need to activate a
// pull-up/pull-down, then you can do it with the gpio program in a script before you start your program.
//
//noinspection GoUnusedExportedFunction
func PullUpDnControl(pin int, pud int) {
	internalPullUpDnControl(pin, pud)
}

//Writes the value HIGH or LOW (1 or 0) to the given pin which must have been previously set as an output.
//
//WiringPi treats any non-zero number as HIGH, however 0 is the only representation of LOW.
//
//noinspection GoUnusedExportedFunction
func DigitalWrite(pin int, mode int) {
	internalDigitalWrite(pin, mode)
}

// Writes the value to the PWM register for the given pin. The Raspberry Pi has one on-board PWM pin, pin 1
// (BMC_GPIO 18, Phys 12) and the range is 0-1024. Other PWM devices may have other PWM ranges.
//
// This function is not able to control the Pi’s on-board PWM when in Sys mode.
//
//noinspection GoUnusedExportedFunction
func PwmWrite(pin int, value int) {
	internalPwmWrite(pin, value)
}

//noinspection GoUnusedExportedFunction
func DigitalRead(pin int) int {
	return internalDigitalRead(pin)
}

//noinspection GoUnusedExportedFunction
func DigitalReadStr(pin int) string {
	if internalDigitalRead(pin) == LOW {
		return "LOW"
	}
	return "HIGH"
}

func GetMode(pin int) int {
	return internalGetMode(pin)
}

//noinspection GoUnusedExportedFunction
func GetModeStr(pin int) string {
	var mode = internalGetMode(pin)

	if mode > len(gpioModes) {
		return "INVALID"
	}

	return gpioModes[GetMode(pin)]
}

//noinspection GoUnusedExportedFunction
func Delay(ms int) {
	internalDelay(ms)
}

//noinspection GoUnusedExportedFunction
func DelayMicroseconds(microSec int) {
	internalDelayMicroseconds(microSec)
}

//noinspection GoUnusedExportedFunction
func WiringISR(pin int, mode int) chan int {
	return internalWiringISR(pin, mode)
}

//noinspection GoUnusedExportedFunction
func IsRaspberryPi() bool {
	_, err := os.Stat("/opt/vc/include/bcm_host.h")
	return !os.IsNotExist(err)
}

//noinspection GoUnusedExportedFunction
func SetupI2C(devId int) int {
	return internalSetupI2C(devId)
}

//noinspection GoUnusedExportedFunction
func I2cRead(fd int) int {
	return internalI2CRead(fd)
}

//noinspection GoUnusedExportedFunction
func MonotonicTime() C.ulonglong {
	var nanoseconds C.ulonglong
	C.monotonic_time()
	return nanoseconds
}

//noinspection GoUnusedExportedFunction
func ConvertMonotonicTimeToUSec(time C.ulonglong) C.ulonglong {
	return time / 1000
}
