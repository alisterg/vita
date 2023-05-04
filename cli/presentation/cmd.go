package presentation

import (
	"fmt"
	"time"

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

		insertVal := core.Entry{
			ItemType: entryType,
			ItemData: RunAddEntryPrompts(prompts),
			ItemDate: time.Now().Unix(),
		}

		client, err := application.GetDbClient()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		application.InsertEntry(client, insertVal)
	},
}
