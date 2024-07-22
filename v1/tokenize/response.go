package tokenize

import "github.com/pci-bridge/sdk-go/pcib"

type GlobalResponse struct {
	TokenID     string       `json:"tokenId"`
	Token       string       `json:"token"`
	TokenExpiry int64        `json:"tokenExpiry"`
	Bin         string       `json:"bin"`
	Last4       string       `json:"last4"`
	CardLength  int32        `json:"cardLength"`
	ExpiryMonth int32        `json:"expiryMonth"`
	ExpiryYear  int32        `json:"expiryYear"`
	PFP         string       `json:"pfp"`
	MFP         string       `json:"mfp"`
	BinData     pcib.BinData `json:"binData"`
}

type Response struct {
	GlobalResponse

	TSID  string `json:"tsid"`
	TSIDK string `json:"tsidk"`

	NameOnCard     string `json:"nameOnCard"`
	EphemeralToken string `json:"ephemeralToken"`

	PFP string `json:"pfp"`
}

type EphemeralResponse struct {
	EphemeralToken string `json:"ephemeralToken"`
}

type WalletResponse struct {
	Token      *GlobalResponse `json:"token"`
	WalletType WalletType      `json:"walletType"`
	WalletData []byte          `json:"walletData"`
}
