package models

type Metadata struct {
	DAOData     DAO    `json:"dao"`
	Description string `json:"description"`
	ExternalURL string `json:"external_url,omitempty"`
	Image       string `json:"image"`
	Name        string `json:"name"`
}
