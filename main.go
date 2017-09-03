package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
)

type Request struct {
	data map[string]interface{}
}

var err error
var session *mgo.Session
var conn = session.DB("getreq").C("requests")

func main() {
	session, err = mgo.Dial("192.168.0.107:27017/getreq")
	if err != nil {
		fmt.Println("Failed connecting to DB")
	} else {
		fmt.Println("Succesfully Connected To Database")
	}
	defer session.Close()

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var reqmap Request
	reqmap, err = json.Marshal(r.Header)
	if err != nil {
		log.Fatal("cannot encode to json ", err)
	}
	//w.Header().Set("Content-type", "application/json")
	//fmt.Fprintf(w, "%s", reqmap)

	err = conn.Insert(reqmap)
	if err != nil {
		fmt.Fprintf(w, "error writing to db")
	}
}
