package configs

import "github.com/spf13/viper"

type Config map[string]any

func App() map[string]any {
	return Config{
		"environment": viper.GetString("APP_ENV"),
	}
}
