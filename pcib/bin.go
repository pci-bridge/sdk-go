package pcib

type BinData struct {
	Scheme      string `json:"scheme"`
	Brand       string `json:"brand"` // e.g. Visa Electron
	SubType     string `json:"subType"`
	CardType    string `json:"cardType"`
	FundingType string `json:"fundingType"`
	Issuer      Issuer `json:"issuer"`
	Business    bool   `json:"business"`
}

type Issuer struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	Phone   string `json:"phone"`
	URL     string `json:"url"`
}
