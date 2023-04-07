package config

type ServerConfiguration struct {
	Port string `mapstructure:"port"`
	Key  string `mapstructure:"key"`
	Mode string `mapstructure:"mode"`
}
