package config

type CommandConfig struct {
	Cmd  string   `yaml:"cmd"`
	Args []string `yaml:"args"`
}

type Config struct {
	In     []*CommandConfig `yaml:"in"`
	Out    []*CommandConfig `yaml:"out"`
	Exclude []string         `yaml:"exclude"`
}
