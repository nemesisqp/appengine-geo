package gokumanga_geo

import (
	"fmt"
	"net/http"
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

func handler(w http.ResponseWriter, r *http.Request) {
	var respBytes []byte
	var err error

	header := w.Header()
	header.Set("X-Content-Type-Options", "nosniff")
	header.Set("X-XSS-Protection", "1; mode=block")
	header.Set("X-Frame-Options", "SAMEORIGIN")
	header.Set("Strict-Transport-Security", "max-age=2592000; includeSubDomains")

	var respObj = &VisitorGeo{
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
	fmt.Fprint(w, string(respBytes))
}
