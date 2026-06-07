package runner

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// Result holds the outcome of a payload execution
type Result struct {
	StatusCode int
	Body       string
	Flags      []string
	Streamed   bool
	Mode       string
}

var flagPattern = regexp.MustCompile(`(?i)(THM|FLAG|CTF)\{[^}]+\}`)

// Run fires a payload at a standard HTTP endpoint
func Run(target, endpoint, field, template string) (*Result, error) {
	url := strings.TrimRight(target, "/") + "/" + strings.TrimLeft(endpoint, "/")

	body, err := json.Marshal(map[string]string{field: template})
	if err != nil {
		return nil, fmt.Errorf("failed to build request body: %w", err)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	result := &Result{StatusCode: resp.StatusCode, Mode: "http"}
	contentType := resp.Header.Get("Content-Type")

	if strings.Contains(contentType, "text/event-stream") {
		result.Streamed = true
		result.Body, err = parseSSE(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("SSE parse failed: %w", err)
		}
	} else {
		raw, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response: %w", err)
		}
		result.Body = string(raw)
	}

	result.Flags = flagPattern.FindAllString(result.Body, -1)
	return result, nil
}

// RunOllama fires a payload directly at an Ollama /api/generate endpoint
func RunOllama(target, model, template string) (*Result, error) {
	url := strings.TrimRight(target, "/") + "/api/generate"

	payload := map[string]interface{}{
		"model":  model,
		"prompt": template,
		"stream": false,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to build request body: %w", err)
	}

	client := &http.Client{Timeout: 120 * time.Second}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var ollamaResp map[string]interface{}
	if err := json.Unmarshal(raw, &ollamaResp); err != nil {
		return nil, fmt.Errorf("failed to parse Ollama response: %w", err)
	}

	responseText, _ := ollamaResp["response"].(string)

	result := &Result{
		StatusCode: resp.StatusCode,
		Body:       responseText,
		Mode:       "ollama",
	}
	result.Flags = flagPattern.FindAllString(result.Body, -1)
	return result, nil
}

// parseSSE reads a server-sent event stream and assembles the full response text
func parseSSE(body io.Reader) (string, error) {
	var assembled strings.Builder
	scanner := bufio.NewScanner(body)

	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "data: ") {
			continue
		}
		raw := strings.TrimPrefix(line, "data: ")
		if raw == "[DONE]" {
			break
		}
		var chunk map[string]interface{}
		if err := json.Unmarshal([]byte(raw), &chunk); err != nil {
			continue
		}
		if append, ok := chunk["append"].(string); ok {
			assembled.WriteString(append)
		}
	}
	return assembled.String(), scanner.Err()
}

// PrintResult displays execution output with flag highlighting
func PrintResult(r *Result) {
	fmt.Printf("\n--- Execution Result ---\n")
	fmt.Printf("  Mode    : %s\n", r.Mode)
	fmt.Printf("  Status  : %d\n", r.StatusCode)
	if r.Streamed {
		fmt.Printf("  Stream  : SSE\n")
	}
	fmt.Printf("  Response:\n%s\n", r.Body)
	if len(r.Flags) > 0 {
		fmt.Println("\n  [!] FLAGS DETECTED:")
		for _, f := range r.Flags {
			fmt.Printf("      >> %s\n", f)
		}
	}
}