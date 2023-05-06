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
	RootCmd.AddCommand(CmdRoutine)

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
// v find movie # lists all movies
// v find movie --n 10 # lists latest 10 movies
// v find movie --search "something" # finds movies with "something" in any field
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

		fmt.Print("Total entries: ")
		PrintCyan(fmt.Sprintf("%d", len(entries)))
		fmt.Println()

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
			fmt.Print("Getting entries: ")
			PrintCyan(fmt.Sprintf("%d", num))
			fmt.Println()

			if len(entries) > num {
				entries = entries[:num]
			}
		}

		for _, entry := range entries {
			PrintEntry(entry)
		}
	},
}

// USAGE
// v routine weekly # runs the 'weekly' routine defined in routines.json
var CmdRoutine = &cobra.Command{
	Use:   "routine",
	Short: "Run a routine of entries",
	Long:  `Run a routine of entries`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		routine := args[0]

		client, err := app.GetDbClient()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		prompts, err := app.LoadPromptsForRoutine(routine)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		entries := make([]core.Entry, 0)

		for _, entryType := range prompts {
			rawEntryType := fmt.Sprintf("%s", entryType)
			prompts, err := app.LoadPrompts(rawEntryType)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}

			entry := core.Entry{
				Uuid:      uuid.New().String(),
				EntryType: rawEntryType,
				Data:      RunAddEntryPrompts(prompts),
				CreatedAt: time.Now().Unix(),
			}
			entries = append(entries, entry)
		}

		app.BulkInsertEntries(client, entries)
		PrintMagenta("Don't forget to add any ad-hoc items!\n")
	},
}
