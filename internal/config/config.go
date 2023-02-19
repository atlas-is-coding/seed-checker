package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	App struct {
		WalletsDir  string `yaml:"wallets_dir" env-required:"true" env-default:"wallets/"`
		WalletsFile string `yaml:"wallets_file" env-required:"true" env-default:"wallets.txt"`
	} `yaml:"app" env-required:"true"`
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}

		if err := cleanenv.ReadConfig("configs/config.yaml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)

			log.Println(help)
			log.Fatalln(err)
		}
	})

	return instance
}
