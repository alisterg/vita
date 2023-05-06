package presentation

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"vita/core"

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

func RunAddEntryPrompts(prompts []interface{}) map[string]string {
	itemData := make(map[string]string)

	for _, rawPrompt := range prompts {
		reader := bufio.NewReader(os.Stdin)

		prompt := fmt.Sprintf("%s", rawPrompt)
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

func PrintEntry(entry *core.Entry) {
	// TODO: instead, sort as defined in entry_types.json
	sorted := core.MapSorter(entry.Data)

	for _, pair := range sorted {
		PrintMagenta(pair.Key)
		fmt.Printf(": %v\n", pair.Value)
	}

	fmt.Println()
}
