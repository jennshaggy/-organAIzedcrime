package cmd

import (
	"fmt"
	"log"

	"github.com/jennshaggy/organAIzedcrime/loader"
	"github.com/spf13/cobra"
)

var tacticsCmd = &cobra.Command{
	Use:   "tactics",
	Short: "Work with ATLAS tactics",
}

var tacticsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all ATLAS tactics",
	Run: func(cmd *cobra.Command, args []string) {
		tactics, _, err := loader.ParseBundle(loader.AtlasDataPath())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Cracking open the playbook. %d ways to ruin someone's Tuesday.\n\n", len(tactics))
		for i, t := range tactics {
			fmt.Printf("%2d. %s\n", i+1, t.Name)
		}
	},
}

func init() {
	tacticsCmd.AddCommand(tacticsListCmd)
	rootCmd.AddCommand(tacticsCmd)
}