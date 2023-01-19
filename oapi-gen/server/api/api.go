package api

import (
	"net/http"
)

type srvApi struct {
}

func NewAPI() *srvApi {
	return &srvApi{}
}

// Make sure we conform to ServerInterface
var _ ServerInterface = (*srvApi)(nil)

func (s *srvApi) GetUser(w http.ResponseWriter, r *http.Request) {
}

func (s *srvApi) PostUser(w http.ResponseWriter, r *http.Request) {
}
