package config

import (
	"bytes"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	GRPC  Server  `yaml:"grpc"`
	HTTP  Server  `yaml:"http"`
	Log   Logging `yaml:"log"`
	DB    SQL     `yaml:"db"`
	Redis Redis   `yaml:"redis"`
}

func New(path, envPrefix string) (Server, error) {
	v := viper.New()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	v.SetEnvPrefix(envPrefix)
	v.SetConfigType("yaml")

	data, err := os.ReadFile(path)
	if err != nil {
		return Server{}, err
	}
	if err := v.ReadConfig(bytes.NewBuffer(data)); err != nil {
		return Server{}, err
	}

	s := Server{}
	if err = v.Unmarshal(&s); err != nil {
		return Server{}, err
	}

	return s, nil
}
