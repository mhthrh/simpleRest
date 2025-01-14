package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"simpleAPI/pkg/handler"
)

var (
	app  = ""
	ip   = ""
	port = ""
)

func init() {
	app = os.Getenv("APP_ENV")
	port = os.Getenv("PORT")
	ip = os.Getenv("IP")
}
func main() {
	log.Println("start application")
	log.Printf("app environment:%s,port:%s", app, port)
	srv := gin.Default()
	srv.Handle("GET", "/", handler.Handler)
	_ = srv.Run(fmt.Sprintf("%s:%s", ip, port))
}
