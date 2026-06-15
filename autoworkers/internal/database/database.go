package database

import (
	"database/sql"
	"fmt"

	"autoworkers/internal/job"

	_ "modernc.org/sqlite"
)

type Database struct {
	DB *sql.DB

}

func Constructor() *Database{
	db, err := sql.Open("sqlite", "./taskforge.db")
	query := `
	CREATE TABLE IF NOT EXISTS Jobs(
	id TEXT,
	type TEXT,
	payload TEXT,
	status INTEGER,
	result TEXT
	)
	`
	fmt.Println("after open", err)
	if err != nil {
		fmt.Println(err)
	}
	err =db.Ping()
	_, err = db.Exec(query)
	if err!=nil{
		 fmt.Println(err)
	}else{

	}
	return &Database{
	DB: db,
	}
}

func (d *Database) SaveJob(j *job.Job){
query := `
INSERT INTO JOBS(ID,TYPE,PAYLOAD,STATUS,RESULT)
VALUES(?,?,?,?,?)
`
res,err := d.DB.Exec(query,j.ID,j.Type,j.Payload,j.Status,j.Result)
if err!=nil{
fmt.Println(err)
}else{
	fmt.Println(res)
}
}

func (d *Database) GetJob(id string) *job.Job{
	query := `
	SELECT id, type, payload, status, result
	FROM Jobs
	WHERE id = ?
	`
	row := d.DB.QueryRow(query,id)
	j := &job.Job{
		
	}
	err := row.Scan(&j.ID,&j.Type,&j.Payload,&j.Status,&j.Result)
	if err!=nil{
		fmt.Println(err)
		return nil
	}
	return j
}

func (d *Database) UpdateJob(j *job.Job){
	query :=`
UPDATE Jobs
SET status = ?, result = ?
WHERE id = ?`
res,err := d.DB.Exec(query,j.Status,j.Result,j.ID)
if err!=nil{
	fmt.Println(err)
}else{
	fmt.Println(res)
}


}