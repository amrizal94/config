package config

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser string
	DBPass string
	DBHost string
	DBName string
	DBPort int
}

func ReadConfig(namefile string) *Config {
	err := godotenv.Load(namefile)
	if err != nil {
		fmt.Println("Error saat baca env", err.Error())
		return nil
	}
	var res Config
	res.DBUser = os.Getenv("DBUSER")
	res.DBPass = os.Getenv("DBPASS")
	res.DBHost = os.Getenv("DBHOST")
	res.DBName = os.Getenv("DBNAME")
	readData := os.Getenv("DBPORT")
	res.DBPort, err = strconv.Atoi(readData)
	if err != nil {
		fmt.Println("Error saat convert", err.Error())
		return nil
	}

	return &res
}

func ConnectSQL(c Config) *sql.DB {
	// format source username:password@tcp(host:port)/databaseName
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Terjadi error", err.Error())
	}

	return db
}
