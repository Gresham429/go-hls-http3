package utils

func ReadUint32(b [3]byte) uint32 {
	return uint32(b[0])<<16 | uint32(b[1])<<8 | uint32(b[2])
}

// BigEndian is the big-endian implementation of the ByteOrder interface.
func ConvertToThreeBytes(n uint32) [3]byte {
	n = n & 0x00ffffff

	b := [3]byte{
		byte(n >> 16),
		byte(n >> 8),
		byte(n),
	}

	return b
}
