package main

import (
	"github.com/micro/go-log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-web"
	"ql/qlweb/handler"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("go.micro.web.qlweb"),
		web.Version("latest"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	router := httprouter.New()
	router.NotFound = http.FileServer(http.Dir("html"))

	// register html handler
	//service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	//service.HandleFunc("/example/call", handler.ExampleCall)
	service.Handle("/", router)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
