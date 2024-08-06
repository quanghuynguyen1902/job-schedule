package config

import (
	"github.com/spf13/viper"
)

// Loader load config from reader into Viper
type Loader interface {
	Load(viper.Viper) (*viper.Viper, error)
}

type Config struct {
	// db
	Cassandra DBConnection

	Kafka Kafka

	Env string
}

type DBConnection struct {
	Host string
	Port string
	User string
	Name string
	Pass string

	SSLMode string
}

type Kafka struct {
	Broker string
}

type ENV interface {
	GetBool(string) bool
	GetString(string) string
}

func Generate(v ENV) *Config {
	return &Config{
		Env: v.GetString("ENV"),

		Cassandra: DBConnection{
			Host:    v.GetString("DB_HOST"),
			Port:    v.GetString("DB_PORT"),
			User:    v.GetString("DB_USER"),
			Name:    v.GetString("DB_NAME"),
			Pass:    v.GetString("DB_PASS"),
			SSLMode: v.GetString("DB_SSL_MODE"),
		},
		Kafka: Kafka{
			Broker: v.GetString("KAFKA_BROKER"),
		},
	}
}

func DefaultConfigLoaders() []Loader {
	var loaders []Loader
	fileLoader := NewFileLoader(".env", ".")
	loaders = append(loaders, fileLoader)
	loaders = append(loaders, NewENVLoader())

	return loaders
}

// LoadConfig load config from loader list
func LoadConfig(loaders []Loader) *Config {
	v := viper.New()
	v.SetDefault("PORT", "8080")
	v.SetDefault("ENV", "local")
	v.SetDefault("ALLOWED_ORIGINS", "*")

	for idx := range loaders {
		newV, err := loaders[idx].Load(*v)

		if err == nil {
			v = newV
		}
	}
	return Generate(v)
}
