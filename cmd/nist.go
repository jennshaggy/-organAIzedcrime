package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

type nistFunction struct {
	ID          string
	Name        string
	Description string
	Categories  []nistCategory
}

type nistCategory struct {
	ID          string
	Name        string
	AIRelevance string
	ATLASIDs    []string
	OWASPIDs    []string
}

var nistFunctions = []nistFunction{
	{
		ID:          "GV",
		Name:        "Govern",
		Description: "Establish policies, processes, and accountability for AI risk management.",
		Categories: []nistCategory{
			{
				ID:          "GV.1",
				Name:        "AI Risk Management Policy",
				AIRelevance: "Define acceptable use, risk tolerance, and accountability for AI systems including LLMs and ML pipelines.",
				ATLASIDs:    []string{},
				OWASPIDs:    []string{"LLM03", "LLM06"},
			},
		},
	},
	{
		ID:          "MP",
		Name:        "Map",
		Description: "Identify and categorize AI risks in context.",
		Categories: []nistCategory{
			{
				ID:          "MP.1.1",
				Name:        "AI System Component Identification",
				AIRelevance: "Inventory all AI components: models, datasets, inference servers, vector databases, pipelines, and third-party dependencies.",
				ATLASIDs:    []string{"AML.T0069", "AML.T0014"},
				OWASPIDs:    []string{"LLM03"},
			},
			{
				ID:          "MP.1.5",
				Name:        "Potential Risk Assessment",
				AIRelevance: "Assess attack surface: misconfigurations, exposed registries, supply chain risks, and prompt injection vectors.",
				ATLASIDs:    []string{"AML.T0010", "AML.T0051", "AML.T0054"},
				OWASPIDs:    []string{"LLM01", "LLM03", "LLM04"},
			},
			{
				ID:          "MP.3.2",
				Name:        "Third-Party AI Resource Risks",
				AIRelevance: "Identify risks from external model weights, datasets, plugins, and infrastructure that may be compromised or malicious.",
				ATLASIDs:    []string{"AML.T0010", "AML.T0010.002", "AML.T0010.003"},
				OWASPIDs:    []string{"LLM03"},
			},
		},
	},
	{
		ID:          "MS",
		Name:        "Measure",
		Description: "Analyze and assess AI risks using quantitative and qualitative methods.",
		Categories: []nistCategory{
			{
				ID:          "MS.2.6",
				Name:        "Intended Function Verification",
				AIRelevance: "Detect unexpected exposed endpoints, Prometheus metrics leakage, and debug interfaces that indicate systems not functioning as intended.",
				ATLASIDs:    []string{"AML.T0069", "AML.T0057"},
				OWASPIDs:    []string{"LLM02", "LLM07"},
			},
			{
				ID:          "MS.2.10",
				Name:        "Privacy Risk Measurement",
				AIRelevance: "Measure risk of training data extraction, system prompt leakage, and sensitive information disclosure through model outputs.",
				ATLASIDs:    []string{"AML.T0024", "AML.T0025", "AML.T0057", "AML.T0056"},
				OWASPIDs:    []string{"LLM02", "LLM07"},
			},
		},
	},
	{
		ID:          "MG",
		Name:        "Manage",
		Description: "Prioritize and address AI risks based on assessment findings.",
		Categories: []nistCategory{
			{
				ID:          "MG.2.2",
				Name:        "Incident Response for AI",
				AIRelevance: "Respond to model poisoning, supply chain compromise, prompt injection attacks, and data exfiltration via AI systems.",
				ATLASIDs:    []string{"AML.T0020", "AML.T0010", "AML.T0051", "AML.T0057"},
				OWASPIDs:    []string{"LLM01", "LLM04"},
			},
			{
				ID:          "MG.3.2",
				Name:        "Bias and Drift Monitoring",
				AIRelevance: "Detect model behavior shifts caused by data poisoning or adversarial manipulation of training pipelines.",
				ATLASIDs:    []string{"AML.T0020", "AML.T0018"},
				OWASPIDs:    []string{"LLM04"},
			},
		},
	},
}

func getNistFunction(id string) *nistFunction {
	for i := range nistFunctions {
		if strings.ToUpper(nistFunctions[i].ID) == strings.ToUpper(id) {
			return &nistFunctions[i]
		}
	}
	return nil
}

var nistCmd = &cobra.Command{
	Use:   "nist",
	Short: "Map NIST AI RMF functions to ATLAS techniques",
}

var nistListCmd = &cobra.Command{
	Use:   "list",
	Short: "List NIST AI RMF 1.0 functions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("=== NIST AI Risk Management Framework 1.0 ===")
		fmt.Println()
		for _, f := range nistFunctions {
			fmt.Printf("  [%s] %s\n", f.ID, f.Name)
			fmt.Printf("       %s\n\n", f.Description)
		}
		fmt.Println("Run: atlas nist get <GV|MP|MS|MG> for categories and ATLAS mappings.")
	},
}

var nistGetCmd = &cobra.Command{
	Use:   "get <function>",
	Short: "Get a NIST AI RMF function with ATLAS mappings (e.g. MP, MS)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := strings.ToUpper(args[0])
		fn := getNistFunction(id)
		if fn == nil {
			fmt.Printf("[!] %s not found. Use GV, MP, MS, or MG.\n", id)
			return
		}

		fmt.Printf("=== [%s] %s ===\n", fn.ID, fn.Name)
		fmt.Printf("%s\n\n", fn.Description)

		for _, cat := range fn.Categories {
			fmt.Printf("--- [%s] %s ---\n", cat.ID, cat.Name)
			fmt.Printf("%s\n\n", cat.AIRelevance)

			if len(cat.ATLASIDs) > 0 {
				fmt.Println("  ATLAS Techniques:")
				for _, id := range cat.ATLASIDs {
					fmt.Printf("    atlas technique get %s\n", id)
				}
			}

			if len(cat.OWASPIDs) > 0 {
				fmt.Println("  OWASP LLM Top 10:")
				for _, id := range cat.OWASPIDs {
					fmt.Printf("    atlas owasp get %s\n", id)
				}
			}
			fmt.Println()
		}

		fmt.Println("[+] Run any mapped technique or OWASP entry for full detail. Happy hunting.")
	},
}

func init() {
	nistCmd.AddCommand(nistListCmd)
	nistCmd.AddCommand(nistGetCmd)
	rootCmd.AddCommand(nistCmd)
}