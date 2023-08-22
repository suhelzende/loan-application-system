package model

type Config struct {
	Service Service
	Logger  Logger
}

type Service struct {
	Name        string
	Port        int
	Description string
}

type Logger struct {
	LogFileName string
	LogLevel    string
}
