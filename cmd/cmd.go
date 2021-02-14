package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strings"
)

func getContexts() []string {
	contexts, err := exec.Command("kubectl", "config", "get-contexts", "--no-headers=true", "-o", "name").Output()
	if err != nil {
		fmt.Println("Can't get context")
		return nil
	}

	return strings.Fields(string(contexts))
}

var rootCmd = &cobra.Command{
	Use:   "kmux",
	Short: "Runs the same command against multiple kubernetes clusters.",
	Long:  `Runs the same command against multiple kubernetes clusters.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		contexts := getContexts()
		fmt.Println(len(contexts))
		fmt.Println(contexts)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
