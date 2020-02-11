// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (

	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)



func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main() {
	router := mux.NewRouter()
	hub := newHub()
	go hub.run()
	router.HandleFunc("/", serveHome)
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	router.HandleFunc("/chanell/status",func(w http.ResponseWriter, r *http.Request) {
		ChannellSwitcher( w, r,hub)
	}).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8078" //localhost
	}

	fmt.Println("Server is running on port 8078")

	err := http.ListenAndServe(":" + port,router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
