package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"vita/application"
	"vita/core"
)

var cmdAdd = &cobra.Command{
	Use:   "add {x}",
	Short: "Add a new item",
	Long:  `Add a new itemm`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		entryType := args[0]

		// TODO: move things into their proper place

		prompts, err := application.GetPrompts(entryType)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		insertVal := core.EntryItem{
			ItemType: entryType,
			ItemData: make(map[string]string),
		}

		for _, rawPrompt := range prompts {
			reader := bufio.NewReader(os.Stdin)

			prompt := fmt.Sprintf("%s", rawPrompt)
			promptFn := color.New(color.FgCyan).SprintFunc()
			fmt.Printf("%s ", promptFn(fmt.Sprintf("%s>", prompt)))

			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}

			input = strings.TrimSuffix(input, "\n")

			insertVal.ItemData[prompt] = input
		}

		fmt.Printf("itemData was %v\n", insertVal)
		client, err := application.GetDbClient()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		application.InsertEntryItem(client, insertVal)
	},
}
