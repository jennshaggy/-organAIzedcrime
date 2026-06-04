package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/jennshaggy/organAIzedcrime/loader"
	"github.com/jennshaggy/organAIzedcrime/payloads"
	"github.com/jennshaggy/organAIzedcrime/renderer"
	"github.com/jennshaggy/organAIzedcrime/runner"
	"github.com/jennshaggy/organAIzedcrime/tools"
	"github.com/spf13/cobra"
)

var showTools bool
var showPayloads bool
var probeTarget string
var probeEndpoint string
var probeField string

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
					if showPayloads {
						payloadList := payloads.Get(query)
						if len(payloadList) == 0 {
							fmt.Println("\n--- Payloads ---")
							fmt.Println("  [-] No payloads mapped for this technique yet.")
						} else {
							fmt.Println("\n--- Payloads ---")
							for _, p := range payloadList {
								fmt.Printf("  [%s]\n", p.Name)
								if p.ATLASGap {
									fmt.Printf("    [!] ATLAS Gap: %s\n", p.GapNote)
								}
								fmt.Printf("    Template : %s\n", p.Template)
								fmt.Printf("    Usage    : %s\n", p.UsageNote)
								fmt.Printf("    Validated: %s\n\n", strings.Join(p.ValidatedAgainst, ", "))
							}
						}
					}
					if probeTarget != "" {
						payloadList := payloads.Get(query)
						if len(payloadList) == 0 {
							fmt.Println("\n--- Probe ---")
							fmt.Println("  [-] No payloads mapped for this technique. Nothing to fire.")
						} else {
							fmt.Printf("\n--- Probe: %d payload(s) against %s%s ---\n", len(payloadList), probeTarget, probeEndpoint)
							for _, p := range payloadList {
								fmt.Printf("\n[>] Firing: %s\n", p.Name)
								result, err := runner.Run(probeTarget, probeEndpoint, probeField, p.Template)
								if err != nil {
									fmt.Printf("    [!] Failed: %s\n", err)
									continue
								}
								runner.PrintResult(result)
							}
							fmt.Println("\n--- Probe complete ---")
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
	techniqueGetCmd.Flags().BoolVarP(&showPayloads, "payloads", "p", false, "Show validated attack payloads")
	techniqueGetCmd.Flags().StringVarP(&probeTarget, "probe", "P", "", "Fire all payloads at target base URL (e.g. http://10.10.10.5)")
	techniqueGetCmd.Flags().StringVarP(&probeEndpoint, "endpoint", "e", "/api/chat_stream", "API endpoint path for probe mode")
	techniqueGetCmd.Flags().StringVarP(&probeField, "field", "f", "message", "JSON field name for probe mode")
	techniqueCmd.AddCommand(techniqueGetCmd)
	rootCmd.AddCommand(techniqueCmd)
}