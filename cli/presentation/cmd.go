package presentation

import (
	"fmt"

	"github.com/spf13/cobra"

	"vita/application"
	"vita/core"
)

var RootCmd = &cobra.Command{Use: "vita"}

func init() {
	RootCmd.AddCommand(CmdAdd)
}

var CmdAdd = &cobra.Command{
	Use:   "add {x}",
	Short: "Add a new item",
	Long:  `Add a new itemm`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		entryType := args[0]

		prompts, err := application.LoadPrompts(entryType)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		insertVal := core.EntryItem{
			ItemType: entryType,
			ItemData: RunAddEntryPrompts(prompts),
		}

		client, err := application.GetDbClient()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		application.InsertEntryItem(client, insertVal)
	},
}
