package main

import (
	"GoRest/Config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gofiber/fiber"
)

func main() {
	//connect to the database
	a := Config.DbURL(Config.BuildDBConfig())
	db, err := sql.Open("sqlserver", a)
	if err != nil {
		log.Fatal(err)
	}
	//execute the query for ftp paths
	rows, err := db.Query("select * from Logins")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			Id       int
			Email    string
			Password string
			MailKey  string
			IsActive bool
		)
		if err := rows.Scan(&Id, &Email, &Password, &MailKey, &IsActive); err != nil {
			panic(err)
		}
		fmt.Printf("%s ", Email)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}

	/////////////////////////////
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World 👋!")
	})

	app.Listen(3000)

}
