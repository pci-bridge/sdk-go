package tokenize

import (
	"encoding/json"
	"errors"
	"github.com/pci-bridge/core/validate"
)

type Request struct {
	MHID          string `json:"mhid"`
	RQID          string `json:"rqid"`
	BPID          string `json:"bpid,omitempty"`
	CorrelationID string `json:"correlationId,omitempty"`

	// TSID If no ID is provided, a new session will be created with the session data
	TSID        string      `json:"tsid,omitempty"`
	SessionData SessionData `json:"sessionData,omitempty"`

	CardNumber  string `json:"cardNumber,omitempty"`
	NameOnCard  string `json:"nameOnCard,omitempty"`
	ExpiryMonth int32  `json:"expiryMonth,omitempty"`
	ExpiryYear  int32  `json:"expiryYear,omitempty"`
	CVV         string `json:"cvv,omitempty"`
	TTLSeconds  int32  `json:"ttlSeconds,omitempty"`
}

func (r Request) Validate() error {
	if !validate.MHID(r.MHID, true) {
		return errors.New("invalid MHID provided")
	}

	if !validate.RQID(r.RQID, true) {
		return errors.New("invalid RQID provided")
	}

	if !validate.BPID(r.BPID, true) {
		return errors.New("invalid BPID provided")
	}

	if !validate.CardNumber(r.CardNumber, true) {
		return errors.New("invalid Card Number provided")
	}

	if !validate.NameOnCard(r.NameOnCard, true) {
		return errors.New("invalid Name On Card provided")
	}

	if !validate.Month(r.ExpiryMonth, true) {
		return errors.New("invalid Expiry Month provided")
	}

	if !validate.Year(r.ExpiryYear, true) {
		return errors.New("invalid Expiry Year provided")
	}

	// Optional Fields

	if !validate.TSID(r.TSID, false) {
		return errors.New("invalid TSID provided")
	}

	if !validate.CorrelationID(r.CorrelationID, false) {
		return errors.New("invalid CorrelationID provided")
	}

	if !validate.CVV(r.CVV, false) {
		return errors.New("invalid CVV provided")
	}

	return r.SessionData.Validate()
}

type EphemeralRequest struct {
	MHID string `json:"mhid"`
	CVV  string `json:"cvv,omitempty"`
}

func (r EphemeralRequest) Validate() error {
	if !validate.MHID(r.MHID, true) {
		return errors.New("invalid MHID provided")
	}

	if !validate.CVV(r.CVV, true) {
		return errors.New("invalid CVV provided")
	}

	return nil
}

var (
	ErrInvalidWalletType = errors.New("invalid WalletType provided")
	ErrInvalidWalletData = errors.New("invalid WalletData provided")
)

type WalletType string

//goland:noinspection GoUnusedConst
const (
	WALLET_TYPE_GOOGLE_PAY WalletType = "GOOGLE_PAY"
	WALLET_TYPE_APPLE_PAY  WalletType = "APPLE_PAY"
)

type WalletRequest struct {
	MHID          string `json:"mhid"`
	RQID          string `json:"rqid"`
	BPID          string `json:"bpid"`
	CorrelationID string `json:"correlationId,omitempty"`

	// TSID If no ID is provided, a new session will be created with the session data
	TSID        string      `json:"tsid,omitempty"`
	SessionData SessionData `json:"sessionData,omitempty"`

	WalletData []byte     `json:"walletData,omitempty"`
	WalletType WalletType `json:"walletType,omitempty"`
}

func (r WalletRequest) Validate() error {
	if !validate.MHID(r.MHID, true) {
		return errors.New("invalid MHID provided")
	}

	if !validate.RQID(r.RQID, true) {
		return errors.New("invalid RQID provided")
	}

	if !validate.BPID(r.BPID, true) {
		return errors.New("invalid BPID provided")
	}

	switch r.WalletType {
	case WALLET_TYPE_GOOGLE_PAY:
	case WALLET_TYPE_APPLE_PAY:
	default:
		return ErrInvalidWalletType
	}

	if len(r.WalletData) == 0 || !json.Valid(r.WalletData) {
		return ErrInvalidWalletData
	}

	// Optional Fields

	if !validate.TSID(r.TSID, false) {
		return errors.New("invalid TSID provided")
	}

	if !validate.CorrelationID(r.CorrelationID, false) {
		return errors.New("invalid CorrelationID provided")
	}

	return r.SessionData.Validate()
}
