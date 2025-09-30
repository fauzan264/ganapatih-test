package response

type RegisterResponse struct {
	ID					int				`json:"id"`
	Username 		string		`json:"username"`
}

type LoginResponse struct {
	Token				string		`json:"token"`
}

type SessionResponse struct {
	ID					int 			`json:"id"`
	Username 		string		`json:"username"`
}

type UserResponse struct {
	ID					int 			`json:"id"`
	Username 		string		`json:"username"`
}