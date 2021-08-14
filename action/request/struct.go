package request

type Response struct {
	Result bool                   `json:"result"`
	Data   map[string]interface{} `json:"data"`
	Msg    string                 `json:"msg"`
}

type MateData struct {
	CID  string
	MD5  string
	Key  string
	Name string
	Size int64
}

type ResultDecrypt struct {
	Key string `json:"key"`
}

type ResultKMSkey struct {
	Key string `json:"key"`
}

type ResultMD5 struct {
	HasFile bool   `json:"hasFile"`
	CID     string `json:"CID"`
	Key     string `json:"key"`
	Md5     string
}
