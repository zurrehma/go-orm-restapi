package main

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	AppName  string
	IPAddr   string
	Port     int
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: "root",
			Password: "mysql12345",
			AppName:  "todo_app",
			IPAddr:   "127.0.0.1",
			Port:     3306,
		},
	}
}
