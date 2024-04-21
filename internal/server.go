package internal

import (
	"fmt"
	"log"
	"net/http"
	
)

func Serve() {
	http.HandleFunc("/rss", rss)
	localPort := Cfg.LocalPort
	log.Println("Starting web server: listening at port ", localPort)
	http.ListenAndServe(":" + localPort, nil)
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
	fmt.Fprint(w, feedString)
}
