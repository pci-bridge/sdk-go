package pcib

type BinData struct {
	Scheme        Scheme  `json:"scheme"`
	Brand         string  `json:"brand"` // e.g. Visa Electron
	SubType       string  `json:"subType"`
	CardType      string  `json:"cardType"`
	FundingType   Funding `json:"fundingType"`
	Issuer        Issuer  `json:"issuer"`
	CleanBankName string  `json:"cleanBankName"`

	BinLength  int `json:"binLength,omitempty"`
	CardLength int `json:"cardLength,omitempty"`

	ProductName string `json:"productName,omitempty"`
	ProductCode string `json:"productCode,omitempty"`
	ComboCard   string `json:"comboCard,omitempty"`

	Currency string `json:"currency,omitempty"`

	Business bool  `json:"business"`
	Consumer *bool `json:"consumer,omitempty"`

	IsPan   *bool `json:"isPan,omitempty"`
	IsToken *bool `json:"isToken,omitempty"`

	PrePaid              *bool `json:"prePaid,omitempty"`
	AccountUpdater       *bool `json:"accountUpdater,omitempty"`
	SupportsTokenization *bool `json:"supportsTokenization,omitempty"`
	Virtual              *bool `json:"virtual,omitempty"`
	Ecommerce            *bool `json:"ecommerce,omitempty"`
	Reloadable           *bool `json:"reloadable,omitempty"`
	BillPayEnabled       *bool `json:"billPayEnabled,omitempty"`
	Regulated            *bool `json:"regulated,omitempty"`
	Alm                  *bool `json:"alm,omitempty"` // Account Level Management
	DomesticOnly         *bool `json:"domesticOnly,omitempty"`
	GamblingBlock        *bool `json:"gamblingBlock,omitempty"`
	Level2               *bool `json:"level2,omitempty"`
	Level3               *bool `json:"level3,omitempty"`
	SharedBin            *bool `json:"sharedBin,omitempty"`

	AuthenticationRequired *bool   `json:"authenticationRequired,omitempty"`
	AuthenticationName     *string `json:"authenticationName,omitempty"` // e.g. PSD2
}

type Issuer struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	Phone   string `json:"phone"`
	URL     string `json:"url"`
}
