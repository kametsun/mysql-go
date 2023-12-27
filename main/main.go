package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type DBConfig struct {
	host     string
	dbType   string
	user     string
	password string
	DB       *sql.DB
	err      error
}

func (dbConfig *DBConfig) connectDB() {
	// Initialize connection string
	connectionStr := fmt.Sprintf("%s:%s@tcp(%s:3306)/?allowNativePasswords=true", dbConfig.user, dbConfig.password, dbConfig.host)
	// Initialize connection object
	dbConfig.DB, dbConfig.err = sql.Open(dbConfig.dbType, connectionStr)
	if dbConfig.err != nil {
		panic(dbConfig.err)
	}
	fmt.Println("Successfully connected to database")
}

func (dbConfig *DBConfig) showDBs() {
	strSql := "SHOW DATABASES;"
	i := 1
	fmt.Printf("SQL: %s\n\n", strSql)
	res, err := dbConfig.DB.Query(strSql)
	if err != nil {
		panic(err)
	}
	for db := ""; res.Next(); {
		res.Scan(&db)
		fmt.Printf("%d: %s\n", i, db)
		i++
	}
}

func (dbConfig *DBConfig) createDB(dbName string) {
	strSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", dbName)
	_, err := dbConfig.DB.Exec(strSql)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successful database creation")
}

func (dbConfig *DBConfig) dropDB(dbName string) {
	strSql := fmt.Sprintf("DROP DATABASE %s;", dbName)
	_, err := dbConfig.DB.Exec(strSql)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successful database deletion")
}

func (dbConfig *DBConfig) selectUseDB(dbName string) {
	strSql := fmt.Sprintf("USE %s;", dbName)
	_, err := dbConfig.DB.Exec(strSql)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successful database selection")
}

func (dbConfig *DBConfig) createTable(tableName string) {
	strSql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s ( id integer );", tableName)
	_, err := dbConfig.DB.Exec(strSql)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successful table creation")
}

func (dbConfig *DBConfig) showTables() {
	i := 1
	strSql := "SHOW TABLES;"
	fmt.Printf("SQL: %s\n\n", strSql)
	res, err := dbConfig.DB.Query(strSql)
	if err != nil {
		log.Fatal(err)
	}
	for table := ""; res.Next(); {
		res.Scan(&table)
		fmt.Printf("%d: %s\n", i, table)
		i++
	}
}

func (dbConfig *DBConfig) dropTable(tableName string) {
	strSql := fmt.Sprintf("DROP TABLE %s;", tableName)
	_, err := dbConfig.DB.Exec(strSql)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successful table deletion")
}
