package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"vita/core"
	"vita/core/entities"

	"github.com/fatih/color"
)

func PrintCyan(text string) {
	colorFn := color.New(color.FgCyan).SprintFunc()
	fmt.Printf("%s", colorFn(text))
}

func PrintMagenta(text string) {
	colorFn := color.New(color.FgMagenta).SprintFunc()
	fmt.Printf("%s", colorFn(text))
}

func RunAddEntryPrompts(prompts []string) map[string]string {
	itemData := make(map[string]string)

	for _, prompt := range prompts {
		reader := bufio.NewReader(os.Stdin)

		PrintCyan(fmt.Sprintf("%s> ", prompt))

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		input = strings.TrimSuffix(input, "\n")

		itemData[prompt] = input
	}

	return itemData
}

func RunUpdateEntryPrompts(existingData map[string]string, prompts []string) map[string]string {
	itemData := make(map[string]string)

	for _, prompt := range prompts {
		reader := bufio.NewReader(os.Stdin)

		PrintCyan(fmt.Sprintf("%s> ", prompt))

		existingValue, exists := existingData[prompt]
		if exists {
			fmt.Printf("(%s)", existingValue)
		}

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		input = strings.TrimSuffix(input, "\n")
		if len(input) > 0 {
			itemData[prompt] = input
		} else {
			itemData[prompt] = existingValue
		}
	}

	return itemData
}

func PrintEntry(entry entities.Entry) {
	// TODO: instead, sort as defined in entry_types.json
	sorted := core.MapSorter(entry.Data)

	for _, pair := range sorted {
		PrintMagenta(pair.Key)
		fmt.Printf(": %v\n", pair.Value)
	}

	fmt.Println()
}
