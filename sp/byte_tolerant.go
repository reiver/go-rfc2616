package sp

// ByteIsSpacingTolerant is a more tolerant version of ByteIsSpacing.
// Where ByteIsSpacing only returns whether the byte is a 'SP' (spacing) character, as defined by IETF RFC-2616:
//
//      SP = <US-ASCII SP, space (32)>
//
// ByteIsSpacingTolerant also allows:
//
//	HT = <US-ASCII HT, horizontal-tab (9)>
func ByteIsSpacingTolerant(value byte) bool {
	return ' ' == value || '\t' == value
}
