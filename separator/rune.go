package separator

// RuneIsSeparator returns whether the rune is a 'separator' according to IETF RFC-2616.
//
//	separators  = "(" | ")" | "<" | ">" | "@"
//	            | "," | ";" | ":" | "\" | <">
//	            | "/" | "[" | "]" | "?" | "="
//	            | "{" | "}" | SP | HT
//
//	SP          = <US-ASCII SP, space (32)>
//	HT          = <US-ASCII HT, horizontal-tab (9)>
func RuneIsSeparator(b rune) bool {
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
