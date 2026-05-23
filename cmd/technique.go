package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/jennshaggy/organAIzedcrime/loader"
	"github.com/jennshaggy/organAIzedcrime/renderer"
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
		_, techniques, err := loader.ParseBundle(loader.AtlasDataPath())
		if err != nil {
			log.Fatal(err)
		}

		for _, t := range techniques {
			for _, ref := range t.ExternalReferences {
				if ref.SourceName == "mitre-atlas" && ref.ExternalID == query {
					text, refs := renderer.CleanDescription(t.Description)
					fmt.Printf("=== [%s] %s ===\n\n", ref.ExternalID, t.Name)
					fmt.Println(text)
					fmt.Println(renderer.FormatRefs(refs))
					if t.IsSubtechnique {
						fmt.Println("Class: Subtechnique. The scalpel, not the hammer.")
					} else {
						fmt.Println("Class: Technique. The whole enchilada.")
					}
					return
				}
			}
		}
		fmt.Printf("[!] ID not in the matrix. Try atlas search.\n")
	},
}

func init() {
	techniqueCmd.AddCommand(techniqueGetCmd)
	rootCmd.AddCommand(techniqueCmd)
}