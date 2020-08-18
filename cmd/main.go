package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"nasa-api/cmd/controllers"
	"nasa-api/config"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const port = ":4000"

var collection *mongo.Collection
var ctx = context.TODO()

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func init() {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("nasaapi").Collection("techport")

	log.Print("MONGO DB connected!", collection)
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
