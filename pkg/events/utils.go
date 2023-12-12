package events

func GetEventSize(data []byte) int16 {
	return int16(uint16(data[1]) | uint16(data[0])<<8)
}

func WriteEventSize(size int16, data []byte) {
	newSize := uint16(size)
	data[0] = byte(newSize >> 8)
	data[1] = byte(newSize)
}

func GetEventType(data []byte) int16 {
	return int16(uint16(data[3]) | uint16(data[2])<<8)
}

func WriteEventType(eventType int16, data []byte) {
	newType := uint16(eventType)
	data[2] = byte(newType >> 8)
	data[3] = byte(newType)
}
