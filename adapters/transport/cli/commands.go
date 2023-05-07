package cli

import (
	"fmt"
	"strings"
	"time"

	"vita/adapters/datasources/dynamo"
	"vita/adapters/datasources/jsonfile"
	"vita/core"
	"vita/core/entities"
	"vita/core/ports"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var entryTypeRepo ports.EntryTypeRepository
var entryRepo ports.EntryRepository
var routineRepo ports.RoutineRepository

var RootCmd = &cobra.Command{Use: "vita"}

func init() {

	// CHANGE THESE IF YOU WISH TO CHANGE THE DATA SOURCE
	entryTypeRepo = jsonfile.EntryTypes{}
	routineRepo = jsonfile.Routines{}
	entryRepo = dynamo.Entries{}

	RootCmd.AddCommand(CmdAdd)
	RootCmd.AddCommand(CmdFind)
	RootCmd.AddCommand(CmdUpdate)
	RootCmd.AddCommand(CmdRoutine)

	CmdFind.Flags().Int("num", 0, "number of entries to return")
	CmdFind.Flags().String("search", "", "search string")
	CmdUpdate.Flags().String("search", "", "search string")
}

// USAGE
// v add movie # runs prompts to insert new movie
var CmdAdd = &cobra.Command{
	Use:   "add {x}",
	Short: "Add a new item",
	Long:  `Add a new itemm`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		entryTypeKey := args[0]

		entryType, err := entryTypeRepo.GetEntryType(entryTypeKey)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		insertVal := entities.Entry{
			Uuid:      uuid.New().String(),
			EntryType: entryTypeKey,
			Data:      RunAddEntryPrompts(entryType.Prompts),
			CreatedAt: time.Now().Unix(),
		}

		err2 := entryRepo.CreateEntry(insertVal)
		if err2 != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Println("Entry created")
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
		entryTypeKey := args[0]

		num, _ := cmd.Flags().GetInt("num")
		search, _ := cmd.Flags().GetString("search")

		entries, err := entryRepo.GetAllEntriesForType(entryTypeKey)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Print("Total entries: ")
		PrintCyan(fmt.Sprintf("%d", len(entries)))
		fmt.Println()

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
// v update movie --search "something" # updates all movies with "something" in any field
// (you need to be explicit)
var CmdUpdate = &cobra.Command{
	Use:   "update",
	Short: "Update a record based on the result of Find",
	Long:  "Update a record based on the result of Find",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		entryTypeKey := args[0]

		search, _ := cmd.Flags().GetString("search")
		if len(strings.TrimSpace(search)) == 0 {
			fmt.Println("You need to provide a search string")
			return
		}

		entryType, err := entryTypeRepo.GetEntryType(entryTypeKey)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		entries, err := entryRepo.GetAllEntriesForType(entryTypeKey)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		filtered := core.SearchEntries(entries, search)

		if len(filtered) == 0 {
			fmt.Println("No entries found")
			return
		}

		for _, entry := range filtered {
			newEntryData := RunUpdateEntryPrompts(entry.Data, entryType.Prompts)
			entry.Data = newEntryData
			err := entryRepo.UpdateEntry(entry)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		}

		fmt.Println("Entry updated")
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
		routineKey := args[0]

		routine, err := routineRepo.GetRoutine(routineKey)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		entries := make([]entities.Entry, 0)

		for _, entryTypeKey := range routine.EntryTypes {
			entryType, err := entryTypeRepo.GetEntryType(entryTypeKey)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}

			entry := entities.Entry{
				Uuid:      uuid.New().String(),
				EntryType: entryTypeKey,
				Data:      RunAddEntryPrompts(entryType.Prompts),
				CreatedAt: time.Now().Unix(),
			}
			entries = append(entries, entry)
		}

		entryRepo.BulkCreateEntries(entries)

		fmt.Print("Entries created:")
		PrintCyan(fmt.Sprintf("%d", len(entries)))
		fmt.Println()

		PrintMagenta("Don't forget to add any ad-hoc items!\n")
	},
}
