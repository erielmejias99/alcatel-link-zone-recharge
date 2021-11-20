package response

import (
	"fmt"
	"github.com/alcatel-link-zone/internal/alcatel/request_and_response"
)

type LoginResponseBody struct {
	Token int `json:"token"`
}


type LoginResponse struct {
	request_and_response.Response
	Result LoginResponseBody `json:"result"`
}

func (l LoginResponse) String() string {
	return fmt.Sprintf("{ token: %d }", l.Result.Token )
}

func (l LoginResponse) GetToken() int {
	return l.Result.Token
}