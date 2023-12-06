package separator

// ByteIsSeparator returns whether the byte is a 'separator' according to IETF RFC-2616.
//
//	separators  = "(" | ")" | "<" | ">" | "@"
//	            | "," | ";" | ":" | "\" | <">
//	            | "/" | "[" | "]" | "?" | "="
//	            | "{" | "}" | SP | HT
//
//	SP          = <US-ASCII SP, space (32)>
//	HT          = <US-ASCII HT, horizontal-tab (9)>
func ByteIsSeparator(b byte) bool {
	switch b {
	case
		'(',  ')',  '<', '>',  '@',
		',',  ';',  ':', '\\', '"',
		'/',  '[',  ']', '?',  '=',
		'{',  '}',  ' ', '\t':
		return true
	default:
		return false
	}
}
