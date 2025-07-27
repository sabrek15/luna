package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sabrek15/luna/internal/ai"
	"github.com/sabrek15/luna/internal/config"
	"github.com/sabrek15/luna/internal/tui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "luna \"<question>\" | <some_command> | luna \"<question about piped Input>\"",
	Short: "Your AI-powered Linux command-Line companion",
	Long: "Luna is CLI tool that uses google's Gemini to answer your questions about Linux, explain command outputs, and debug scripts directly in your terminal",
	Args: cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {

		if err := config.LoadConfig(); err != nil {
			return fmt.Errorf("failed to load configuration: %w", err);
		}

		if config.Cfg.APIKey == "" {
			return errors.New("API Key is not found. Please set it using 'luna config --set-key YOUR-API-KEY");
		}

		question := strings.Join(args, " ");
		if question == "" {
			return errors.New("a qeustion is required");
		}

		stat, _ := os.Stdin.Stat();
		var pipedInput string
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			inputBytes, err := io.ReadAll(os.Stdin);
			if err != nil {
				return fmt.Errorf("failed to read from pipe: %w", err);
			}
			pipedInput = string(inputBytes);
		}
		
		var prompt string;
		if pipedInput != "" {
			prompt = fmt.Sprintf("You are 'luna', a helpful AI assistant for Linux users. A user has piped the following content from their terminal. Please answer their question about it.\n\n--- Piped Content ---\n%s\n--- End Piped Content ---\n\nUser's Question: %s\n\nAnswer:", pipedInput, question);
		} else {
			prompt = fmt.Sprintf("You are 'luna', a helpful AI assistant for Linux users. Please answer the user's question concisely and provide executable commands in markdown code blocks where appropriate.\n\nUser's Question: %s\n\nAnswer:", question);
		}

		fmt.Println("🌙 Luna is thinking...")

		response, err := ai.GenerateContent(config.Cfg.APIKey, prompt);
		if err != nil {
			return fmt.Errorf("failed to get response from Gemini: %w", err);
		}

		tui.RenderMarkdown(response);

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1);
	}
}