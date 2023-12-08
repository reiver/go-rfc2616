package sp

// ByteIs return whether the byte is a 'SP' (spacing) character, as defined by IETF RFC-2616.
//
//	SP = <US-ASCII SP, space (32)>
func ByteIs(value byte) bool {
	return ' ' == value
}
