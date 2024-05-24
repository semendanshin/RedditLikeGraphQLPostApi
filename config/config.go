package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env         string   `yaml:"env" env-required:"true"` // dev, test, prod
	Server      Server   `yaml:"server"`
	UseDatabase *bool    `yaml:"use_database" env-required:"false" env-default:"false"`
	Postgres    Postgres `yaml:"postgres"`
}

type Server struct {
	Address string        `yaml:"address" env-required:"true"`
	Timeout time.Duration `yaml:"timeout" env-required:"true"`
}

type Postgres struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	Name string `yaml:"name"`
}

func MustParseConfig(path string) Config {
	var cfg Config

	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}

func FetchPath() string {
	var path string
	flag.StringVar(&path, "config", "", "path to config file")

	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}

	if path == "" {
		path = "config/local.yaml"
	}

	return path
}
