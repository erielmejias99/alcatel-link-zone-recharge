package request

import (
	"github.com/alcatel-link-zone/internal/alcatel/request_and_response"
)

type LoginRequestParam struct{
	Username string `json:"UserName"`
	Password string `json:"Password"`
}


type LoginRequest struct{
	Params LoginRequestParam `json:"params"`
	request_and_response.Request
}

func NewLoginRequest( username string, password string ) *LoginRequest {
	request := &LoginRequest{
		Params: LoginRequestParam{
			Username: username,
			Password: password,
		},
	}
	request.Method = "Login"
	request.JsonRPC = "2.0"
	request.Id = "1.1"
	return request
}