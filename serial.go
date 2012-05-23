package serial

// #include <termios.h>
// #include "serial.h"
import "C"

import (
	"os"
)

const (
	BAUD_50     = C.B50
	BAUD_75     = C.B75
	BAUD_110    = C.B110
	BAUD_134    = C.B134
	BAUD_150    = C.B150
	BAUD_200    = C.B200
	BAUD_300    = C.B300
	BAUD_600    = C.B600
	BAUD_1200   = C.B1200
	BAUD_1800   = C.B1800
	BAUD_2400   = C.B2400
	BAUD_4800   = C.B4800
	BAUD_9600   = C.B9600
	BAUD_19200  = C.B19200
	BAUD_38400  = C.B38400
	BAUD_7200   = C.B7200
	BAUD_14400  = C.B14400
	BAUD_28800  = C.B28800
	BAUD_57600  = C.B57600
	BAUD_76800  = C.B76800
	BAUD_115200 = C.B115200
	BAUD_230400 = C.B230400
)

func New(port string, baud int) (*os.File, error) {
	fd := int(C.serial_open(C.CString(port), C.int(baud)))
	if fd == -1 {
		return nil, os.ErrPermission
	}
	return os.NewFile(uintptr(fd), port), nil
}
