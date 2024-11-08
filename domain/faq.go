package domain



type FaqRequest struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}