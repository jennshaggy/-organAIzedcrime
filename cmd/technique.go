package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/jennshaggy/organAIzedcrime/loader"
	"github.com/spf13/cobra"
)

var techniqueCmd = &cobra.Command{
	Use:   "technique",
	Short: "Work with a single ATLAS technique",
}

var techniqueGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a technique by ATLAS ID (e.g. AML.T0010)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := strings.ToUpper(args[0])
		_, techniques, err := loader.ParseBundle("ATLAS.json")
		if err != nil {
			log.Fatal(err)
		}

		for _, t := range techniques {
			for _, ref := range t.ExternalReferences {
				if ref.SourceName == "mitre-atlas" && ref.ExternalID == query {
					fmt.Printf("=== [%s] %s ===\n\n", ref.ExternalID, t.Name)
					fmt.Printf("%s\n\n", t.Description)
					if t.IsSubtechnique {
						fmt.Println("Type: Subtechnique")
					} else {
						fmt.Println("Type: Technique")
					}
					return
				}
			}
		}
		fmt.Printf("No technique found with ID: %s\n", query)
	},
}

func init() {
	techniqueCmd.AddCommand(techniqueGetCmd)
	rootCmd.AddCommand(techniqueCmd)
}