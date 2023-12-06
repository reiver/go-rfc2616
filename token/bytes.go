package token

import (
	"sourcecode.social/reiver/go-rfc2616/char"
	"sourcecode.social/reiver/go-rfc2616/ctl"
	"sourcecode.social/reiver/go-rfc2616/separator"
)

// Bytes returns the 'token', as defiend by IETF RFC-2616.
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
func Bytes(p []byte) (method []byte, rest []byte, ok bool) {
	if 0 == len(p) {
		return nil, nil, false
	}

	{
		p0 := p[0]

		if !char.ByteIsChar(p0) {
			return nil, p, false
		}

		if separator.ByteIsSeparator(p0) {
			return nil, p, false
		}

		if ctl.ByteIsControlCharacter(p0) {
			return nil, p, false
		}
	}

	{
		var i int
		var b byte

		for  i,b = range p {
			if !char.ByteIsChar(b) || separator.ByteIsSeparator(b) || ctl.ByteIsControlCharacter(b) {
				return p[:i], p[i:], true
			}
		}

		return p, nil, true
	}
}
