package core

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type AppConfig struct {
	App        AppInfo          `mapstructure:"app"`
	DB         DatabaseConfig   `mapstructure:"db"`
	Redis      RedisConfig      `mapstructure:"redis"`
	Monitoring MonitoringConfig `mapstructure:"monitoring"`
	JWT        JWTConfig        `mapstructure:"jwt"`
	Storage    StorageConfig    `mapstructure:"storage"`
}

type AppInfo struct {
	Name string `mapstructure:"name"`
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Env  string `mapstructure:"env"`
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

type MonitoringConfig struct {
	Jaeger     JaegerConfig     `mapstructure:"jaeger"`
	Prometheus PrometheusConfig `mapstructure:"prometheus"`
}

type JaegerConfig struct {
	Endpoint       string `mapstructure:"endpoint"`
	ServiceName    string `mapstructure:"service_name"`
	ServiceVersion string `mapstructure:"service_version"`
	Environment    string `mapstructure:"environment"`
}

type PrometheusConfig struct {
	Enabled     bool   `mapstructure:"enabled"`
	MetricsPath string `mapstructure:"metrics_path"`
}

type JWTConfig struct {
	Secret          string `mapstructure:"secret"`
	AccessTokenTTL  int    `mapstructure:"access_token_ttl"`  // in minutes
	RefreshTokenTTL int    `mapstructure:"refresh_token_ttl"` // in minutes
	Issuer          string `mapstructure:"issuer"`
}

type StorageConfig struct {
	S3 S3Config `mapstructure:"s3"`
}

type S3Config struct {
	Enabled         bool   `mapstructure:"enabled"`
	Bucket          string `mapstructure:"bucket"`
	Region          string `mapstructure:"region"`
	AccessKeyID     string `mapstructure:"access_key_id"`
	SecretAccessKey string `mapstructure:"secret_access_key"`
	Endpoint        string `mapstructure:"endpoint"`
	UsePathStyle    bool   `mapstructure:"use_path_style"`
	PublicURLBase   string `mapstructure:"public_url_base"`
}

var Config AppConfig

func InitConfig(configFlag string) {
	var configPath string

	if _, err := os.Stat("go.mod"); err == nil {
		projectRoot, err := findProjectRoot()
		if err != nil {
			panic(fmt.Errorf("failed to get project root: %w", err))
		}
		configPath = filepath.Join(projectRoot, "configs")
	} else {
		configPath = "/app/configs"
	}

	configFile := "config.yaml"
	if configFlag != "" {
		configFile = fmt.Sprintf("config.%s.yaml", configFlag)
	}

	fullConfigPath := filepath.Join(configPath, configFile)

	if _, err := os.Stat(fullConfigPath); os.IsNotExist(err) && configFlag != "" {
		fmt.Printf("⚠️ Config file %s not found, falling back to config.yaml\n", configFile)
		fullConfigPath = filepath.Join(configPath, "config.yaml")
	}

	viper.SetEnvPrefix("SAFALIFE")
	viper.AutomaticEnv()

	if err := viper.BindEnv("db.host", "SAFALIFE_DB_HOST"); err != nil {
		panic(fmt.Errorf("failed to bind env SAFALIFE_DB_HOST: %w", err))
	}
	if err := viper.BindEnv("db.user", "SAFALIFE_DB_USER"); err != nil {
		panic(fmt.Errorf("failed to bind env SAFALIFE_DB_USER: %w", err))
	}
	if err := viper.BindEnv("db.password", "SAFALIFE_DB_PASSWORD"); err != nil {
		panic(fmt.Errorf("failed to bind env SAFALIFE_DB_PASSWORD: %w", err))
	}
	if err := viper.BindEnv("db.name", "SAFALIFE_DB_NAME"); err != nil {
		panic(fmt.Errorf("failed to bind env SAFALIFE_DB_NAME: %w", err))
	}
	if err := viper.BindEnv("db.port", "SAFALIFE_DB_PORT"); err != nil {
		panic(fmt.Errorf("failed to bind env SAFALIFE_DB_PORT: %w", err))
	}
	if err := viper.BindEnv("db.ssl_mode", "SAFALIFE_DB_SSL_MODE"); err != nil {
		panic(fmt.Errorf("failed to bind env SAFALIFE_DB_SSL_MODE: %w", err))
	}
	if err := viper.BindEnv("db.timezone", "SAFALIFE_DB_TIMEZONE"); err != nil {
		panic(fmt.Errorf("failed to bind env SAFALIFE_DB_TIMEZONE: %w", err))
	}

	if err := viper.BindEnv("redis.host", "SAFALIFE_REDIS_HOST"); err != nil {
		panic(fmt.Errorf("failed to bind env SAFALIFE_REDIS_HOST: %w", err))
	}
	if err := viper.BindEnv("redis.port", "SAFALIFE_REDIS_PORT"); err != nil {
		panic(fmt.Errorf("failed to bind env SAFALIFE_REDIS_PORT: %w", err))
	}
	if err := viper.BindEnv("redis.password", "SAFALIFE_REDIS_PASSWORD"); err != nil {
		panic(fmt.Errorf("failed to bind env SAFALIFE_REDIS_PASSWORD: %w", err))
	}
	if err := viper.BindEnv("redis.db", "SAFALIFE_REDIS_DB"); err != nil {
		panic(fmt.Errorf("failed to bind env SAFALIFE_REDIS_DB: %w", err))
	}

	if err := viper.BindEnv("app.name", "SAFALIFE_APP_NAME"); err != nil {
		panic(fmt.Errorf("failed to bind env SAFALIFE_APP_NAME: %w", err))
	}
	if err := viper.BindEnv("app.host", "SAFALIFE_APP_HOST"); err != nil {
		panic(fmt.Errorf("failed to bind env SAFALIFE_APP_HOST: %w", err))
	}
	if err := viper.BindEnv("app.port", "SAFALIFE_APP_PORT"); err != nil {
		panic(fmt.Errorf("failed to bind env SAFALIFE_APP_PORT: %w", err))
	}
	if err := viper.BindEnv("app.env", "SAFALIFE_APP_ENV"); err != nil {
		panic(fmt.Errorf("failed to bind env SAFALIFE_APP_ENV: %w", err))
	}

	if err := viper.BindEnv("monitoring.jaeger.endpoint", "SAFALIFE_MONITORING_JAEGER_ENDPOINT"); err != nil {
		panic(fmt.Errorf("failed to bind env SAFALIFE_MONITORING_JAEGER_ENDPOINT: %w", err))
	}
	if err := viper.BindEnv("monitoring.jaeger.service_name", "SAFALIFE_MONITORING_JAEGER_SERVICE_NAME"); err != nil {
		panic(fmt.Errorf("failed to bind env SAFALIFE_MONITORING_JAEGER_SERVICE_NAME: %w", err))
	}
	if err := viper.BindEnv("monitoring.jaeger.service_version", "SAFALIFE_MONITORING_JAEGER_SERVICE_VERSION"); err != nil {
		panic(fmt.Errorf("failed to bind env SAFALIFE_MONITORING_JAEGER_SERVICE_VERSION: %w", err))
	}
	if err := viper.BindEnv("monitoring.jaeger.environment", "SAFALIFE_MONITORING_JAEGER_ENVIRONMENT"); err != nil {
		panic(fmt.Errorf("failed to bind env SAFALIFE_MONITORING_JAEGER_ENVIRONMENT: %w", err))
	}

	if err := viper.BindEnv("monitoring.prometheus.enabled", "SAFALIFE_PROMETHEUS_ENABLED"); err != nil {
		panic(fmt.Errorf("failed to bind env SAFALIFE_PROMETHEUS_ENABLED: %w", err))
	}
	if err := viper.BindEnv("monitoring.prometheus.metrics_path", "SAFALIFE_PROMETHEUS_METRICS_PATH"); err != nil {
		panic(fmt.Errorf("failed to bind env SAFALIFE_PROMETHEUS_METRICS_PATH: %w", err))
	}

	// Storage S3 bindings
	if err := viper.BindEnv("storage.s3.enabled", "SAFALIFE_STORAGE_S3_ENABLED"); err != nil {
		panic(fmt.Errorf("failed to bind env SAFALIFE_STORAGE_S3_ENABLED: %w", err))
	}
	if err := viper.BindEnv("storage.s3.bucket", "SAFALIFE_STORAGE_S3_BUCKET"); err != nil {
		panic(fmt.Errorf("failed to bind env SAFALIFE_STORAGE_S3_BUCKET: %w", err))
	}
	if err := viper.BindEnv("storage.s3.region", "SAFALIFE_STORAGE_S3_REGION"); err != nil {
		panic(fmt.Errorf("failed to bind env SAFALIFE_STORAGE_S3_REGION: %w", err))
	}
	if err := viper.BindEnv("storage.s3.access_key_id", "SAFALIFE_STORAGE_S3_ACCESS_KEY_ID"); err != nil {
		panic(fmt.Errorf("failed to bind env SAFALIFE_STORAGE_S3_ACCESS_KEY_ID: %w", err))
	}
	if err := viper.BindEnv("storage.s3.secret_access_key", "SAFALIFE_STORAGE_S3_SECRET_ACCESS_KEY"); err != nil {
		panic(fmt.Errorf("failed to bind env SAFALIFE_STORAGE_S3_SECRET_ACCESS_KEY: %w", err))
	}
	if err := viper.BindEnv("storage.s3.endpoint", "SAFALIFE_STORAGE_S3_ENDPOINT"); err != nil {
		panic(fmt.Errorf("failed to bind env SAFALIFE_STORAGE_S3_ENDPOINT: %w", err))
	}
	if err := viper.BindEnv("storage.s3.use_path_style", "SAFALIFE_STORAGE_S3_USE_PATH_STYLE"); err != nil {
		panic(fmt.Errorf("failed to bind env SAFALIFE_STORAGE_S3_USE_PATH_STYLE: %w", err))
	}
	if err := viper.BindEnv("storage.s3.public_url_base", "SAFALIFE_STORAGE_S3_PUBLIC_URL_BASE"); err != nil {
		panic(fmt.Errorf("failed to bind env SAFALIFE_STORAGE_S3_PUBLIC_URL_BASE: %w", err))
	}

	viper.SetConfigFile(fullConfigPath)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found. Relying on environment variables.")
		} else {
			panic(fmt.Errorf("error loading config file %s: %w", fullConfigPath, err))
		}
	} else {
		fmt.Printf("Loaded config from: %s\n", fullConfigPath)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		panic(fmt.Errorf("unable to decode config into struct: %w", err))
	}
}

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
