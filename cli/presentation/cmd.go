package presentation

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cobra"

	"vita/app"
	"vita/core"
)

var RootCmd = &cobra.Command{Use: "vita"}

func init() {
	RootCmd.AddCommand(CmdAdd)
	RootCmd.AddCommand(CmdFind)

	CmdFind.Flags().Int("num", 0, "number of entries to return")
	CmdFind.Flags().String("search", "", "search string")
}

// USAGE
// v add movie # runs prompts to insert new movie
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
			Uuid:      uuid.New().String(),
			EntryType: entryType,
			Data:      RunAddEntryPrompts(prompts),
			CreatedAt: time.Now().Unix(),
		}

		client, err := app.GetDbClient()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		err2 := app.InsertEntry(client, insertVal)
		if err2 != nil {
			fmt.Printf("Error: %v\n", err2)
		}
	},
}

// USAGE
// v show movie # lists all movies
// v show movie --n 10 # lists latest 10 movies
// v show movie --search "something" # finds movies with "something" in any field
var CmdFind = &cobra.Command{
	Use:   "find",
	Short: "Find items",
	Long:  `Find items`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		entryType := args[0]

		client, err := app.GetDbClient()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		entries, err := app.GetAllEntriesForType(client, entryType)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Printf("Total entries: %d entries\n", len(entries))

		num, _ := cmd.Flags().GetInt("num")
		search, _ := cmd.Flags().GetString("search")

		if len(strings.TrimSpace(search)) > 0 {
			filtered := core.SearchEntries(entries, search)

			if len(filtered) == 0 {
				fmt.Println("No entries found")
				return
			}

			for _, entry := range filtered {
				PrintEntry(entry)
			}

			return
		}

		if num > 0 {
			fmt.Printf("Getting %d entries\n", num)
			entries = entries[:num]
		}

		for _, entry := range entries {
			PrintEntry(entry)
		}
	},
}
