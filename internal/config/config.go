package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)


type Config struct {
	APIKey string `mapstructure:"api_key"`
	Model string `mapstructure:"model"`
}

var Cfg Config;

func LoadConfig() error {
	viper.SetDefault("model", "gemini-2.5-flash")

	home, err := os.UserHomeDir();
	if err != nil {
		return fmt.Errorf("could not get user home directory: %w", err);
	}

	configPath := filepath.Join(home, ".config", "luna");
	viper.AddConfigPath(configPath);
	viper.SetConfigName("config");
	viper.SetConfigType("yaml");

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if err := os.MkdirAll(configPath, os.ModePerm); err != nil {
				return fmt.Errorf("could not create config directory: %w", err);
			}

			if err := viper.SafeWriteConfig(); err != nil {
				return fmt.Errorf("could not create config file: %w", err);
			}
		} else {
			return fmt.Errorf("error in reading config file: %w", err);
		}
	}

	err = viper.Unmarshal(&Cfg);
	if err != nil {
		return fmt.Errorf("unable to decode into struct: %w", err);
	}

	return nil;
}

func SaveConfig() error {
	home, err := os.UserHomeDir();
	if err != nil {
		return fmt.Errorf("could not get user home directory: %w", err);	
	}
	configPath := filepath.Join(home, ".config", "luna");
	configFile := filepath.Join(configPath, "config.yaml");

	if err := os.MkdirAll(configPath, os.ModePerm); err != nil {
		return fmt.Errorf("could not create config directory: %w", err);
	}

	return viper.WriteConfigAs(configFile);
}