package main

import (
	"autoworkers/internal/api"
	"autoworkers/internal/queue"
	"autoworkers/internal/store"
	"autoworkers/internal/worker"
)

func main(){
	s := store.Constructor()
	q := queue.Constructor()
	w :=worker.Constructor(q, s)
	server := api.Constructor(q,s)
	go worker.Workers(w)
	server.Start()

}