package config

type AppConfig struct {
	DBHost    string `env:"DB_HOST" envDefault:"localhost"`
	DBPort    string `env:"DB_PORT" envDefault:"5432"`
	DBUser    string `env:"DB_USER" envDefault:"postgres"`
	DBPass    string `env:"DB_PASS" envDefault:"123456"`
	DBName    string `env:"DB_NAME" envDefault:"postgres"`
	DBSchema  string `env:"DB_SCHEMA" envDefault:"public"`
	LogFormat string `env:"LOG_FORMAT" envDefault:"text"`
	LogLevel  string `env:"LOG_LEVEL" envDefault:"debug"`
	LogOutput string `env:"LOG_OUTPUT" envDefault:"file://logs/metadata.log"`
}
