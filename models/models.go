package models

type Log struct {
	Timestamp string `json:"timestamp"` // JSON kalitiga mos keluvchi struct maydoni
	Level     string `json:"level"`
	Message   string `json:"message"`
	UserID    string `json:"user_id"`
	Service   string `json:"service_name"`
	Error     string `json:"error"`
}

type Filter struct {
	TimestampFrom string `json:"timestamp_from"`
	TimestampTo   string `json:"timestamp_to"`
	Level         string `json:"level"`
	UserID        string `json:"user_id"`
	Service       string `json:"service_name"`
}

type SuccessResponse struct {
	Status  string
	Message string
}

type UnsuccessResponse struct {
	Status  string
	Message string
	Error   string
}
