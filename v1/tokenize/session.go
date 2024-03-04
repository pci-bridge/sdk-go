package tokenize

import (
	"errors"
	"github.com/pci-bridge/core/validate"
)

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

func (r SessionData) Validate() error {
	if !validate.Language(r.Language, false) {
		return errors.New("invalid language provided")
	}
	if !validate.Timezone(r.Timezone, false) {
		return errors.New("invalid timezone provided")
	}
	if !validate.UserAgent(r.UserAgent, false) {
		return errors.New("invalid userAgent provided")
	}
	return nil
}
