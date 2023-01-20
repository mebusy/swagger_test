package api

import (
	// "log"
	"encoding/json"
	"net/http"
)

func (s *srvApi) GetUser(w http.ResponseWriter, r *http.Request) {
	var userget ResUserGet

	userget.Errcode = ERR_NO_ERROR
	userget.Data = &UserOut{Id: 12, Name: "cy"}

	json.NewEncoder(w).Encode(userget)
}

func (s *srvApi) PostUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(RES_NOERROR)
}
