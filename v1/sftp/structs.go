package sftp

import (
	"github.com/pci-bridge/sdk-go/v1/tokenize"
)

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
	Data           string             `json:"data"`
	InputTokenHVT  string             `json:"inputHvt"`
	InputTokenBPID string             `json:"inputBpid"`
	UpgradeToken   *tokenize.Response `json:"upgradeToken"`
}
