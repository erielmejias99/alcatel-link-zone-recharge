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

type Connection struct {

}

func (c Connection) GetConnectionState() ( *response.GetConnectionStatusResponse, error ) {

	requestStruct := request_and_response.Request{ Method: "GetConnectionState", JsonRPC: "2.0", Id: "3.1" }
	requestBytes, err := json.Marshal( requestStruct )
	if err != nil {
		return nil, err
	}

	resp, err := http.Post( BaseURL, "application/json", bytes.NewBuffer( requestBytes ) )
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	getConnectionStatus := &response.GetConnectionStatusResponse{}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal( respBytes, getConnectionStatus )
	if err != nil {
		return nil, err
	}
	if !getConnectionStatus.IsOk() {
		return nil, getConnectionStatus.Err
	}
	return getConnectionStatus, nil
}

func (c Connection) disconnect() ( bool, error ) {

	requestStruct := request_and_response.Request{ Method: "DisConnect", JsonRPC: "2.0", Id: "3.3" }
	requestBytes, err := json.Marshal( requestStruct )
	if err != nil {
		return false, err
	}

	resp, err := http.Post( requestStruct.GetUrl(BaseURL), "application/json", bytes.NewBuffer( requestBytes ) )
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	disconnectResp := &request_and_response.Response{}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	err = json.Unmarshal( respBytes, disconnectResp )
	if err != nil {
		return false, err
	}
	if !disconnectResp.IsOk() {
		return false, disconnectResp.Err
	}
	return true, nil
}

func (c Connection) connect() ( bool, error ) {

	requestStruct := request_and_response.Request{ Method: "Connect", JsonRPC: "2.0", Id: "3.2" }
	requestBytes, err := json.Marshal( requestStruct )
	if err != nil {
		return false, err
	}

	resp, err := http.Post( requestStruct.GetUrl(BaseURL), "application/json", bytes.NewBuffer( requestBytes ) )
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	connectResp := &request_and_response.Response{}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	err = json.Unmarshal( respBytes, connectResp )
	if err != nil {
		return false, err
	}
	if !connectResp.IsOk() {
		return false, connectResp.Err
	}
	return true, nil
}

func (c Connection) setNetworkType(networkType request.SetNetworkType ) error {


	requestStruct := request.NewNetworkSettingsRequest( networkType )
	requestBytes, err := json.Marshal( requestStruct )
	if err != nil {
		return err
	}

	resp, err := http.Post( requestStruct.GetUrl(BaseURL), "application/json", bytes.NewBuffer( requestBytes ) )
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	setNetworkResponse := &request_and_response.Response{}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal( respBytes, setNetworkResponse )
	if err != nil {
		return err
	}
	if !setNetworkResponse.IsOk() {
		return setNetworkResponse.Err
	}
	return nil
}