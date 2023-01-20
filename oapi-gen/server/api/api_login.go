package api

import (
	// "log"
	"encoding/json"
	"net/http"
)

func (s *srvApi) PostLogin(w http.ResponseWriter, r *http.Request) {
	var login ResLogin
	login.Data = &LoginOut{Token: "apikey-123456"}
	json.NewEncoder(w).Encode(login)
}
