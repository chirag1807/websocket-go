package route

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/chirag1807/websocket-go/api/model/request"
	"github.com/chirag1807/websocket-go/api/repository"
	socketio "github.com/googollee/go-socket.io"
	"github.com/jackc/pgx/v5"
)

func SocketEvents(server *socketio.Server, conn *pgx.Conn) {
	chatRepository := repository.NewChatRepo(conn)

	server.OnConnect("/", func(c socketio.Conn) error {
		log.Println("Connection Made Successfully.", c.ID())
		return nil
	})

	server.OnEvent("/", "message", func(s socketio.Conn, msg map[string]interface{}) {
		log.Println("message:", msg)
		// s.Emit("") => this will not work in postman, only work in client side (in react or something)
		server.BroadcastToNamespace("/", msg["receiver"].(string), msg)
		msg["sender"], _ = strconv.ParseInt(msg["sender"].(string), 10, 64)
		msg["receiver"], _ = strconv.ParseInt(msg["receiver"].(string), 10, 64)
		response, _ := json.Marshal(msg)
		var chat request.Chat
		json.Unmarshal(response, &chat)
		chatRepository.AddToSingleChat(chat)
	})

	server.OnEvent("/", "room-chat", func (s socketio.Conn, msg map[string]interface{})  {
		log.Println("message:", msg)
		server.BroadcastToRoom("/", "Room1", "room-chat", msg)
		msg["roomid"], _ = strconv.ParseInt(msg["roomid"].(string), 10, 64)
		msg["sender"], _ = strconv.ParseInt(msg["sender"].(string), 10, 64)
		response, _ := json.Marshal(msg)
		var chat request.RoomChat
		json.Unmarshal(response, &chat)
		chatRepository.AddToRoomChat(chat)
	})

	server.OnEvent("/", "join-room", func (s socketio.Conn, roomName string)  {
		server.JoinRoom("/", roomName, s)
	})

	server.OnEvent("/", "leave-room", func (s socketio.Conn, roomName string)  {
		server.LeaveRoom("/", roomName, s)
	})

	server.OnError("/", func(c socketio.Conn, err error) {
		log.Fatal(err)
	})

	server.OnDisconnect("/", func(c socketio.Conn, s string) {
		log.Println("disconnected:", s)
	})
}
