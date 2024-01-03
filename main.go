package main

import (
	"inlove-app-server/db"
	"inlove-app-server/router"
	"log"
	"os"
)

//go:generate go run github.com/steebchen/prisma-client-go generate
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
