package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	host     = "localhost"
	dbType   = "mysql"
	user     = "golang"
	password = "go-pass"
)

func connectDB() *sql.DB {
	// Initialize connection string
	connectionStr := fmt.Sprintf("%s:%s@tcp(%s:3306)/?allowNativePasswords=true", user, password, host)
	// Initialize connection object
	db, err := sql.Open(dbType, connectionStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to database")
	return db
}

func createDB(db *sql.DB, dbName string) {
	strSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", dbName)
	_, err := db.Exec(strSql)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successful database creation")
}

func dropDB(db *sql.DB, dbName string) {
	strSql := fmt.Sprintf("DROP DATABASE %s;", dbName)
	_, err := db.Exec(strSql)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successful database deletion")
}

func selectUseDB(db *sql.DB, dbName string) {
	strSql := fmt.Sprintf("USE %s;", dbName)
	_, err := db.Exec(strSql)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successful database selection")
}

func createTable(db *sql.DB, tableName string) {
	strSql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s ( id integer );", tableName)
	_, err := db.Exec(strSql)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successful table selection")
}

func dropTable(db *sql.DB, tableName string) {
	strSql := fmt.Sprintf("DROP TABLE %s;", tableName)
	_, err := db.Exec(strSql)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successful table deletion")
}

func main() {
	db := connectDB()
	createDB(db, "testdayo")
	selectUseDB(db, "testdayo")
	createTable(db, "tests")
	dropTable(db, "tests")
	dropDB(db, "testdayo")
}
