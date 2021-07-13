package main

import (
	"context"
	"database/sql"
	"fmt"
	"genLink/proto/api"
	"log"
	"testing"
)

// установить соединение с базой

const (
	user     string = "postgres"
	password string = "1643"
	database string = "linkdatabase"
	host     string = "localhost"
	port     string = "5432"
	sslmode  string = "disable"
)

var ShortLink string
var LongLink = "test10.com"

type testCase struct {
	testlink      string // тестовая ссылка
	testshortlink string
	isErr         bool // наличие ошибок
}

func db() (*sql.DB, error) {
	connData := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		user, password, database, host, port, sslmode)
	db, err := sql.Open("postgres", connData)
	if err != nil {
		return nil, err
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
		db.QueryRow("select longlink from "+table+" where shortlink = $1", resShortLink.ShortURL).Scan(&resLongLink)
		//resLongLink, err := server.Get(context.Background(), resShortLink)
		//if err != nil {
		//	t.Fatal("Err Create method: ", err)
		//}

		if tc.testlink != resLongLink && tc.isErr == false {
			t.Error("Values not equal")
		}
		if tc.testshortlink != "Ссылка не сгенерирована, пустой URL" && tc.isErr == true {
			t.Error("Expected error")
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
	//fmt.Printf("Тeстовая короткая: %v\n", ShortLink)
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
