//go:build !windows

package keyboard

import (
	"os"
	"syscall"
	"unsafe"
)

func SetRawMode() {
	var termios syscall.Termios
	syscall.Syscall(syscall.SYS_IOCTL, os.Stdin.Fd(), uintptr(syscall.TCGETS), uintptr(unsafe.Pointer(&termios)))
	termios.Lflag &^= syscall.ECHO | syscall.ICANON
	syscall.Syscall(syscall.SYS_IOCTL, os.Stdin.Fd(), uintptr(syscall.TCSETS), uintptr(unsafe.Pointer(&termios)))
}

func ResetRawMode() {
	var termios syscall.Termios
	syscall.Syscall(syscall.SYS_IOCTL, os.Stdin.Fd(), uintptr(syscall.TCGETS), uintptr(unsafe.Pointer(&termios)))
	termios.Lflag |= syscall.ECHO | syscall.ICANON
	syscall.Syscall(syscall.SYS_IOCTL, os.Stdin.Fd(), uintptr(syscall.TCSETS), uintptr(unsafe.Pointer(&termios)))
}
