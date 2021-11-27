package response

import "github.com/alcatel-link-zone/internal/alcatel/request_and_response"

type NetworkStatus int

const (
	Disconnected NetworkStatus = 1
	Connected NetworkStatus = 2
)

type GetConnectionStatusResult struct {
	ConnectionStatus NetworkStatus `json:"ConnectionStatus"`
	IPv4Adrress      string        `json:"IPv4Adrress"`
}

type GetConnectionStatusResponse struct {
	request_and_response.Response
	Result GetConnectionStatusResult `json:"result"`
}