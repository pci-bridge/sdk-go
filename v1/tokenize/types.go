package tokenize

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

type WalletType string

//goland:noinspection GoUnusedConst
const (
	WALLET_TYPE_GOOGLE_PAY WalletType = "GOOGLE_PAY"
	WALLET_TYPE_APPLE_PAY  WalletType = "APPLE_PAY"
)
