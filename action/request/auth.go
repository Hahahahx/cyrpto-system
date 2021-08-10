package request

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"crypto-system/internal/context"
)

func VerifyMD5(c *context.Request, md5 string) ResultMD5 {
	params := url.Values{}
	Url, err := url.Parse(c.App.Config.Server.URL("auth/verify"))

	c.App.Logger.Error(err)

	params.Set("md5", md5)

	Url.RawQuery = params.Encode()
	resp, err := http.Get(Url.String())

	c.App.Logger.Log(Url.String())

	defer resp.Body.Close()

	var res Response
	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &res)

	if !res.Result {
		c.App.Logger.Error(errors.New("http request error"))
	}

	var data ResultMD5
	_ = json.Unmarshal(body, &data)

	data.Md5 = md5

	return data
}

func GetKMSKey(c *context.Request) string {

	Url, err := url.Parse(c.App.Config.Server.URL("auth/key"))
	c.App.Logger.Error(err)
	resp, err := http.Get(Url.String())
	c.App.Logger.Error(err)
	c.App.Logger.Log(Url.String())
	defer resp.Body.Close()

	var res Response
	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &res)

	if !res.Result {
		c.App.Logger.Error(errors.New("http request error"))
	}

	return res.Data["key"]

}

func DecryptFile(c *context.Request, cid string) string {

	urlValues := url.Values{}
	urlValues.Add("cid", cid)
	urlValues.Add("publicKey", string(c.App.Crypto.PublicKey))
	resp, err := http.PostForm(c.App.Config.Server.URL("auth/decrypt"), urlValues)

	c.App.Logger.Error(err)
	c.App.Logger.Log(c.App.Config.Server.URL("auth/decrypt"))
	body, _ := ioutil.ReadAll(resp.Body)
	var res Response
	_ = json.Unmarshal(body, &res)

	if !res.Result {
		c.App.Logger.Error(errors.New("http request error"))
	}
	return res.Data["key"]
}

func UploadFile(c *context.Request, md *MateData) bool {

	urlValues := url.Values{}
	urlValues.Add("CID", md.CID)
	urlValues.Add("FingerPrint", md.MD5)
	urlValues.Add("Key", md.Key)
	resp, err := http.PostForm(c.App.Config.Server.URL("auth/upload"), urlValues)
	c.App.Logger.Error(err)
	c.App.Logger.Log(c.App.Config.Server.URL("auth/upload"))

	body, _ := ioutil.ReadAll(resp.Body)
	var res Response
	_ = json.Unmarshal(body, &res)

	if !res.Result {
		c.App.Logger.Error(errors.New("http request error"))
	}
	return res.Result
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
