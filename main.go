package main

import (
	"log"
	"net/http"
	"time"
	apiTools "github.com/go-psql-orm/apitools"
	"github.com/gorilla/mux"
)


func main() {
	//Cria o roteador
	router := mux.NewRouter()
	spa := apiTools.SpaHandler{StaticPath: "src", IndexPath: "index.html"}
	
	apiTools.HandlersHandle(router, spa)
	
	//Dados do servidor
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	//Encerra a execução se tiver erro
	log.Fatal(srv.ListenAndServe())
}