package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var sampleCmd = &cobra.Command{
	Use:   "hello",
	Short: "Print hello",
	Long:  "Print hello",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello!")
	},
}
