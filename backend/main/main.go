package main

import (
	"flag"

	"MyBlog/server"
)

func main() {
	var initData bool
	flag.BoolVar(&initData, "init-data", false, "initialize data")
	flag.Parse()

	port := ":2333"
	
	server.Start(initData, port)
}
