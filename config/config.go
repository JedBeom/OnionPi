package config

type Config struct {
	DB struct {
		User     string `json:"db_name"`
		Password string `json:"db_password"`
		Database string `json:"db_name"`
	}

	Server struct {
		Host           string `json:"server_host"`
		Port           string `json:"server_port"`
		SessCookieName string `json:"server_sess_cookie"`
	}
}
