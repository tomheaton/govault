package database

import (
	"fmt"
)

func InsertIntoTable(databaseName string, tableName string, columnNames []string, values []string) {
	database, exists := Databases[databaseName]
	if !exists {
		fmt.Println("Error: Database does not exist.")
		return
	}

	table, exists := database.Tables[tableName]
	if !exists {
		fmt.Println("Error: Table does not exist.")
		return
	}

	if len(columnNames) != len(values) {
		fmt.Println("Error: Column count does not match value count.")
		return
	}

	row := make(map[string]string)
	for i, columnName := range columnNames {
		row[columnName] = values[i]
	}
	table.Rows = append(table.Rows, row)
	database.Tables[tableName] = table
	Databases[databaseName] = database

	fmt.Println("Data inserted into table:", tableName)
}
