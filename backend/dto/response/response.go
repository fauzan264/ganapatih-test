package response

type Response struct {
	Message 	string 				`json:"message,omitempty"`
	Data			any						`json:"data,omitempty"`
	Token			string				`json:"token,omitempty"`
}