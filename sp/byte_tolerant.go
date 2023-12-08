package sp

// ByteIsTolerant is a more tolerant version of ByteIs.
// Where ByteIs only returns whether the byte is a 'SP' (spacing) character, as defined by IETF RFC-2616:
//
//      SP = <US-ASCII SP, space (32)>
//
// ByteIsTolerant also allows:
//
//	HT = <US-ASCII HT, horizontal-tab (9)>
func ByteIsTolerant(value byte) bool {
	return ' ' == value || '\t' == value
}
