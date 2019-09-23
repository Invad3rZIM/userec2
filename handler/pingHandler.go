package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Uploads video to s3 bucket and also adds metadata to datastore.
func (h *Handler) PingHandler(w http.ResponseWriter, r *http.Request) {
	h.enableCors(&w)

	fmt.Println("xx", r.Body)
	json.NewEncoder(w).Encode("pong")
	w.WriteHeader(http.StatusOK)
	return
}

//Uploads video to s3 bucket and also adds metadata to datastore.
func (h *Handler) KillVids(w http.ResponseWriter, r *http.Request) {
	h.enableCors(&w)
	h.database.KillVids()

	json.NewEncoder(w).Encode("kill")
	w.WriteHeader(http.StatusOK)
	return
}
