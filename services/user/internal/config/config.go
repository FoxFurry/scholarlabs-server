package config

type Config struct {
	Environment string `env:"ENV"`

	DBHost string `env:"DB_HOST"`
	DBPort string `env:"DB_PORT"`
	DBPass string `env:"DB_PASS"`
	DBUser string `env:"DB_USER"`
	DBName string `env:"DB_NAME"`

	Host       string `env:"HOST"`
	SigningKey string `env:"SIGNING_KEY"`

	MarketplaceServiceBaseURL string `env:"MARKETPLACE_BASE_URL"`
	UserServiceBaseURL        string `env:"USER_BASE_URL"`
	CourseServiceBaseURL      string `env:"COURSE_BASE_URL"`

	TokenSecret string `env:"TOKEN_SECRET"`
}
