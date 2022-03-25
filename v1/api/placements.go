package api

type CreatePlacementRequest struct {
	LocationMatch string `json:"locationMatch"`
}
type CreatePlacementResponse struct {
	Token        string `json:"token"`
	Verification string `json:"verification"`
	Expiry       int64  `json:"expiry"`
}

type ListPlacementsResponse struct {
	Placements []ListPlacementResponse
}
type ListPlacementResponse struct {
	Locations []string
}

type RetrievePlacementResponse struct {
	Expiry        int64  `json:"expiry"`
	Created       int64  `json:"created"`
	Status        string `json:"status"`
	LocationMatch string `json:"locationMatch"`
}

type UpdatePlacementRequest struct {
	AddLocations []string
	RemLocations []string
}
