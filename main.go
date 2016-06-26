package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"Coolpy/Cors"
	"Coolpy/Account"
	"encoding/json"
	"Coolpy/BasicAuth"
)

func main() {
	router := httprouter.New()
	router.GET("/:uid", Basicauth.Auth(Index))
	router.POST("/", IndexPost)

	if err := http.ListenAndServe(":8080", Cors.CORS(router)); err != nil {
		log.Fatal(err)
	}
}

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	p,err := Account.Get(ps.ByName("uid"))
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(p)
}

func IndexPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var p Account.Person
	err := decoder.Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Fatal(err)
		return
	}
	p.CreateUkey()
	err = Account.CreateOrReplace(&p)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Fatal(err)
		return
	}
	json.NewEncoder(w).Encode(p)
}