package cache

import (
	"errors"

	"hypeman-videouploader/constants"
	"hypeman-videouploader/metadata"
)

//TimeCache holds data caches for the various time durations...
type Cache struct {
	recentUploads map[string]*metadata.Metadata
}

func NewCache() *Cache {
	return &Cache{
		recentUploads: make(map[string]*metadata.Metadata),
	}
}

//Implement Redis here later!
func (c *Cache) Add(data *metadata.Metadata) {
	c.recentUploads[data.Videoname] = data
}

//Retrieve tries to get the video from the cache, returning an error if it can't
func (c *Cache) Retrieve(video string) (*metadata.Metadata, error) {
	if v, ok := c.recentUploads[video]; ok {
		return v, nil
	}

	return nil, errors.New(constants.NotInCacheError)
}
