package requestline_test

import (
	"testing"

	"bytes"

	"sourcecode.social/reiver/go-rfc2616/requestline"
)

func TestBytesTolerant(t *testing.T) {

	tests := []struct{
		Value []byte
		ExpectedMethod []byte
		ExpectedRequestURI []byte
		ExpectedVersion []byte
		ExpectedRest []byte
		ExpectedOK bool
	}{
		{
			Value: nil,
			ExpectedMethod: nil,
			ExpectedRequestURI: nil,
			ExpectedVersion: nil,
			ExpectedRest: nil,
			ExpectedOK: false,
		},
		{
			Value: []byte(nil),
			ExpectedMethod: nil,
			ExpectedRequestURI: nil,
			ExpectedVersion: nil,
			ExpectedRest: nil,
			ExpectedOK: false,
		},



		{
			Value: []byte{},
			ExpectedMethod: nil,
			ExpectedRequestURI: nil,
			ExpectedVersion: nil,
			ExpectedRest: nil,
			ExpectedOK: false,
		},
		{
			Value: []byte(""),
			ExpectedMethod: nil,
			ExpectedRequestURI: nil,
			ExpectedVersion: nil,
			ExpectedRest: nil,
			ExpectedOK: false,
		},



		{
			Value:        []byte("GET"),
			ExpectedMethod: nil,
			ExpectedRequestURI: nil,
			ExpectedVersion: nil,
			ExpectedRest: []byte("GET"),
			ExpectedOK: false,
		},
		{
			Value:        []byte("GET "),
			ExpectedMethod: nil,
			ExpectedRequestURI: nil,
			ExpectedVersion: nil,
			ExpectedRest: []byte("GET "),
			ExpectedOK: false,
		},
		{
			Value:          []byte("GET /"),
			ExpectedMethod: []byte("GET"),
			ExpectedRequestURI: []byte("/"),
			ExpectedVersion: nil,
			ExpectedRest: nil,
			ExpectedOK: true,
		},

		{
			Value:          []byte("GET /\n"),
			ExpectedMethod: []byte("GET"),
			ExpectedRequestURI: []byte("/"),
			ExpectedVersion: nil,
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("GET /\r\n"),
			ExpectedMethod: []byte("GET"),
			ExpectedRequestURI: []byte("/"),
			ExpectedVersion: nil,
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("GET /\nHost: example.com\n\n"),
			ExpectedMethod: []byte("GET"),
			ExpectedRequestURI: []byte("/"),
			ExpectedVersion: nil,
			ExpectedRest:          []byte("Host: example.com\n\n"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("GET /\r\nHost: example.com\r\n\r\n"),
			ExpectedMethod: []byte("GET"),
			ExpectedRequestURI: []byte("/"),
			ExpectedVersion: nil,
			ExpectedRest:            []byte("Host: example.com\r\n\r\n"),
			ExpectedOK: true,
		},

		{
			Value:          []byte("GET / "),
			ExpectedMethod: []byte("GET"),
			ExpectedRequestURI: []byte("/"),
			ExpectedVersion: nil,
			ExpectedRest: nil,
			ExpectedOK: true,
		},

		{
			Value:          []byte("GET / \n"),
			ExpectedMethod: []byte("GET"),
			ExpectedRequestURI: []byte("/"),
			ExpectedVersion: nil,
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("GET / \r\n"),
			ExpectedMethod: []byte("GET"),
			ExpectedRequestURI: []byte("/"),
			ExpectedVersion: nil,
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("GET / \nHost: example.com\n"),
			ExpectedMethod: []byte("GET"),
			ExpectedRequestURI: []byte("/"),
			ExpectedVersion: nil,
			ExpectedRest:           []byte("Host: example.com\n"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("GET / \r\nHost: example.com\r\n"),
			ExpectedMethod: []byte("GET"),
			ExpectedRequestURI: []byte("/"),
			ExpectedVersion: nil,
			ExpectedRest:             []byte("Host: example.com\r\n"),
			ExpectedOK: true,
		},

		{
			Value:          []byte("GET / HTTP/1.1"),
			ExpectedMethod: []byte("GET"),
			ExpectedRequestURI: []byte("/"),
			ExpectedVersion:      []byte("HTTP/1.1"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("GET / HTTP/1.1\n"),
			ExpectedMethod: []byte("GET"),
			ExpectedRequestURI: []byte("/"),
			ExpectedVersion:      []byte("HTTP/1.1"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("GET / HTTP/1.1\r\n"),
			ExpectedMethod: []byte("GET"),
			ExpectedRequestURI: []byte("/"),
			ExpectedVersion:      []byte("HTTP/1.1"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:          []byte("GET / HTTP/1.1\nHost: example.com\n\n"),
			ExpectedMethod: []byte("GET"),
			ExpectedRequestURI: []byte("/"),
			ExpectedVersion:      []byte("HTTP/1.1"),
			ExpectedRest:                   []byte("Host: example.com\n\n"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("GET / HTTP/1.1\r\nHost: example.com\r\n\r\n"),
			ExpectedMethod: []byte("GET"),
			ExpectedRequestURI: []byte("/"),
			ExpectedVersion:      []byte("HTTP/1.1"),
			ExpectedRest:                     []byte("Host: example.com\r\n\r\n"),
			ExpectedOK: true,
		},



		{
			Value:          []byte("PUNCH /apple/banana/cherry.php HTTP/1.1\r\nHost: example.com\r\n\r\n"),
			ExpectedMethod: []byte("PUNCH"),
			ExpectedRequestURI:   []byte("/apple/banana/cherry.php"),
			ExpectedVersion:                               []byte("HTTP/1.1"),
			ExpectedRest:                                              []byte("Host: example.com\r\n\r\n"),
			ExpectedOK: true,
		},
	}

	for testNumber, test := range tests {

		actualMethod, actualRequestURI, actualVersion, actualRest, actualOK := requestline.BytesTolerant(test.Value)

		{
			expected := test.ExpectedOK
			actual   := actualOK

			if expected != actual {
				t.Errorf("For test #%d, the actual ok-result is not what was expected.", testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:   %t", actual)
				t.Logf("VALUE; %q (%#v)", test.Value, test.Value)
				continue
			}
		}

		{
			expected := test.ExpectedMethod
			actual   := actualMethod

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual method is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q (%#v)", expected, expected)
				t.Logf("ACTUAL:   %q (%#v)", actual, actual)
				t.Logf("VALUE; %q (%#v)", test.Value, test.Value)
				continue
			}
		}

		{
			expected := test.ExpectedRequestURI
			actual   := actualRequestURI

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual request-uri is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q (%#v)", expected, expected)
				t.Logf("ACTUAL:   %q (%#v)", actual, actual)
				t.Logf("VALUE; %q (%#v)", test.Value, test.Value)
				continue
			}
		}

		{
			expected := test.ExpectedVersion
			actual   := actualVersion

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual version is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q (%#v)", expected, expected)
				t.Logf("ACTUAL:   %q (%#v)", actual, actual)
				t.Logf("VALUE; %q (%#v)", test.Value, test.Value)
				continue
			}
		}

		{
			expected := test.ExpectedRest
			actual   := actualRest

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual rest is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q (%#v)", expected, expected)
				t.Logf("ACTUAL:   %q (%#v)", actual, actual)
				t.Logf("VALUE; %q (%#v)", test.Value, test.Value)
				continue
			}
		}
	}
}
