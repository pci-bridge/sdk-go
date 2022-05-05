package tokenize

import "github.com/pci-bridge/sdk-go/pcib"

type Response struct {
	TSID           string       `json:"tsid"`
	TSIDK          string       `json:"tsidk"`
	TokenID        string       `json:"tokenId"`
	Token          string       `json:"token"`
	TokenExpiry    int64        `json:"tokenExpiry"`
	EphemeralToken string       `json:"ephemeralToken"`
	Bin            string       `json:"bin"`
	Last4          string       `json:"last4"`
	ExpiryMonth    int32        `json:"expiryMonth"`
	ExpiryYear     int32        `json:"expiryYear"`
	NameOnCard     int32        `json:"nameOnCard"`
	PFP            string       `json:"pfp"`
	MFP            string       `json:"mfp"`
	BinData        pcib.BinData `json:"binData"`
}
