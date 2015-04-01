package gokumanga_geo

import (
	"fmt"
	"net/http"
	"strings"
	"encoding/json"
)

type VisitorGeo struct {
	Country     string `json:"country"`
	Region      string `json:"region"`
	City        string `json:"city"`
	CityLatLong string `json:"citylatlong"`
}

func init() {
	http.HandleFunc("/", handler)
}

func AddSafeHeaders(w *http.ResponseWriter) {
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	w.Header().Set("X-Frame-Options", "SAMEORIGIN")
	w.Header().Set("Strict-Transport-Security", "max-age=2592000; includeSubDomains")
}

func handler(w http.ResponseWriter, r *http.Request) {
	var respBytes []byte
	var err error

	AddSafeHeaders(&w)

	var respObj = &Response{
		Country: r.Header.Get("X-AppEngine-Country"),
		Region: r.Header.Get("X-AppEngine-Region"),
		City: r.Header.Get("X-AppEngine-City"),
		CityLatLong: r.Header.Get("X-AppEngine-CityLatLong"),
	}

	if respBytes, err = json.Marshal(respObj); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(respBytes)
}
