package keyboard

import (
	"os"
	"syscall"
	"unsafe"
)

const (
	KeyArrowLeft  = "\x1b[D"
	KeyArrowRight = "\x1b[C"
	KeyArrowUp    = "\x1b[A"
	KeyArrowDown  = "\x1b[B"
	KeyEsc        = "\x1b"
)

// setRawMode puts the terminal into raw mode
func SetRawMode() {
	var termios syscall.Termios

	syscall.Syscall(syscall.SYS_IOCTL, os.Stdin.Fd(), uintptr(syscall.TCGETS), uintptr(unsafe.Pointer(&termios)))

	termios.Lflag &^= syscall.ECHO | syscall.ICANON
	syscall.Syscall(syscall.SYS_IOCTL, os.Stdin.Fd(), uintptr(syscall.TCSETS), uintptr(unsafe.Pointer(&termios)))
}

// resetRawMode resets the terminal to its original mode
func ResetRawMode() {
	var termios syscall.Termios

	syscall.Syscall(syscall.SYS_IOCTL, os.Stdin.Fd(), uintptr(syscall.TCGETS), uintptr(unsafe.Pointer(&termios)))

	termios.Lflag |= syscall.ECHO | syscall.ICANON
	syscall.Syscall(syscall.SYS_IOCTL, os.Stdin.Fd(), uintptr(syscall.TCSETS), uintptr(unsafe.Pointer(&termios)))
}
