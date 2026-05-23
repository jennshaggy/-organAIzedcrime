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
		tactics, techniques, err := loader.ParseBundle(loader.AtlasDataPath())
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
					phaseName:   strings.ToLower(strings.ReplaceAll(t.Name, " ", "-")),
				}
				break
			}
		}

		if found == nil {
			fmt.Printf("[!] That tactic doesn't exist in this timeline. Yet.\n")
			return
		}

		fmt.Printf("=== [TACTIC] %s ===\n\n", found.name)
		fmt.Printf("%s\n\n", found.description)
		fmt.Println("--- Tools of the trade ---")

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