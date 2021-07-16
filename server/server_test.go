package main

import (
	"context"
	"database/sql"
	"genLink/proto/api"
	"log"
	"testing"
)

// установить соединение с базой

const (
	login   string = "root"
	pass    string = "0000"
	mysqldb string = "0.0.0.0:3306"
)

var ShortLink string
var LongLink = "test.com"

type testCase struct {
	testlink      string // тестовая ссылка
	testshortlink string
	isErr         bool // наличие ошибок
}

func db() (*sql.DB, error) {

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
	return db, nil
}

func TestCreate(t *testing.T) {
	//тест цепляется у базе
	db, err := db()
	if err != nil {
		log.Fatal(err)
	}

	var gen Generator
	gen.SetIsLong(10).SetAsciiSlice(LatinaBig, LatinaBig, Numbers)

	server := Server{DB: db, Gen: gen}

	testCases := []testCase{
		//
		testCase{
			testlink: LongLink,
			isErr:    false,
		},
		testCase{
			testlink:      "",
			testshortlink: "",
			isErr:         true,
		},
	}
	for i, tc := range testCases {
		resShortLink, err := server.Create(context.Background(), &api.URL{URL: tc.testlink})
		if err != nil {
			t.Fatal("Err Create method: ", err)
		}
		//сохраняем результат в глобальную переменную, чтобы использовать ее в следующем тесте
		if i == 0 {
			ShortLink = resShortLink.ShortURL
		}

		var resLongLink string
		_ = db.QueryRow("select longlink from "+table+" where shortlink = ?", resShortLink.ShortURL).Scan(&resLongLink)

		if tc.testlink != resLongLink && tc.isErr == false {
			t.Error("Values not equal")
		}
		if tc.testshortlink != "" && tc.isErr == true {
			t.Error("Expected empty ShortURL")
		}
	}
}

func TestGet(t *testing.T) {
	db, err := db()
	if err != nil {
		log.Fatal(err)
	}
	server := Server{DB: db}

	testCases := []testCase{
		testCase{
			testshortlink: ShortLink,
			testlink:      LongLink,
			isErr:         false,
		},
	}
	for _, tc := range testCases {
		resLongLink, err := server.Get(context.Background(), &api.ShortURL{ShortURL: tc.testshortlink})
		if resLongLink.URL != tc.testlink {
			t.Error("Values not equal in Get ")
		}
		if err != nil && tc.isErr == false {
			t.Error("Unexpected error in Get")
		}

	}
}
