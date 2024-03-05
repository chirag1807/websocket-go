package route

import (
	"github.com/chirag1807/websocket-go/api/controller"
	"github.com/chirag1807/websocket-go/api/repository"
	"github.com/chirag1807/websocket-go/api/service"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func Routes(conn *pgx.Conn) *chi.Mux {
	r := chi.NewRouter()

	userRepository := repository.NewAuthRepo(conn)
	userService := service.NewAuthService(userRepository)
	authController := controller.NewAuthController(userService)

	chatRepository := repository.NewChatRepo(conn)
	chatService := service.NewChatService(chatRepository)
	chatController := controller.NewChatController(chatService)

	r.Route("/api/user", func(r chi.Router) {
		r.Post("/registration", authController.UserRegistration)
		r.Post("/login", authController.UserLogin)
		r.Post("/add-to-chat", chatController.AddToSingleChat)
		r.Get("/chat-history/{ID}", chatController.GetAllChats)
		r.Get("/add-to-room-chat", chatController.AddToRoomChat)
	})

	r.Route("/api/room", func(r chi.Router) {
		r.Post("/create-room", chatController.CreateRoom)
		r.Put("/add-member-to-room", chatController.AddMemberToRoom)
		r.Put("/remove-member-to-room", chatController.RemoveMemberFromRoom)
	})

	return r
}
