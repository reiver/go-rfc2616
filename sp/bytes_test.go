package sp_test

import (
	"testing"

	"bytes"

	"sourcecode.social/reiver/go-rfc2616/sp"
)

func TestBytes(t *testing.T) {

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
			Value:        []byte("apple banana cherry"),
			ExpectedResult: nil,
			ExpectedRest: []byte("apple banana cherry"),
			ExpectedOK: false,
		},

		{
			Value:          []byte(" apple banana cherry"),
			ExpectedResult: []byte(" "),
			ExpectedRest:    []byte("apple banana cherry"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("  apple banana cherry"),
			ExpectedResult: []byte("  "),
			ExpectedRest:     []byte("apple banana cherry"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("   apple banana cherry"),
			ExpectedResult: []byte("   "),
			ExpectedRest:      []byte("apple banana cherry"),
			ExpectedOK: true,
		},
		{
			Value:          []byte("    apple banana cherry"),
			ExpectedResult: []byte("    "),
			ExpectedRest:       []byte("apple banana cherry"),
			ExpectedOK: true,
		},

		{
			Value:        []byte("\tapple banana cherry"),
			ExpectedResult: nil,
			ExpectedRest: []byte("\tapple banana cherry"),
			ExpectedOK: false,
		},
		{
			Value:        []byte("\t\tapple banana cherry"),
			ExpectedResult: nil,
			ExpectedRest: []byte("\t\tapple banana cherry"),
			ExpectedOK: false,
		},
		{
			Value:        []byte("\t\t\tapple banana cherry"),
			ExpectedResult: nil,
			ExpectedRest: []byte("\t\t\tapple banana cherry"),
			ExpectedOK: false,
		},
		{
			Value:        []byte("\t\t\t\tapple banana cherry"),
			ExpectedResult: nil,
			ExpectedRest: []byte("\t\t\t\tapple banana cherry"),
			ExpectedOK: false,
		},



		{
			Value:        []byte("\t\t     apple banana cherry"),
			ExpectedResult: nil,
			ExpectedRest: []byte("\t\t     apple banana cherry"),
			ExpectedOK: false,
		},
		{
			Value:          []byte("  \t\t\t\t\tapple banana cherry"),
			ExpectedResult: []byte("  "),
			ExpectedRest:     []byte("\t\t\t\t\tapple banana cherry"),
			ExpectedOK: true,
		},
	}

	for testNumber, test := range tests {

		actualResult, actualRest, actualOK := sp.Bytes(test.Value)

		{
			expected := test.ExpectedOK
			actual   := actualOK

			if expected != actual {
				t.Errorf("For test #%d, the actual ok-result is not what was expected.", testNumber)
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
				t.Errorf("For test #%d, the actual result is not what was expected.", testNumber)
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
				t.Errorf("For test #%d, the actual result is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q (%#v)", expected, expected)
				t.Logf("ACTUAL:   %q (%#v)", actual, actual)
				t.Logf("VALUE: %q (%#v)", test.Value, test.Value)
				continue
			}
		}

	}
}
