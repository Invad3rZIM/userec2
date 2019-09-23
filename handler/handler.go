package handler

import (
	"errors"
	"fmt"
	"hypeman-userec2/cache"
	"hypeman-userec2/database"
	"net/http"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	database *database.Database
	cache    *cache.Cache
}

func NewHandler(client *mongo.Client, redis *redis.Client) *Handler {
	return &Handler{
		database: database.NewDatabase(client, redis),
		cache:    cache.NewCache(),
	}
}

//EnableCors enables cors (Cross Organizational Resource Sharing)
func (h *Handler) enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

//VerifyBody is a helper function to ensure all http requests contain the requisite fields returns error if fields missing
func (h *Handler) verifyBody(body map[string]interface{}, str ...string) error {
	for _, s := range str {
		fmt.Println(s)
		if _, ok := body[s]; !ok {
			return errors.New("error: missing field: " + s)
		}
	}
	return nil
}
