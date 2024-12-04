package tokenize

import (
	"testing"

	"github.com/pci-bridge/sdk-go/validate"
)

func expectValidationResult(t *testing.T, p validate.Validator, expect bool, message string) {
	if err := p.Validate(); (err == nil) != expect {
		t.Errorf(message+", got %s", err)
	}
}
