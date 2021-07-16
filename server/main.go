package main

import (
	"database/sql"
	"fmt"
	"genLink/proto/api"
	//
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"net"
)

func main() {

	const (
		login   string = "root"
		pass    string = "0000"
		mysqldb string = "127.0.0.1:3306"
	)

	connData := login + ":" + pass + "@(" + mysqldb + ")/"
	db, err := sql.Open("mysql", connData)

	_, err = db.Exec("create database if not exists linkdb")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("use linkdb")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("create table if not exists links( id int auto_increment primary key, longlink  varchar(250) not null, shortlink varchar(30)  not null)")
	if err != nil {
		panic(err)
	}
	//создаем генератор с нужной конфигурацией
	var gen Generator
	gen.SetIsLong(10).SetAsciiSlice(LatinaBig, LatinaBig, Numbers)

	// создаем объект, который реализует интерфейс GenLinkServer,
	// передаем в него обект для работы с базой
	s := &Server{DB: db, Gen: gen}
	srv := grpc.NewServer()
	api.RegisterGenLinkServer(srv, s)

	fmt.Println("Подключился к бд, запустил grpc server")
	l, err := net.Listen("tcp", "127.0.0.1:9080")
	if err != nil {
		panic(err)
	}

	err = srv.Serve(l)
	if err != nil {
		panic(err)
	}
}
