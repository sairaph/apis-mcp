//go:build windows

package cache

import (
	"syscall"
	"unsafe"
)

var getDiskFreeSpaceEx = syscall.NewLazyDLL("kernel32.dll").NewProc("GetDiskFreeSpaceExW")

func freeBytes(path string) (uint64, error) {
	wide, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return 0, err
	}
	var available uint64
	result, _, callErr := getDiskFreeSpaceEx.Call(uintptr(unsafe.Pointer(wide)), uintptr(unsafe.Pointer(&available)), 0, 0)
	if result == 0 {
		return 0, callErr
	}
	return available, nil
}
