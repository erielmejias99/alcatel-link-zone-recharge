package alcatel

import (
	"bytes"
	"encoding/json"
	"github.com/alcatel-link-zone/internal/alcatel/request_and_response"
	"github.com/alcatel-link-zone/internal/alcatel/request_and_response/request"
	"github.com/alcatel-link-zone/internal/alcatel/request_and_response/response"
	"io/ioutil"
	"net/http"
	"time"
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
	if ussdResponse.Err != nil {
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

	// Send code to the Network
	sendUssdResponse, err := a.sendUSSD( code, codeType )
	if _, ok := err.(*request_and_response.Error); ok && sendUssdResponse != nil {
		return nil, sendUssdResponse.Err
	}
	if err != nil {
		return nil, err
	}

	time.Sleep( 5 * time.Second )

	// Get the response from the Network
	getUssdResponse, err := a.getUssdResponse()
	if _, ok := err.(*request_and_response.Error); ok && getUssdResponse != nil {
		return nil, sendUssdResponse.Err
	}
	if err != nil {
		return nil, err
	}

	return getUssdResponse, nil


}