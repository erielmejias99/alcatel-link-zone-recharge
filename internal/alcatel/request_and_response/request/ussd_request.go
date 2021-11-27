package request

import "github.com/alcatel-link-zone/internal/alcatel/request_and_response"

type UssdTypes int

const (
	Default UssdTypes = 1
	Idk UssdTypes = 2
)

type SendUssdRequestParam struct {
	UssdContent string    `json:"UssdContent"`
	UssdType    UssdTypes `json:"UssdType"`
}

type SendUssdRequest struct {
	Params * SendUssdRequestParam `json:"params"`
	*request_and_response.Request
}

func NewSendUssdRequest( content string, ussdType UssdTypes ) *SendUssdRequest {
	request := &SendUssdRequest{
		Params: &SendUssdRequestParam{
			content,
			ussdType,
		},
	}
	request.Request = request_and_response.NewRequest( "2.0", "SendUSSD", "8.1" )
	return request
}
