//go:build !windows
// +build !windows

package render

import (
	"os"
	"syscall"
	"unsafe"
)

func getTerminalSizeUnix() (int, int, error) {
	type unixWinsize struct {
		Row    uint16
		Col    uint16
		Xpixel uint16
		Ypixel uint16
	}

	ws := &unixWinsize{}
	retCode, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		os.Stdout.Fd(),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)),
	)

	if int(retCode) == -1 {
		return 0, 0, errno
	}
	return int(ws.Row), int(ws.Col), nil
}
