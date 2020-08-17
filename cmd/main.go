package main

import (
	"log"
	"net/http"

	"nasa-api/config"
	"nasa-api/controllers"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

const port = ":4000"

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {

	ev, err := config.LoadEnv()
	if err != nil {
		log.Fatal("Bad Envs", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	pong, err := client.Ping().Result()

	if err != nil {
		log.Panic(err)
	} else {
		log.Println("Redis is connected", pong)
	}

	router := mux.NewRouter()
	router.Use(commonMiddleware)
	router.HandleFunc("/techport", controllers.TechPort(ev)).Methods("GET")

	log.Printf("Server running %1s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
