API-клиент для serials-now
==========================
[![Build Status](https://travis-ci.org/iamsalnikov/serials-now-api.svg?branch=travis)](https://travis-ci.org/iamsalnikov/serials-now-api)
[![GitHub license](https://img.shields.io/github/license/iamsalnikov/serials-now-api.svg)](https://github.com/iamsalnikov/serials-now-api/blob/master/LICENSE.md)

## Принцип работы

Основной принцип работы с клиентом такой:

1. Создаем экземпляр клиента
2. Создаем экземпляр эндойнта
3. Выполняем запрос и получаем результат

Разберем на примере поиска:

```go
package main

import (
	"github.com/iamsalnikov/serials-now-api/search"
	"github.com/iamsalnikov/serials-now-api"
	"log"
)

func main() {
	searchEndpoint := search.NewEndpoint()

	client, err := serials_now_api.NewClient("https://serials.ru/")
	if err != nil {
	    log.Fatalln(err)
	}

	err := client.Send(searchEndpoint)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(searchEndpoint.Serials[1])
}

```

У каждого эндпоинта может быть свой набор данных.