package request

import "github.com/alcatel-link-zone/internal/alcatel"

type LoginRequestParam struct{
	username string `json:"UserName"`
	password string `json:"Password"`
}


type LoginRequest struct{
	Params LoginRequestParam `json:"params"`
	alcatel.Request
}

func NewLoginRequest( username string, password string ) *LoginRequest {
	request := &LoginRequest{
		Params: LoginRequestParam{
			username: username,
			password: password,
		},
	}
	request.Method = "Login"
	request.JsonRPC = "2.0"
	request.Id = "1.1"
	return request
}