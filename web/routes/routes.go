package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const AppPrefix = "/test/v1"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Result  interface{} `json: "Result"`
}

func RoutesHandler(router *mux.Router) {

	router.HandleFunc(fmt.Sprintf("%s/healthcheck", AppPrefix), healthcheck)

}

func healthcheck(w http.ResponseWriter, r *http.Request) {

	resp := Response{Success: true, Message: "App is Running"}
	json_resp, _ := json.Marshal(resp)
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(json_resp)

}
