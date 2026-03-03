//go:build windows
// +build windows

package util

import "golang.org/x/sys/windows"

// DiskUsage returns total and free bytes of the filesystem that holds path.
func DiskUsage(path string) (total, free uint64, err error) {
	p, err := windows.UTF16PtrFromString(path)
	if err != nil {
		return 0, 0, err
	}
	var avail, tot, freeBytes uint64
	if err = windows.GetDiskFreeSpaceEx(p, &avail, &tot, &freeBytes); err != nil {
		return 0, 0, err
	}
	return tot, freeBytes, nil
}
