package tokenize

import "testing"

func expectValidationResult(t *testing.T, p validator, expect bool, message string) {
	if err := p.Validate(); (err == nil) != expect {
		t.Errorf(message+", got %s", err)
	}
}
