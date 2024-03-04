package tokenize

import (
	"github.com/pci-bridge/sdk-go/validate"
	"testing"
)

func expectValidationResult(t *testing.T, p validate.Validator, expect bool, message string) {
	if err := p.Validate(); (err == nil) != expect {
		t.Errorf(message+", got %s", err)
	}
}
