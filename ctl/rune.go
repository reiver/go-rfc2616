package ctl

// RuneIsControl returns whether a rune is a control character, as defined by IETF RFC-2616.
//
//	CTL = <any US-ASCII control character (octets 0 - 31) and DEL (127)>
func RuneIsControlCharacter(value rune) bool {
	switch {
	case 0 <= value && value <= 31:
		return true
	case 127 == value:
		return true
	default:
		return false
	}
}
