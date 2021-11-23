package request_and_response

import "fmt"

type Request struct {
	JsonRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Id      string `json:"id"`
}

func NewRequest( jsonRPC, method, id string ) *Request{
	return &Request{
		JsonRPC: jsonRPC,
		Method: method,
		Id: id,
	}
}


type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e Error) Error() string  {
	return fmt.Sprintf("{ code: \"%s\", message:\"%s\" }", e.Code, e.Message )
}

type Response struct {
	JsonRPC string `json:"jsonrpc"`
	Id      string `json:"id"`
	Err     *Error `json:"error"`
}

func (r Response) IsOk() bool {
	return r.Err == nil
}

func (r Request) GetUrl(base string) string {
	return fmt.Sprintf( "%s?api=%s", base, r.Method )
}