package util

import (
	"log"
	"encoding/json"
	"net/http"
	"net/http/httputil"
)

type PayloadItem map[string]interface{}

// PrintDump prints the dump of request, optionally writing it in the response
func PrintDump(w http.ResponseWriter, r *http.Request, write bool) {
	dump, _ := httputil.DumpRequest(r, true)
	log.Printf("%v", string(dump))
	if write == true {
		w.Write(dump)
	}
}

// Decode into a map[string]interface{} the JSON in the POST Request
func DecodePostJSON(r *http.Request, logging bool) (PayloadItem, error) {
	var err error
	var payLoad PayloadItem
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&payLoad)
	if logging == true {
		log.Printf("Parsed body:%v", payLoad)
	}

	return payLoad, err
}