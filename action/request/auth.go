package request

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"crypto-system/internal/context"
)

func VerifyMD5(md5 string) map[string]interface{} {
	params := url.Values{}
	Url, err := url.Parse(context.App.Config.Server.URL("auth/verify"))

	context.App.Logger.Error(err)

	params.Set("md5", md5)

	Url.RawQuery = params.Encode()
	resp, err := http.Get(Url.String())

	context.App.Logger.Error(err)
	context.App.Logger.Log(Url.String())

	defer resp.Body.Close()

	var res Response
	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &res)

	if !res.Result {
		context.App.Logger.Error(errors.New("http request error"))
	}

	return res.Data
}

func GetKMSKey() string {

	Url, err := url.Parse(context.App.Config.Server.URL("auth/key"))
	context.App.Logger.Error(err)
	resp, err := http.Get(Url.String())
	context.App.Logger.Error(err)
	context.App.Logger.Log(Url.String())
	defer resp.Body.Close()

	var res Response
	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &res)

	if !res.Result {
		context.App.Logger.Error(errors.New("http request error"))
	}

	return res.Data["key"].(string)

}

func DecryptFile(cid string) string {

	urlValues := url.Values{}
	urlValues.Add("cid", cid)
	urlValues.Add("publicKey", string(context.App.Crypto.PublicKey))
	resp, err := http.PostForm(context.App.Config.Server.URL("auth/decrypt"), urlValues)

	context.App.Logger.Error(err)
	context.App.Logger.Log(context.App.Config.Server.URL("auth/decrypt"))
	body, _ := ioutil.ReadAll(resp.Body)
	var res Response
	_ = json.Unmarshal(body, &res)

	if !res.Result {
		context.App.Logger.Error(errors.New("http request error"))
	}
	return res.Data["key"].(string)
}

func UploadFile(md *MateData) bool {

	urlValues := url.Values{}
	urlValues.Add("CID", md.CID)
	urlValues.Add("Key", md.Key)
	urlValues.Add("MD5", md.MD5)
	urlValues.Add("Name", md.Name)
	urlValues.Add("Encrypt", strconv.FormatBool(md.Encrypt))
	urlValues.Add("Size", strconv.Itoa(int(md.Size)))
	resp, err := http.PostForm(context.App.Config.Server.URL("auth/upload"), urlValues)
	context.App.Logger.Error(err)
	context.App.Logger.Log(context.App.Config.Server.URL("auth/upload"))

	body, _ := ioutil.ReadAll(resp.Body)
	var res Response
	_ = json.Unmarshal(body, &res)

	if !res.Result {
		context.App.Logger.Error(errors.New("http request error"))
	}
	return res.Result
}

func GetList() []MateData {

	Url, err := url.Parse(context.App.Config.Server.URL("auth/list"))
	context.App.Logger.Error(err)
	resp, err := http.Get(Url.String())
	context.App.Logger.Error(err)
	context.App.Logger.Log(Url.String())
	defer resp.Body.Close()

	var res Response
	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &res)

	if !res.Result {
		context.App.Logger.Error(errors.New("http request error"))
	}

	data, _ := res.Data["list"].(string)

	var files []MateData
	json.Unmarshal([]byte(data), &files)

	return files
}
