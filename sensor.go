//go:build !linux

package ina260

import "errors"

type S struct{}

func New(addr uint8, bus int) (sensor S, err error) {
	err = errors.New("ina260: Only supported on linux")
	return
}
