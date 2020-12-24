package main

import (
	"encoding/json"
	"fmt"
	"github.com/segmentio/ksuid"
	"net/http"
	"time"
)

type Responce struct {
	RequestId string `json:requestId`
	Error interface{} `json:"error,omitempty""`
	Result interface{} `json:"result,omitempty"`
}
func sayok(w http.ResponseWriter, req *http.Request) {
	delay := req.FormValue("delay")
	connection := req.FormValue("connection")
	if delay != "" {
		intetval, err := time.ParseDuration(delay)
		if err == nil {
			<-time.After(intetval)
		}
	}

	resp:=&Responce{}
	resp.RequestId=ksuid.New().String()
	resp.Result = req.Header
	bytes,err:= json.Marshal(resp)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	if connection != "" {
		w.Header().Set("Connection",connection)
	}

	w.Write(bytes)

}

func main() {
	fmt.Println("hello world")
	http.HandleFunc("/v1/ok", sayok)

	http.ListenAndServe(":8000", nil)

}
