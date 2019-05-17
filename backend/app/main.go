package main

import (
	"HBlog/db"
	"HBlog/server"
	"flag"
	"log"
)

func main() {
	var addr string
	var initData bool
	flag.StringVar(&addr, "addr", ":1323", "server listens at this addr")
	flag.Parse()

	err := db.InitializeGlobalDB("127.0.0.1")
	if err != nil {
		log.Panic(err)
	}

	s := server.NewServer(addr)
	err = s.Init()
	if err != nil {
		log.Panic(err)
	}
	s.Start()
}
