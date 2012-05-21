package serial

// #include <termios.h>
// #include "serial.h"
import "C"

import (
	"os"
)

const (
	BAUD_115200 = C.B115200
	BAUD_9600   = C.B9600
)

func New(port string, baud int) (*os.File, error) {
	fd := int(C.serial_open(C.CString(port), C.int(baud)))
	if fd == -1 {
		return nil, os.ErrPermission
	}
	return os.NewFile(uintptr(fd), port), nil
}
