package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"github.com/tbruyelle/hipchat-go/hipchat"
)

type RoomConfig struct {
	token *hipchat.OAuthAccessToken
	hc *hipchat.Client
	name string
}

type Context struct {
	baseURL string
	static string
	rooms map[string]*RoomConfig
}

func main() {
	var (
		port = flag.String("port", "8080", "web server port")
		static = flag.String("static", "./static/", "static folder")
		baseURL = flag.String("baseurl", os.Getenv("BASE_URL"), "local base url")
	)

	flag.Parse()

	c := &Context{
		baseURL: *baseURL,
		static: *static,
		rooms: make(map[string]*RoomConfig),
	}

	log.Printf("Base HipCHat integration v0.10 - running on port:%v", *port)

	r := c.Routes()
	http.Handle("/", r)
	http.ListenAndServe(":"+*port, nil)
}