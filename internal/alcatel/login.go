package alcatel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/alcatel-link-zone/internal/alcatel/request_and_response"
	"github.com/alcatel-link-zone/internal/alcatel/request_and_response/request"
	"github.com/alcatel-link-zone/internal/alcatel/request_and_response/response"
	"io/ioutil"
	"net/http"
)

type Login struct {
	token int
}

func (a Login) Login( password string ) ( *response.LoginResponse,error ) {
	requestStruct := request.NewLoginRequest( "admin", password )
	loginRequestData, err := json.Marshal( requestStruct )
	if err != nil {
		return nil, err
	}

	resp, err := http.Post( requestStruct.GetUrl(BaseURL), "application/json", bytes.NewBuffer( loginRequestData ) )
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
	if !loginResp.IsOk() {
		return nil, loginResp.Err
	}

	a.token = loginResp.GetToken()
	return loginResp, nil
}

func (a Login) Logout() ( bool, error ) {
	r := request_and_response.Request{
		JsonRPC: "2.0",
		Method: "Logout",
		Id: "1.2",
	}
	logoutBytes, err := json.Marshal( r )
	if err != nil {
		return false, err
	}

	resp, err := http.Post( BaseURL, "text", bytes.NewBuffer( logoutBytes ) )
	if err != nil{
		return false, err
	}

	defer resp.Body.Close();
	respBytes,err := ioutil.ReadAll( resp.Body )
	if err != nil {
		return false, err
	}

	logoutResponse := &request_and_response.Response{}
	err = json.Unmarshal( respBytes, logoutResponse )
	if err != nil{
		return false, err
	}
	if logoutResponse.IsOk(){
		return true, nil
	}

	return false, logoutResponse.Err
}

func (a Login) GetLoginState() (*response.LoginStateResponse, error ){

	r := &request_and_response.Request{
		JsonRPC: "2.0",
		Method: "GetLoginState",
		Id: "1.3",
	}
	requestBytes, _ := json.Marshal(r)

	resp, err := http.Post( fmt.Sprintf("%s?api=%s", BaseURL, r.Method ), "text", bytes.NewBuffer( requestBytes ) )
	if err != nil {
		return &response.LoginStateResponse{}, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	loginStateResponse := &response.LoginStateResponse{}
	err = json.Unmarshal( respBytes, loginStateResponse )
	if err != nil {
		return nil, err
	}
	return loginStateResponse, nil
}