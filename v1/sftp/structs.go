package sftp

import "github.com/pci-bridge/sdk-go/pcib"

type baseRequest struct {
	MHID string `json:"mhid"`
	RQID string `json:"rqid"`
}

type FileInfo struct {
	Filename string `json:"filename"`
	IsDir    bool   `json:"isDir"`
	FileSize int64  `json:"fileSize"`
	Modified int64  `json:"modified"`
}

type File struct {
	Lines []*FileLine `json:"lines"`
}

type FileLine struct {
	Data           string         `json:"data"`
	InputTokenHVT  string         `json:"inputHvt"`
	InputTokenBPID string         `json:"inputBpid"`
	UpgradeToken   *TokenResponse `json:"upgradeToken"`
}

type TokenResponse struct {
	TokenID        string       `json:"tokenId"`
	Token          string       `json:"token"`
	TokenExpiry    int64        `json:"tokenExpiry"`
	EphemeralToken string       `json:"ephemeralToken"`
	Bin            string       `json:"bin"`
	Last4          string       `json:"last4"`
	CardLength     int32        `json:"cardLength"`
	ExpiryMonth    int32        `json:"expiryMonth"`
	ExpiryYear     int32        `json:"expiryYear"`
	PFP            string       `json:"pfp"`
	MFP            string       `json:"mfp"`
	BinData        pcib.BinData `json:"binData"`
}
