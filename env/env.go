package env

import "github.com/spf13/viper"

type Config struct {
	HttpPort        string
	HttpHost        string
	UserApiURL      string
	AuthAPIEndPoint string
}

var Settings *Config

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")

	viper.BindEnv("HTTP_PORT")
	viper.BindEnv("HTTP_HOST")
	viper.BindEnv("USER_API_URL")
	viper.BindEnv("SETTINGS_API_URL")
	viper.BindEnv("AUTH_API_END_POINT")
	Settings = &Config{
		HttpPort:        viper.GetString("HTTP_PORT"),
		HttpHost:        viper.GetString("HTTP_HOST"),
		UserApiURL:      viper.GetString("USER_API_URL"),
		AuthAPIEndPoint: viper.GetString("AUTH_API_END_POINT"),
	}
}
