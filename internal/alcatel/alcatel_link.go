package alcatel

import (
	"bytes"
	"encoding/json"
	"github.com/alcatel-link-zone/internal/alcatel/request_and_response/request"
	"github.com/alcatel-link-zone/internal/alcatel/request_and_response/response"
	"io/ioutil"
	"net/http"
)

type Alcatel struct {
	token int
}

func (a Alcatel) Login( password string ) ( *response.LoginResponse,error ) {
	loginRequestData, err := json.Marshal( request.NewLoginRequest( "admin", password ) )
	if err != nil {
		return nil, err
	}

	resp, err := http.Post( BaseURL, "application/json", bytes.NewBuffer( loginRequestData ) )
	if err != nil{
		return nil, err
	}

	defer resp.Body.Close();
	respBytes,err := ioutil.ReadAll( resp.Body )

	loginResp := &response.LoginResponse{}
	err = json.Unmarshal( respBytes, loginResp )
	if err != nil{
		return nil, err
	}
	if loginResp.Err != nil {
		return nil, loginResp.Err
	}

	a.token = loginResp.GetToken()
	return loginResp, nil
}


