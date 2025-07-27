package ai

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/sabrek15/luna/internal/config"
	"google.golang.org/api/option"
)


func GenerateContent(apikey, prompt string) (string, error) {
	client, err := genai.NewClient(context.Background(), option.WithAPIKey(apikey));
	if err != nil {
		return "", fmt.Errorf("could not create genai Client: %w", err);
	}
	defer client.Close();

	model := client.GenerativeModel(config.Cfg.Model);
	resp, err := model.GenerateContent(context.Background(), genai.Text(prompt));
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %w", err);
	}

	var responseBuilder strings.Builder;
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				if txt, ok := part.(genai.Text); ok {
					responseBuilder.WriteString(string(txt));
				}
			}
		}
	}

	if responseBuilder.Len() == 0 {
		return "", fmt.Errorf("received empty response from the API")
	}
	return responseBuilder.String(), nil
}
