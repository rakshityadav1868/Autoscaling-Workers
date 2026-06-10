package main

import (
	"autoworkers/internal/job"
	"autoworkers/internal/queue"
	"autoworkers/internal/store"
	"autoworkers/internal/worker"
	"fmt"
	"time"
)

func main(){
	s := store.Constructor()
	q := queue.Constructor()
	w :=worker.Constructor(q, s)
	go worker.Workers(w)
	testjob := &job.Job{
	ID: "job-1",
	Type: "test",
	Payload: "hello world",
	Status: job.Pending,
	}
	store.Create(testjob,s)
	queue.Enqueue(q,testjob)
	
	time.Sleep(2* time.Second)
	fmt.Println(testjob.Status)
	fmt.Println(testjob.Result)

}