package main

import (
	"api/exif"
	"flag"
	"fmt"
	"golib/client"
	"golib/server"
	"golib/server/http"
	"log"
	"time"
)

func main() {
	var (
		addr   string
		geoURL string
		token  string
	)
	flag.StringVar(&addr, "a", "0.0.0.0:3000", "Addr")
	flag.StringVar(&geoURL, "g", "", "GEO URL")
	flag.StringVar(&token, "t", "", "Token")
	flag.Parse()

	var geoClient *client.HTTP
	if geoURL != "" && token != "" {
		geoClient, _ = client.NewHTTP(geoURL, map[string]string{"token": token}, "")
	}
	exif.Initialize(geoClient)

	fmt.Println("Serve HTTP on", addr)
	log.Fatal(server.ServeHTTP(addr, http.NewServer(), 5*time.Second))
}
