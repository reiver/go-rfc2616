package method

import (
	"sourcecode.social/reiver/go-rfc2616/token"
)

// Bytes returns the 'Method', as defiend by IETF RFC-2616.
//
//	 Method          = "OPTIONS"
//	                 | "GET"
//	                 | "HEAD"
//	                 | "POST"
//	                 | "PUT"
//	                 | "DELETE"
//	                 | "TRACE"
//	                 | "CONNECT"
//	                 | extension-method
//
//	extension-method = token
//
// Which, in practice, can be simplified to:
//
//	 Method          = token
//
// And 'token' is defined as:
//
//	token            = 1*<any CHAR except CTLs or separators>
//
//	CHAR           = <any US-ASCII character (octets 0 - 127)>
//
//	CTL            = <any US-ASCII control character (octets 0 - 31) and DEL (127)>
//
//	separators     = "(" | ")" | "<" | ">" | "@"
//	               | "," | ";" | ":" | "\" | <">
//	               | "/" | "[" | "]" | "?" | "="
//	               | "{" | "}" | SP | HT
//
//	SP             = <US-ASCII SP, space (32)>
//
//	HT	       = <US-ASCII HT, horizontal-tab (9)>
func Bytes(p []byte) ([]byte, []byte, bool) {
	return token.Bytes(p)
}
