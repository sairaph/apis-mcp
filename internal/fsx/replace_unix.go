//go:build !windows

package fsx

import "os"

// Replace atomically publishes source at destination, replacing an existing file.
func Replace(source, destination string) error {
	return os.Rename(source, destination)
}
