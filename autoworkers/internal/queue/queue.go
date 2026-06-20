package queue

import "autoworkers/internal/job"

type Queue struct{
	JobId chan string
}

func Constructor() *Queue{
	channel := make(chan string,100)
	p := &Queue{
		JobId: channel,
	}
	return p

}

func  Enqueue(p *Queue, x *job.Job){
	p.JobId <- x.ID

}

func Dequeue(p *Queue) string{
	value := <- p.JobId
	return value
	
}

func Close(p *Queue){
	close(p.JobId)
}