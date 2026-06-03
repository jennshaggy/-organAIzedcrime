package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/jennshaggy/organAIzedcrime/loader"
	"github.com/jennshaggy/organAIzedcrime/payloads"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search <keyword>",
	Short: "Search tactics, techniques, and payloads by keyword",
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
			fmt.Println("  [-] No tactics took the bait.")
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
			fmt.Println("  [-] No techniques surfaced. Recon harder.")
		}

		payloadHits := 0
		fmt.Println("\n=== Payloads ===")
		for _, p := range payloads.All() {
			if strings.Contains(strings.ToLower(p.Name), query) ||
				strings.Contains(strings.ToLower(p.Template), query) ||
				strings.Contains(strings.ToLower(p.UsageNote), query) ||
				strings.Contains(strings.ToLower(p.TechniqueID), query) {
				gap := ""
				if p.ATLASGap {
					gap = " [ATLAS gap]"
				}
				fmt.Printf("  [%s] → %s%s\n", p.Name, p.TechniqueID, gap)
				payloadHits++
			}
		}
		if payloadHits == 0 {
			fmt.Println("  [-] No payloads matched. The library is still growing.")
		}

		fmt.Printf("\n[+] %d tactic(s). %d technique(s). %d payload(s). Happy hunting.\n", tacticHits, techHits, payloadHits)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}