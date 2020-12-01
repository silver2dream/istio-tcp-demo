package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"gorm.io/gorm"

	conf "server/conf"
)

var config conf.Conf
var protocol conf.Protocol

type Player struct {
	gorm.Model
	Name     string
	Account  string
	Password string
}

func echoHandler(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Hello World"))
}

var router = mux.NewRouter()

func init() {
	var confFile string
	flag.StringVar(&confFile, "c", os.Args[1], "config file")
	flag.Parse()

	err := conf.ConfParser(confFile, &protocol, &config)
	if err != nil {
		log.Fatalf("parser config failed:", err.Error())
	}
}

func main() {
	router.HandleFunc("/echo", echoHandler)

	fmt.Println("server start.")
	fmt.Println(config.Srv.Port)
	http.ListenAndServe(config.Srv.Port, router)

}
