package token_test

import (
	"testing"

	"bytes"

	"sourcecode.social/reiver/go-rfc2616/token"
)

func TestBytes(t *testing.T) {

	tests := []struct{
		Value []byte
		ExpectedToken []byte
		ExpectedRest []byte
		ExpectedOK bool
	}{
		{
			Value:         []byte("CONNECT /once/twice/thrice/fource.txt HTTP/1.1\r\n"),
			ExpectedToken: []byte("CONNECT"),
			ExpectedRest:         []byte(" /once/twice/thrice/fource.txt HTTP/1.1\r\n"),
			ExpectedOK: true,
		},
		{
			Value:         []byte("DELETE /once/twice/thrice/fource.txt HTTP/1.1\r\n"),
			ExpectedToken: []byte("DELETE"),
			ExpectedRest:        []byte(" /once/twice/thrice/fource.txt HTTP/1.1\r\n"),
			ExpectedOK: true,
		},
		{
			Value:         []byte("GET /once/twice/thrice/fource.txt HTTP/1.1\r\n"),
			ExpectedToken: []byte("GET"),
			ExpectedRest:     []byte(" /once/twice/thrice/fource.txt HTTP/1.1\r\n"),
			ExpectedOK: true,
		},
		{
			Value:         []byte("HEAD /once/twice/thrice/fource.txt HTTP/1.1\r\n"),
			ExpectedToken: []byte("HEAD"),
			ExpectedRest:      []byte(" /once/twice/thrice/fource.txt HTTP/1.1\r\n"),
			ExpectedOK: true,
		},
		{
			Value:         []byte("PATCH /once/twice/thrice/fource.txt HTTP/1.1\r\n"),
			ExpectedToken: []byte("PATCH"),
			ExpectedRest:       []byte(" /once/twice/thrice/fource.txt HTTP/1.1\r\n"),
			ExpectedOK: true,
		},
		{
			Value:         []byte("POST /once/twice/thrice/fource.txt HTTP/1.1\r\n"),
			ExpectedToken: []byte("POST"),
			ExpectedRest:      []byte(" /once/twice/thrice/fource.txt HTTP/1.1\r\n"),
			ExpectedOK: true,
		},
		{
			Value:         []byte("PUT /once/twice/thrice/fource.txt HTTP/1.1\r\n"),
			ExpectedToken: []byte("PUT"),
			ExpectedRest:     []byte(" /once/twice/thrice/fource.txt HTTP/1.1\r\n"),
			ExpectedOK: true,
		},
		{
			Value:         []byte("TRACE /once/twice/thrice/fource.txt HTTP/1.1\r\n"),
			ExpectedToken: []byte("TRACE"),
			ExpectedRest:       []byte(" /once/twice/thrice/fource.txt HTTP/1.1\r\n"),
			ExpectedOK: true,
		},



		{
			Value:         []byte("KICK /apple/banana/cherry.php HTTP/1.1\r\n"),
			ExpectedToken: []byte("KICK"),
			ExpectedRest:      []byte(" /apple/banana/cherry.php HTTP/1.1\r\n"),
			ExpectedOK: true,
		},
		{
			Value:         []byte("PUNCH /apple/banana/cherry.php HTTP/1.1\r\n"),
			ExpectedToken: []byte("PUNCH"),
			ExpectedRest:       []byte(" /apple/banana/cherry.php HTTP/1.1\r\n"),
			ExpectedOK: true,
		},



		{
			Value: nil,
			ExpectedToken: nil,
			ExpectedRest: nil,
			ExpectedOK: false,
		},
		{
			Value: []byte(nil),
			ExpectedToken: nil,
			ExpectedRest: nil,
			ExpectedOK: false,
		},



		{
			Value: []byte(""),
			ExpectedToken: nil,
			ExpectedRest: nil,
			ExpectedOK: false,
		},



		{
			Value:        []byte{0},
			ExpectedToken: nil,
			ExpectedRest: []byte{0},
			ExpectedOK: false,
		},
		{
			Value:        []byte{1},
			ExpectedToken: nil,
			ExpectedRest: []byte{1},
			ExpectedOK: false,
		},
		{
			Value:        []byte{2},
			ExpectedToken: nil,
			ExpectedRest: []byte{2},
			ExpectedOK: false,
		},
		{
			Value:        []byte{3},
			ExpectedToken: nil,
			ExpectedRest: []byte{3},
			ExpectedOK: false,
		},
		{
			Value:        []byte{4},
			ExpectedToken: nil,
			ExpectedRest: []byte{4},
			ExpectedOK: false,
		},
		{
			Value:        []byte{5},
			ExpectedToken: nil,
			ExpectedRest: []byte{5},
			ExpectedOK: false,
		},
		{
			Value:        []byte{6},
			ExpectedToken: nil,
			ExpectedRest: []byte{6},
			ExpectedOK: false,
		},
		{
			Value:        []byte{7},
			ExpectedToken: nil,
			ExpectedRest: []byte{7},
			ExpectedOK: false,
		},
		{
			Value:        []byte{8},
			ExpectedToken: nil,
			ExpectedRest: []byte{8},
			ExpectedOK: false,
		},
		{
			Value:        []byte{9},
			ExpectedToken: nil,
			ExpectedRest: []byte{9},
			ExpectedOK: false,
		},
		{
			Value:        []byte{10},
			ExpectedToken: nil,
			ExpectedRest: []byte{10},
			ExpectedOK: false,
		},
		{
			Value:        []byte{11},
			ExpectedToken: nil,
			ExpectedRest: []byte{11},
			ExpectedOK: false,
		},
		{
			Value:        []byte{12},
			ExpectedToken: nil,
			ExpectedRest: []byte{12},
			ExpectedOK: false,
		},
		{
			Value:        []byte{13},
			ExpectedToken: nil,
			ExpectedRest: []byte{13},
			ExpectedOK: false,
		},
		{
			Value:        []byte{14},
			ExpectedToken: nil,
			ExpectedRest: []byte{14},
			ExpectedOK: false,
		},
		{
			Value:        []byte{15},
			ExpectedToken: nil,
			ExpectedRest: []byte{15},
			ExpectedOK: false,
		},
		{
			Value:        []byte{16},
			ExpectedToken: nil,
			ExpectedRest: []byte{16},
			ExpectedOK: false,
		},
		{
			Value:        []byte{17},
			ExpectedToken: nil,
			ExpectedRest: []byte{17},
			ExpectedOK: false,
		},
		{
			Value:        []byte{18},
			ExpectedToken: nil,
			ExpectedRest: []byte{18},
			ExpectedOK: false,
		},
		{
			Value:        []byte{19},
			ExpectedToken: nil,
			ExpectedRest: []byte{19},
			ExpectedOK: false,
		},
		{
			Value:        []byte{20},
			ExpectedToken: nil,
			ExpectedRest: []byte{20},
			ExpectedOK: false,
		},
		{
			Value:        []byte{21},
			ExpectedToken: nil,
			ExpectedRest: []byte{21},
			ExpectedOK: false,
		},
		{
			Value:        []byte{22},
			ExpectedToken: nil,
			ExpectedRest: []byte{22},
			ExpectedOK: false,
		},
		{
			Value:        []byte{23},
			ExpectedToken: nil,
			ExpectedRest: []byte{23},
			ExpectedOK: false,
		},
		{
			Value:        []byte{24},
			ExpectedToken: nil,
			ExpectedRest: []byte{24},
			ExpectedOK: false,
		},
		{
			Value:        []byte{25},
			ExpectedToken: nil,
			ExpectedRest: []byte{25},
			ExpectedOK: false,
		},
		{
			Value:        []byte{26},
			ExpectedToken: nil,
			ExpectedRest: []byte{26},
			ExpectedOK: false,
		},
		{
			Value:        []byte{27},
			ExpectedToken: nil,
			ExpectedRest: []byte{27},
			ExpectedOK: false,
		},
		{
			Value:        []byte{28},
			ExpectedToken: nil,
			ExpectedRest: []byte{28},
			ExpectedOK: false,
		},
		{
			Value:        []byte{29},
			ExpectedToken: nil,
			ExpectedRest: []byte{29},
			ExpectedOK: false,
		},
		{
			Value:        []byte{30},
			ExpectedToken: nil,
			ExpectedRest: []byte{30},
			ExpectedOK: false,
		},
		{
			Value:        []byte{31},
			ExpectedToken: nil,
			ExpectedRest: []byte{31},
			ExpectedOK: false,
		},
		{
			Value:        []byte{127},
			ExpectedToken: nil,
			ExpectedRest: []byte{127},
			ExpectedOK: false,
		},



		{
			Value:        []byte("("),
			ExpectedToken: nil,
			ExpectedRest: []byte("("),
			ExpectedOK: false,
		},
		{
			Value:        []byte(")"),
			ExpectedToken: nil,
			ExpectedRest: []byte(")"),
			ExpectedOK: false,
		},
		{
			Value:        []byte("<"),
			ExpectedToken: nil,
			ExpectedRest: []byte("<"),
			ExpectedOK: false,
		},
		{
			Value:        []byte(">"),
			ExpectedToken: nil,
			ExpectedRest: []byte(">"),
			ExpectedOK: false,
		},
		{
			Value:        []byte("@"),
			ExpectedToken: nil,
			ExpectedRest: []byte("@"),
			ExpectedOK: false,
		},
		{
			Value:       []byte(","),
			ExpectedToken: nil,
			ExpectedRest: []byte(","),
			ExpectedOK: false,
		},
		{
			Value:        []byte(";"),
			ExpectedToken: nil,
			ExpectedRest: []byte(";"),
			ExpectedOK: false,
		},
		{
			Value:        []byte(":"),
			ExpectedToken: nil,
			ExpectedRest: []byte(":"),
			ExpectedOK: false,
		},
		{
			Value:        []byte(`\`),
			ExpectedToken: nil,
			ExpectedRest: []byte(`\`),
			ExpectedOK: false,
		},
		{
			Value:        []byte(`"`),
			ExpectedToken: nil,
			ExpectedRest: []byte(`"`),
			ExpectedOK: false,
		},
		{
			Value:        []byte("/"),
			ExpectedToken: nil,
			ExpectedRest: []byte("/"),
			ExpectedOK: false,
		},
		{
			Value:        []byte("["),
			ExpectedToken: nil,
			ExpectedRest: []byte("["),
			ExpectedOK: false,
		},
		{
			Value:        []byte("]"),
			ExpectedToken: nil,
			ExpectedRest: []byte("]"),
			ExpectedOK: false,
		},
		{
			Value:        []byte("?"),
			ExpectedToken: nil,
			ExpectedRest: []byte("?"),
			ExpectedOK: false,
		},
		{
			Value:        []byte("="),
			ExpectedToken: nil,
			ExpectedRest: []byte("="),
			ExpectedOK: false,
		},
		{
			Value:        []byte("{"),
			ExpectedToken: nil,
			ExpectedRest: []byte("{"),
			ExpectedOK: false,
		},
		{
			Value:        []byte("}"),
			ExpectedToken: nil,
			ExpectedRest: []byte("}"),
			ExpectedOK: false,
		},
		{
			Value:        []byte(" "),
			ExpectedToken: nil,
			ExpectedRest: []byte(" "),
			ExpectedOK: false,
		},
		{
			Value:        []byte("\t"),
			ExpectedToken: nil,
			ExpectedRest: []byte("\t"),
			ExpectedOK: false,
		},



		{
			Value:        []byte{0x00},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x00},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x01},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x01},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x02},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x02},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x03},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x03},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x04},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x04},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x05},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x05},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x06},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x06},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x07},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x07},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x08},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x08},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x09},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x09},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x0A},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x0A},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x0B},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x0B},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x0C},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x0C},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x0D},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x0D},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x0E},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x0E},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x0F},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x0F},
			ExpectedOK: false,
		},


		{
			Value:        []byte{0x10},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x10},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x11},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x11},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x12},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x12},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x13},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x13},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x14},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x14},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x15},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x15},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x16},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x16},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x17},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x17},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x18},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x18},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x19},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x19},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x1A},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x1A},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x1B},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x1B},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x1C},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x1C},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x1D},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x1D},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x1E},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x1E},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x1F},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x1F},
			ExpectedOK: false,
		},


		{
			Value:        []byte{0x20},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x20},
			ExpectedOK: false,
		},
		{
			Value:         []byte{0x21}, // == '!'
			ExpectedToken: []byte{0x21},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:        []byte{0x22},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x22},
			ExpectedOK: false,
		},
		{
			Value:         []byte{0x23}, // == '#'
			ExpectedToken: []byte{0x23},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x24}, // == '$'
			ExpectedToken: []byte{0x24},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x25}, // == '%'
			ExpectedToken: []byte{0x25},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x26}, // == '&'
			ExpectedToken: []byte{0x26},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x27}, // == '\''
			ExpectedToken: []byte{0x27},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:        []byte{0x28},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x28},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x29},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x29},
			ExpectedOK: false,
		},
		{
			Value:         []byte{0x2A}, // == '*'
			ExpectedToken: []byte{0x2A},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x2B}, // == '+'
			ExpectedToken: []byte{0x2B},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:        []byte{0x2C},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x2C},
			ExpectedOK: false,
		},
		{
			Value:         []byte{0x2D}, // == '-'
			ExpectedToken: []byte{0x2D},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x2E}, // == '.'
			ExpectedToken: []byte{0x2E},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:        []byte{0x2F},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x2F},
			ExpectedOK: false,
		},


		{
			Value:         []byte{0x30}, // == '0'
			ExpectedToken: []byte{0x30},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x31}, // == '1'
			ExpectedToken: []byte{0x31},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x32}, // == '2'
			ExpectedToken: []byte{0x32},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x33}, // == '3'
			ExpectedToken: []byte{0x33},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x34}, // == '4'
			ExpectedToken: []byte{0x34},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x35}, // == '5'
			ExpectedToken: []byte{0x35},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x36}, // == '6'
			ExpectedToken: []byte{0x36},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x37}, // == '7'
			ExpectedToken: []byte{0x37},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x38}, // == '8'
			ExpectedToken: []byte{0x38},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x39}, // == '9'
			ExpectedToken: []byte{0x39},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:        []byte{0x3A},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x3A},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x3B},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x3B},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x3C},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x3C},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x3D},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x3D},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x3E},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x3E},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x3F},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x3F},
			ExpectedOK: false,
		},


		{
			Value:        []byte{0x40},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x40},
			ExpectedOK: false,
		},
		{
			Value:         []byte{0x41}, // == 'A'
			ExpectedToken: []byte{0x41},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x42}, // == 'B'
			ExpectedToken: []byte{0x42},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x43}, // == 'C'
			ExpectedToken: []byte{0x43},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x44}, // == 'D'
			ExpectedToken: []byte{0x44},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x45}, // == 'E'
			ExpectedToken: []byte{0x45},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x46}, // == 'F'
			ExpectedToken: []byte{0x46},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x47}, // == 'G'
			ExpectedToken: []byte{0x47},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x48}, // == 'H'
			ExpectedToken: []byte{0x48},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x49}, // == 'I'
			ExpectedToken: []byte{0x49},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x4A}, // == 'J'
			ExpectedToken: []byte{0x4A},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x4B}, // == 'K'
			ExpectedToken: []byte{0x4B},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x4C}, // == 'L'
			ExpectedToken: []byte{0x4C},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x4D}, // == 'M'
			ExpectedToken: []byte{0x4D},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x4E}, // == 'N'
			ExpectedToken: []byte{0x4E},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x4F}, // == 'O'
			ExpectedToken: []byte{0x4F},
			ExpectedRest: nil,
			ExpectedOK: true,
		},


		{
			Value:         []byte{0x50}, // == 'P'
			ExpectedToken: []byte{0x50},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x51}, // == 'Q'
			ExpectedToken: []byte{0x51},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x52}, // == 'R'
			ExpectedToken: []byte{0x52},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x53}, // == 'S'
			ExpectedToken: []byte{0x53},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x54}, // == 'T'
			ExpectedToken: []byte{0x54},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x55}, // == 'U'
			ExpectedToken: []byte{0x55},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x56}, // == 'V'
			ExpectedToken: []byte{0x56},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x57}, // == 'W'
			ExpectedToken: []byte{0x57},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x58}, // == 'X'
			ExpectedToken: []byte{0x58},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x59}, // == 'Y'
			ExpectedToken: []byte{0x59},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x5A}, // == 'Z'
			ExpectedToken: []byte{0x5A},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:        []byte{0x5B},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x5B},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x5C},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x5C},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x5D},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x5D},
			ExpectedOK: false,
		},
		{
			Value:         []byte{0x5E}, // == '^'
			ExpectedToken: []byte{0x5E},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x5F}, // == '_'
			ExpectedToken: []byte{0x5F},
			ExpectedRest: nil,
			ExpectedOK: true,
		},


		{
			Value:         []byte{0x60}, // == '`'
			ExpectedToken: []byte{0x60},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x61}, // == 'a'
			ExpectedToken: []byte{0x61},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x62}, // == 'b'
			ExpectedToken: []byte{0x62},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x63}, // == 'c'
			ExpectedToken: []byte{0x63},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x64}, // == 'd'
			ExpectedToken: []byte{0x64},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x65}, // == 'e'
			ExpectedToken: []byte{0x65},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x66}, // == 'f'
			ExpectedToken: []byte{0x66},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x67}, // == 'g'
			ExpectedToken: []byte{0x67},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x68}, // == 'h'
			ExpectedToken: []byte{0x68},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x69}, // == 'i'
			ExpectedToken: []byte{0x69},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x6A}, // == 'j'
			ExpectedToken: []byte{0x6A},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x6B}, // == 'k'
			ExpectedToken: []byte{0x6B},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x6C}, // == 'l'
			ExpectedToken: []byte{0x6C},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x6D}, // == 'm'
			ExpectedToken: []byte{0x6D},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x6E}, // == 'n'
			ExpectedToken: []byte{0x6E},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x6F}, // == 'o'
			ExpectedToken: []byte{0x6F},
			ExpectedRest: nil,
			ExpectedOK: true,
		},


		{
			Value:         []byte{0x70}, // == 'p'
			ExpectedToken: []byte{0x70},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x71}, // == 'q'
			ExpectedToken: []byte{0x71},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x72}, // == 'r'
			ExpectedToken: []byte{0x72},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x73}, // == 's'
			ExpectedToken: []byte{0x73},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x74}, // == 't'
			ExpectedToken: []byte{0x74},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x75}, // == 'u'
			ExpectedToken: []byte{0x75},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x76}, // == 'v'
			ExpectedToken: []byte{0x76},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x77}, // == 'w'
			ExpectedToken: []byte{0x77},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x78}, // == 'x'
			ExpectedToken: []byte{0x78},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x79}, // == 'y'
			ExpectedToken: []byte{0x79},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:         []byte{0x7A}, // == 'z'
			ExpectedToken: []byte{0x7A},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:        []byte{0x7B},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x7B},
			ExpectedOK: false,
		},
		{
			Value:         []byte{0x7C}, // == '|'
			ExpectedToken: []byte{0x7C},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:        []byte{0x7D},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x7D},
			ExpectedOK: false,
		},
		{
			Value:         []byte{0x7E}, // == '~'
			ExpectedToken: []byte{0x7E},
			ExpectedRest: nil,
			ExpectedOK: true,
		},
		{
			Value:        []byte{0x7F},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x7F},
			ExpectedOK: false,
		},


		{
			Value:        []byte{0x80},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x80},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x81},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x81},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x82},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x82},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x83},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x83},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x84},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x84},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x85},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x85},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x86},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x86},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x87},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x87},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x88},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x88},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x89},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x89},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x8A},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x8A},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x8B},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x8B},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x8C},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x8C},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x8D},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x8D},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x8E},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x8E},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x8F},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x8F},
			ExpectedOK: false,
		},


		{
			Value:        []byte{0x90},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x90},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x91},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x91},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x92},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x92},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x93},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x93},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x94},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x94},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x95},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x95},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x96},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x96},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x97},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x97},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x98},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x98},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x99},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x99},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x9A},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x9A},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x9B},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x9B},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x9C},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x9C},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x9D},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x9D},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x9E},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x9E},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0x9F},
			ExpectedToken: nil,
			ExpectedRest: []byte{0x9F},
			ExpectedOK: false,
		},


		{
			Value:        []byte{0xA0},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xA0},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xA1},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xA1},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xA2},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xA2},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xA3},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xA3},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xA4},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xA4},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xA5},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xA5},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xA6},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xA6},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xA7},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xA7},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xA8},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xA8},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xA9},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xA9},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xAA},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xAA},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xAB},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xAB},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xAC},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xAC},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xAD},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xAD},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xAE},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xAE},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xAF},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xAF},
			ExpectedOK: false,
		},


		{
			Value:        []byte{0xB0},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xB0},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xB1},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xB1},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xB2},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xB2},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xB3},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xB3},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xB4},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xB4},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xB5},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xB5},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xB6},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xB6},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xB7},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xB7},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xB8},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xB8},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xB9},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xB9},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xBA},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xBA},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xBB},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xBB},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xBC},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xBC},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xBD},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xBD},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xBE},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xBE},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xBF},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xBF},
			ExpectedOK: false,
		},


		{
			Value:        []byte{0xC0},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xC0},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xC1},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xC1},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xC2},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xC2},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xC3},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xC3},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xC4},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xC4},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xC5},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xC5},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xC6},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xC6},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xC7},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xC7},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xC8},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xC8},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xC9},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xC9},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xCA},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xCA},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xCB},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xCB},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xCC},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xCC},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xCD},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xCD},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xCE},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xCE},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xCF},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xCF},
			ExpectedOK: false,
		},


		{
			Value:        []byte{0xD0},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xD0},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xD1},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xD1},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xD2},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xD2},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xD3},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xD3},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xD4},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xD4},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xD5},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xD5},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xD6},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xD6},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xD7},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xD7},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xD8},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xD8},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xD9},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xD9},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xDA},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xDA},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xDB},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xDB},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xDC},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xDC},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xDD},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xDD},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xDE},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xDE},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xDF},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xDF},
			ExpectedOK: false,
		},


		{
			Value:        []byte{0xE0},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xE0},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xE1},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xE1},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xE2},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xE2},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xE3},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xE3},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xE4},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xE4},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xE5},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xE5},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xE6},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xE6},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xE7},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xE7},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xE8},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xE8},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xE9},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xE9},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xEA},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xEA},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xEB},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xEB},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xEC},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xEC},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xED},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xED},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xEE},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xEE},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xEF},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xEF},
			ExpectedOK: false,
		},


		{
			Value:        []byte{0xF0},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xF0},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xF1},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xF1},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xF2},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xF2},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xF3},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xF3},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xF4},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xF4},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xF5},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xF5},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xF6},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xF6},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xF7},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xF7},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xF8},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xF8},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xF9},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xF9},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xFA},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xFA},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xFB},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xFB},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xFC},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xFC},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xFD},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xFD},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xFE},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xFE},
			ExpectedOK: false,
		},
		{
			Value:        []byte{0xFF},
			ExpectedToken: nil,
			ExpectedRest: []byte{0xFF},
			ExpectedOK: false,
		},
	}

	for testNumber, test := range tests {

		actualToken, actualRest, actualOK := token.Bytes(test.Value)

		{
			expected := test.ExpectedOK
			actual   := actualOK

			if expected != actual {
				t.Errorf("For test #%d, the actual ok-result is not what was expected.", testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:   %t", actual)
				t.Logf("VALUE: %q", test.Value)
				t.Logf("EXPECTED-TOKEN: %q (%#v)", test.ExpectedToken, test.ExpectedToken)
				t.Logf("EXPECTED-REST: %q (%#v)", test.ExpectedRest, test.ExpectedRest)
				continue
			}
		}

		{
			expected := test.ExpectedToken
			actual   := actualToken

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual token is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("VALUE: %q", test.Value)
				continue
			}
		}

		{
			expected := test.ExpectedRest
			actual   := actualRest

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual rest is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("VALUE: %q", test.Value)
				continue
			}
		}
	}
}
