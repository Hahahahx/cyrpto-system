package request

type Response struct {
	Result bool              `json:"result"`
	Data   map[string]string `json:"data"`
	Msg    string            `json:"msg"`
}
