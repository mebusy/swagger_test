package api

import (
	// "log"
	"fmt"
	"net/http"
	"server/utils"
)

func (s *srvApi) Get(w http.ResponseWriter, r *http.Request, params GetParams) {
	if params.Api != nil {
		switch *params.Api {
		case "info":
			utils.InfoHandler(w, r)
		default:
			fmt.Fprint(w, "ok")
		}
	} else {
		fmt.Fprint(w, "ok")
	}
}
