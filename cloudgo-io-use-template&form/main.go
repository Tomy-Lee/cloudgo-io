package main

import (
	"os"
	"net/http"
	"fmt"
	"html/template"
	"log"
	//直接调用自己文件夹下的service包
	"cloudgo-inout/service"
	flag "github.com/spf13/pflag"
)

const (
	PORT string = "8080"
)


func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	server := service.NewServer()
	server.Run(":" + port)
}
