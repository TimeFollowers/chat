package config

type Server struct {
	App   App   `mapstructure:"app" json:"app" yaml:"app"`
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}
