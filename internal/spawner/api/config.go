package api

type Config struct {
	Address string `toml:"address"`
}

func NewConfig() *Config {
	return &Config{
		Address: ":8080",
	}
}
