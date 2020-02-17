package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{

	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func service(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Services")
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(p))
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func wsEndPoint(w http.ResponseWriter, r *http.Request) {
	//Allow any connection
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client Connected successfully..!!")
	err = ws.WriteMessage(1, []byte("Hi Client"))

	if err != nil {
		log.Println(err)
	}
	//listen for client message
	reader(ws)

}

func setUpRouters() {

	http.HandleFunc("/", homepage)
	http.HandleFunc("/services", service)
}

func main() {

	fmt.Println("Websockets")
	setUpRouters()
	log.Fatal(http.ListenAndServe(":8090", nil))

}
