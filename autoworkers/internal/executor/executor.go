package executor

import (
	"autoworkers/internal/job"
	"errors"
	"time"
)

func Execute(j *job.Job) (string,error){
	time.Sleep(2 * time.Second)
	if j.Type=="fail"{
		err := errors.New("failed")
		return "", err
	}else{
		return "success", nil
	}
}
