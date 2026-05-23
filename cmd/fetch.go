package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/jennshaggy/organAIzedcrime/loader"
	"github.com/spf13/cobra"
)

const atlasURL = "https://raw.githubusercontent.com/mitre-atlas/atlas-navigator-data/main/dist/stix-atlas.json"

func spinner(done chan bool) {
	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	i := 0
	for {
		select {
		case <-done:
			fmt.Printf("\r") // clear spinner line
			return
		default:
			fmt.Printf("\r%s  Pulling latest ATLAS intel...", frames[i%len(frames)])
			i++
			time.Sleep(80 * time.Millisecond)
		}
	}
}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Download the latest ATLAS data from GitHub",
	Run: func(cmd *cobra.Command, args []string) {
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

		path := loader.AtlasDataPath()
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Could not create directory: %v\n", err)
			os.Exit(1)
		}

		out, err := os.Create(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not create file: %v\n", err)
			os.Exit(1)
		}
		defer out.Close()

		done := make(chan bool)
		go spinner(done)

		bytes, err := io.Copy(out, resp.Body)
		done <- true

		if err != nil {
			fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("✓  %.1f KB — ATLAS intel locked and loaded.\n", float64(bytes)/1024)
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}