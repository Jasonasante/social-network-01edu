package main

import (
	"fmt"
	"log"
	"net/http"
	"social-network/backend/functions"
	"social-network/backend/websocket"
)

func main() {
	// Create tabless
	functions.CreateSqlTables()
	// Serve files within static and public
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	// Handle websocket connections.
	hub := websocket.NewHub()
	go hub.Run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.ServeWs(hub, w, r)
	})

	// Endpoint handlers
	http.HandleFunc("/", functions.Homepage)
	http.HandleFunc("/login", functions.Login)
	http.HandleFunc("/logout", functions.Logout)
	http.HandleFunc("/register", functions.Register)
	http.HandleFunc("/api/user", functions.GetUserFromSessions)
	http.HandleFunc("/api/users", functions.UsersApi)
	http.HandleFunc("/api/followers", functions.FollowersApi)
	http.HandleFunc("/profile", functions.Profile)
	// http.HandleFunc("/public-profiles", functions.DynamicPath)
	http.HandleFunc("/create-chat", functions.CreateChat)
	http.HandleFunc("/edit-chatroom", functions.EditChatroom)
	http.HandleFunc("/get-chat", functions.Chat)
	http.HandleFunc("/ws/chat", functions.ServeWs)
	http.HandleFunc("/ws/user", functions.ServeWs)

	http.HandleFunc("/create-post", functions.CreatePost)
	go functions.H.Run()
	go functions.SqlExec.ExecuteStatements()

	fmt.Printf("SOCIAL-NETWORK serving at http://localhost:8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
