package response

import "github.com/alcatel-link-zone/internal/alcatel/request_and_response"

type UssdSentResponse struct {
	request_and_response.Response
	Result UssdSentResponseResult `json:"result"`
}

type UssdSentResponseResult struct {
	UssdType int `json:"UssdType"`
	SendState int `json:"SendState"`
	UssdContent string `json:"UssdContent"`
	UssdContentLen int `json:"UssdContentLen"`
}

