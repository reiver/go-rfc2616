package method

import (
	"sourcecode.social/reiver/go-rfc2616/sp"
)

func BytesTolerant(p []byte) (result []byte, rest []byte, ok bool) {

	if len(p) <= 0 {
		return nil, nil, false
	}

	{
		p0 := p[0]

		if sp.ByteIsTolerant(p0) || '\r' == p0 || '\n' == p0 {
			return nil, p, false
		}
	}

	{
		for i,b := range p {
			if sp.ByteIsTolerant(b) || '\r' == b || '\n' == b {
				return p[:i], p[i:], true
			}
		}

		return p, nil, true
	}
}
