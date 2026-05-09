package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/jennshaggy/organAIzedcrime/loader"
	"github.com/spf13/cobra"
)

var tacticCmd = &cobra.Command{
	Use:   "tactic",
	Short: "Work with a single ATLAS tactic",
}

var tacticGetCmd = &cobra.Command{
	Use:   "get <name>",
	Short: "Get a tactic and its techniques",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := strings.ToLower(strings.Join(args, " "))
		tactics, techniques, err := loader.ParseBundle("ATLAS.json")
		if err != nil {
			log.Fatal(err)
		}

		var found *struct {
			name        string
			description string
			phaseName   string
		}

		for _, t := range tactics {
			if strings.ToLower(t.Name) == query {
				found = &struct {
					name        string
					description string
					phaseName   string
				}{
					name:        t.Name,
					description: t.Description,
				}
				// Derive phase name from tactic name for matching techniques
				found.phaseName = strings.ToLower(strings.ReplaceAll(t.Name, " ", "-"))
				break
			}
		}

		if found == nil {
			fmt.Printf("No tactic found matching: %s\n", query)
			return
		}

		fmt.Printf("=== %s ===\n\n", found.name)
		fmt.Printf("%s\n\n", found.description)
		fmt.Println("--- Techniques ---")

		for _, tech := range techniques {
			for _, phase := range tech.KillChainPhases {
				if phase.PhaseName == found.phaseName {
					id := ""
					for _, ref := range tech.ExternalReferences {
						if ref.SourceName == "mitre-atlas" {
							id = ref.ExternalID
						}
					}
					fmt.Printf("  [%s] %s\n", id, tech.Name)
				}
			}
		}
	},
}

func init() {
	tacticCmd.AddCommand(tacticGetCmd)
	rootCmd.AddCommand(tacticCmd)
}