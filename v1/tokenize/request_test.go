package tokenize

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func TestTokenizeRequestValidate(t *testing.T) {
	r := &Request{}
	expectValidationResult(t, r, false, "Expected validation to fail, with no required fields")

	r.MHID = "MHID1234"
	expectValidationResult(t, r, false, "Expected validation to fail")

	r.RQID = fmt.Sprintf("%s", sha256.Sum256([]byte(strconv.Itoa(rand.Int()))))
	expectValidationResult(t, r, false, "Expected validation to fail")

	r.BPID = fmt.Sprintf("%x", rand.Int63())
	expectValidationResult(t, r, false, "Expected validation to fail")

	r.CardNumber = "4111111111111111"
	expectValidationResult(t, r, false, "Expected validation to fail")

	r.NameOnCard = "John Smith"
	expectValidationResult(t, r, false, "Expected validation to fail")

	r.ExpiryMonth = 1
	expectValidationResult(t, r, false, "Expected validation to fail")

	r.ExpiryYear = 2021
	expectValidationResult(t, r, true, "Expected validation to pass")

	// optional fields
	r.TSID = "x"
	expectValidationResult(t, r, false, "Expected validation to fail")
	r.TSID = fmt.Sprintf("%s", sha256.Sum256([]byte(strconv.Itoa(rand.Int()))))
	expectValidationResult(t, r, true, "Expected validation to pass")

	r.CorrelationID = "x"
	expectValidationResult(t, r, false, "Expected validation to fail")
	r.CorrelationID = fmt.Sprintf("%x", rand.Int63())
	expectValidationResult(t, r, true, "Expected validation to pass")

	r.CVV = "12"
	expectValidationResult(t, r, false, "Expected validation to fail")
	r.CVV = "123"
	expectValidationResult(t, r, true, "Expected validation to pass")
}

func TestTokenizeWalletRequestValidate(t *testing.T) {
	r := WalletRequest{}
	expectValidationResult(t, r, false, "Expected validation to fail, with no required fields")

	r.MHID = "MHID1234"
	expectValidationResult(t, r, false, "Expected validation to fail")

	r.RQID = fmt.Sprintf("%s", sha256.Sum256([]byte(strconv.Itoa(rand.Int()))))
	expectValidationResult(t, r, false, "Expected validation to fail")

	r.BPID = fmt.Sprintf("%x", rand.Int63())
	expectValidationResult(t, r, false, "Expected validation to fail")

	r.WalletType = ""
	expectValidationResult(t, r, false, "Expected validation to fail")
	r.WalletType = "GOOGLE_PAY"

	r.WalletData = []byte("")
	expectValidationResult(t, r, false, "Expected validation to fail")
	r.WalletData = []byte("{}")

	expectValidationResult(t, r, true, "Expected validation to pass")

	// optional fields
	r.TSID = "x"
	expectValidationResult(t, r, false, "Expected validation to fail")
	r.TSID = fmt.Sprintf("%s", sha256.Sum256([]byte(strconv.Itoa(rand.Int()))))
	expectValidationResult(t, r, true, "Expected validation to pass")

	r.CorrelationID = "x"
	expectValidationResult(t, r, false, "Expected validation to fail")
	r.CorrelationID = fmt.Sprintf("%x", rand.Int63())
	expectValidationResult(t, r, true, "Expected validation to pass")
}

func TestTokenizeSessionDataValidate(t *testing.T) {
	if err := (SessionData{}).Validate(); err != nil {
		t.Errorf("Expected validation to pass, with no required fields, got %s", err)
	}

	d := SessionData{
		Language:  "x",
		Timezone:  "x",
		UserAgent: "x",
	}
	expectValidationResult(t, d, false, "Expected validation to fail")

	d.Language = "EN"
	expectValidationResult(t, d, false, "Expected validation to fail")

	d.Timezone = "America/New_York"
	expectValidationResult(t, d, false, "Expected validation to fail")

	d.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36"
	expectValidationResult(t, d, true, "Expected validation to pass")
}

func TestTokenizeEphemeralRequestValidate(t *testing.T) {
	r := EphemeralRequest{}
	expectValidationResult(t, r, false, "Expected validation to fail, with no required fields")

	r.MHID = "MHID1234"
	expectValidationResult(t, r, false, "Expected validation to fail")

	r.CVV = "12"
	expectValidationResult(t, r, false, "Expected validation to fail")

	r.CVV = "123"
	expectValidationResult(t, r, true, "Expected validation to pass")
}
