package request

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

type Response struct {
	Result bool                   `json:"result"`
	Data   map[string]interface{} `json:"data"`
	Msg    string                 `json:"msg"`
}

type MateData struct {
	CID     string
	Key     string `string:"none"`
	MD5     string `string:"none"`
	Name    string
	Size    int64
	Encrypt bool      `string:"none"`
	Ctime   time.Time `string:"none"`
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

func (m *MateData) ToString() string {
	return ToString(m)
}

func ToString(item interface{}) string {

	v := reflect.ValueOf(item).Elem()
	k := reflect.TypeOf(item).Elem()
	var buildKeys strings.Builder

	buildValues := []interface{}{}
	for i := 0; i < k.NumField(); i++ {

		f := k.Field(i)
		if f.Tag.Get("string") == "none" {
			continue
		}
		buildKeys.WriteString(f.Name + " : %v ")
		val := v.Field(i).Interface()
		buildValues = append(buildValues, val)
		if i != k.NumField()-1 {
			buildKeys.WriteString("\t")
		}
	}

	return fmt.Sprintf(buildKeys.String(), buildValues...)
}
