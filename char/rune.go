package char

// RuneIsChar returns whether the rune is a 'CHAR', as defined by IETF RFC-2616.
//
//	CHAR = <any US-ASCII character (octets 0 - 127)>
func RuneIsChar(value rune) bool {
	return value <= 127
}
