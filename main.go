package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"

	"github.com/Durga-chikkala/unique-request-counter/handler"
	"github.com/Durga-chikkala/unique-request-counter/service"
	redisStore "github.com/Durga-chikkala/unique-request-counter/store/redis"
	"github.com/Durga-chikkala/unique-request-counter/writer"
	"github.com/Durga-chikkala/unique-request-counter/writer/kafka"
	"github.com/Durga-chikkala/unique-request-counter/writer/logfile"
)

func main() {
	logFile, err := os.OpenFile("requests.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	defer logFile.Close()

	log.SetOutput(logFile)

	// router initialization
	router := mux.NewRouter()

	// DB initialization
	redisDB := NewRedisDB()

	// writer initialization
	countWriter := InitializeWriter()

	// layer initialization
	store := redisStore.New(redisDB)
	svc := service.New(&store, countWriter)
	httpHandler := handler.New(&svc)

	// endpoint
	router.HandleFunc("/api/verve/accept", httpHandler.Get).Methods(http.MethodGet)

	// Start a concurrent goroutine to log unique request counts periodically
	go svc.LogUniqueRequestCount()

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}

}

func InitializeWriter() writer.CountWriter {
	switch os.Getenv("WRITER") {
	case "LOGFILE":
		return logfile.Logfile{}
	case "KAFKA":
		return kafka.Kafka{}
	}

	return logfile.Logfile{}
}

func NewRedisDB() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password set
		DB:       0,                // Use default DB
	})

	return rdb
}
