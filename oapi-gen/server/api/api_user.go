package api

import (
	// "log"
	"encoding/json"
	"net/http"
)

func (s *srvApi) GetUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(NoError)
}

func (s *srvApi) PostUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(NoError)
}
