package request

type Response struct {
	Result bool                   `json:"result"`
	Data   map[string]interface{} `json:"data"`
	Msg    string                 `json:"msg"`
}
