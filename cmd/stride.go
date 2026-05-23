package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

type strideCategory struct {
	ID         string
	Name       string
	Property   string
	AIManifest string
	ATLASIDs   []string
	OWASPIDs   []string
}

var strideCategories = []strideCategory{
	{
		ID:       "S",
		Name:     "Spoofing",
		Property: "Authenticity",
		AIManifest: "Adversaries impersonate trusted model sources, inject fabricated content into RAG " +
			"pipelines, or deploy lookalike models to deceive downstream systems and users.",
		ATLASIDs: []string{"AML.T0010", "AML.T0010.003", "AML.T0070", "AML.T0052"},
		OWASPIDs: []string{"LLM03", "LLM08"},
	},
	{
		ID:       "T",
		Name:     "Tampering",
		Property: "Integrity",
		AIManifest: "Adversaries poison training data, manipulate model weights, or corrupt vector " +
			"databases to embed backdoors or shift model decision boundaries over time.",
		ATLASIDs: []string{"AML.T0020", "AML.T0018", "AML.T0043", "AML.T0010.002", "AML.T0070"},
		OWASPIDs: []string{"LLM04", "LLM08"},
	},
	{
		ID:       "R",
		Name:     "Repudiation",
		Property: "Non-repudiability",
		AIManifest: "AI systems lack explainability — model decisions cannot be traced, versioning is " +
			"often absent, and audit trails for training data and inference are rarely complete.",
		ATLASIDs: []string{"AML.T0024", "AML.T0025"},
		OWASPIDs: []string{"LLM02"},
	},
	{
		ID:       "I",
		Name:     "Information Disclosure",
		Property: "Confidentiality",
		AIManifest: "Models leak training data, system prompts, credentials, or internal configuration " +
			"through outputs. Competitors reconstruct proprietary models via systematic API queries.",
		ATLASIDs: []string{"AML.T0057", "AML.T0056", "AML.T0024", "AML.T0025", "AML.T0069"},
		OWASPIDs: []string{"LLM02", "LLM07"},
	},
	{
		ID:       "D",
		Name:     "Denial of Service",
		Property: "Availability",
		AIManifest: "Adversaries flood inference APIs with expensive prompts, trigger maximum-length " +
			"responses, or abuse agentic tool-calling loops to exhaust compute budgets without downtime.",
		ATLASIDs: []string{"AML.T0034", "AML.T0029"},
		OWASPIDs: []string{"LLM10"},
	},
	{
		ID:       "E",
		Name:     "Elevation of Privilege",
		Property: "Authorization",
		AIManifest: "Adversaries jailbreak LLMs to bypass content restrictions, escalate to privileged " +
			"tool access, or exploit agentic pipelines to execute unauthorized actions at scale.",
		ATLASIDs: []string{"AML.T0054", "AML.T0051", "AML.T0053", "AML.T0080"},
		OWASPIDs: []string{"LLM01", "LLM06"},
	},
}

func getStrideCategory(id string) *strideCategory {
	for i := range strideCategories {
		if strideCategories[i].ID == id {
			return &strideCategories[i]
		}
	}
	return nil
}

var strideCmd = &cobra.Command{
	Use:   "stride",
	Short: "Map STRIDE threat categories to ATLAS techniques",
}

var strideListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all STRIDE categories with AI manifestations",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("=== STRIDE for AI Systems ===")
		fmt.Println()
		for _, c := range strideCategories {
			fmt.Printf("  [%s] %s — %s\n", c.ID, c.Name, c.Property)
		}
		fmt.Println()
		fmt.Println("Run: atlas stride get <S|T|R|I|D|E> for full detail and ATLAS mappings.")
	},
}

var strideGetCmd = &cobra.Command{
	Use:   "get <category>",
	Short: "Get a STRIDE category with ATLAS and OWASP mappings (e.g. S, T, E)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := strings.ToUpper(args[0])
		cat := getStrideCategory(id)
		if cat == nil {
			fmt.Printf("[!] %s is not a valid STRIDE category. Use S, T, R, I, D, or E.\n", id)
			return
		}

		fmt.Printf("=== [%s] %s ===\n", cat.ID, cat.Name)
		fmt.Printf("Security Property: %s\n\n", cat.Property)
		fmt.Printf("AI Manifestation:\n%s\n\n", cat.AIManifest)

		fmt.Println("--- ATLAS Techniques ---")
		for _, id := range cat.ATLASIDs {
			fmt.Printf("  atlas technique get %s\n", id)
		}

		fmt.Println("\n--- OWASP LLM Top 10 ---")
		for _, id := range cat.OWASPIDs {
			fmt.Printf("  atlas owasp get %s\n", id)
		}

		fmt.Println("\n[+] Run any mapped technique or OWASP entry for full detail. Happy hunting.")
	},
}

func init() {
	strideCmd.AddCommand(strideListCmd)
	strideCmd.AddCommand(strideGetCmd)
	rootCmd.AddCommand(strideCmd)
}