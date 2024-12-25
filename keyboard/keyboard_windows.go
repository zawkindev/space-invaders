//go:build windows

package keyboard

import (
	"os"

	"golang.org/x/sys/windows"
)

func SetRawMode() {
	// Windows implementation using windows.SetConsoleMode
	handle := windows.Handle(os.Stdin.Fd())
	var mode uint32
	windows.GetConsoleMode(handle, &mode)
	windows.SetConsoleMode(handle, mode & ^uint32(windows.ENABLE_ECHO_INPUT) & ^uint32(windows.ENABLE_LINE_INPUT))
}

func ResetRawMode() {
	handle := windows.Handle(os.Stdin.Fd())
	var mode uint32
	windows.GetConsoleMode(handle, &mode)
	windows.SetConsoleMode(handle, mode|uint32(windows.ENABLE_ECHO_INPUT)|uint32(windows.ENABLE_LINE_INPUT))
}
