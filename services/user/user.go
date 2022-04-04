package user

import "encoding/json"

// LoginResponse struct
type LoginResponse struct {
	Data DataLoginResponse
}

type DataLoginResponse struct {
	Ticket string
}

type User struct {
	Id        string `json:"id"`
	UserName  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

//Decode Json Response
func (l *LoginResponse) Decode(text []byte) error {
	return json.Unmarshal(text, l)
}
