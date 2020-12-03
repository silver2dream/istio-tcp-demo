package content

import (
	"fmt"
	"net/http"
	"server/conf"
	database "server/db"
	"time"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

type Http struct {
	config conf.Conf
}

func (h *Http) echoHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("receive from:", request.RemoteAddr)
	name := database.GetTestingSQLService()
	response.Write([]byte(name + ":" + time.Now().Format("2006-01-02 15:04:05")))
}

func (h *Http) Start() {
	router.HandleFunc("/echo", h.echoHandler)

	fmt.Println("http server start.")
	fmt.Println(h.config.Proto.Port)
	http.ListenAndServe(h.config.Proto.Port, router)
}
