package database

type Table struct {
	Name    string
	Columns map[string]string
	Rows    []map[string]string
}

type Database struct {
	Name   string
	Tables map[string]Table
}

var Databases = make(map[string]Database)
