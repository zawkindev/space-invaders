//go:build windows
// +build windows

package render

import (
	"os"
	"unsafe"

	"golang.org/x/sys/windows"
)

func getTerminalSizeWindows() (int, int, error) {
	kernel32 := windows.NewLazyDLL("kernel32.dll")
	getConsoleScreenBufferInfo := kernel32.NewProc("GetConsoleScreenBufferInfo")

	type windowsCoord struct {
		X int16
		Y int16
	}

	type windowsSmallRect struct {
		Left   int16
		Top    int16
		Right  int16
		Bottom int16
	}

	type windowsConsoleScreenBufferInfo struct {
		Size              windowsCoord
		CursorPosition    windowsCoord
		Attributes        uint16
		Window            windowsSmallRect
		MaximumWindowSize windowsCoord
	}

	var csbi windowsConsoleScreenBufferInfo
	handle := windows.Handle(os.Stdout.Fd())

	ret, _, err := getConsoleScreenBufferInfo.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(&csbi)),
	)

	if ret == 0 {
		return 0, 0, err
	}

	height := csbi.Window.Bottom - csbi.Window.Top + 1
	width := csbi.Window.Right - csbi.Window.Left + 1

	return int(height), int(width), nil
}
