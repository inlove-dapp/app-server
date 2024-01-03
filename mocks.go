package main

// prisma
//go:generate go run github.com/steebchen/prisma-client-go generate

// mocks
//go:generate mockgen -source=prisma/db/db_gen.go -destination=mocks/db_gen.go -package=mocks
