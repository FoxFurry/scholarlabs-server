package config

type Config struct {
	Environment string `env:"ENV" envDefault:"development"`

	Host string `env:"HOST" envDefault:"0.0.0.0:8000"`

	UserServiceBaseURL        string `env:"USER_BASE_URL" envDefault:"localhost:22000"`
	CourseServiceBaseURL      string `env:"COURSE_BASE_URL" envDefault:"127.0.0.1:22001"`
	EnvironmentServiceBaseURL string `env:"ENVIRONMENT_BASE_URL" envDefault:"127.0.0.1:22002"`

	TokenSecret string `env:"TOKEN_SECRET"`
}
