package config

import (
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type AppConfig struct {
	v         *viper.Viper
	Name      string          `required:"true" yaml:"name"`
	Port      int             `required:"true" yaml:"port"`
	LogLevel  string          `required:"true" yaml:"loglevel"`
	PersonApi PersonApiConfig `required:"true" yaml:"personApi"`
}

type PersonApiConfig struct {
	BaseUrl string        `required:"true" yaml:"baseUrl"`
	Timeout time.Duration `required:"true" yaml:"timeout"`
}

func (c *AppConfig) readAppConfig() {
	v := viper.New()

	v.SetTypeByDefaultValue(true)
	v.SetConfigFile(os.Getenv("CONFIG_FILE_PATH"))
	c.v = v

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(c); err != nil {
		panic(err)
	}
}

func NewConfiguration() *AppConfig {
	applicationConfig := &AppConfig{}
	applicationConfig.readAppConfig()
	applicationConfig.v.WatchConfig()
	applicationConfig.v.OnConfigChange(func(in fsnotify.Event) {
		applicationConfig.readAppConfig()
	})

	return applicationConfig
}
