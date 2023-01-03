package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var nestCmd = &cobra.Command{
	Use:   "nest",
	Short: "nested commnad",
	Long:  "nested command",
}

var nestSampleCmd = &cobra.Command{
	Use:   "bye",
	Short: "Print bye",
	Long:  "Print bye",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Bye!")
	},
}
