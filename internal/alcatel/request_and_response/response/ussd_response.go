package response

import "github.com/alcatel-link-zone/internal/alcatel/request_and_response"

type SendUssdResponseResult struct {

}

type SendUssdResponse struct {
	request_and_response.Response
	Result  SendUssdResponseResult `json:"result"`
}

