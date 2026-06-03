package cmd

import (
	"fmt"
	"strings"

	"github.com/jennshaggy/organAIzedcrime/payloads"
	"github.com/spf13/cobra"
)

var payloadsCmd = &cobra.Command{
	Use:   "payloads",
	Short: "Browse validated attack payloads",
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
	Short: "Get full detail on a payload by name",
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
	},
}

func init() {
	payloadsCmd.AddCommand(payloadsListCmd)
	payloadsCmd.AddCommand(payloadsGetCmd)
	rootCmd.AddCommand(payloadsCmd)
}