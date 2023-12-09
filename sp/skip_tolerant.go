package sp

func SkipTolerant(p []byte) []byte {
	length := len(p)

	if length <= 0 {
		return nil
	}

	for i,b := range p {
		if !ByteIsTolerant(b) {
			return p[i:]
		}
	}

	return nil
}
