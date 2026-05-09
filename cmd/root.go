package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var banner = "\n" +
	`,.  ,           .` + "\n" +
	`                    /  \ |           |         o` + "\n" +
	`,-. ;-. ,-: ,-: ;-. |--| | ,-, ,-. ,-| ,-. ;-. . ;-.-. ,-.` + "\n" +
	`| | |   | | | | | | |  | |  /  |-' | | |   |   | | | | |-'` + "\n" +
	"`-' '   `-| `-` ' ' '  ' ' '-' `-' `-' `-' '   ' ' ' ' `-'" + "\n" +
	"        `-'\n"

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