package sp

// RuneIsTolerant is a more tolerant version of RuneIs.
// Where RuneIs only returns whether the rune is a 'SP' (spacing) character, as defined by IETF RFC-2616:
//
//      SP = <US-ASCII SP, space (32)>
//
// RuneIsTolerant also allows:
//
//	HT = <US-ASCII HT, horizontal-tab (9)>
func RuneIsTolerant(value rune) bool {
	return ' ' == value || '\t' == value
}
