package main

import (
	"autoworkers/internal/api"
	"autoworkers/internal/database"
	"autoworkers/internal/redis"
	"autoworkers/internal/store"
	"autoworkers/internal/worker"
	"fmt"
)

func main(){
	s := store.Constructor()
	d := database.Constructor()
	r := redis.Constructor()
	w :=worker.Constructor(r, s,d)
	fmt.Println(d)
	server := api.Constructor(r,s,d)
	go worker.Workers(w)
	server.Start()

}