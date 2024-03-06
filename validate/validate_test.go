package validate

import (
	"testing"
)

func TestMHID(t *testing.T) {
	tests := []struct {
		input    string
		strict   bool
		expected bool
	}{
		{"", true, false},
		{"", false, true},
		{"12345678", true, true},
		{"123456789", true, false},
		{"123456789", false, false},
		{"1234567", true, false},
		{"1234567", false, false},
	}

	for _, test := range tests {
		if MHID(test.input, test.strict) != test.expected {
			t.Errorf("MHID(%q, %v) = %v, expected %v", test.input, test.strict, !test.expected, test.expected)
		}
	}
}

func TestRQID(t *testing.T) {
	tests := []struct {
		input  string
		strict bool
		expect bool
	}{
		{"", true, false},
		{"", false, true},
	}

	for _, test := range tests {
		if RQID(test.input, test.strict) != test.expect {
			t.Errorf("RQID(%s, %v) = %v, expected %v", test.input, test.strict, !test.expect, test.expect)
		}
	}
}
func TestCorrelationID(t *testing.T) {
	tests := []struct {
		input  string
		strict bool
		expect bool
	}{
		{"", true, false},
		{"", false, true},
	}

	for _, test := range tests {
		if CorrelationID(test.input, test.strict) != test.expect {
			t.Errorf("CorrelationID(%s, %v) = %v, expected %v", test.input, test.strict, !test.expect, test.expect)
		}
	}
}
func TestTSID(t *testing.T) {
	tests := []struct {
		input  string
		strict bool
		expect bool
	}{
		{"", true, false},
		{"", false, true},
	}

	for _, test := range tests {
		if TSID(test.input, test.strict) != test.expect {
			t.Errorf("TSID(%s, %v) = %v, expected %v", test.input, test.strict, !test.expect, test.expect)
		}
	}
}
func TestMUID(t *testing.T) {
	tests := []struct {
		input  string
		strict bool
		expect bool
	}{
		{"", true, false},
		{"", false, true},
	}

	for _, test := range tests {
		if MUID(test.input, test.strict) != test.expect {
			t.Errorf("MUID(%s, %v) = %v, expected %v", test.input, test.strict, !test.expect, test.expect)
		}
	}
}
func TestBPID(t *testing.T) {
	tests := []struct {
		input  string
		strict bool
		expect bool
	}{
		{"", true, false},
		{"", false, true},
	}

	for _, test := range tests {
		if BPID(test.input, test.strict) != test.expect {
			t.Errorf("BPID(%s, %v) = %v, expected %v", test.input, test.strict, !test.expect, test.expect)
		}
	}
}
func TestECID(t *testing.T) {
	tests := []struct {
		input  string
		strict bool
		expect bool
	}{
		{"", true, false},
		{"", false, true},
	}

	for _, test := range tests {
		if ECID(test.input, test.strict) != test.expect {
			t.Errorf("ECID(%s, %v) = %v, expected %v", test.input, test.strict, !test.expect, test.expect)
		}
	}
}
func TestLVT(t *testing.T) {
	tests := []struct {
		input  string
		strict bool
		expect bool
	}{
		{"", true, false},
		{"", false, true},
	}

	for _, test := range tests {
		if LVT(test.input, test.strict) != test.expect {
			t.Errorf("LVT(%s, %v) = %v, expected %v", test.input, test.strict, !test.expect, test.expect)
		}
	}
}
func TestHVT(t *testing.T) {
	tests := []struct {
		input  string
		strict bool
		expect bool
	}{
		{"", true, false},
		{"", false, true},
	}

	for _, test := range tests {
		if HVT(test.input, test.strict) != test.expect {
			t.Errorf("HVT(%s, %v) = %v, expected %v", test.input, test.strict, !test.expect, test.expect)
		}
	}
}
func TestLanguage(t *testing.T) {
	tests := []struct {
		input  string
		strict bool
		expect bool
	}{
		{"", true, false},
		{"", false, true},
	}

	for _, test := range tests {
		if Language(test.input, test.strict) != test.expect {
			t.Errorf("Language(%s, %v) = %v, expected %v", test.input, test.strict, !test.expect, test.expect)
		}
	}
}
func TestTimezone(t *testing.T) {
	tests := []struct {
		input  string
		strict bool
		expect bool
	}{
		{"", true, false},
		{"", false, true},
	}

	for _, test := range tests {
		if Timezone(test.input, test.strict) != test.expect {
			t.Errorf("Timezone(%s, %v) = %v, expected %v", test.input, test.strict, !test.expect, test.expect)
		}
	}
}
func TestUserAgent(t *testing.T) {
	tests := []struct {
		input  string
		strict bool
		expect bool
	}{
		{"", true, false},
		{"", false, true},
	}

	for _, test := range tests {
		if UserAgent(test.input, test.strict) != test.expect {
			t.Errorf("UserAgent(%s, %v) = %v, expected %v", test.input, test.strict, !test.expect, test.expect)
		}
	}
}
func TestCVV(t *testing.T) {
	tests := []struct {
		input  string
		strict bool
		expect bool
	}{
		{"", true, false},
		{"", false, true},
		{"123", true, true},
		{"1234", true, true},
		{"12345", true, false},
		{"12", true, false},
		{"an", true, false},
		{"abc", true, false},
		{"abcdef", true, false},
	}

	for _, test := range tests {
		if CVV(test.input, test.strict) != test.expect {
			t.Errorf("CVV(%s, %v) = %v, expected %v", test.input, test.strict, !test.expect, test.expect)
		}
	}
}
func TestNameOnCard(t *testing.T) {
	tests := []struct {
		input  string
		strict bool
		expect bool
	}{
		{"", true, false},
		{"", false, true},
	}

	for _, test := range tests {
		if NameOnCard(test.input, test.strict) != test.expect {
			t.Errorf("NameOnCard(%s, %v) = %v, expected %v", test.input, test.strict, !test.expect, test.expect)
		}
	}
}
func TestCardNumber(t *testing.T) {
	tests := []struct {
		input  string
		strict bool
		expect bool
	}{
		{"", true, false},
		{"", false, true},
		{"4111 **** **** 1111", true, false},
		{"4111 **** **** 1111", false, false},
	}

	for _, test := range tests {
		if CardNumber(test.input, test.strict) != test.expect {
			t.Errorf("CardNumber(%s, %v) = %v, expected %v", test.input, test.strict, !test.expect, test.expect)
		}
	}
}

func TestMonth(t *testing.T) {
	tests := []struct {
		input  int32
		strict bool
		expect bool
	}{
		{0, true, false},
		{0, false, true},
	}

	for _, test := range tests {
		if Month(test.input, test.strict) != test.expect {
			t.Errorf("Month(%d, %v) = %v, expected %v", test.input, test.strict, !test.expect, test.expect)
		}
	}
}
func TestYear(t *testing.T) {
	tests := []struct {
		input  int32
		strict bool
		expect bool
	}{
		{0, true, false},
		{0, false, true},
	}

	for _, test := range tests {
		if Year(test.input, test.strict) != test.expect {
			t.Errorf("Year(%d, %v) = %v, expected %v", test.input, test.strict, !test.expect, test.expect)
		}
	}
}
