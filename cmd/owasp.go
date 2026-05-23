package cmd

import (
	"fmt"
	"strings"

	"github.com/jennshaggy/organAIzedcrime/owasp"
	"github.com/spf13/cobra"
)

var owaspCmd = &cobra.Command{
	Use:   "owasp",
	Short: "Query OWASP LLM Top 10 2025",
}

var owaspListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all OWASP LLM Top 10 2025 entries",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("=== OWASP LLM Top 10 2025 ===\n")
		for _, e := range owasp.Top10 {
			fmt.Printf("  [%s] %s\n", e.ID, e.Name)
		}
		fmt.Println("\nRun: atlas owasp get <ID> for full detail and ATLAS mappings.")
	},
}

var owaspGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get an OWASP LLM Top 10 entry with ATLAS mappings (e.g. LLM01)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := strings.ToUpper(args[0])
		entry := owasp.Get(id)
		if entry == nil {
			fmt.Printf("[!] %s not found. Run atlas owasp list to see all entries.\n", id)
			return
		}

		fmt.Printf("=== [%s] %s ===\n\n", entry.ID, entry.Name)
		fmt.Printf("%s\n\n", entry.Description)
		fmt.Println("--- ATLAS Technique Mappings ---")
		for _, id := range entry.ATLASIDs {
			fmt.Printf("  atlas technique get %s\n", id)
		}
		fmt.Println("\n[+] Run any mapped technique for full detail. Happy hunting.")
	},
}

func init() {
	owaspCmd.AddCommand(owaspListCmd)
	owaspCmd.AddCommand(owaspGetCmd)
	rootCmd.AddCommand(owaspCmd)
}