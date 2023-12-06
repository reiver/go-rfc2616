package sp

// RuneIsSpacingTolerant is a more tolerant version of RuneIsSpacing.
// Where RuneIsSpacing only returns whether the rune is a 'SP' (spacing) character, as defined by IETF RFC-2616:
//
//      SP = <US-ASCII SP, space (32)>
//
// RuneIsSpacingTolerant also allows:
//
//	HT = <US-ASCII HT, horizontal-tab (9)>
func RuneIsSpacingTolerant(value rune) bool {
	return ' ' == value || '\t' == value
}
