package main

import (
	"context"
	"database/sql"
	"fmt"
	"genLink/proto/api"
)

type Server struct {
	DB *sql.DB
	Gen Generator
}

const (
	table  string = "links"
)

func (s *Server) Create(ctx context.Context, longLink *api.URL) (shortLink *api.ShortURL, err error) {
	// генерирование ссылки
	s.Gen.GenerateShortLink()
	// выполняем запрос на поиск длинной ссылки
	var sl string
	s.DB.QueryRow("select shortlink from "+table+" where longlink = $1", longLink.URL).Scan(&sl)
	if sl!= "" {
		if sl != s.Gen.genLink{
			fmt.Println(longLink.GetURL() + " найден " + sl)
			return &api.ShortURL{ShortURL: sl}, nil
		}
	}
	// если соответствий не найдено добавляем в таблицу новое значение
	_, err = s.DB.Exec("insert into " + table + " (longlink, shortlink) values ($1,$2)", longLink.GetURL() , s.Gen.genLink)
	if err != nil {
		return nil, err
	}
	fmt.Println(longLink.GetURL() + " сгенерирован " + s.Gen.genLink)
	return &api.ShortURL{ShortURL: s.Gen.genLink}, nil
}

func (s *Server) Get(ctx context.Context, shortLink *api.ShortURL)(*api.URL, error){
	var ll string
	s.DB.QueryRow("select longlink from "+table+" where shortlink = $1", shortLink.ShortURL).Scan(&ll)
	if ll!= "" {
			fmt.Println(shortLink.ShortURL + " найден " + ll)
			return &api.URL{URL: ll}, nil
	}
	return &api.URL{URL: "empty"}, nil
}
