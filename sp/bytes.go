package sp

func Bytes(p []byte) (result []byte, rest []byte, ok bool) {
	if len(p) <= 0 {
		return nil, nil, false
	}

	{
		p0 := p[0]

		if !ByteIs(p0) {
			return nil, p, false
		}
	}

	for i,b := range p {
		if !ByteIs(b) {
			return p[:i], p[i:], true
		}
	}

	return p, nil, true
}
