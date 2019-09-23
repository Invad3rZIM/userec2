package database

import (
	"context"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"

	"hypeman-videouploader/constants"
)

type Database struct {
	mongo *mongo.Client
	redis *redis.Client
}

func NewDatabase(mongo *mongo.Client, redis *redis.Client) *Database {
	return &Database{mongo: mongo, redis: redis}
}

func (d *Database) KillVids() {
	x := []string{constants.METADATA, "VideoDislikes", "VideoLikes", "VideoLaughs", "VideoViews"}

	for _, a := range x {
		videos := d.mongo.Database(constants.HYPEMAN).Collection(a)
		filter := bson.M{}

		videos.DeleteMany(context.TODO(), filter)

	}

	d.redis.FlushAll()
	d.ClearS3()
}
