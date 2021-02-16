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

var filters []string

func applyFilters(contexts []string) (ret []string) {
	for _, f := range filters {
		for _, c := range contexts {
			if strings.Contains(c, f) {
				ret = append(ret, c)
			}
		}
	}
	return ret
}

func runCommand(args []string, context string, ch chan<- []string) {
	arguments := append(args, "--context", context)
	// yay passing an array as variadic parameter

	ret, err := exec.Command("kubectl", arguments...).Output()
	if err != nil {
		fmt.Println("Couldn't run command.")
		fmt.Printf("%s", err)
	}
	ch <- append([]string{context + "\n"}, string(ret))
}

var rootCmd = &cobra.Command{
	Use:   "kmux",
	Short: "Runs the same command against multiple kubernetes clusters.",
	Long:  `Runs the same command against multiple kubernetes clusters.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		contexts := applyFilters(getContexts())

		fmt.Println(contexts)

		ch := make(chan []string, len(contexts))
		for _, c := range contexts {
			go runCommand(args, c, ch)
		}

		for _, _ = range contexts {
			fmt.Println(<-ch)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringSliceVar(&filters, "filter", []string{}, "Strings by which to filter contexts")
}
