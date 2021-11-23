package alcatel

import (
	"bytes"
	"encoding/json"
	"github.com/alcatel-link-zone/internal/alcatel/request_and_response"
	"github.com/alcatel-link-zone/internal/alcatel/request_and_response/response"
	"io/ioutil"
	"net/http"
)

type System struct {

}

func (s *System) GetSystemStatus() (*response.GetSystemStatusResponse, error){

	requestStruct := request_and_response.Request{
		Method:  "GetSystemStatus",
		Id:      "13.4",
		JsonRPC: "2.0",
	}

	requestBytes, _ := json.Marshal( requestStruct )

	req, err := http.Post(requestStruct.GetUrl(BaseURL), "application/json", bytes.NewBuffer( requestBytes ) )
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	responseStruct := &response.GetSystemStatusResponse{}
	responseBytes, err := ioutil.ReadAll( req.Body )
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal( responseBytes, responseStruct )
	if err != nil {
		return nil, err
	}
	if responseStruct.IsOk(){
		return responseStruct, nil
	}
	return nil, responseStruct.Err
}