package requestline

import (
	"sourcecode.social/reiver/go-rfc2616/eol"
	rfc2616method     "sourcecode.social/reiver/go-rfc2616/method"
	rfc2616requesturi "sourcecode.social/reiver/go-rfc2616/requesturi"
	rfc2616version    "sourcecode.social/reiver/go-rfc2616/version"
	"sourcecode.social/reiver/go-rfc2616/sp"
)

func BytesTolerant(p []byte) (method []byte, requesturi []byte, version []byte, rest []byte, ok bool) {

	if len(p) < 3 {
		return nil, nil, nil, p, false
	}

	var pp []byte = p

	{
		result, rest, ok := rfc2616method.BytesTolerant(pp)
		if !ok {
			return nil, nil, nil, p, false
		}

		method = result
		pp = rest

		if len(pp) <= 0 {
			return nil, nil, nil, p, false
		}
	}

	{
		pp0 := pp[0]
		if !sp.ByteIs(pp0) {
			return nil, nil, nil, p, false
		}

		pp = pp[1:]
	}

	{
		pp = sp.SkipTolerant(pp)

		if len(pp) <= 0 {
			return nil, nil, nil, p, false
		}
	}

	{
		result, rest, ok := rfc2616requesturi.BytesTolerant(pp)
		if !ok {
			return nil, nil, nil, p, false
		}

		requesturi = result
		pp = rest

		if len(pp) <= 0 {
			return method, requesturi, nil, pp, true
		}
	}

	{
		pp0 := pp[0]

		switch {
		case sp.ByteIs(pp0):
			pp = pp[1:]

			// Nothing (else) here.
		default:
			_, rest, ok := eol.BytesTolerant(pp)
			if !ok {
				return nil, nil, nil, p, false
			}

			return method, requesturi, nil, rest, true
		}
	}

	{
		pp = sp.SkipTolerant(pp)

		if len(pp) <= 0 {
			return method, requesturi, nil, pp, true
		}
	}

	{
		result, rest, ok := rfc2616version.BytesTolerant(pp)
		if !ok {
			return method, requesturi, nil, pp, true
		}

		version = result
		pp = rest

		if len(pp) <= 0 {
			return method, requesturi, version, pp, true
		}
	}

	{
		pp = sp.SkipTolerant(pp)

		if len(pp) <= 0 {
			return method, requesturi, version, pp, true
		}
	}

	{
		_, rest, ok := eol.BytesTolerant(pp)
		if !ok {
			return nil, nil, nil, p, false
		}

		pp = rest
	}

	return method, requesturi, version, pp, true
}
