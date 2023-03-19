package config

type Config struct {
	Environment string `env:"ENV"`

	GatewayHost string `env:"GATEWAY_HOST"`
	SigningKey  string `env:"SIGNING_KEY"`

	MarketplaceServiceBaseURL string `env:"MARKETPLACE_BASE_URL"`
	UserServiceBaseURL        string `env:"USER_BASE_URL"`
	CourseServiceBaseURL      string `env:"COURSE_BASE_URL"`
}
