package Configration

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func SetUpDB()*sql.DB{
	driver:=os.Getenv("DBDRIVER")
	user:=os.Getenv("DBUSER")
	pass:=os.Getenv("DBPASS")
	dbname:=os.Getenv("DBNAME")
	db,err:=sql.Open(driver,fmt.Sprintf("%s:%s@/%s",user,pass,dbname))
	if err!=nil{
		panic(err)
	}
	return db
}
