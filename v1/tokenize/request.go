package tokenize

type Request struct {
	MHID          string `json:"mhid"`
	RQID          string `json:"rqid"`
	BPID          string `json:"bpid,omitempty"`
	CorrelationID string `json:"correlationId,omitempty"`

	// TSID If no ID is provided, a new session will be created with the session data
	TSID        string      `json:"tsid,omitempty"`
	SessionData SessionData `json:"sessionData,omitempty"`

	CardNumber  string `json:"cardNumber,omitempty"`
	NameOnCard  string `json:"nameOnCard,omitempty"`
	ExpiryMonth int32  `json:"expiryMonth,omitempty"`
	ExpiryYear  int32  `json:"expiryYear,omitempty"`
	CVV         string `json:"cvv,omitempty"`
	TTLSeconds  int32  `json:"ttlSeconds,omitempty"`
}

type WalletRequest struct {
	MHID          string `json:"mhid"`
	RQID          string `json:"rqid"`
	BPID          string `json:"bpid"`
	CorrelationID string `json:"correlationId,omitempty"`

	// TSID If no ID is provided, a new session will be created with the session data
	TSID        string      `json:"tsid,omitempty"`
	SessionData SessionData `json:"sessionData,omitempty"`

	WalletData string     `json:"walletData,omitempty"`
	WalletType WalletType `json:"walletType,omitempty"`
}
