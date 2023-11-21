package db_ris

// StopPlace represents one given result from API like https://apis.deutschebahn.com/db-api-marketplace/apis/ris-stations/v1/stop-places/by-key
type StopPlace struct {
	StationID string   `json:"stationID"`
	EvaNumber string   `json:"evaNumber"`
	Location  Location `json:"position"`
}

// Location represents the Location of the StopPlace
type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

// StopPlaceQueryResult represents the result from the API, containing one or more StopPlaces
type StopPlaceQueryResult struct {
	StopPlaces []StopPlace
}

// Key represents a key given from https://apis.deutschebahn.com/db-api-marketplace/apis/ris-stations/v1/stop-places/{evaNumber}/keys
type Key struct {
	Type string `json:"type"`
	Key  string `json:"key"`
}

// T represents the result of https://apis.deutschebahn.com/db-api-marketplace/apis/ris-stations/v1/stop-places/{evaNumber}/keys
type T struct {
	Keys []Key `json:"keys"`
}
