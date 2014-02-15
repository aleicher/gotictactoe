package main

import "encoding/json"
import "net/http"
import "log"
import "os"

import "github.com/gorilla/mux"

func router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", IndexHandler).Methods("GET")
	router.HandleFunc("/games", GamesHandler)
	router.HandleFunc("/players", PlayersHandler)
	router.HandleFunc("/games/{gameId}", GameHandler)
	router.HandleFunc("/games/{gameId}/moves", MovesHandler)
	router.HandleFunc("/games/{gameId}/moves/{moveId}", MoveHandler)
	return router
}

func routerHandler(router *mux.Router) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		router.ServeHTTP(res, req)
	}
}

func IndexHandler(res http.ResponseWriter, req *http.Request) {
	data, _ := json.Marshal("{'hello':'wercker!'}")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Write(data)
}

func GamesHandler(res http.ResponseWriter, req *http.Request) {
}

func GameHandler(res http.ResponseWriter, req *http.Request) {
}

func PlayersHandler(res http.ResponseWriter, req *http.Request) {
}

func MovesHandler(res http.ResponseWriter, req *http.Request) {
}

func MoveHandler(res http.ResponseWriter, req *http.Request) {
}

func main() {
	handler := routerHandler(router())
	err := http.ListenAndServe(":"+os.Getenv("PORT"), handler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
