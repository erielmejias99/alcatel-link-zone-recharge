package request

import "github.com/alcatel-link-zone/internal/alcatel/request_and_response"

var SetUssdEndRequest = request_and_response.Request{
	JsonRPC: "2.0",
	Method: "SetUSSDEnd",
	Id: "8.3",
}


var GetUssdSentRequest = request_and_response.Request{
	JsonRPC: "2.0",
	Method: "GetUSSDSendResult",
	Id: "8.2",
}


