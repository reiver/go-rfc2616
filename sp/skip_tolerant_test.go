package sp_test

import (
	"testing"

	"bytes"

	"sourcecode.social/reiver/go-rfc2616/sp"
)

func TestSkipTolerant(t *testing.T) {

	tests := []struct{
		Value []byte
		Expected []byte
	}{
		{
			Value: nil,
			Expected: nil,
		},
		{
			Value: []byte(nil),
			Expected: nil,
		},



		{
			Value: []byte{},
			Expected: nil,
		},
		{
			Value: []byte(""),
			Expected: nil,
		},



		{
			Value: []byte(" "),
			Expected: nil,
		},
		{
			Value: []byte("  "),
			Expected: nil,
		},
		{
			Value: []byte("   "),
			Expected: nil,
		},
		{
			Value: []byte("    "),
			Expected: nil,
		},
		{
			Value: []byte("     "),
			Expected: nil,
		},
		{
			Value: []byte("      "),
			Expected: nil,
		},
		{
			Value: []byte("       "),
			Expected: nil,
		},

		{
			Value: []byte("                                        "),
			Expected: nil,
		},



		{
			Value: []byte("\t"),
			Expected: nil,
		},
		{
			Value: []byte("\t\t"),
			Expected: nil,
		},
		{
			Value: []byte("\t\t\t"),
			Expected: nil,
		},
		{
			Value: []byte("\t\t\t\t"),
			Expected: nil,
		},
		{
			Value: []byte("\t\t\t\t\t"),
			Expected: nil,
		},
		{
			Value: []byte("\t\t\t\t\t\t"),
			Expected: nil,
		},
		{
			Value: []byte("\t\t\t\t\t\t\t"),
			Expected: nil,
		},

		{
			Value: []byte("\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t"),
			Expected: nil,
		},



		{
			Value: []byte("\t  \t         \t  \t\t"),
			Expected: nil,
		},



		{
			Value:   []byte(" apple "),
			Expected: []byte("apple "),
		},
		{
			Value:  []byte("\tbanana"),
			Expected: []byte("banana"),
		},
		{
			Value: []byte("\t\t cherry\r\n"),
			Expected:   []byte("cherry\r\n"),
		},



		{
			Value: []byte("  \t\t 😈 :-)\r\n"),
			Expected:     []byte("😈 :-)\r\n"),
		},



		{
			Value:    []byte("once :-)\r\n"),
			Expected: []byte("once :-)\r\n"),
		},
		{
			Value:    []byte("twice 😈 *-*\n"),
			Expected: []byte("twice 😈 *-*\n"),
		},
		{
			Value:    []byte("thrice 🙂"),
			Expected: []byte("thrice 🙂"),
		},
		{
			Value:    []byte("fource 👾"),
			Expected: []byte("fource 👾"),
		},
	}

	for testNumber, test := range tests {

		actual := sp.SkipTolerant(test.Value)

		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual result is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q (%#v)", expected, expected)
			t.Logf("ACTUAL:   %q (%#v)", actual, actual)
			t.Logf("VALUE: %q (%#v)", test.Value, test.Value)
			continue
		}

	}
}
