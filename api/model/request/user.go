package request

type User struct {
	Name string `json:"name"`
	Bio string `json:"bio"`
	Email string `json:"email"`
	Password string `json:"password"`
}
