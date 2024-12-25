//go:build windows
// +build windows

package render

func GetTerminalSize() (int, int, error) {
	return getTerminalSizeWindows()
}
