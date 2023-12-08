package method

func BytesTolerant(p []byte) (result []byte, rest []byte, ok bool) {

	if len(p) <= 0 {
		return nil, nil, false
	}

	{
		p0 := p[0]

		if ' ' == p0 || '\t' == p0 || '\r' == p0 || '\n' == p0 {
			return nil, p, false
		}
	}

	{
		for i,b := range p {
			if ' ' == b || '\t' == b || '\r' == b || '\n' == b {
				return p[:i], p[i:], true
			}
		}

		return p, nil, true
	}
}
