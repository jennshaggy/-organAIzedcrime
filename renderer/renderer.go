package renderer

import (
	"fmt"
	"regexp"
	"strings"
)

// inlineLink matches [text](/path) and [text][ref] patterns
var inlineLink = regexp.MustCompile(`\[([^\]]+)\]\([^)]+\)`)

// refLink matches [\[N\]][refname] patterns
var refLink = regexp.MustCompile(`\[\\?\[?[^\]]*\\?\]?\]\[[^\]]+\]`)

// refDefinition matches [refname]: URL "optional title" at end of text
var refDefinition = regexp.MustCompile(`(?m)^\[([^\]]+)\]:\s+(https?://\S+)(?:\s+"[^"]*")?$`)

// CleanDescription strips markdown link syntax from a description string
// and returns the cleaned text plus a slice of reference URLs.
func CleanDescription(raw string) (text string, refs []string) {
	// Extract reference definitions first
	matches := refDefinition.FindAllStringSubmatch(raw, -1)
	seen := map[string]bool{}
	for _, m := range matches {
		url := m[2]
		if !seen[url] {
			refs = append(refs, url)
			seen[url] = true
		}
	}

	// Remove reference definition lines
	text = refDefinition.ReplaceAllString(raw, "")

	// Replace [text](/path) with just text
	text = inlineLink.ReplaceAllString(text, "$1")

	// Remove [\[N\]][refname] citation markers
	text = refLink.ReplaceAllString(text, "")

	// Clean up extra blank lines left behind
	for strings.Contains(text, "\n\n\n") {
		text = strings.ReplaceAll(text, "\n\n\n", "\n\n")
	}

	text = strings.TrimSpace(text)
	return
}

// FormatRefs returns a formatted reference block for terminal output
func FormatRefs(refs []string) string {
	if len(refs) == 0 {
		return ""
	}
	var sb strings.Builder
	sb.WriteString("\n--- References ---\n")
	for i, url := range refs {
		fmt.Fprintf(&sb, "  [%d] %s\n", i+1, url)
	}
	return sb.String()
}