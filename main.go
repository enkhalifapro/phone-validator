package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/enkhalifapro/phone-validator/api"
	"github.com/enkhalifapro/phone-validator/phones"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := httprouter.New()

	r.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Origin", "*")
		}

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})

	db, err := sql.Open("sqlite3", "./sample.db")
	if err != nil {
		panic(err)
	}

	phoneMgr := phones.New(db)

	// register country routes
	countryHandler := api.NewCountryHandler(phoneMgr)
	r.GET("/countries", countryHandler.List)

	fmt.Println("Server started at port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
