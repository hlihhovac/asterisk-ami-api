package main

import (
	"github.com/hlihhovac/asterisk-ami-api/tree/master/internal/platform/ami"
	"github.com/hlihhovac/asterisk-ami-api/tree/master/internal/utils/config"
	"log"
	"net/http"
)

func main() {

	var conf = config.GetConfig()
	var err error

	srv := &http.Server{
		Addr: conf.General.Listen,
		//ReadTimeout:  api.HTTP_TIMEOUT,
		//WriteTimeout: api.HTTP_TIMEOUT,
		Handler: api.NewHandler(),
	}

	if err = srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
