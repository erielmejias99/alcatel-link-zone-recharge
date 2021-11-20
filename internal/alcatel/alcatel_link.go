package alcatel

import (
	"bytes"
	"encoding/json"
	"github.com/alcatel-link-zone/internal/alcatel/request"
	"github.com/alcatel-link-zone/internal/alcatel/response"
	"net/http"
)

type Alcatel struct {
	token string
}

func (a Alcatel) Login( password string ) ( *response.LoginResponse,error ) {
	loginRequestData, err := json.Marshal( request.NewLoginRequest( "admin", password ) )
	if err != nil {
		return nil, err
	}

	resp, err := http.Post( BaseURL + "", "application/json", bytes.NewBuffer( loginRequestData ) )
	if err != nil{
		return nil, err
	}

	defer resp.Body.Close();

	loginResp := response.LoginResponse{}
	json.NewDecoder( resp.Body ).Decode(resp)

	a.token = loginResp.Token
	return &loginResp, nil
}


