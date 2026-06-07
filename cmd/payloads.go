package cmd

import (
	"fmt"
	"strings"

	"github.com/jennshaggy/organAIzedcrime/payloads"
	"github.com/jennshaggy/organAIzedcrime/runner"
	"github.com/spf13/cobra"
)

var targetURL string
var endpointPath string
var fieldName string
var ollamaTarget string
var ollamaModel string

var payloadsCmd = &cobra.Command{
	Use:   "payloads",
	Short: "Browse and execute validated attack payloads",
}

var payloadsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all validated payloads across all techniques",
	Run: func(cmd *cobra.Command, args []string) {
		all := payloads.All()
		if len(all) == 0 {
			fmt.Println("[!] No payloads in the library yet.")
			return
		}
		fmt.Printf("=== Payload Library — %d validated payloads ===\n\n", len(all))
		for _, p := range all {
			fmt.Printf("  [%s]\n", p.Name)
			fmt.Printf("    Technique : %s\n", p.TechniqueID)
			if p.ATLASGap {
				fmt.Printf("    [!] ATLAS Gap: %s\n", p.GapNote)
			}
			fmt.Printf("    Validated : %s\n\n", strings.Join(p.ValidatedAgainst, ", "))
		}
	},
}

var payloadsGetCmd = &cobra.Command{
	Use:   "get <name>",
	Short: "Get full detail on a payload by name, optionally execute it",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		p := payloads.GetByName(args[0])
		if p == nil {
			fmt.Printf("[!] Payload '%s' not found. Try atlas payloads list.\n", args[0])
			return
		}
		fmt.Printf("=== [%s] ===\n\n", p.Name)
		fmt.Printf("  Technique : %s\n", p.TechniqueID)
		if p.ATLASGap {
			fmt.Printf("  [!] ATLAS Gap: %s\n", p.GapNote)
		}
		fmt.Printf("  Template  : %s\n", p.Template)
		fmt.Printf("  Usage     : %s\n", p.UsageNote)
		fmt.Printf("  Validated : %s\n", strings.Join(p.ValidatedAgainst, ", "))

		if ollamaTarget != "" {
			if ollamaModel == "" {
				fmt.Println("[!] --ollama requires --model to specify the Ollama model name.")
				return
			}
			fmt.Printf("\n[>] Firing at Ollama: %s (model: %s)\n", ollamaTarget, ollamaModel)
			result, err := runner.RunOllama(ollamaTarget, ollamaModel, p.Template)
			if err != nil {
				fmt.Printf("[!] Execution failed: %s\n", err)
				return
			}
			runner.PrintResult(result)
		} else if targetURL != "" {
			fmt.Printf("\n[>] Firing at %s%s (field: %s)\n", targetURL, endpointPath, fieldName)
			result, err := runner.Run(targetURL, endpointPath, fieldName, p.Template)
			if err != nil {
				fmt.Printf("[!] Execution failed: %s\n", err)
				return
			}
			runner.PrintResult(result)
		}
	},
}

func init() {
	payloadsGetCmd.Flags().StringVarP(&targetURL, "target", "T", "", "Target base URL (e.g. http://10.10.10.5)")
	payloadsGetCmd.Flags().StringVarP(&endpointPath, "endpoint", "e", "/api/chat_stream", "API endpoint path")
	payloadsGetCmd.Flags().StringVarP(&fieldName, "field", "f", "message", "JSON field name for the payload")
	payloadsGetCmd.Flags().StringVarP(&ollamaTarget, "ollama", "o", "", "Ollama base URL (e.g. http://10.10.10.5:11434)")
	payloadsGetCmd.Flags().StringVarP(&ollamaModel, "model", "m", "", "Ollama model name (e.g. challenge:latest)")
	payloadsCmd.AddCommand(payloadsListCmd)
	payloadsCmd.AddCommand(payloadsGetCmd)
	rootCmd.AddCommand(payloadsCmd)
}