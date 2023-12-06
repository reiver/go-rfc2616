package char

// ByteIsChar returns whether the byte is a 'CHAR', as defined by IETF RFC-2616.
//
//	CHAR = <any US-ASCII character (octets 0 - 127)>
func ByteIsChar(value byte) bool {
	return value <= 127
}
