//go:build linux
package ina260

import "github.com/d2r2/go-i2c"

type S struct {
	i2c *i2c.I2C
}

func New(addr uint8, bus int) (sensor S, err error) {
	d, err := i2c.NewI2C(addr, bus)
	if err != nil {
		return
	}

	sensor.i2c = d
	return
}

func (s S) ReadData(readVoltage, readCurrent, readPower bool) (voltage, current, power float64, err error) {
	var v uint16

	if readVoltage {
		v, err = s.i2c.ReadRegU16BE(REGBusVoltage)
		if err != nil {
			return
		}

		voltage = 0.00125 * float64(v) // 1.25mv/bit
	}

	if readCurrent {
		v, err = s.i2c.ReadRegU16BE(REGCurrent)
		if err != nil {
			return
		}

		//  Fix 2's complement
		if v&(1<<15) > 0 {
			v -= 65535
		}

		current = 0.00125 * float64(v) // 1.25mA/bit
	}

	if readPower {
		v, err = s.i2c.ReadRegU16BE(REGPower)
		if err != nil {
			return
		}

		power = 0.01 * float64(v) // 10mW/bit
	}

	return
}

func (s S) ManufacturerID() (uint16, error) {
	return s.i2c.ReadRegU16BE(REGManufacturerID)
}

func (s S) DieID() (id uint16, revision uint8, err error) {
	v, err := s.i2c.ReadRegU16BE(REGDieID)
	if err != nil {
		return
	}

	id = v >> 4
	revision = uint8(v & 0x000F)
	return
}

func (s S) Close() error {
	return s.i2c.Close()
}
