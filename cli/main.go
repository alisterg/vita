package main

import (
	"vita/commands"
)

func main() {
	commands.Execute()

	// example fetching and printing arbitrary data

	// tableName := "lifedata"

	// client, err := application.GetDbClient()
	// if err != nil {
	// 	fmt.Println("Failed to create DynamoDB client")
	// 	return
	// }

	// result, err := application.QueryTable(client, tableName, "type", "book")
	// if err != nil {
	// 	fmt.Println("Failed to query table")
	// 	return
	// }

	// for _, item := range result.Items {
	// 	for k, v := range item {
	// 		fmt.Printf("Result was %v: %v \n", k, v)
	// 	}
	// }
}
