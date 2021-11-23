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
	"time"
)

type Alcatel struct {
	token int
	ussdType request.UssdTypes
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
	if !loginResp.IsOk() {
		return nil, loginResp.Err
	}

	a.token = loginResp.GetToken()
	return loginResp, nil
}

func (a Alcatel) Logout() ( bool, error ) {
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

func (a Alcatel) GetLoginState() (*response.LoginStateResponse, error ){

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

func (a Alcatel) sendUSSD(code string, codeType request.UssdTypes ) ( *response.SendUssdResponse, error ) {

	requestBytes, err := json.Marshal( request.NewSendUssdRequest( code, codeType ) )
	if err != nil {
		return nil, err
	}

	resp, err := http.Post( BaseURL, "application/json", bytes.NewBuffer( requestBytes ) )
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ussdResponse := &response.SendUssdResponse{}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal( respBytes, ussdResponse )
	if err != nil {
		return nil, err
	}
	if !ussdResponse.IsOk() {
		return nil, ussdResponse.Err
	}

	return ussdResponse, nil
}

func (a Alcatel) getUssdResponse() ( *response.UssdSentResponse, error ) {

	requestBytes, err := json.Marshal( request.GetUssdSentRequest )
	if err != nil {
		return nil, err
	}

	resp, err := http.Post( BaseURL, "application/json", bytes.NewBuffer( requestBytes ) )
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	getUssdSentResponse := &response.UssdSentResponse{}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal( respBytes, getUssdSentResponse )
	if err != nil {
		return nil, err
	}
	if getUssdSentResponse.Err != nil {
		return nil, getUssdSentResponse.Err
	}
	return getUssdSentResponse, nil
}

func (a Alcatel) SendUssdCode( code string, codeType request.UssdTypes ) ( *response.UssdSentResponse,error ) {

	// Set NetWork mode to 3G. Etecsa doesn't allow ussd code in 4G Mode


	// Send code to the Network
	sendUssdResponse, err := a.sendUSSD( code, codeType )
	if sendUssdResponse != nil && !sendUssdResponse.IsOk(){
		return nil, sendUssdResponse.Err
	}
	if err != nil {
		return nil, err
	}

	ussdState := 0
	tries := 1
	getUssdResponse, err := &response.UssdSentResponse{}, nil

	for ;ussdState != 2 && tries <= 3; {
		getUssdResponse, err = a.getUssdResponse()
		if getUssdResponse != nil{
			ussdState = getUssdResponse.Result.SendState
		}
		time.Sleep( 5 * time.Second )
		tries ++
	}

	// Get the response from the Network
	if _, ok := err.(*request_and_response.Error); ok && getUssdResponse != nil {
		return nil, sendUssdResponse.Err
	}
	if err != nil {
		return nil, err
	}
	return getUssdResponse, nil

}

func (a Alcatel) setUssdEnd() error {
	r := request_and_response.Request{
		JsonRPC: "2.0",
		Method: "SetUSSDEnd",
		Id: "8.3",
	}
	requestBytes, err := json.Marshal( r )
	if err != nil {
		return err
	}

	resp, err := http.Post( r.GetUrl(BaseURL),"application/json", bytes.NewBuffer(requestBytes) )
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll( resp.Body )
	if err != nil {
		return err
	}
	setUssdEndResponse := &request_and_response.Response{}
	err = json.Unmarshal( respBytes, setUssdEndResponse )
	if err != nil {
		return err
	}
	if !setUssdEndResponse.IsOk(){
		return setUssdEndResponse.Err
	}
	return nil
}