package config

import "github.com/spf13/viper"


type Config interface {
	GetString(string) string
	GetInt(string) int
}

type cfg struct {
	vp *viper.Viper
}

func NewConfig() Config {


}