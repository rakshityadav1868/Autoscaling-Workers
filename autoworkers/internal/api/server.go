package api

import (
	"autoworkers/internal/database"
	"autoworkers/internal/redis"
	"autoworkers/internal/store"
	"net/http"
)

type ApiServer struct{
	apistore *store.Store
	apiredis *redis.Redis
	count int
	apidatabase *database.Database
}

func Constructor(r *redis.Redis, s *store.Store, d *database.Database) *ApiServer{
p := &ApiServer{
	apistore: s,
	apiredis: r,
	count: 0,
	apidatabase: d,
}
return p
}


func (a *ApiServer) Start(){
	http.HandleFunc("POST /jobs", a.SubmitJob)
	http.HandleFunc("GET /jobs", a.GetAllJobs)
	http.HandleFunc("/jobs/",a.GetJob)
	http.ListenAndServe(":8080",nil)
}