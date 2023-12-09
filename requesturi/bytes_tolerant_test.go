package requesturi_test

import (
	"testing"

	"bytes"

	"sourcecode.social/reiver/go-rfc2616/requesturi"
)

func TestBytesTolerant(t *testing.T) {

	tests := []struct{
		Value []byte
		ExpectedResult []byte
		ExpectedRest []byte
		ExpectedOK bool
	}{
		{
			Value: nil,
			ExpectedResult: nil,
			ExpectedRest: nil,
			ExpectedOK: false,
		},
		{
			Value: []byte(nil),
			ExpectedResult: nil,
			ExpectedRest: nil,
			ExpectedOK: false,
		},



		{
			Value: []byte{},
			ExpectedResult: nil,
			ExpectedRest: nil,
			ExpectedOK: false,
		},
		{
			Value: []byte(""),
			ExpectedResult: nil,
			ExpectedRest: nil,
			ExpectedOK: false,
		},



		{
			Value:        []byte(" "),
			ExpectedResult: nil,
			ExpectedRest: []byte(" "),
			ExpectedOK: false,
		},
		{
			Value:        []byte("\t"),
			ExpectedResult: nil,
			ExpectedRest: []byte("\t"),
			ExpectedOK: false,
		},
		{
			Value:        []byte("\r"),
			ExpectedResult: nil,
			ExpectedRest: []byte("\r"),
			ExpectedOK: false,
		},
		{
			Value: []byte("\n"),
			ExpectedResult: nil,
			ExpectedRest: []byte("\n"),
			ExpectedOK: false,
		},



		{
			Value:          []byte("/"),
			ExpectedResult: []byte("/"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("/one/two/three/four.php"),
			ExpectedResult: []byte("/one/two/three/four.php"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("/once/twice/thrice/fource.php?apple=1&banana=2.0&cherry=three"),
			ExpectedResult: []byte("/once/twice/thrice/fource.php?apple=1&banana=2.0&cherry=three"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("http://example.com/"),
			ExpectedResult: []byte("http://example.com/"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("https://example.com/.well-known/webfinger?resource=acct:joeblow@example.com"),
			ExpectedResult: []byte("https://example.com/.well-known/webfinger?resource=acct:joeblow@example.com"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("urn:sha1:3I42H3S6NNFQ2MSVX7XZKYAYSCX5QBYJ"),
			ExpectedResult: []byte("urn:sha1:3I42H3S6NNFQ2MSVX7XZKYAYSCX5QBYJ"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("joeblow@example.com"),
			ExpectedResult: []byte("joeblow@example.com"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("acct:joeblow@example.com"),
			ExpectedResult: []byte("acct:joeblow@example.com"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("joeblow@example.com/the/path.md"),
			ExpectedResult: []byte("joeblow@example.com/the/path.md"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("^joeblow@example.com/this/is/the/way.html"),
			ExpectedResult: []byte("^joeblow@example.com/this/is/the/way.html"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte(`()<>@,;:\"/[]?={}`),
			ExpectedResult: []byte(`()<>@,;:\"/[]?={}`),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("😈"),
			ExpectedResult: []byte(`😈`),
			ExpectedRest: nil,
			ExpectedOK: true,
		},



		{
			Value:          []byte("/ "),
			ExpectedResult: []byte("/"),
			ExpectedRest:          []byte(" "),
			ExpectedOK: true,
		},
		{
			Value:          []byte("/one/two/three/four.php "),
			ExpectedResult: []byte("/one/two/three/four.php"),
			ExpectedRest:         []byte(" "),
			ExpectedOK: true,
		},
		{
			Value:          []byte("/once/twice/thrice/fource.php?apple=1&banana=2.0&cherry=three "),
			ExpectedResult: []byte("/once/twice/thrice/fource.php?apple=1&banana=2.0&cherry=three"),
			ExpectedRest:      []byte(" "),
			ExpectedOK: true,
		},
		{
			Value:          []byte("http://example.com/ "),
			ExpectedResult: []byte("http://example.com/"),
			ExpectedRest:       []byte(" "),
			ExpectedOK: true,
		},
		{
			Value:          []byte("https://example.com/.well-known/webfinger?resource=acct:joeblow@example.com "),
			ExpectedResult: []byte("https://example.com/.well-known/webfinger?resource=acct:joeblow@example.com"),
			ExpectedRest:        []byte(" "),
			ExpectedOK: true,
		},
		{
			Value:          []byte("urn:sha1:3I42H3S6NNFQ2MSVX7XZKYAYSCX5QBYJ "),
			ExpectedResult: []byte("urn:sha1:3I42H3S6NNFQ2MSVX7XZKYAYSCX5QBYJ"),
			ExpectedRest:       []byte(" "),
			ExpectedOK: true,
		},
		{
			Value:          []byte("joeblow@example.com "),
			ExpectedResult: []byte("joeblow@example.com"),
			ExpectedRest:      []byte(" "),
			ExpectedOK: true,
		},
		{
			Value:          []byte("acct:joeblow@example.com "),
			ExpectedResult: []byte("acct:joeblow@example.com"),
			ExpectedRest:        []byte(" "),
			ExpectedOK: true,
		},
		{
			Value:          []byte("joeblow@example.com/the/path.md "),
			ExpectedResult: []byte("joeblow@example.com/the/path.md"),
			ExpectedRest:       []byte(" "),
			ExpectedOK: true,
		},
		{
			Value:          []byte("^joeblow@example.com/this/is/the/way.html "),
			ExpectedResult: []byte("^joeblow@example.com/this/is/the/way.html"),
			ExpectedRest:     []byte(" "),
			ExpectedOK: true,
		},
		{
			Value:          []byte(`()<>@,;:\"/[]?={} `),
			ExpectedResult: []byte(`()<>@,;:\"/[]?={}`),
			ExpectedRest:                    []byte(` `),
			ExpectedOK: true,
		},
		{
			Value:          []byte("😈 "),
			ExpectedResult: []byte(`😈`),
			ExpectedRest:     []byte(" "),
			ExpectedOK: true,
		},



		{
			Value:          []byte("/\t"),
			ExpectedResult: []byte("/"),
			ExpectedRest:          []byte("\t"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("/one/two/three/four.php\t"),
			ExpectedResult: []byte("/one/two/three/four.php"),
			ExpectedRest:         []byte("\t"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("/once/twice/thrice/fource.php?apple=1&banana=2.0&cherry=three\t"),
			ExpectedResult: []byte("/once/twice/thrice/fource.php?apple=1&banana=2.0&cherry=three"),
			ExpectedRest:      []byte("\t"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("http://example.com/\t"),
			ExpectedResult: []byte("http://example.com/"),
			ExpectedRest:       []byte("\t"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("https://example.com/.well-known/webfinger?resource=acct:joeblow@example.com\t"),
			ExpectedResult: []byte("https://example.com/.well-known/webfinger?resource=acct:joeblow@example.com"),
			ExpectedRest:        []byte("\t"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("urn:sha1:3I42H3S6NNFQ2MSVX7XZKYAYSCX5QBYJ\t"),
			ExpectedResult: []byte("urn:sha1:3I42H3S6NNFQ2MSVX7XZKYAYSCX5QBYJ"),
			ExpectedRest:       []byte("\t"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("joeblow@example.com\t"),
			ExpectedResult: []byte("joeblow@example.com"),
			ExpectedRest:      []byte("\t"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("acct:joeblow@example.com\t"),
			ExpectedResult: []byte("acct:joeblow@example.com"),
			ExpectedRest:        []byte("\t"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("joeblow@example.com/the/path.md\t"),
			ExpectedResult: []byte("joeblow@example.com/the/path.md"),
			ExpectedRest:       []byte("\t"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("^joeblow@example.com/this/is/the/way.html\t"),
			ExpectedResult: []byte("^joeblow@example.com/this/is/the/way.html"),
			ExpectedRest:     []byte("\t"),
			ExpectedOK: true,
		},
		{
			Value:          []byte(`()<>@,;:\"/[]?={} `),
			ExpectedResult: []byte(`()<>@,;:\"/[]?={}`),
			ExpectedRest:                    []byte(` `),
			ExpectedOK: true,
		},
		{
			Value:          []byte("😈\t"),
			ExpectedResult: []byte(`😈`),
			ExpectedRest:     []byte("\t"),
			ExpectedOK: true,
		},



		{
			Value:          []byte("/\r"),
			ExpectedResult: []byte("/"),
			ExpectedRest:          []byte("\r"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("/one/two/three/four.php\r"),
			ExpectedResult: []byte("/one/two/three/four.php"),
			ExpectedRest:         []byte("\r"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("/once/twice/thrice/fource.php?apple=1&banana=2.0&cherry=three\r"),
			ExpectedResult: []byte("/once/twice/thrice/fource.php?apple=1&banana=2.0&cherry=three"),
			ExpectedRest:      []byte("\r"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("http://example.com/\r"),
			ExpectedResult: []byte("http://example.com/"),
			ExpectedRest:       []byte("\r"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("https://example.com/.well-known/webfinger?resource=acct:joeblow@example.com\r"),
			ExpectedResult: []byte("https://example.com/.well-known/webfinger?resource=acct:joeblow@example.com"),
			ExpectedRest:        []byte("\r"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("urn:sha1:3I42H3S6NNFQ2MSVX7XZKYAYSCX5QBYJ\r"),
			ExpectedResult: []byte("urn:sha1:3I42H3S6NNFQ2MSVX7XZKYAYSCX5QBYJ"),
			ExpectedRest:       []byte("\r"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("joeblow@example.com\r"),
			ExpectedResult: []byte("joeblow@example.com"),
			ExpectedRest:      []byte("\r"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("acct:joeblow@example.com\r"),
			ExpectedResult: []byte("acct:joeblow@example.com"),
			ExpectedRest:        []byte("\r"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("joeblow@example.com/the/path.md\r"),
			ExpectedResult: []byte("joeblow@example.com/the/path.md"),
			ExpectedRest:       []byte("\r"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("^joeblow@example.com/this/is/the/way.html\r"),
			ExpectedResult: []byte("^joeblow@example.com/this/is/the/way.html"),
			ExpectedRest:     []byte("\r"),
			ExpectedOK: true,
		},
		{
			Value:          []byte(`()<>@,;:\"/[]?={} `),
			ExpectedResult: []byte(`()<>@,;:\"/[]?={}`),
			ExpectedRest:                    []byte(` `),
			ExpectedOK: true,
		},
		{
			Value:          []byte("😈\r"),
			ExpectedResult: []byte(`😈`),
			ExpectedRest:     []byte("\r"),
			ExpectedOK: true,
		},



		{
			Value:          []byte("/\n"),
			ExpectedResult: []byte("/"),
			ExpectedRest:          []byte("\n"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("/one/two/three/four.php\n"),
			ExpectedResult: []byte("/one/two/three/four.php"),
			ExpectedRest:         []byte("\n"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("/once/twice/thrice/fource.php?apple=1&banana=2.0&cherry=three\n"),
			ExpectedResult: []byte("/once/twice/thrice/fource.php?apple=1&banana=2.0&cherry=three"),
			ExpectedRest:      []byte("\n"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("http://example.com/\n"),
			ExpectedResult: []byte("http://example.com/"),
			ExpectedRest:       []byte("\n"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("https://example.com/.well-known/webfinger?resource=acct:joeblow@example.com\n"),
			ExpectedResult: []byte("https://example.com/.well-known/webfinger?resource=acct:joeblow@example.com"),
			ExpectedRest:        []byte("\n"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("urn:sha1:3I42H3S6NNFQ2MSVX7XZKYAYSCX5QBYJ\n"),
			ExpectedResult: []byte("urn:sha1:3I42H3S6NNFQ2MSVX7XZKYAYSCX5QBYJ"),
			ExpectedRest:       []byte("\n"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("joeblow@example.com\n"),
			ExpectedResult: []byte("joeblow@example.com"),
			ExpectedRest:      []byte("\n"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("acct:joeblow@example.com\n"),
			ExpectedResult: []byte("acct:joeblow@example.com"),
			ExpectedRest:        []byte("\n"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("joeblow@example.com/the/path.md\n"),
			ExpectedResult: []byte("joeblow@example.com/the/path.md"),
			ExpectedRest:       []byte("\n"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("^joeblow@example.com/this/is/the/way.html\n"),
			ExpectedResult: []byte("^joeblow@example.com/this/is/the/way.html"),
			ExpectedRest:     []byte("\n"),
			ExpectedOK: true,
		},
		{
			Value:          []byte(`()<>@,;:\"/[]?={} `),
			ExpectedResult: []byte(`()<>@,;:\"/[]?={}`),
			ExpectedRest:                    []byte(` `),
			ExpectedOK: true,
		},
		{
			Value:          []byte("😈\n"),
			ExpectedResult: []byte(`😈`),
			ExpectedRest:     []byte("\n"),
			ExpectedOK: true,
		},
	}

	for testNumber, test := range tests {

		actualResult, actualRest, actualOK := requesturi.BytesTolerant(test.Value)

		{
			expected := test.ExpectedOK
			actual   := actualOK

			if expected != actual {
				t.Errorf("For test #%d, the actual ok-result is not what was expected." , testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:   %t", actual)
				t.Logf("VALUE: %q (%#v)", test.Value, test.Value)
				continue
			}
		}

		{
			expected := test.ExpectedResult
			actual   := actualResult

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual result is not what was expected." , testNumber)
				t.Logf("EXPECTED: %q (%#v)", expected, expected)
				t.Logf("ACTUAL:   %q (%#v)", actual, actual)
				t.Logf("VALUE: %q (%#v)", test.Value, test.Value)
				continue
			}
		}

		{
			expected := test.ExpectedRest
			actual   := actualRest

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual rest is not what was expected." , testNumber)
				t.Logf("EXPECTED: %q (%#v)", expected, expected)
				t.Logf("ACTUAL:   %q (%#v)", actual, actual)
				t.Logf("VALUE: %q (%#v)", test.Value, test.Value)
				continue
			}
		}
	}
}
