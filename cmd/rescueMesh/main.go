package main

import (
	"encoding/json"
	"net/http"
	"sync"
	"fmt"
	"github.com/akshayjha21/RescueMesh.git/internal/p2p"
)

type IncomingMsg struct {
	Message string `json:"message"`
}

var messageArr []string
var messagemu sync.Mutex

//cr variable for chat room
//enable cors
//store message
//getMessage
//pushMessage

var cr *p2p.ChatRoom

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                   // Allow all origins to access
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // Allowed HTTP methods
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")       // Allowed headers
}
func StoreMessage(message string) {
	messagemu.Lock()
	defer messagemu.Unlock()
	messageArr = append(messageArr, message)
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Only Get Method supported", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(messageArr)
	if err != nil {
		http.Error(w, "failed to encode messages", http.StatusInternalServerError)
		return
	}
}
func PostMessage(w http.ResponseWriter, r *http.Request){
	enableCORS(w)
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Only POST Method supported", http.StatusBadRequest)
		return

	}
	var msg_post IncomingMsg
	err:=json.NewDecoder(r.Body).Decode(&msg_post)
	if err != nil || msg_post.Message == "" {
		http.Error(w, "failed to decode", http.StatusBadRequest)
		return
	}

	err_pub := cr.Publish(msg_post.Message)
	StoreMessage(msg_post.Message)

	if err_pub != nil {
		fmt.Println("Sending message failed trying again...")
		http.Error(w, "failed to publish", http.StatusInternalServerError)

		return
	}
	w.WriteHeader(http.StatusOK)

}
