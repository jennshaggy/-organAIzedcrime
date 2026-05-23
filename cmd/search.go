package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/jennshaggy/organAIzedcrime/loader"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search <keyword>",
	Short: "Search tactics and techniques by keyword",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := strings.ToLower(strings.Join(args, " "))
		tactics, techniques, err := loader.ParseBundle(loader.AtlasDataPath())
		if err != nil {
			log.Fatal(err)
		}

		tacticHits := 0
		fmt.Println("=== Tactics ===")
		for _, t := range tactics {
			if strings.Contains(strings.ToLower(t.Name), query) ||
				strings.Contains(strings.ToLower(t.Description), query) {
				id := ""
				for _, ref := range t.ExternalReferences {
					if ref.SourceName == "mitre-atlas" {
						id = ref.ExternalID
					}
				}
				fmt.Printf("  [%s] %s\n", id, t.Name)
				tacticHits++
			}
		}
		if tacticHits == 0 {
			fmt.Println("  No tactics matched.")
		}

		techHits := 0
		fmt.Println("\n=== Techniques ===")
		for _, t := range techniques {
			if strings.Contains(strings.ToLower(t.Name), query) ||
				strings.Contains(strings.ToLower(t.Description), query) {
				id := ""
				for _, ref := range t.ExternalReferences {
					if ref.SourceName == "mitre-atlas" {
						id = ref.ExternalID
					}
				}
				sub := ""
				if t.IsSubtechnique {
					sub = " [subtechnique]"
				}
				fmt.Printf("  [%s] %s%s\n", id, t.Name, sub)
				techHits++
			}
		}
		if techHits == 0 {
			fmt.Println("  No techniques matched.")
		}

		fmt.Printf("\n%d tactic(s) and %d technique(s) matched \"%s\"\n", tacticHits, techHits, query)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}