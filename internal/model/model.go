package model

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type Patient struct {
	ID            int
	Name          string
	Email         string
	Address       string
	PhoneNumber   string
	DocumentPhoto string
}
