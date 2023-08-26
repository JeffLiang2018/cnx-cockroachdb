package main
import "fmt"

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Hello World, CNX CockroachDB!!!")
	host := "10.164.0.189"
	port := 26257
	user := "pvos_dia_db_rw"
	password := "c0f47960-9c78-4039-9760-c0ff390b3259_pvos_dia_db_rw"
	database := "pvos_dia_db"
	sslrootcert := "./ca.crt"
	sslcert := "./client.root.crt"
	sslkey := "./client.root.key"
	dns := fmt.Sprintf("host=%s user=%s password=%s database=%s port=%d sslmode=require sslrootcert=%s sslcert=%s sslkey=%s",
		host, user, password, database, port, sslrootcert, sslcert, sslkey)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("error happened when creating CockroachDB connection, err: %v", err))
	}
	if db != nil {
		fmt.Println("CockroachDB connection is created successfully")
	}
	fmt.Println("CNX CockroachDB is connected successfully!!!!!")
	rawDB, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("error happened when getting rawDB before closing connection, err: %v", err))
	}
	rawDB.Close()
	fmt.Println("CockroachDB connection is closed successfully")
}