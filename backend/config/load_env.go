package config

import "github.com/spf13/viper"

type Config struct {
	DBHost string `mapstructure:"MYSQL_HOST"`
	DBPort string `mapstructure:"MYSQL_PORT"`
	DBName string `mapstructure:"MYSQL_DATABASE"`
	DBUser string `mapstructure:"MYSQL_USER"`
	DBPass string `mapstructure:"MYSQL_PASSWORD"`
	UploadPath string `mapstructure:"UPLOAD_PATH"`
}

func LoadConfig(path string) (Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}
