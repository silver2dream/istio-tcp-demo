package main

import (
	"flag"
	"log"
	"os"

	"gorm.io/gorm"

	conf "server/conf"
	"server/content"
)

var config conf.Conf
var parserErr error
var protocol conf.Protocol

type Player struct {
	gorm.Model
	Name     string
	Account  string
	Password string
}

func init() {
	var confFile string
	flag.StringVar(&confFile, "c", os.Args[1], "config file")
	flag.Parse()

	config, parserErr = conf.ConfParser(confFile, &protocol)
	if parserErr != nil {
		log.Fatalf("parser config failed:", parserErr.Error())
	}
}

func main() {
	IFactory := &content.ContentFactory{}
	IContent := IFactory.Create(config)
	if IContent == nil {
		panic("content is nil.")
	}
	IContent.Start()
}
