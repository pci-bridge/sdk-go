package pcib

type BinData struct {
	Scheme      Scheme  `json:"scheme"`
	Brand       string  `json:"brand"` // e.g. Visa Electron
	SubType     string  `json:"subType"`
	CardType    string  `json:"cardType"`
	FundingType Funding `json:"fundingType"`
	Issuer      Issuer  `json:"issuer"`

	Business bool `json:"business"`
	Consumer bool `json:"consumer"`

	Length        int    `json:"length"`
	PrePaid       bool   `json:"prePaid"`
	Reloadable    bool   `json:"reloadable"`
	NetworkToken  bool   `json:"networkToken"`
	Refreshable   bool   `json:"refreshable"`
	Alm           bool   `json:"alm"` // Account Level Management
	DomesticOnly  bool   `json:"domesticOnly"`
	GamblingBlock bool   `json:"gamblingBlock"`
	Level2        bool   `json:"level2"`
	Level3        bool   `json:"level3"`
	Currency      string `json:"currency"`
	SharedBin     bool   `json:"sharedBin"`

	AuthenticationRequired bool   `json:"authenticationRequired"`
	AuthenticationName     string `json:"authenticationName"` // e.g. PSD2
}

type Issuer struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	Phone   string `json:"phone"`
	URL     string `json:"url"`
}
