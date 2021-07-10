package main

import (
	"database/sql"
	"fmt"
	"genLink/proto/api"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"net"
)

func main() {

	const (
		user     string = "postgres"
		password string = "1643"
		database string = "linkdatabase"
		host     string = "localhost"
		port     string = "5432"
		sslmode  string = "disable"
	)

	connData := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		user, password, database, host, port, sslmode)
	db, err := sql.Open("postgres", connData)
	if err != nil {
		fmt.Println(err)
		//log.Fatal(err)
	}

	//создаем генератор с нужной конфигурацией
	var gen Generator
	gen.SetIsLong(10).SetAsciiSlice(LatinaBig, LatinaBig, Numbers)

	// создаем объект, который реализует интерфейс GenLinkServer,
	// передаем в него обект для работы с базой
	s := &Server{DB:db, Gen: gen}
	srv := grpc.NewServer()
	api.RegisterGenLinkServer(srv,s)

	l,err:= net.Listen("tcp",":8080" )
	if err != nil{
		panic(err)
	}

	err = srv.Serve(l)
	if err != nil{
		panic(err)
	}
}
