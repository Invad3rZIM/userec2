package main

import (
	"context"
	"net/http"

	"hypeman-videouploader/constants"
	"hypeman-videouploader/handler"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	log.Info(constants.START_MESSAGE)

	redis := redis.NewClient(&redis.Options{
		Addr:     constants.REDIS_HOST,
		Password: constants.REDIS_PASS,
		DB:       0,
	})

	clientOptions := options.Client().ApplyURI(constants.MONGO_URL)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Info(constants.MONGO_CONNECTED)

	err = client.Ping(context.Background(), nil)

	//make tmp directory
	handler.CreateDirIfNotExist("./tmp")

	h := handler.NewHandler(client, redis)

	http.HandleFunc("/user/profile/upload", h.ProfileUploadHandler)
	http.HandleFunc("/ping", h.PingHandler)

	port := ":443"

	log.Info(constants.SERVING_MESSAGE, port)

	//err = http.ListenAndServeTLS(port, "hypeapi.crt.pem", "hypeapi.key.pem", nil)
	err = http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal(err)
	}
}
