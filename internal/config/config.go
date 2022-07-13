package config

import (
	"github.com/spf13/viper"
	"time"
)

const (
	// Default values
	defaultHTTPPort               = "8000"
	defaultHTTPRWTimeout          = 10 * time.Second
	defaultHTTPMaxHeaderMegabytes = 1

	EnvLocal = "local"
	Prod     = "prod"
)

type (
	Config struct {
		HTTP HTTPConfig
		JWT  JWTConfig
	}

	HTTPConfig struct {
		Host               string        `mapstructure:"host"`
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
	}

	JWTConfig struct {
	}
)

// Init the configuration
func Init(configDir, env string) (*Config, error) {
	setDefaults()
	if err := parseConfiguration(configDir, env); err != nil {
		return nil, err
	}

	var cfg Config

	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}
	setFromEnv(&cfg)

	return &cfg, nil
}

// Set defaults values of configuration
func setDefaults() {
	viper.SetDefault("http.port", defaultHTTPPort)
	viper.SetDefault("http.max_header_megabytes", defaultHTTPMaxHeaderMegabytes)
	viper.SetDefault("http.timeouts.read", defaultHTTPRWTimeout)
	viper.SetDefault("http.timeouts.write", defaultHTTPRWTimeout)
}

// parse the config file
func parseConfiguration(dir, env string) error {
	viper.AddConfigPath(dir)
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if env == EnvLocal {
		return nil
	}

	viper.SetConfigName(env)
	return viper.MergeInConfig()
}

// unmarshal data from config file to env
func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}

	return nil
}

// set values from env to structure
func setFromEnv(cfg *Config) {}
