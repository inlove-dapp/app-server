package main

import (
	"inlove-app-server/db"
	"inlove-app-server/router"
	"log"
)

//go:generate go run github.com/steebchen/prisma-client-go generate
func main() {
	route := router.Router()
	db.GetDB()

	err := route.Run()
	if err != nil {
		log.Fatal(err)
	}
}
