package handler

import (
	"encoding/json"
	"fmt"
	"hypeman-userec2/constants"
	"hypeman-userec2/metadata"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var sess *session.Session

func init() {
	var err error
	// Create a single AWS session (we can re use this if we're uploading many files)
	sess, err = session.NewSession(&aws.Config{Region: aws.String(constants.S3_REGION)})

	if err != nil {
		log.Fatal(err)
	}
}

//Uploads video to s3 bucket and also adds metadata to datastore.
func (h *Handler) ProfileUploadHandler(w http.ResponseWriter, r *http.Request) {
	h.enableCors(&w)

	file, header, err := r.FormFile("video")

	if err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	username := r.FormValue("username")

	videoname := username + "/" + header.Filename

	if data, err := h.cache.Retrieve(videoname); err == nil {
		json.NewEncoder(w).Encode(data)
		w.WriteHeader(http.StatusOK)
		return
	}

	//ensure video not already in the database
	if err := h.database.CheckExists(videoname); err != nil {
		json.NewEncoder(w).Encode(err.Error())
		w.WriteHeader(http.StatusOK)
		return
	}

	tagCount, err := strconv.Atoi(r.FormValue("tagcount"))

	if err != nil {
		fmt.Fprintln(w, err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	tags := []string{}

	for i := 0; i < tagCount; i++ {
		tags = append(tags, r.FormValue(fmt.Sprint("tag", i)))
	}

	filePath := "./" + header.Filename

	//2 Concurrent Threads
	var metadata *metadata.Metadata

	done := make(chan bool)

	//create tmp file
	tmpFile, _ := os.Create(filePath)
	io.Copy(tmpFile, file)

	h.AddFileToS3(sess, header.Filename, username)

	//delete tmp file
	os.Remove(filePath)

	done <- true

	json.NewEncoder(w).Encode(metadata)
	w.WriteHeader(http.StatusOK)
	return
}
