package presentation

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"vita/app"
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

		prompts, err := app.LoadPrompts(entryType)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		insertVal := core.Entry{
			EntryType: entryType,
			Data:      RunAddEntryPrompts(prompts),
			CreatedAt: time.Now().Unix(),
		}

		client, err := app.GetDbClient()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		app.InsertEntry(client, insertVal)
	},
}
