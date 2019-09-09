package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	getMapping("/hello")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	//println(r.URL.Path[1:])
	bytes, e := ioutil.ReadAll(r.Body)
	if e != nil {

	}
	request := myRequest{}
	json.Unmarshal(bytes, request)
	println(request.String())
}

func getMapping(endpoint string) {
	http.HandleFunc(endpoint, handler)
}

func (request myRequest) String() string{
	return request.Ok
}

type myRequest struct {
	Ok string `json:"hey"`
}
