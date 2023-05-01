package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(add)
}

var add = &cobra.Command{
	Use:   "add {x}",
	Short: "Add a new item",
	Long:  `Add a new itemm`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("args: %v", args)
	},
}
