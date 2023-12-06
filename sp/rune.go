package sp

// RuneIsSpacing return whether the rune is a 'SP' (spacing) character, as defined by IETF RFC-2616.
//
//	SP = <US-ASCII SP, space (32)>
func RuneIsSpacing(value rune) bool {
	return ' ' == value
}
