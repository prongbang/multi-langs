package iot

type Config struct {
	Url   string `json:"url"`
	Token string `json:"token"`
}

func (c *Config) Connect() {

}
