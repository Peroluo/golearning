package main

import (
	initrouter "golearning/init"
	"net/http"
	"time"
)

func main() {
	Router := initrouter.InitRouter()
	s := &http.Server{
		Addr:           ":8888",
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}
