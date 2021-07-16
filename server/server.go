package main

import (
	"context"
	"database/sql"
	"fmt"
	"genLink/proto/api"
	"log"
	"os"
)

type Server struct {
	DB  *sql.DB
	Gen Generator
}

const (
	table string = "linkdb.links"
)

func (s *Server) Create(ctx context.Context, longLink *api.URL) (shortLink *api.ShortURL, err error) {

	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	if longLink.URL == "" {
		return &api.ShortURL{ShortURL: ""}, nil
	}

	// выполняем запрос на поиск длинной ссылки
	var sl string
	s.DB.QueryRow("select shortlink from "+table+" where longlink = ?", longLink.URL).Scan(&sl)
	if sl != "" {
		fmt.Println(longLink.GetURL() + " найден " + sl)
		return &api.ShortURL{ShortURL: sl}, nil
	}
	// генерирование ссылки
	s.Gen.GenerateShortLink()
	// если соответствий не найдено добавляем в таблицу новое значение
	_, err = s.DB.Query("insert into "+table+" (longlink, shortlink) values (?,?)", longLink.GetURL(), s.Gen.genLink)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	fmt.Println(longLink.GetURL() + " сгенерирован " + s.Gen.genLink)
	return &api.ShortURL{ShortURL: s.Gen.genLink}, nil
}

func (s *Server) Get(ctx context.Context, shortLink *api.ShortURL) (*api.URL, error) {
	var ll string
	s.DB.QueryRow("select longlink from "+table+" where shortlink = ?", shortLink.ShortURL).Scan(&ll)
	if ll != "" {
		fmt.Println(shortLink.ShortURL + " найден " + ll)
		return &api.URL{URL: ll}, nil
	}
	return &api.URL{URL: "empty"}, nil
}
