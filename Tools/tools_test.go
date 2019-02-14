package Tools

import (
	"fmt"
	"net/url"
	"testing"
)

var v = url.Values{"a": []string{"a1", "a2"}, "b": []string{"b1", "b2"}}

func TestMakeStr(t *testing.T) {
	fmt.Println(MakeStr(v))
	//a=a1&a=a2&b=b1&b=b2
}

func TestMakeStr2(t *testing.T) {
	fmt.Println(MakeStr2(v))
	//a=a1a=a2b=b1b=b2
}

func TestMakeStr3(t *testing.T) {
	fmt.Println(MakeStr3(v))
	//a|a1|a|a2|b|b1|b|b2
}

func TestGetPara(t *testing.T) {
	fmt.Println(GetPara(v, "a"))
	//a1
}

func TestJoinPara(t *testing.T) {
	fmt.Println(JoinPara(v))
	//a=a1&b=b1
}

func TestMakeVal(t *testing.T) {
	fmt.Println(MakeVal(v, "sss"))
	//a1sssa2sssb1sssb2
}

func TestMd5(t *testing.T) {
	fmt.Println(Md5("123456"))
	//e10adc3949ba59abbe56e057f20f883e
}

func TestHmac(t *testing.T) {
	fmt.Println(Hmac("123", "123456"))
	//[108 64 93 20 4 103 65 152 28 244 203 135 151 253 190 198 239 245 237 178]
}

func TestSha1(t *testing.T) {
	fmt.Println(Sha1("123456"))
	//7c4a8d09ca3762af61e59520943dc26494f8941b
}

func TestHttpGet(t *testing.T) {

}
