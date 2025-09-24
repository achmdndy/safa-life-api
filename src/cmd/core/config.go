package core

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type AppConfig struct {
	AppName    string           `mapstructure:"app_name"`
	Port       int              `mapstructure:"port"`
	DB         DatabaseConfig   `mapstructure:"db"`
	Redis      RedisConfig      `mapstructure:"redis"`
	Server     ServerConfig     `mapstructure:"server"`
	Monitoring MonitoringConfig `mapstructure:"monitoring"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	Port     int    `mapstructure:"port"`
	SSLMode  string `mapstructure:"ssl_mode"`
	TimeZone string `mapstructure:"timezone"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	Env  string `mapstructure:"env"`
}

type MonitoringConfig struct {
	Jaeger JaegerConfig `mapstructure:"jaeger"`
}

type JaegerConfig struct {
	Endpoint       string `mapstructure:"endpoint"`
	ServiceName    string `mapstructure:"service_name"`
	ServiceVersion string `mapstructure:"service_version"`
	Environment    string `mapstructure:"environment"`
}

var Config AppConfig

func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("go.mod not found")
		}
		dir = parent
	}
}

func InitConfig(configFlag string) {
	configFile := "config.yaml"

	// If config flag is provided, use it to determine the config file
	if configFlag != "" {
		configFile = fmt.Sprintf("config.%s.yaml", configFlag)
	}

	projectRoot, err := findProjectRoot()
	if err != nil {
		panic(fmt.Errorf("failed to get working directory: %w", err))
	}

	// Look for config file in the current directory's configs folder
	configPath := filepath.Join(projectRoot, "configs", configFile)

	// Check if the specific config file exists, fallback to default if not
	if _, err := os.Stat(configPath); os.IsNotExist(err) && configFlag != "" {
		fmt.Printf("⚠️  Config file %s not found, falling back to config.yaml\n", configFile)
		configPath = filepath.Join(projectRoot, "configs", "config.yaml")
	}

	// Set environment variable prefix for automatic env binding
	viper.SetEnvPrefix("SAFA")
	viper.AutomaticEnv()

	// Bind environment variables to config keys
	viper.BindEnv("db.host", "SAFA_DB_HOST")
	viper.BindEnv("db.user", "SAFA_DB_USER")
	viper.BindEnv("db.password", "SAFA_DB_PASSWORD")
	viper.BindEnv("db.name", "SAFA_DB_NAME")
	viper.BindEnv("db.port", "SAFA_DB_PORT")
	viper.BindEnv("db.ssl_mode", "SAFA_DB_SSL_MODE")
	viper.BindEnv("db.timezone", "SAFA_DB_TIMEZONE")

	viper.BindEnv("redis.host", "SAFA_REDIS_HOST")
	viper.BindEnv("redis.port", "SAFA_REDIS_PORT")
	viper.BindEnv("redis.password", "SAFA_REDIS_PASSWORD")
	viper.BindEnv("redis.db", "SAFA_REDIS_DB")

	viper.BindEnv("server.host", "SAFA_SERVER_HOST")
	viper.BindEnv("server.port", "SAFA_SERVER_PORT")
	viper.BindEnv("server.env", "SAFA_SERVER_ENV")

	viper.BindEnv("app_name", "SAFA_APP_NAME")
	viper.BindEnv("port", "SAFA_APP_PORT")

	// Bind monitoring environment variables
	viper.BindEnv("monitoring.jaeger.endpoint", "JAEGER_ENDPOINT")
	viper.BindEnv("monitoring.jaeger.service_name", "OTEL_SERVICE_NAME")
	viper.BindEnv("monitoring.jaeger.service_version", "OTEL_SERVICE_VERSION")
	viper.BindEnv("monitoring.jaeger.environment", "SAFA_SERVER_ENV")

	// Bind prometheus environment variables
	viper.BindEnv("monitoring.prometheus.enabled", "PROMETHEUS_ENABLED")
	viper.BindEnv("monitoring.prometheus.metrics_path", "PROMETHEUS_METRICS_PATH")

	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("error loading config file %s: %w", configPath, err))
	}

	fmt.Printf("✅ Loaded config from: %s\n", configPath)

	if err := viper.Unmarshal(&Config); err != nil {
		panic(fmt.Errorf("unable to decode config into struct: %w", err))
	}
}
