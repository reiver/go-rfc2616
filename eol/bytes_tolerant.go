package eol

func BytesTolerant(p []byte) (result []byte, rest []byte, ok bool) {

	length := len(p)

	if length <= 0 {
		return nil, p, false
	}

	var p0 byte = p[0]

	switch p0 {
	case '\n':
		return p[:1], p[1:], true
	case '\r':
		// Nothing here.
	default:
		return nil, p, false
	}

	// If we got here then p0 == '\r'

	if length < 2 {
		return nil, p, false
	}

	var p1 byte = p[1]

	switch p1 {
	case '\n':
		return p[:2], p[2:], true
	default:
		return nil, p, false
	}
}
