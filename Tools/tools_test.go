package Tools

import (
	"fmt"
	"net/url"
	"testing"
)

var v = url.Values{"a": []string{"a1", "a2"}, "b": []string{"b1", "b2"}}

func TestMakeStr(t *testing.T) {
	fmt.Println(MakeStr(v))
}

func TestMakeStr2(t *testing.T) {

}

func TestMakeStr3(t *testing.T) {

}

func TestGetPara(t *testing.T) {

}

func TestJoinPara(t *testing.T) {

}

func TestMakeVal(t *testing.T) {

}

func TestMd5(t *testing.T) {

}

func TestHmac(t *testing.T) {

}

func TestSha1(t *testing.T) {

}

func TestHttpGet(t *testing.T) {

}
