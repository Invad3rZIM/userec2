package database

import (
	"context"
	"errors"

	"hypeman-userec2/constants"

	"gopkg.in/mgo.v2/bson"
)

func (db *Database) CheckExists(videoname string) error {
	filter := bson.M{"videoname": videoname}
	count, _ := db.mongo.Database(constants.HYPEMAN).Collection(constants.METADATA).CountDocuments(context.TODO(), filter)

	if count > 0 {
		return errors.New("Video already in the database")
	}

	return nil
}
