package internal

import (
	"fmt"
	"log"
	"net/http"
	
)

func Serve() {
	http.HandleFunc("/rss", rss)
	log.Println("Starting web server: listening at port 8090")
	http.ListenAndServe(":8090", nil)
}

func rss(w http.ResponseWriter, req *http.Request) {
	log.Println("Received request: Querying Omnivore...")

	feedString, err := getFeed()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/rss+xml")
	fmt.Fprintf(w, feedString)
}
