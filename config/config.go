// Package config ...
package config

// Config ...
type Config struct {
	DbURL string
}

func SetupConfig() Config {
	return Config{
		DbURL: "mongodb://localhost:27017/mongodb",
	}
}
