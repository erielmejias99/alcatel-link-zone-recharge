package alcatel

type Request struct {
	JsonRPC string `json:"jsonrpc"`
	Method string `json:"method"`
	Id string `json:"id"`
}