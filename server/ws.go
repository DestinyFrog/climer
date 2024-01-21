package server

import (
	"fmt"
	"net/http"
	"encoding/json"
	"time"
	"log"

	"climer/models"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader {
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

var clients []websocket.Conn

func OpenWSConnection(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	clients = append(clients, *conn)

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("err: %s\n", err.Error())
			return
		}

		fmt.Printf( "%s send: %s (%d)\n", conn.RemoteAddr(), string(msg), msgType )
		if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		}
	}
}

func UpdateConnection() {
	crudeData, err := models.GetData()
	if (err != nil) {
		fmt.Printf( "error: Não foi possível receber informações do Arduino\n" )
		return
	}

	codedData, err := json.Marshal(crudeData)
	if err != nil {
		fmt.Printf( "error: Dados do Arduino improcessaveis\n" )
		return
	}

	log.Println( "Data Update" )

	for idx, client := range clients {
		if err = client.WriteMessage(1, codedData); err != nil {
			log.Printf( "error: Socket {%d} Desconectado\n", idx )
			clients = append(clients[:idx], clients[idx+1:]...)
			continue
		}
	}
}

func UpdateRoutine() {
	for {
		time.Sleep(5 * time.Second)

		if ( len( clients ) > 0) {
			go UpdateConnection()
		}
	}
}