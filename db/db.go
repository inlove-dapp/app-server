package db

import (
	"inlove-app-server/prisma/db"
	"log"
	"sync"
)

var (
	client *db.PrismaClient
	once   sync.Once
)

// GetDB method is responsible for establishing a prisma connection and returning a pointer to it.
func GetDB() *db.PrismaClient {
	once.Do(func() {
		client = db.NewClient()
		err := client.Prisma.Connect()
		if err != nil {
			log.Fatal(err)
			//panic(err)
		}
	})
	return client
}
