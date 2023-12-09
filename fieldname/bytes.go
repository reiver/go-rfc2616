package fieldname

import (
	"sourcecode.social/reiver/go-rfc2616/token"
)

func Bytes(p []byte) ([]byte, []byte, bool) {
	return token.Bytes(p)
}
