package config

type Config struct {
	Environment string `env:"ENV" envDefault:"development"`

	DBHost string `env:"DB_HOST" envDefault:"localhost"`
	DBPort string `env:"DB_PORT" envDefault:"3306"`
	DBPass string `env:"DB_PASS" envDefault:"test"`
	DBUser string `env:"DB_USER" envDefault:"user"`
	DBName string `env:"DB_NAME" envDefault:"user"`

	Host string `env:"HOST" envDefault:"0.0.0.0:22000"`

	TokenSecret string `env:"TOKEN_SECRET"`
}
