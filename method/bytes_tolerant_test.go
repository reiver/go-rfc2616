package method_test

import (
	"testing"

	"bytes"

	"sourcecode.social/reiver/go-rfc2616/method"
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
			Value:          []byte("CONNECT"),
			ExpectedResult: []byte("CONNECT"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("DELETE"),
			ExpectedResult: []byte("DELETE"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("GET"),
			ExpectedResult: []byte("GET"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("HEAD"),
			ExpectedResult: []byte("HEAD"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("PATCH"),
			ExpectedResult: []byte("PATCH"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("POST"),
			ExpectedResult: []byte("POST"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("PUT"),
			ExpectedResult: []byte("PUT"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("TRACE"),
			ExpectedResult: []byte("TRACE"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("/GET"),
			ExpectedResult: []byte("/GET"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("/W"),
			ExpectedResult: []byte("/W"),
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
			Value:          []byte("CONNECT "),
			ExpectedResult: []byte("CONNECT"),
			ExpectedRest:          []byte(" "),
			ExpectedOK: true,
		},
		{
			Value:          []byte("DELETE "),
			ExpectedResult: []byte("DELETE"),
			ExpectedRest:         []byte(" "),
			ExpectedOK: true,
		},
		{
			Value:          []byte("GET "),
			ExpectedResult: []byte("GET"),
			ExpectedRest:      []byte(" "),
			ExpectedOK: true,
		},
		{
			Value:          []byte("HEAD "),
			ExpectedResult: []byte("HEAD"),
			ExpectedRest:       []byte(" "),
			ExpectedOK: true,
		},
		{
			Value:          []byte("PATCH "),
			ExpectedResult: []byte("PATCH"),
			ExpectedRest:        []byte(" "),
			ExpectedOK: true,
		},
		{
			Value:          []byte("POST "),
			ExpectedResult: []byte("POST"),
			ExpectedRest:       []byte(" "),
			ExpectedOK: true,
		},
		{
			Value:          []byte("PUT "),
			ExpectedResult: []byte("PUT"),
			ExpectedRest:      []byte(" "),
			ExpectedOK: true,
		},
		{
			Value:          []byte("TRACE "),
			ExpectedResult: []byte("TRACE"),
			ExpectedRest:        []byte(" "),
			ExpectedOK: true,
		},
		{
			Value:          []byte("/GET "),
			ExpectedResult: []byte("/GET"),
			ExpectedRest:       []byte(" "),
			ExpectedOK: true,
		},
		{
			Value:          []byte("/W "),
			ExpectedResult: []byte("/W"),
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
			Value:          []byte("CONNECT\t"),
			ExpectedResult: []byte("CONNECT"),
			ExpectedRest:          []byte("\t"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("DELETE\t"),
			ExpectedResult: []byte("DELETE"),
			ExpectedRest:         []byte("\t"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("GET\t"),
			ExpectedResult: []byte("GET"),
			ExpectedRest:      []byte("\t"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("HEAD\t"),
			ExpectedResult: []byte("HEAD"),
			ExpectedRest:       []byte("\t"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("PATCH\t"),
			ExpectedResult: []byte("PATCH"),
			ExpectedRest:        []byte("\t"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("POST\t"),
			ExpectedResult: []byte("POST"),
			ExpectedRest:       []byte("\t"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("PUT\t"),
			ExpectedResult: []byte("PUT"),
			ExpectedRest:      []byte("\t"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("TRACE\t"),
			ExpectedResult: []byte("TRACE"),
			ExpectedRest:        []byte("\t"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("/GET\t"),
			ExpectedResult: []byte("/GET"),
			ExpectedRest:       []byte("\t"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("/W\t"),
			ExpectedResult: []byte("/W"),
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
			Value:          []byte("CONNECT\r"),
			ExpectedResult: []byte("CONNECT"),
			ExpectedRest:          []byte("\r"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("DELETE\r"),
			ExpectedResult: []byte("DELETE"),
			ExpectedRest:         []byte("\r"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("GET\r"),
			ExpectedResult: []byte("GET"),
			ExpectedRest:      []byte("\r"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("HEAD\r"),
			ExpectedResult: []byte("HEAD"),
			ExpectedRest:       []byte("\r"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("PATCH\r"),
			ExpectedResult: []byte("PATCH"),
			ExpectedRest:        []byte("\r"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("POST\r"),
			ExpectedResult: []byte("POST"),
			ExpectedRest:       []byte("\r"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("PUT\r"),
			ExpectedResult: []byte("PUT"),
			ExpectedRest:      []byte("\r"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("TRACE\r"),
			ExpectedResult: []byte("TRACE"),
			ExpectedRest:        []byte("\r"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("/GET\r"),
			ExpectedResult: []byte("/GET"),
			ExpectedRest:       []byte("\r"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("/W\r"),
			ExpectedResult: []byte("/W"),
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
			Value:          []byte("CONNECT\n"),
			ExpectedResult: []byte("CONNECT"),
			ExpectedRest:          []byte("\n"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("DELETE\n"),
			ExpectedResult: []byte("DELETE"),
			ExpectedRest:         []byte("\n"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("GET\n"),
			ExpectedResult: []byte("GET"),
			ExpectedRest:      []byte("\n"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("HEAD\n"),
			ExpectedResult: []byte("HEAD"),
			ExpectedRest:       []byte("\n"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("PATCH\n"),
			ExpectedResult: []byte("PATCH"),
			ExpectedRest:        []byte("\n"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("POST\n"),
			ExpectedResult: []byte("POST"),
			ExpectedRest:       []byte("\n"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("PUT\n"),
			ExpectedResult: []byte("PUT"),
			ExpectedRest:      []byte("\n"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("TRACE\n"),
			ExpectedResult: []byte("TRACE"),
			ExpectedRest:        []byte("\n"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("/GET\n"),
			ExpectedResult: []byte("/GET"),
			ExpectedRest:       []byte("\n"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("/W\n"),
			ExpectedResult: []byte("/W"),
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

		actualResult, actualRest, actualOK := method.BytesTolerant(test.Value)

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
