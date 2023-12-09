package version_test

import (
	"testing"

	"bytes"

	"sourcecode.social/reiver/go-rfc2616/version"
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
			ExpectedOK: true,
		},
		{
			Value: []byte(nil),
			ExpectedResult: nil,
			ExpectedRest: nil,
			ExpectedOK: true,
		},



		{
			Value: []byte{},
			ExpectedResult: nil,
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value: []byte(""),
			ExpectedResult: nil,
			ExpectedRest: nil,
			ExpectedOK: true,
		},



		{
			Value:        []byte(" "),
			ExpectedResult: nil,
			ExpectedRest: []byte(" "),
			ExpectedOK: true,
		},
		{
			Value:        []byte("\t"),
			ExpectedResult: nil,
			ExpectedRest: []byte("\t"),
			ExpectedOK: true,
		},
		{
			Value:        []byte("\r"),
			ExpectedResult: []byte(""),
			ExpectedRest: []byte("\r"),
			ExpectedOK: true,
		},
		{
			Value: []byte("\n"),
			ExpectedResult: []byte(""),
			ExpectedRest: []byte("\n"),
			ExpectedOK: true,
		},



		{
			Value:          []byte("HTTP/1.1"),
			ExpectedResult: []byte("HTTP/1.1"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("HTTP/1"),
			ExpectedResult: []byte("HTTP/1"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("finger/1"),
			ExpectedResult: []byte("finger/1"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("finger"),
			ExpectedResult: []byte("finger"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("1.22.333"),
			ExpectedResult: []byte("1.22.333"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("v12.23.34"),
			ExpectedResult: []byte("v12.23.34"),
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
			Value:          []byte("HTTP/1.1 "),
			ExpectedResult: []byte("HTTP/1.1"),
			ExpectedRest:          []byte(" "),
			ExpectedOK: true,
		},
		{
			Value:          []byte("HTTP/1 "),
			ExpectedResult: []byte("HTTP/1"),
			ExpectedRest:         []byte(" "),
			ExpectedOK: true,
		},
		{
			Value:          []byte("finger/1 "),
			ExpectedResult: []byte("finger/1"),
			ExpectedRest:      []byte(" "),
			ExpectedOK: true,
		},
		{
			Value:          []byte("finger "),
			ExpectedResult: []byte("finger"),
			ExpectedRest:       []byte(" "),
			ExpectedOK: true,
		},
		{
			Value:          []byte("1.22.333 "),
			ExpectedResult: []byte("1.22.333"),
			ExpectedRest:        []byte(" "),
			ExpectedOK: true,
		},
		{
			Value:          []byte("v12.23.34 "),
			ExpectedResult: []byte("v12.23.34"),
			ExpectedRest:       []byte(" "),
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
			Value:          []byte("HTTP/1.1\t"),
			ExpectedResult: []byte("HTTP/1.1"),
			ExpectedRest:          []byte("\t"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("HTTP/1\t"),
			ExpectedResult: []byte("HTTP/1"),
			ExpectedRest:         []byte("\t"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("finger/1\t"),
			ExpectedResult: []byte("finger/1"),
			ExpectedRest:      []byte("\t"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("finger\t"),
			ExpectedResult: []byte("finger"),
			ExpectedRest:       []byte("\t"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("1.22.333\t"),
			ExpectedResult: []byte("1.22.333"),
			ExpectedRest:        []byte("\t"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("v12.23.34\t"),
			ExpectedResult: []byte("v12.23.34"),
			ExpectedRest:       []byte("\t"),
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
			Value:          []byte("HTTP/1.1\r"),
			ExpectedResult: []byte("HTTP/1.1"),
			ExpectedRest:          []byte("\r"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("HTTP/1\r"),
			ExpectedResult: []byte("HTTP/1"),
			ExpectedRest:         []byte("\r"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("finger/1\r"),
			ExpectedResult: []byte("finger/1"),
			ExpectedRest:      []byte("\r"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("finger\r"),
			ExpectedResult: []byte("finger"),
			ExpectedRest:       []byte("\r"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("1.22.333\r"),
			ExpectedResult: []byte("1.22.333"),
			ExpectedRest:        []byte("\r"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("v12.23.34\r"),
			ExpectedResult: []byte("v12.23.34"),
			ExpectedRest:       []byte("\r"),
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
			Value:          []byte("HTTP/1.1\n"),
			ExpectedResult: []byte("HTTP/1.1"),
			ExpectedRest:          []byte("\n"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("HTTP/1\n"),
			ExpectedResult: []byte("HTTP/1"),
			ExpectedRest:         []byte("\n"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("finger/1\n"),
			ExpectedResult: []byte("finger/1"),
			ExpectedRest:      []byte("\n"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("finger\n"),
			ExpectedResult: []byte("finger"),
			ExpectedRest:       []byte("\n"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("1.22.333\n"),
			ExpectedResult: []byte("1.22.333"),
			ExpectedRest:        []byte("\n"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("v12.23.34\n"),
			ExpectedResult: []byte("v12.23.34"),
			ExpectedRest:       []byte("\n"),
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

		actualResult, actualRest, actualOK := version.BytesTolerant(test.Value)

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
