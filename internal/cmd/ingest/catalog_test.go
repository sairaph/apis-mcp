//go:build dev

package main

func init() {
	excludeBuiltinFromJobIndex = true
}
