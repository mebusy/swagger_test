package api

import (
	// "log"
	"encoding/json"
	"net/http"
)

func (s *srvApi) PostLogin(w http.ResponseWriter, r *http.Request) {
	var login ResLogin

	login.Data = nil

	json.NewEncoder(w).Encode(login)
}
