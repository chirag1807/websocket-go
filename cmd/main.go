package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/chirag1807/websocket-go/api/route"
	"github.com/chirag1807/websocket-go/config"
	"github.com/chirag1807/websocket-go/constants"
	"github.com/chirag1807/websocket-go/db"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"github.com/jackc/pgx/v5"
)

func main() {
	config.LoadEnv("../.config/")
	conn, err := db.DBConnection()
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close(context.Background())

	server := flag.String("server", "", "http,websocket")
	flag.Parse()

	if *server == "http" {
		StartHTTPServer(conn)
	} else if *server == "websocket" {
		StartWebSocketServer(conn)
	} else {
		log.Println("invalid server. Available server: http or websocket")
	}

}

func StartHTTPServer(conn *pgx.Conn) {
	r := route.Routes(conn)

	log.Println("Server started on port no. " + constants.PORT_NO)
	log.Fatal(http.ListenAndServe(constants.PORT_NO, r))
}

func StartWebSocketServer(conn *pgx.Conn) {
	server := socketio.NewServer(&engineio.Options{
		PingTimeout:  60 * time.Second,
		PingInterval: 30 * time.Second,
		Transports: []transport.Transport{
			&polling.Transport{
				Client: &http.Client{
					Timeout: time.Minute,
				},
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
			},
			&websocket.Transport{
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
			},
		},
	})

	route.SocketEvents(server, conn)

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
		defer server.Close()
	}()

	http.Handle("/socket.io/", server)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
