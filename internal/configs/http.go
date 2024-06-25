package configs

type HTTP struct {
	Address string `yaml:"address"`
}

func (s *HTTP) Run() error {
	return nil
}
