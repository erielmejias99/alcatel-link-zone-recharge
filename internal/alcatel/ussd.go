package alcatel

import (
	"bytes"
	"encoding/json"
	"github.com/alcatel-link-zone/internal/alcatel/request_and_response"
	"github.com/alcatel-link-zone/internal/alcatel/request_and_response/request"
	"github.com/alcatel-link-zone/internal/alcatel/request_and_response/response"
	"io/ioutil"
	"net/http"
)

type Ussd struct {
	ussdType int
	ussdRunning bool
}


func (u Ussd) getUssdResponse() ( *response.UssdSentResponse, error ) {

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
	u.ussdRunning = true
	return getUssdSentResponse, nil
}

func (u Ussd) SetUssdEnd() error {
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
	u.ussdType = 0
	return nil
}

func (u Ussd) sendUSSD(code string ) ( *response.SendUssdResponse, error ) {

	requestStruct := request.NewSendUssdRequest( code, request.UssdTypes(u.ussdType) )
	requestBytes, err := json.Marshal( requestStruct )
	if err != nil {
		return nil, err
	}

	resp, err := http.Post( requestStruct.GetUrl(BaseURL), "application/json", bytes.NewBuffer( requestBytes ) )
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
