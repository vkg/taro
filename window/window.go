package window

import (
	"fmt"
	"syscall"
	"unsafe"
)

type winsize struct {
	row uint16
	col uint16
}

// GetSize returns (row, col)
func GetSize() (int, int, error) {
	ws := &winsize{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		return 0, 0, fmt.Errorf("Syscall failed ErrNo: %v", errno)
	}
	return int(ws.row), int(ws.col), nil
}
