package database

import (
	"database/sql"
	"fmt"
)

type Database struct {
	DB *sql.DB

}

func Constructor() *Database{
	db, err := sql.Open("sqlite", "./taskforge.db")
	if err != nil {
		fmt.Println(err)
	}
	err =db.Ping()
	if err!=nil{
		 fmt.Println(err)
	}else{

	}
	return &Database{
	DB: db,
	}
}