package request

import (
	"github.com/alcatel-link-zone/internal/alcatel/request_and_response"
)

type SetNetworkType int

const (
	Automatic SetNetworkType = 0
	Net2G SetNetworkType = 1
	Net3G SetNetworkType = 2
	Net4G SetNetworkType = 3
)

type SetNetworkSettingsParams struct {
	NetworkMode SetNetworkType `json:""`
	NetselectionMode int `json:""`
}

type SetNetworkSettingsRequest struct {
	request_and_response.Request
	Params *SetNetworkSettingsParams `json:"params"`
}

func NewNetworkSettingsRequest( networkType SetNetworkType ) *SetNetworkSettingsRequest {
	request := &SetNetworkSettingsRequest{
		Params: &SetNetworkSettingsParams{
			NetworkMode: networkType,
			NetselectionMode: 0,
		},
	}
	request.JsonRPC = "2.0"
	request.Method = "SetNetworkSettings"
	request.Id = "4.7"

	return request
}