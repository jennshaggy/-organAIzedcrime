package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var banner = `
                                                   d8888 8888888                        888                  d8b                        
                                                  d88888   888                          888                  Y8P                        
                                                 d88P888   888                          888                                             
 .d88b.  888d888  .d88b.   8888b.  88888b.      d88P 888   888   88888888  .d88b.   .d88888  .d8888b 888d888 888 88888b.d88b.   .d88b.  
d88""88b 888P"   d88P"88b     "88b 888 "88b    d88P  888   888      d88P  d8P  Y8b d88" 888 d88P"    888P"   888 888 "888 "88b d8P  Y8b 
888  888 888     888  888 .d888888 888  888   d88P   888   888     d88P   88888888 888  888 888      888     888 888  888  888 88888888 
Y88..88P 888     Y88b 888 888  888 888  888  d8888888888   888    d88P    Y8b.     Y88b 888 Y88b.    888     888 888  888  888 Y8b.     
 "Y88P"  888      "Y88888 "Y888888 888  888 d88P     888 8888888 88888888  "Y8888   "Y88888  "Y8888P 888     888 888  888  888  "Y8888  
                      888                                                                                                               
                 Y8b d88P                                                                                                               
                  "Y88P"
`

var rootCmd = &cobra.Command{
	Use:   "atlas",
	Short: "OrganAIzedCrime: Query MITRE ATLAS AI/ML threat data",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println(banner)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}