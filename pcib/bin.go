package pcib

type BinData struct {
	Scheme        Scheme  `json:"scheme"`
	Brand         string  `json:"brand"` // e.g. Visa Electron
	SubType       string  `json:"subType"`
	CardType      string  `json:"cardType"`
	FundingType   Funding `json:"fundingType"`
	Issuer        Issuer  `json:"issuer"`
	CleanBankName string  `json:"cleanBankName"`

	BinLength  int `json:"binLength"`
	CardLength int `json:"cardLength"`

	ProductName string `json:"productName"`
	ProductCode string `json:"productCode"`
	ComboCard   string `json:"comboCard"`

	Currency string `json:"currency"`

	Business bool `json:"business"`
	Consumer bool `json:"consumer"`

	IsPan   bool `json:"isPan"`
	IsToken bool `json:"isToken"`

	Length               int  `json:"length"`
	PrePaid              bool `json:"prePaid"`
	AccountUpdater       bool `json:"accountUpdater"`
	SupportsTokenization bool `json:"supportsTokenization"`
	Virtual              bool `json:"virtual"`
	Ecommerce            bool `json:"ecommerce"`
	Reloadable           bool `json:"reloadable"`
	BillPayEnabled       bool `json:"billPayEnabled"`
	Regulated            bool `json:"regulated"`
	Alm                  bool `json:"alm"` // Account Level Management
	DomesticOnly         bool `json:"domesticOnly"`
	GamblingBlock        bool `json:"gamblingBlock"`
	Level2               bool `json:"level2"`
	Level3               bool `json:"level3"`
	SharedBin            bool `json:"sharedBin"`

	AuthenticationRequired bool   `json:"authenticationRequired"`
	AuthenticationName     string `json:"authenticationName"` // e.g. PSD2
}

type Issuer struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	Phone   string `json:"phone"`
	URL     string `json:"url"`
}
