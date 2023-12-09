package eol_test

import (
	"testing"

	"bytes"

	"sourcecode.social/reiver/go-rfc2616/eol"
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
			Value         : []byte("\n"),
			ExpectedResult: []byte("\n"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value         : []byte("\r\n"),
			ExpectedResult: []byte("\r\n"),
			ExpectedRest: nil,
			ExpectedOK: true,
		},



		{
			Value         : []byte("\nHost: example.com\n"),
			ExpectedResult: []byte("\n"),
			ExpectedRest:     []byte("Host: example.com\n"),
			ExpectedOK: true,
		},
		{
			Value         : []byte("\r\nHost: example.com\r\n"),
			ExpectedResult: []byte("\r\n"),
			ExpectedRest:       []byte("Host: example.com\r\n"),
			ExpectedOK: true,
		},



		{
			Value:        []byte("\r"),
			ExpectedResult: nil,
			ExpectedRest: []byte("\r"),
			ExpectedOK: false,
		},
	}

	for testNumber, test := range tests {

		actualResult, actualRest, actualOK := eol.BytesTolerant(test.Value)

		{
			expected := test.ExpectedOK
			actual   := actualOK

			if expected != actual {
				t.Errorf("For test #%d, the actual ok-result is not what wad expected.", testNumber)
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
				t.Errorf("For test #%d, the actual result is not what wad expected.", testNumber)
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
				t.Errorf("For test #%d, the actual rest is not what wad expected.", testNumber)
				t.Logf("EXPECTED: %q (%#v)", expected, expected)
				t.Logf("ACTUAL:   %q (%#v)", actual, actual)
				t.Logf("VALUE: %q (%#v)", test.Value, test.Value)
				continue
			}
		}
	}
}
