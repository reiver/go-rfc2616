package ctl

// ByteIsControl returns whether a byte is a control character, as defined by IETF RFC-2616.
//
//	CTL = <any US-ASCII control character (octets 0 - 31) and DEL (127)>
func ByteIsControlCharacter(value byte) bool {
	switch {
	case 0 <= value && value <= 31:
		return true
	case 127 == value:
		return true
	default:
		return false
	}
}
