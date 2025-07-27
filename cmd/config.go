package cmd

import (
	"fmt"
	"strings"

	"github.com/sabrek15/luna/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	keyToSet string
	modelToSet string
	showConfig	bool
);

var configCmd = &cobra.Command{
	Use: "config",
	Short: "Manage configuration for luna",
	Long: "Set or view configuration values, such as API key and Model",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := config.LoadConfig(); err != nil {
			return fmt.Errorf("failed to load configuration: %w", err)
		}

		actionTaken := false

		// --- Corrected Logic: Handle each setting independently ---
		if keyToSet != "" {
			viper.Set("api_key", keyToSet)
			fmt.Println("✔ API Key updated.")
			actionTaken = true
		}

		if modelToSet != "" {
			viper.Set("model", modelToSet)
			fmt.Println("✔ Model updated.")
			actionTaken = true
		}

		// If we updated any value, save the entire config file once.
		if keyToSet != "" || modelToSet != "" {
			if err := config.SaveConfig(); err != nil {
				return fmt.Errorf("failed to save configuration: %w", err)
			}
		}

		// --- Improved Show Logic ---
		if showConfig {
			fmt.Println("--- Luna Configuration ---")
			
			// API Key display
			apiKey := config.Cfg.APIKey
			if apiKey == "" {
				fmt.Println("api_key: Not set. Use 'luna config --set-key YOUR_KEY'")
			} else {
				if len(apiKey) > 8 {
					maskedKey := apiKey[:4] + strings.Repeat("*", len(apiKey)-8) + apiKey[len(apiKey)-4:]
					fmt.Printf("api_key: %s\n", maskedKey)
				} else {
					fmt.Println("api_key: Set (but too short to mask).")
				}
			}

			// Model display (relies on the default value we set in config.go)
			fmt.Printf("model:   %s\n", config.Cfg.Model)
			fmt.Println("--------------------------")
			actionTaken = true
		}

		// If the user ran `luna config` with no flags, show the help menu.
		if !actionTaken {
			return cmd.Help()
		}
		return nil;
	},
}

func init() {
	rootCmd.AddCommand(configCmd);
	configCmd.Flags().StringVar(&keyToSet, "set-key", "", "set your Google AI Studio API Key");
	configCmd.Flags().StringVar(&modelToSet, "set-model", "", "set you google model for your companion");
	configCmd.Flags().BoolVar(&showConfig, "show", false, "show the current API Key");
}