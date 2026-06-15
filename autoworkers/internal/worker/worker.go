package worker

import (
	"autoworkers/internal/executor"
	"autoworkers/internal/job"
	"autoworkers/internal/queue"
	"autoworkers/internal/store"
	"fmt"
	"autoworkers/internal/database"
)
type Worker struct{
	queue *queue.Queue
	store *store.Store
	database *database.Database
}
func Constructor(queue *queue.Queue,store *store.Store, database *database.Database) *Worker{
	s := &Worker{
		queue: queue,
		store: store,
		database: database,
	}
	return s

}

func Workers(m *Worker){
	for{
		jobId := queue.Dequeue(m.queue)
		jobobj := store.Get(m.store,jobId)
		if jobobj==nil{
			fmt.Println("No job found")
			continue
		}else{
			
			jobobj.Status = job.Running
			store.UpdateStatus(jobobj,m.store)
			result := executor.Execute(jobobj)
			jobobj.Result = result
			jobobj.Status = job.Completed
			store.UpdateStatus(jobobj,m.store)
			m.database.UpdateJob(jobobj)
		}
	}

}