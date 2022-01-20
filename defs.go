package ina260

const (
	PCAAutoincrementOFF           = 0x00
	PCAAutoincrementALL           = 0x80
	PCAAutoincrementIndividual    = 0xA0
	PCAAutoincrementControl       = 0xC0
	PCAAutoincrementControlGlobal = 0xE0

	REGConfig         = 0x00
	REGCurrent        = 0x01
	REGBusVoltage     = 0x02
	REGPower          = 0x03
	REGMaskEnable     = 0x06
	REGAlert          = 0x07
	REGManufacturerID = 0xFE
	REGDieID          = 0xFF

	RST     = 15
	AVG2    = 11
	AVG1    = 10
	AVG0    = 9
	VBUSCT2 = 8
	VBUSCT1 = 7
	VBUSCT0 = 6
	ISHCT2  = 5
	ISHCT1  = 4
	ISHCT0  = 3
	MODE3   = 2
	MODE2   = 1
	MODE1   = 0

	OCL  = 15
	UCL  = 14
	BOL  = 13
	BUL  = 12
	POL  = 11
	CNVR = 10
	AFF  = 4
	CVRF = 3
	OVF  = 2
	APOL = 1
	LEN  = 0
)
