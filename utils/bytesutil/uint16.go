package bytesutil

// ToUInt16 converts data byte slice to uint16.
// If len(data) < 2, the most significant byte
// is assumed as zero.
// If len(data) > 2, truncates data removing
// the most significant bytes until len == 2.
func ToUInt16(data []byte) uint16 {
	var result uint16

	length := len(data)

	if length >= 2 {
		result = uint16(data[length-2])<<8 | uint16(data[length-1])
	} else if length == 1 {
		result = uint16(0)<<8 | uint16(data[0])
	} else {
		result = 0
	}

	return result
}

// AppendUInt16 receives a byte slice and append a uint16
// to it.
func AppendUInt16(buffer []byte, data uint16) []byte {
	buffer = append(buffer, byte(data>>8))
	buffer = append(buffer, byte(data&0x00ff))
	return buffer
}
