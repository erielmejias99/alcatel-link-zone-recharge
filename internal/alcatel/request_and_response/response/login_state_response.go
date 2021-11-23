package response

import "github.com/alcatel-link-zone/internal/alcatel/request_and_response"

type LoginState int

const (
	LoggedOut LoginState = 0
	LoggedIn LoginState = 1
	LogInTimesUseOut LoginState = 2
	LogInStatePasswordWrong LoginState = 3
)

type LoginStateResponseResult struct {
	State               LoginState `json:"State"`
	LoginRemainingTimes int `json:"LoginRemainingTimes"`
	LockedRemainingTime int `json:"LockedRemainingTime"`
}

type LoginStateResponse struct {
	request_and_response.Response
	Result LoginStateResponseResult `json:"result"`
}