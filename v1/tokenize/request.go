package tokenize

import "github.com/pci-bridge/sdk-go/pcib"

type Request struct {
	MHID string `json:"mhid"`
	RQID string `json:"rqid"`
	TSID string `json:"tsid,omitempty"`
	// If no TSID is provided, a new session will be created with the session data
	SessionData   SessionData `json:"sessionData,omitempty"`
	CorrelationID string      `json:"correlationId,omitempty"`
	BPID          string      `json:"bpid,omitempty"`
	CardNumber    string      `json:"cardNumber,omitempty"`
	NameOnCard    string      `json:"nameOnCard,omitempty"`
	ExpiryMonth   int32       `json:"expiryMonth,omitempty"`
	ExpiryYear    int32       `json:"expiryYear,omitempty"`
	CVV           string      `json:"cvv,omitempty"`
	TTLSeconds    int32       `json:"ttlSeconds,omitempty"`
}

type Dimension struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type SessionData struct {
	ColorDepth         int       `json:"colorDepth"`
	JavaEnabled        bool      `json:"javaEnabled"`
	JavascriptEnabled  bool      `json:"javascriptEnabled"`
	Language           string    `json:"language"`
	Screen             Dimension `json:"screen"`
	ScreenAvailable    Dimension `json:"screenAvailable"`
	WindowInner        Dimension `json:"windowInner"`
	WindowOuter        Dimension `json:"windowOuter"`
	TimezoneOffsetMins int       `json:"timezoneOffsetMins"`
	Timezone           string    `json:"timezone"`
	UserAgent          string    `json:"userAgent"`
	CookiesEnabled     bool      `json:"cookiesEnabled"`
	IsTouch            bool      `json:"isTouch"`
}

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
	PFP            string       `json:"pfp"`
	BinData        pcib.BinData `json:"binData"`
}
