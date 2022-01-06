package galileo

type tagDesc struct {
	Len  uint
	Type string
}

var tagsTable = map[byte]tagDesc{
	// iron version
	0x01: {1, "uint"},
	// firmware version
	0x02: {1, "uint"},
	// IMEI
	0x03: {15, "string"},
	// device id
	0x04: {2, "uint"},
	// archive record number
	0x10: {2, "uint"},
	// Дата и время
	0x20: {4, "time"},
	// Coordinates in degrees, number of satellites,
	// the sign of the correctness of the determination of coordinates and
	// coordinate source
	0x30: {9, "coord"},
	// Speed ​​in km / h is directed in degrees
	0x33: {4, "speed"},
	// height, m.
	0x34: {2, "int"},
	// One of the values: 1. HDOP (divided by 10) - if the source of GPS coordinates
	// module, 2 error in meters if the source of the gsm network (multiply by 10)
	0x35: {1, "uint"},
	// Device status
	0x40: {2, "bitstring"},
	// Supply voltage, mV
	0x41: {2, "uint"},
	// Battery voltage, mV
	0x42: {2, "uint"},
	// Input statuses
	0x45: {2, "bitstring"},
	// Output statuses
	0x46: {2, "bitstring"},	
	// The value at the input is 0.
	// Depending on the settings, one of the options: voltage,
	// number of pulses, frequency Hz
	0x50: {2, "uint"},
	// Value at input 1.
	// Depending on the settings, one of the options: voltage,
	// number of pulses, frequency Hz
	0x51: {2, "uint"},
	// Value at input 2.
	// Depending on the settings, one of the options: voltage,
	// number of pulses, frequency Hz
	0x52: {2, "uint"},
	// Value at input 3.
// Depending on the settings, one of the options: voltage,
// number of pulses, frequency Hz	
	0x53: {2, "uint"},

// Value at input 4.
// Depending on the settings, one of the options: voltage,
// number of pulses, frequency Hz
	0x54: {2, "uint"},
// Value at input 5.
// Depending on the settings, one of the options: voltage,
// number of pulses, frequency Hz
	0x55: {2, "uint"},
	// Value at input 6.
// Depending on the settings, one of the options: voltage,
// number of pulses, frequency Hz
	0x56: {2, "uint"},
// Value at input 7.
// Depending on the settings, one of the options: voltage,
// number of pulses, frequency Hz
	0x57: {2, "uint"},
	// RS485 [0] FLS with address 0
	0x60: {2, "uint"},
	// RS485[1] ДУТ с адресом 1
	0x61: {2, "uint"},
	// RS485[2] ДУТ с адресом 2
	0x62: {2, "uint"},
	// RS485 [1] FLS with address 1
	//0x63: {2, "uint"},
	//// RS485 [4] FLS with address 2
	//0x64: {2, "uint"},
}
