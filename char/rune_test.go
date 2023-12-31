package char_test

import (
	"testing"

	"sourcecode.social/reiver/go-rfc2616/char"
)

func TestRuneIsChar(t *testing.T) {

	tests := []struct{
		Value rune
		Expected bool
	}{
		{
			Value: 0x00,
			Expected: true,
		},
		{
			Value: 0x01,
			Expected: true,
		},
		{
			Value: 0x02,
			Expected: true,
		},
		{
			Value: 0x03,
			Expected: true,
		},
		{
			Value: 0x04,
			Expected: true,
		},
		{
			Value: 0x05,
			Expected: true,
		},
		{
			Value: 0x06,
			Expected: true,
		},
		{
			Value: 0x07,
			Expected: true,
		},
		{
			Value: 0x08,
			Expected: true,
		},
		{
			Value: 0x09,
			Expected: true,
		},
		{
			Value: 0x0A,
			Expected: true,
		},
		{
			Value: 0x0B,
			Expected: true,
		},
		{
			Value: 0x0C,
			Expected: true,
		},
		{
			Value: 0x0D,
			Expected: true,
		},
		{
			Value: 0x0E,
			Expected: true,
		},
		{
			Value: 0x0F,
			Expected: true,
		},


		{
			Value: 0x10,
			Expected: true,
		},
		{
			Value: 0x11,
			Expected: true,
		},
		{
			Value: 0x12,
			Expected: true,
		},
		{
			Value: 0x13,
			Expected: true,
		},
		{
			Value: 0x14,
			Expected: true,
		},
		{
			Value: 0x15,
			Expected: true,
		},
		{
			Value: 0x16,
			Expected: true,
		},
		{
			Value: 0x17,
			Expected: true,
		},
		{
			Value: 0x18,
			Expected: true,
		},
		{
			Value: 0x19,
			Expected: true,
		},
		{
			Value: 0x1A,
			Expected: true,
		},
		{
			Value: 0x1B,
			Expected: true,
		},
		{
			Value: 0x1C,
			Expected: true,
		},
		{
			Value: 0x1D,
			Expected: true,
		},
		{
			Value: 0x1E,
			Expected: true,
		},
		{
			Value: 0x1F,
			Expected: true,
		},


		{
			Value: 0x20,    // == ' ' // I.e., space
			Expected: true, // <----------
		},
		{
			Value: 0x21,
			Expected: true,
		},
		{
			Value: 0x22,
			Expected: true,
		},
		{
			Value: 0x23,
			Expected: true,
		},
		{
			Value: 0x24,
			Expected: true,
		},
		{
			Value: 0x25,
			Expected: true,
		},
		{
			Value: 0x26,
			Expected: true,
		},
		{
			Value: 0x27,
			Expected: true,
		},
		{
			Value: 0x28,
			Expected: true,
		},
		{
			Value: 0x29,
			Expected: true,
		},
		{
			Value: 0x2A,
			Expected: true,
		},
		{
			Value: 0x2B,
			Expected: true,
		},
		{
			Value: 0x2C,
			Expected: true,
		},
		{
			Value: 0x2D,
			Expected: true,
		},
		{
			Value: 0x2E,
			Expected: true,
		},
		{
			Value: 0x2F,
			Expected: true,
		},


		{
			Value: 0x30,
			Expected: true,
		},
		{
			Value: 0x31,
			Expected: true,
		},
		{
			Value: 0x32,
			Expected: true,
		},
		{
			Value: 0x33,
			Expected: true,
		},
		{
			Value: 0x34,
			Expected: true,
		},
		{
			Value: 0x35,
			Expected: true,
		},
		{
			Value: 0x36,
			Expected: true,
		},
		{
			Value: 0x37,
			Expected: true,
		},
		{
			Value: 0x38,
			Expected: true,
		},
		{
			Value: 0x39,
			Expected: true,
		},
		{
			Value: 0x3A,
			Expected: true,
		},
		{
			Value: 0x3B,
			Expected: true,
		},
		{
			Value: 0x3C,
			Expected: true,
		},
		{
			Value: 0x3D,
			Expected: true,
		},
		{
			Value: 0x3E,
			Expected: true,
		},
		{
			Value: 0x3F,
			Expected: true,
		},


		{
			Value: 0x40,
			Expected: true,
		},
		{
			Value: 0x41,
			Expected: true,
		},
		{
			Value: 0x42,
			Expected: true,
		},
		{
			Value: 0x43,
			Expected: true,
		},
		{
			Value: 0x44,
			Expected: true,
		},
		{
			Value: 0x45,
			Expected: true,
		},
		{
			Value: 0x46,
			Expected: true,
		},
		{
			Value: 0x47,
			Expected: true,
		},
		{
			Value: 0x48,
			Expected: true,
		},
		{
			Value: 0x49,
			Expected: true,
		},
		{
			Value: 0x4A,
			Expected: true,
		},
		{
			Value: 0x4B,
			Expected: true,
		},
		{
			Value: 0x4C,
			Expected: true,
		},
		{
			Value: 0x4D,
			Expected: true,
		},
		{
			Value: 0x4E,
			Expected: true,
		},
		{
			Value: 0x4F,
			Expected: true,
		},


		{
			Value: 0x50,
			Expected: true,
		},
		{
			Value: 0x51,
			Expected: true,
		},
		{
			Value: 0x52,
			Expected: true,
		},
		{
			Value: 0x53,
			Expected: true,
		},
		{
			Value: 0x54,
			Expected: true,
		},
		{
			Value: 0x55,
			Expected: true,
		},
		{
			Value: 0x56,
			Expected: true,
		},
		{
			Value: 0x57,
			Expected: true,
		},
		{
			Value: 0x58,
			Expected: true,
		},
		{
			Value: 0x59,
			Expected: true,
		},
		{
			Value: 0x5A,
			Expected: true,
		},
		{
			Value: 0x5B,
			Expected: true,
		},
		{
			Value: 0x5C,
			Expected: true,
		},
		{
			Value: 0x5D,
			Expected: true,
		},
		{
			Value: 0x5E,
			Expected: true,
		},
		{
			Value: 0x5F,
			Expected: true,
		},


		{
			Value: 0x60,
			Expected: true,
		},
		{
			Value: 0x61,
			Expected: true,
		},
		{
			Value: 0x62,
			Expected: true,
		},
		{
			Value: 0x63,
			Expected: true,
		},
		{
			Value: 0x64,
			Expected: true,
		},
		{
			Value: 0x65,
			Expected: true,
		},
		{
			Value: 0x66,
			Expected: true,
		},
		{
			Value: 0x67,
			Expected: true,
		},
		{
			Value: 0x68,
			Expected: true,
		},
		{
			Value: 0x69,
			Expected: true,
		},
		{
			Value: 0x6A,
			Expected: true,
		},
		{
			Value: 0x6B,
			Expected: true,
		},
		{
			Value: 0x6C,
			Expected: true,
		},
		{
			Value: 0x6D,
			Expected: true,
		},
		{
			Value: 0x6E,
			Expected: true,
		},
		{
			Value: 0x6F,
			Expected: true,
		},


		{
			Value: 0x70,
			Expected: true,
		},
		{
			Value: 0x71,
			Expected: true,
		},
		{
			Value: 0x72,
			Expected: true,
		},
		{
			Value: 0x73,
			Expected: true,
		},
		{
			Value: 0x74,
			Expected: true,
		},
		{
			Value: 0x75,
			Expected: true,
		},
		{
			Value: 0x76,
			Expected: true,
		},
		{
			Value: 0x77,
			Expected: true,
		},
		{
			Value: 0x78,
			Expected: true,
		},
		{
			Value: 0x79,
			Expected: true,
		},
		{
			Value: 0x7A,
			Expected: true,
		},
		{
			Value: 0x7B,
			Expected: true,
		},
		{
			Value: 0x7C,
			Expected: true,
		},
		{
			Value: 0x7D,
			Expected: true,
		},
		{
			Value: 0x7E,
			Expected: true,
		},
		{
			Value: 0x7F,
			Expected: true,
		},


		{
			Value: 0x80,
			Expected: false,
		},
		{
			Value: 0x81,
			Expected: false,
		},
		{
			Value: 0x82,
			Expected: false,
		},
		{
			Value: 0x83,
			Expected: false,
		},
		{
			Value: 0x84,
			Expected: false,
		},
		{
			Value: 0x85,
			Expected: false,
		},
		{
			Value: 0x86,
			Expected: false,
		},
		{
			Value: 0x87,
			Expected: false,
		},
		{
			Value: 0x88,
			Expected: false,
		},
		{
			Value: 0x89,
			Expected: false,
		},
		{
			Value: 0x8A,
			Expected: false,
		},
		{
			Value: 0x8B,
			Expected: false,
		},
		{
			Value: 0x8C,
			Expected: false,
		},
		{
			Value: 0x8D,
			Expected: false,
		},
		{
			Value: 0x8E,
			Expected: false,
		},
		{
			Value: 0x8F,
			Expected: false,
		},


		{
			Value: 0x90,
			Expected: false,
		},
		{
			Value: 0x91,
			Expected: false,
		},
		{
			Value: 0x92,
			Expected: false,
		},
		{
			Value: 0x93,
			Expected: false,
		},
		{
			Value: 0x94,
			Expected: false,
		},
		{
			Value: 0x95,
			Expected: false,
		},
		{
			Value: 0x96,
			Expected: false,
		},
		{
			Value: 0x97,
			Expected: false,
		},
		{
			Value: 0x98,
			Expected: false,
		},
		{
			Value: 0x99,
			Expected: false,
		},
		{
			Value: 0x9A,
			Expected: false,
		},
		{
			Value: 0x9B,
			Expected: false,
		},
		{
			Value: 0x9C,
			Expected: false,
		},
		{
			Value: 0x9D,
			Expected: false,
		},
		{
			Value: 0x9E,
			Expected: false,
		},
		{
			Value: 0x9F,
			Expected: false,
		},


		{
			Value: 0xA0,
			Expected: false,
		},
		{
			Value: 0xA1,
			Expected: false,
		},
		{
			Value: 0xA2,
			Expected: false,
		},
		{
			Value: 0xA3,
			Expected: false,
		},
		{
			Value: 0xA4,
			Expected: false,
		},
		{
			Value: 0xA5,
			Expected: false,
		},
		{
			Value: 0xA6,
			Expected: false,
		},
		{
			Value: 0xA7,
			Expected: false,
		},
		{
			Value: 0xA8,
			Expected: false,
		},
		{
			Value: 0xA9,
			Expected: false,
		},
		{
			Value: 0xAA,
			Expected: false,
		},
		{
			Value: 0xAB,
			Expected: false,
		},
		{
			Value: 0xAC,
			Expected: false,
		},
		{
			Value: 0xAD,
			Expected: false,
		},
		{
			Value: 0xAE,
			Expected: false,
		},
		{
			Value: 0xAF,
			Expected: false,
		},


		{
			Value: 0xB0,
			Expected: false,
		},
		{
			Value: 0xB1,
			Expected: false,
		},
		{
			Value: 0xB2,
			Expected: false,
		},
		{
			Value: 0xB3,
			Expected: false,
		},
		{
			Value: 0xB4,
			Expected: false,
		},
		{
			Value: 0xB5,
			Expected: false,
		},
		{
			Value: 0xB6,
			Expected: false,
		},
		{
			Value: 0xB7,
			Expected: false,
		},
		{
			Value: 0xB8,
			Expected: false,
		},
		{
			Value: 0xB9,
			Expected: false,
		},
		{
			Value: 0xBA,
			Expected: false,
		},
		{
			Value: 0xBB,
			Expected: false,
		},
		{
			Value: 0xBC,
			Expected: false,
		},
		{
			Value: 0xBD,
			Expected: false,
		},
		{
			Value: 0xBE,
			Expected: false,
		},
		{
			Value: 0xBF,
			Expected: false,
		},


		{
			Value: 0xC0,
			Expected: false,
		},
		{
			Value: 0xC1,
			Expected: false,
		},
		{
			Value: 0xC2,
			Expected: false,
		},
		{
			Value: 0xC3,
			Expected: false,
		},
		{
			Value: 0xC4,
			Expected: false,
		},
		{
			Value: 0xC5,
			Expected: false,
		},
		{
			Value: 0xC6,
			Expected: false,
		},
		{
			Value: 0xC7,
			Expected: false,
		},
		{
			Value: 0xC8,
			Expected: false,
		},
		{
			Value: 0xC9,
			Expected: false,
		},
		{
			Value: 0xCA,
			Expected: false,
		},
		{
			Value: 0xCB,
			Expected: false,
		},
		{
			Value: 0xCC,
			Expected: false,
		},
		{
			Value: 0xCD,
			Expected: false,
		},
		{
			Value: 0xCE,
			Expected: false,
		},
		{
			Value: 0xCF,
			Expected: false,
		},


		{
			Value: 0xD0,
			Expected: false,
		},
		{
			Value: 0xD1,
			Expected: false,
		},
		{
			Value: 0xD2,
			Expected: false,
		},
		{
			Value: 0xD3,
			Expected: false,
		},
		{
			Value: 0xD4,
			Expected: false,
		},
		{
			Value: 0xD5,
			Expected: false,
		},
		{
			Value: 0xD6,
			Expected: false,
		},
		{
			Value: 0xD7,
			Expected: false,
		},
		{
			Value: 0xD8,
			Expected: false,
		},
		{
			Value: 0xD9,
			Expected: false,
		},
		{
			Value: 0xDA,
			Expected: false,
		},
		{
			Value: 0xDB,
			Expected: false,
		},
		{
			Value: 0xDC,
			Expected: false,
		},
		{
			Value: 0xDD,
			Expected: false,
		},
		{
			Value: 0xDE,
			Expected: false,
		},
		{
			Value: 0xDF,
			Expected: false,
		},


		{
			Value: 0xE0,
			Expected: false,
		},
		{
			Value: 0xE1,
			Expected: false,
		},
		{
			Value: 0xE2,
			Expected: false,
		},
		{
			Value: 0xE3,
			Expected: false,
		},
		{
			Value: 0xE4,
			Expected: false,
		},
		{
			Value: 0xE5,
			Expected: false,
		},
		{
			Value: 0xE6,
			Expected: false,
		},
		{
			Value: 0xE7,
			Expected: false,
		},
		{
			Value: 0xE8,
			Expected: false,
		},
		{
			Value: 0xE9,
			Expected: false,
		},
		{
			Value: 0xEA,
			Expected: false,
		},
		{
			Value: 0xEB,
			Expected: false,
		},
		{
			Value: 0xEC,
			Expected: false,
		},
		{
			Value: 0xED,
			Expected: false,
		},
		{
			Value: 0xEE,
			Expected: false,
		},
		{
			Value: 0xEF,
			Expected: false,
		},


		{
			Value: 0xF0,
			Expected: false,
		},
		{
			Value: 0xF1,
			Expected: false,
		},
		{
			Value: 0xF2,
			Expected: false,
		},
		{
			Value: 0xF3,
			Expected: false,
		},
		{
			Value: 0xF4,
			Expected: false,
		},
		{
			Value: 0xF5,
			Expected: false,
		},
		{
			Value: 0xF6,
			Expected: false,
		},
		{
			Value: 0xF7,
			Expected: false,
		},
		{
			Value: 0xF8,
			Expected: false,
		},
		{
			Value: 0xF9,
			Expected: false,
		},
		{
			Value: 0xFA,
			Expected: false,
		},
		{
			Value: 0xFB,
			Expected: false,
		},
		{
			Value: 0xFC,
			Expected: false,
		},
		{
			Value: 0xFD,
			Expected: false,
		},
		{
			Value: 0xFE,
			Expected: false,
		},
		{
			Value: 0xFF,
			Expected: false,
		},
	}

	for testNumber, test := range tests {

		actual := char.RuneIsChar(test.Value)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual result is not what was expected.", testNumber)
			t.Logf("EXPECTED: %t", expected)
			t.Logf("ACTUAL:   %t", actual)
			t.Logf("VALUE: %q (0x%02X)", test.Value, test.Value)
			continue
		}
	}
}
