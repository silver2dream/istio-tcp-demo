package content

import (
	"client/conf"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Http struct {
	config conf.Conf
}

func (h *Http) Start() {

	for {
		res, err := http.Get(h.config.Host)
		if err != nil {
			panic(err)
		}

		if res.StatusCode != 200 {
			log.Fatalf("Unexpected response status code: %v", res.StatusCode)
		}
		defer res.Body.Close()
		sitemap, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", sitemap)
		time.Sleep(time.Duration(h.config.Interval) * time.Second)
	}
}
