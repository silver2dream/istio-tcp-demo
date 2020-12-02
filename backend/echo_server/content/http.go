package content

import (
	"fmt"
	"net/http"
	"server/conf"
	"time"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

type Http struct {
	config conf.Conf
}

func (h *Http) echoHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("receive from:", request.RemoteAddr)
	response.Write([]byte(time.Now().Format("2006-01-02 15:04:05")))
}

func (h *Http) Start() {
	router.HandleFunc("/echo", h.echoHandler)

	fmt.Println("http server start.")
	fmt.Println(h.config.Srv.Port)
	http.ListenAndServe(h.config.Srv.Port, router)
}
