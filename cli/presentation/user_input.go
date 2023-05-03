package presentation

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func RunAddEntryPrompts(prompts []interface{}) map[string]string {
	itemData := make(map[string]string)

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

		itemData[prompt] = input
	}

	return itemData
}
