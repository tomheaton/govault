package database

import (
	"fmt"
)

func CreateDatabase(name string) {
	if _, exists := Databases[name]; exists {
		fmt.Println("Error: Database already exists.")
		return
	}
	Databases[name] = Database{Name: name, Tables: make(map[string]Table)}
	fmt.Println("Database created:", name)
}

func CreateTable(databaseName string, tableName string, columns map[string]string) {
	database, exists := Databases[databaseName]
	if !exists {
		fmt.Println("Error: Database does not exist.")
		return
	}
	if _, exists := database.Tables[tableName]; exists {
		fmt.Println("Error: Table already exists.")
		return
	}
	database.Tables[tableName] = Table{Name: tableName, Columns: columns, Rows: make([]map[string]string, 0)}
	Databases[databaseName] = database
	fmt.Println("Table created:", tableName)
}
