package api

type CreateEndpointRequest struct {
	CredentialType string            `json:"credentialType"`
	Configuration  map[string]string `json:"configuration"`
}
type CreateEndpointResponse struct {
	CredentialID string `json:"credentialID"`
	AccessToken  string `json:"accessToken"`
}

type ListEndpointsResponse struct {
	Endpoints []ListEndpointResponse `json:"serviceAccounts"`
}

type ListEndpointResponse struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Status string `json:"status"`
}

type RetrieveEndpointResponse struct {
	Type   string
	Status string
}
