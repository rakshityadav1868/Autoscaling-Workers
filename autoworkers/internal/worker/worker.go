package worker

import (
	"autoworkers/internal/database"
	"autoworkers/internal/executor"
	"autoworkers/internal/job"
	"autoworkers/internal/redis"
	"autoworkers/internal/store"
	"fmt"
)
type Worker struct{
	id int
	redisqueue *redis.Redis
	store *store.Store
	database *database.Database
}
func Constructor(id int ,redisqueue *redis.Redis,store *store.Store, database *database.Database) *Worker{
	s := &Worker{
		id: id,
		redisqueue: redisqueue,
		store: store,
		database: database,
	}
	return s

}

func Workers(m *Worker){
	for{
		jobId := m.redisqueue.Dequeue()
		fmt.Printf("Worker %d processing %s\n",m.id,jobId)
		jobobj := store.Get(m.store,jobId)
		if jobobj==nil{
			fmt.Println("No job found")
			continue
		}else{
			
			jobobj.Status = job.Running
			store.UpdateStatus(jobobj,m.store)
			m.database.UpdateJob(jobobj)
			result ,err := executor.Execute(jobobj)
			if err!=nil{
				jobobj.Status = job.Failed
				Error := err.Error()
				jobobj.Error = Error
				store.UpdateStatus(jobobj,m.store)
				m.database.UpdateJob(jobobj)
				continue
			}else{

				jobobj.Result = result
				jobobj.Status = job.Completed
				store.UpdateStatus(jobobj,m.store)
				m.database.UpdateJob(jobobj)
			}
		}
	}

}