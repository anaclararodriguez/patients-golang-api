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
	ContactNumber string
	DocumentPhoto string
}

type Config struct {
	MysqlPassword string
	MysqlUser     string
	MysqlHost     string
	MysqlPort     string
	MysqlDb       string
	ServerPort    string
	FromAddress   string
	SmtpServer    string
	SmtpPort      string
	SmtpUser      string
	SmtpPass      string
}
