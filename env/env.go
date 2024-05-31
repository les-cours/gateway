package env

import "github.com/spf13/viper"

type Config struct {
	HttpPort        string
	HttpHost        string
	UserApiURL      string
	LearningApiURL  string
	AuthAPIEndPoint string
	OrgsApiURL      string
	PaymentApiURL   string
}

var Settings *Config

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")

	viper.BindEnv("HTTP_PORT")
	viper.BindEnv("HTTP_HOST")
	viper.BindEnv("USER_API_URL")
	viper.BindEnv("LEARNING_API_URL")
	viper.BindEnv("AUTH_API_END_POINT")
	viper.BindEnv("ORGS_API_URL")
	viper.BindEnv("PAYMENT_API_URL")

	Settings = &Config{
		HttpPort:        viper.GetString("HTTP_PORT"),
		HttpHost:        viper.GetString("HTTP_HOST"),
		UserApiURL:      viper.GetString("USER_API_URL"),
		LearningApiURL:  viper.GetString("LEARNING_API_URL"),
		AuthAPIEndPoint: viper.GetString("AUTH_API_END_POINT"),
		OrgsApiURL:      viper.GetString("ORGS_API_URL"),
		PaymentApiURL:   viper.GetString("PAYMENT_API_URL"),
	}
}
