package main

import (
	"log"
	"net/http"
	"encoding/json"
)

type ChanellResponse struct {
	Type 		string  `json:"type"`
	ChanellID 	string 	`json:"chanell_id"`
	Message 	*Message `json:"message,omitempty"`

}

type Message struct {
	To      string `json:"to"`
	Name    string `json:"name"`
	Message string `json:"message"`
}


func ChannellSwitcher(w http.ResponseWriter, r *http.Request,hub *Hub){
	log.Println(r.RequestURI, r.Method)
	channel := ChanellResponse{}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&channel)
	if err != nil {
		log.Println(r.RequestURI, r.Method, err.Error())
	}

	/// here  will find client by token or another way
	client := &Client{}

	if channel.Type == "sub"{
			client.hub.register <-client
			serveWs(hub,w,r)
	} else if channel.Type == "unsub"{
		client.hub.unregister <- client
		client.conn.Close()
	}else if channel.Type =="event"{

	}else{
		w.WriteHeader(400)
	}


}

