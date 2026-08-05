package config

import "github.com/spf13/viper"

type Config struct {
	DBHost string `mapstructure:"MYSQL_HOST"`
	DBPort string `mapstructure:"MYSQL_PORT"`
	DBName string `mapstructure:"MYSQL_DB"`
	DBUser string `mapstructure:"MYSQL_USER"`
	DBPass string `mapstructure:"MYSQL_PASSWORD"`
	UploadPath string `mapstructure:"UPLOAD_PATH"`
	JWTSecret  string `mapstructure:"JWT_SECRET"`
	GoEnv              string `mapstructure:"GO_ENV"`
	AdminSeedPassword  string `mapstructure:"ADMIN_SEED_PASSWORD"`
	Port               string `mapstructure:"PORT"`
	AllowedOrigins     string `mapstructure:"ALLOWED_ORIGINS"`
	GoogleClientID     string `mapstructure:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret string `mapstructure:"GOOGLE_CLIENT_SECRET"`
	GoogleRedirectURL  string `mapstructure:"GOOGLE_REDIRECT_URL"`
	VAPIDPublicKey     string `mapstructure:"VAPID_PUBLIC_KEY"`
	VAPIDPrivateKey    string `mapstructure:"VAPID_PRIVATE_KEY"`
	VAPIDSubject       string `mapstructure:"VAPID_SUBJECT"`
}

func LoadConfig(path string) (Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	// Bind each config key to the environment explicitly. viper.AutomaticEnv
	// makes values readable via viper.Get*, but it does NOT surface through
	// Unmarshal — so without these binds, a deployment that supplies config
	// purely via environment variables (e.g. compose env_file, no app.env file)
	// would unmarshal into an empty struct. Binding fixes that while keeping the
	// app.env file working in development.
	for _, key := range []string{
		"MYSQL_HOST", "MYSQL_PORT", "MYSQL_DB", "MYSQL_USER", "MYSQL_PASSWORD",
		"UPLOAD_PATH", "JWT_SECRET", "GO_ENV", "ADMIN_SEED_PASSWORD",
		"PORT", "ALLOWED_ORIGINS",
		"GOOGLE_CLIENT_ID", "GOOGLE_CLIENT_SECRET", "GOOGLE_REDIRECT_URL",
		"VAPID_PUBLIC_KEY", "VAPID_PRIVATE_KEY", "VAPID_SUBJECT",
	} {
		_ = viper.BindEnv(key)
	}

	// A missing app.env is fine in production, where config is injected via the
	// environment. Only a malformed/unreadable file is fatal.
	if err := viper.ReadInConfig(); err != nil {
		if _, notFound := err.(viper.ConfigFileNotFoundError); !notFound {
			return Config{}, err
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}
