package response

import (
	"github.com/alcatel-link-zone/internal/alcatel/request_and_response"
)
type NetworkType int

const (
	NoService NetworkType = 0
	Net2G NetworkType = 1
	Net2G2 NetworkType = 2
	Net3G NetworkType = 3
	Net3G2 NetworkType = 4
	Net3GPlus NetworkType = 5
	Net3GPlus2 NetworkType = 6
	Net4G NetworkType = 7
	Net4GPlus NetworkType = 8
)

func (n NetworkType) Value() int {
	return int(n)
}

var NetWorkLiteral = []string{"NO_SERVICE", "2G", "2G", "3G", "3G", "3G", "3G+", "3G+", "4G", "4G+" }

func GetNetworkStringLiteral( networkType NetworkType ) string {
	if networkType.Value() >= 0 && networkType.Value() < len(NetWorkLiteral) {
		return NetWorkLiteral[ networkType ]
	}
	return NetWorkLiteral[0]
}

type GetSystemStatusResult struct {
	NetworkType      NetworkType `json:"NetworkType"`
	SignalStrength   int         `json:"SignalStrength"`
	WlanState        int         `json:"WlanState"`
	ConnectionStatus int         `json:"ConnectionStatus"`
	SmsState         int         `json:"SmsState"`
	NetworkName      string      `json:"NetworkName"`
}

type GetSystemStatusResponse struct {
	request_and_response.Response
	Result GetSystemStatusResult `json:"result"`
}