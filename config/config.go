package config

import (
	"flag"
	"github.com/spf13/viper"
	"os"
	"strings"
	"sync"
)

var (
	cfg  *Config
	once sync.Once
)

type Config struct {
	AllowedOrigins     string `mapstructure:"ALLOWED_ORIGINS"`
	AppName            string `mapstructure:"APP_NAME"`
	AppPort            string `mapstructure:"APP_PORT"`
	DBHost             string `mapstructure:"DB_HOST"`
	DBName             string `mapstructure:"DB_NAME"`
	DBPassword         string `mapstructure:"DB_PASSWORD"`
	DBPort             string `mapstructure:"DB_PORT"`
	DBUsername         string `mapstructure:"DB_USERNAME"`
	JwtSecret          string `mapstructure:"JWT_SECRET"`
	JwtExpiredTime     int    `mapstructure:"JWT_EXPIRED_TIME"`
	SaltPassword       string `mapstructure:"SALT_PASSWORD"`
	ServerReadTimeout  int    `mapstructure:"SERVER_READ_TIMEOUT"`
	ServerWriteTimeout int    `mapstructure:"SERVER_WRITE_TIMEOUT"`
}

func Load() {
	once.Do(func() {

		v := viper.New()
		v.AutomaticEnv()

		v.AddConfigPath(".")
		v.SetConfigType("env")
		v.SetConfigName(".env")

		err := v.ReadInConfig()
		if err != nil {
			panic(err)
		}

		config := new(Config)
		err = v.Unmarshal(config)
		if err != nil {
			panic(err)
		}

		cfg = config
	})
}

func Get() *Config {
	if strings.HasSuffix(os.Args[0], ".test") || flag.Lookup("test.v") != nil {
		return &Config{
			SaltPassword: "salty-password",
		}
	}
	return cfg
}
