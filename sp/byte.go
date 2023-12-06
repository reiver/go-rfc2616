package sp

// ByteIsSpacing return whether the byte is a 'SP' (spacing) character, as defined by IETF RFC-2616.
//
//	SP = <US-ASCII SP, space (32)>
func ByteIsSpacing(value byte) bool {
	return ' ' == value
}
