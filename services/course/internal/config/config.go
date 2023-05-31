package config

type Config struct {
	Environment string `env:"ENV" envDefault:"development"`

	DBHost string `env:"DB_HOST" envDefault:"localhost"`
	DBPort string `env:"DB_PORT" envDefault:"3306"`
	DBPass string `env:"DB_PASS" envDefault:"test"`
	DBUser string `env:"DB_USER" envDefault:"scholarlabs"`
	DBName string `env:"DB_NAME" envDefault:"scholarlabs"`

	Host string `env:"HOST" envDefault:"0.0.0.0:22001"`

	SpacesKey string `env:"SPACES_KEY"`
	SpacesSec string `env:"SPACES_SEC"`

	TokenSecret string `env:"TOKEN_SECRET"`
}
