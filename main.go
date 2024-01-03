package main

import (
	"log"
	"os"
	"vercel-go/db"
	"vercel-go/router"
)

func main() {
	route := router.Router()
	databaseUrl := os.Getenv("DATABASE_URL")
	db.GetDB()

	log.Printf("Starting server... %s", databaseUrl)

	err := route.Run()
	if err != nil {
		log.Fatal(err)
	}
}
