package chef 

type Config struct {
	UseFastHttp bool
}

func NewConfig() *Config {
	return &Config{}
}