package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	// sqlDb, err := sql.Open("postgres", "postgres://ptntmqtillzjye:4c40f520658ff4c0ef5e2884c2c5443ea516a1dd508f88c4887bb810ec076b4f@ec2-3-92-119-83.compute-1.amazonaws.com:5432/d7mct601f3o8be")
	sqlDb, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	Db, err = gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDb,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Db.AutoMigrate(&models.User{})

	// createUsersSql := "CREATE TABLE IF NOT EXISTS users (id varchar(10), first_name varchar(10), last_name varchar(10), email varchar(30), password varchar(20))"

	// var createTableList []string = []string{createUsersSql}
	// for _, v := range createTableList {
	// 	res, err := sqlDb.Exec(v)
	// 	if err != nil {
	// 		fmt.Println(err)

	// 	}
	// 	fmt.Println(res)
	// }

}