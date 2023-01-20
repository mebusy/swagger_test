package api

import (
	// "log"
	"encoding/json"
	"fmt"
	"net/http"
	"server/utils"
)

func (s *srvApi) GetUser(w http.ResponseWriter, r *http.Request) {
	var userget ResUserGet

	userget.Errcode = ERR_NO_ERROR
	userget.Data = &UserOut{Id: 12, Name: "cy"}

	json.NewEncoder(w).Encode(userget)
}

func (s *srvApi) PostUser(w http.ResponseWriter, r *http.Request) {
	body, _ := utils.GetRequestBody(r)
	// even though validator can ensure body is valid
	// we still check it as a protection
	// if err != nil {
	// 	errmsg := err.Error()
	// 	json.NewEncoder(w).Encode(ResError{Errcode: ERR_INVALID_BODY, Errmsg: &errmsg})
	// 	return
	// }

	// unmarshal body to PostUserJSONRequestBody
	var reqbody PostUserJSONRequestBody
	_ = json.Unmarshal(body, &reqbody)
	// if err != nil {
	// 	errmsg := err.Error()
	// 	json.NewEncoder(w).Encode(ResError{Errcode: ERR_INVALID_BODY, Errmsg: &errmsg})
	// 	return
	// }
	fmt.Printf("%+v\n", reqbody)

	json.NewEncoder(w).Encode(RES_NOERROR)
}
