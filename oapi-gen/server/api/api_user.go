package api

import (
	"log"
	"net/http"
)

func (s *srvApi) GetUser(w http.ResponseWriter, r *http.Request) {
	log.Println("GetUser")
}

func (s *srvApi) PostUser(w http.ResponseWriter, r *http.Request) {
}
