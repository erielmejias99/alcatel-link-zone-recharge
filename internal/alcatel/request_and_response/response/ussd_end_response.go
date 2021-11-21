package response

import "github.com/alcatel-link-zone/internal/alcatel/request_and_response"

type SetUssdEndResponse struct {
	request_and_response.Response
	Result SetUssdEndResponseResult `json:"result"`
}

type SetUssdEndResponseResult struct {
		
}


