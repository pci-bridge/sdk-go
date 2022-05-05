package tokenize

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
