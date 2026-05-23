package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/jennshaggy/organAIzedcrime/loader"
	"github.com/jennshaggy/organAIzedcrime/renderer"
	"github.com/jennshaggy/organAIzedcrime/tools"
	"github.com/spf13/cobra"
)

var showTools bool

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

					if showTools {
						toolList := tools.Get(query)
						if len(toolList) == 0 {
							fmt.Println("\n--- Tools ---")
							fmt.Println("  [-] No tools mapped for this technique yet.")
						} else {
							fmt.Println("\n--- Tools ---")
							for _, tool := range toolList {
								fmt.Printf("  [%s]\n", tool.Name)
								fmt.Printf("    %s\n", tool.Description)
								fmt.Printf("    $ %s\n\n", tool.Usage)
							}
						}
					}
					return
				}
			}
		}
		fmt.Printf("[!] ID not in the matrix. Try atlas search.\n")
	},
}

func init() {
	techniqueGetCmd.Flags().BoolVarP(&showTools, "tools", "t", false, "Show associated security tools")
	techniqueCmd.AddCommand(techniqueGetCmd)
	rootCmd.AddCommand(techniqueCmd)
}