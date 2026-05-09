package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

const atlasURL = "https://raw.githubusercontent.com/mitre-atlas/atlas-navigator-data/main/dist/stix-atlas.json"

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Download the latest ATLAS data from GitHub",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fetching latest ATLAS data...")

		resp, err := http.Get(atlasURL)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Request failed: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Fprintf(os.Stderr, "Bad response: %s\n", resp.Status)
			os.Exit(1)
		}

		out, err := os.Create("ATLAS.json")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not create file: %v\n", err)
			os.Exit(1)
		}
		defer out.Close()

		bytes, err := io.Copy(out, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Done. %.1f KB saved to ATLAS.json\n", float64(bytes)/1024)
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}