package main

import (
	"log"
	"suppliers/api"
	db "suppliers/db/postgres"
)

func main() {
	store, err := db.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("%+v\n", store)
	server := api.NewAPIServer(":3000", store)
	server.Run()
}
