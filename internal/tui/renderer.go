package tui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/glamour"
)

func RenderMarkdown(text string) {
	r, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(100),
	);

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating markdown renderer: %v", err);
		
		fmt.Print(text);
		return
	}

	out, err := r.Render(text);
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error rendering markdown: %v", err);
		
		fmt.Print(text);
		return
	}

	print(out);
}