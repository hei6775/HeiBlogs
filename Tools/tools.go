package Tools

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
)

func MakeStr(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		prefix := k + "="
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(prefix)
			buf.WriteString(v)
		}
	}
	return buf.String()
}

func MakeStr2(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		prefix := k + "="
		for _, v := range vs {
			//			if buf.Len() > 0 {
			//				buf.WriteByte('&')
			//			}
			buf.WriteString(prefix)
			buf.WriteString(v)
		}
	}
	return buf.String()
}

func MakeStr3(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		prefix := k + "|"
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('|')
			}
			buf.WriteString(prefix)
			buf.WriteString(v)
		}
	}
	return buf.String()
}

//Get参数
func GetPara(Form url.Values, Name string) string {

	if len(Form[Name]) > 0 {
		return Form[Name][0]
	}

	return ""

}

//拼接参数
func JoinPara(Form url.Values) string {

	ValStr := url.Values{}
	for id, obj := range Form {

		//log.Release("JoinPara id:%v obj:%v", id, obj)

		if id == "sign" {
			continue
		}

		if len(obj) == 0 {
			continue
		}

		ValStr.Set(id, obj[0])
	}

	SignStr := MakeStr(ValStr)

	return SignStr

}

//拼接数值
func MakeVal(v url.Values, para string) string {
	if v == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]

		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteString(para)
			}

			buf.WriteString(v)
		}
	}
	return buf.String()
}

func Md5(data string) string {
	md5 := md5.New()
	md5.Write([]byte(data))
	md5Data := md5.Sum([]byte(""))
	return hex.EncodeToString(md5Data)
}

func Hmac(key, data string) []byte {
	hmac := hmac.New(sha1.New, []byte(key))
	hmac.Write([]byte(data))
	return hmac.Sum([]byte(""))
}

func Sha1(data string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(data))
	return hex.EncodeToString(sha1.Sum([]byte("")))
}

//HttpGet请求
func HttpGet(UrlStr string) string {

	resp, err := http.Get(UrlStr)
	if err != nil {
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return string(body)

}
