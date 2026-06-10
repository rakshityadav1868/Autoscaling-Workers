package api

import (
	"autoworkers/internal/job"
	"autoworkers/internal/queue"
	"autoworkers/internal/store"
	"fmt"
	"net/http"
)

func (a *ApiServer) SubmitJob(w http.ResponseWriter, r *http.Request){
	fmt.Println(a.apistore)
	fmt.Println(a.apiqueue)
	testjob := &job.Job{
		ID: "job-1",
		Type: "test",
		Payload: "hello world",
		Status: job.Pending,
	}
	store.Create(testjob,a.apistore)
	queue.Enqueue(a.apiqueue,testjob)
	fmt.Fprintln(w,testjob.Status)
	fmt.Fprintln(w,testjob.ID)

}
