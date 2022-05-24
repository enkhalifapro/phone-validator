package main

import (
	"database/sql"
	"fmt"
	"github.com/mattn/go-sqlite3"
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

	sql.Register("sqlite3_with_extensions",
		&sqlite3.SQLiteDriver{
			Extensions: []string{
				"sqlite3_mod_regexp",
			},
		})
	db, err := sql.Open("sqlite3_with_extensions", "./sample.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	phoneMgr := phones.New(db)

	// register country routes
	phoneHandler := api.NewPhoneHandler(phoneMgr)

	r.GET("/phones", phoneHandler.ListPhones)
	r.GET("/phones/:countryName", phoneHandler.GetPhonesByCountry)
	r.GET("/countries", phoneHandler.ListCountries)

	fmt.Println("Server started at port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
