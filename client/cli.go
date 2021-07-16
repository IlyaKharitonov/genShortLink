package main

import (
	"context"
	"fmt"
	"genLink/proto/api"
	"google.golang.org/grpc"
	"log"
	"os"
)

func Сhoice(cli api.GenLinkClient) {

	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	var start string
	const Create = "1"
	const Get = "2"
	const Exit = "3"

	fmt.Print("\nЧто будем делать?\n" +
		"1. Генерировать сокращенную ссылку\n" +
		"2. Использовать короткую ссылку для получения первичной ссылки\n" +
		"3. Выйти\n" +
		"Выбери пункт:")
	fmt.Scan(&start)

	var link string
	switch start {
	case Create:
		fmt.Print("Вставь ссылку:")
		fmt.Scan(&link)
		result, err := cli.Create(context.Background(), &api.URL{URL: link})
		if err != nil {
			log.Println(err)
			log.Fatalf("Error in grpcserver, method Create: %s", err)
		}
		fmt.Printf("Короткая ссылка: %s\n", result.ShortURL)
		Сhoice(cli)
	case Get:
		fmt.Print("Вставь короткую ссылку:")
		fmt.Scan(&link)
		result, err := cli.Get(context.Background(), &api.ShortURL{ShortURL: link})
		if err != nil {
			log.Println(err)

			log.Fatalf("Error in grpcserver, method Get: %s", err)
		}
		if result.URL != "empty" {
			fmt.Printf("Найдена длинная ссылка: %s\n", result.URL)
		} else {
			fmt.Print("Этой ссылке ничего не соответствует :(\n")
		}
		Сhoice(cli)
	case Exit:
		fmt.Print("\nПрощай!")
		os.Exit(0)
	default:
		fmt.Print("\nТакой команды нет. Попробуй еще!\n")
		Сhoice(cli)
	}

}

func main() {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	fmt.Print("Привет, это сокращатель ссылок!\n")
	//создаем соединение с grpc сервером
	conn, err := grpc.Dial("127.0.0.1:9080", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		panic(err)
	}
	//создаем клиента с помощью сгенерированного кода
	cli := api.NewGenLinkClient(conn)
	//вызываем функцию, которая общается с нами и вызывает удаленные функции
	Сhoice(cli)
}
