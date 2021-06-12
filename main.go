package main

import (
	"log"

	"github.com/bo278252518/geek02/server/http"
)

func main() {
	svr := http.New(":8080")
	svr.RegisterHandler()
	if err := svr.Run(); err != nil {
		log.Print(err)
	}
}
