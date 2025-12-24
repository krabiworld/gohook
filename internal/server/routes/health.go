package routes

import (
	"fmt"
	"log"
	"net/http"
)

func Health(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if _, err := fmt.Fprint(w, `{"status":"ok"}`); err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}
