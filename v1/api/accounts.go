package api

import "github.com/pci-bridge/sdk-go/pcib"

type CreateServiceAccountRequest struct {
	Permission pcib.Permission `json:"permission"`
}

type CreateServiceAccountResponse struct {
	Token        string `json:"token"`
	Verification string `json:"verification"`
	Expiry       int64  `json:"expiry"`
}

type ListServiceAccountsResponse struct {
	ServiceAccounts []ListServiceAccountResponse `json:"serviceAccounts"`
}

type ListServiceAccountResponse struct {
	Token      string          `json:"token"`
	Permission pcib.Permission `json:"permission"`
	Expiry     int64           `json:"expiry"`
}

type RetrieveServiceAccountResponse struct {
	Permission pcib.Permission `json:"permission"`
	Expiry     int64           `json:"expiry"`
	Created    int64           `json:"created"`
	Status     string          `json:"status"`
}

type UpdateServiceAccountRequest struct {
	Permission pcib.Permission `json:"permission"`
}
