package dto

type PromptResponse struct {
	ID           int    `json:"ID"`
	SaveLocation string `json:"saveLocation"`
	Status       string `json:"status"`
}
