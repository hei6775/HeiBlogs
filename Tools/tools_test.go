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

func TestHttpPost(t *testing.T) {
	tarurl := "http://www.jyhdyx.com/jxy_pay/xiao7"

	game_key := "8a584ffe27b58e5bc3c4dc3a7fe418ba"
	scret := "E19C5E9825FE965A060985CA898ACC68"

	game_orderid := "5c6524ac48d88d1e16a1a43b"
	game_price := "6.00"
	user_id := "1"
	str := fmt.Sprintf("game_key=%v&game_orderid=%v&game_price=%v&user_id=%v%v",
		game_key, game_orderid, game_price, user_id, scret)
	encryp_data := Md5(str)

	extends_data := "zhi_fu_tou_chuan_can_shu"
	game_area := "1"
	game_group := "1"
	game_role_id := "1"
	subject := "1"
	xiao7_goid := "1"
	str2 := fmt.Sprintf("encryp_data=%v&extends_data=%v&game_area=%v&game_group=%v&game_orderid=%v&game_price=%v&game_role_id=%v&subject=%v&user_id=%v&xiao7_goid=%v",
		encryp_data, extends_data, game_area, game_group, game_orderid, game_price, game_role_id, subject, user_id, xiao7_goid)
	str3 := fmt.Sprintf("%v%v", str2, scret)
	sign_data := Md5(str3)
	fmt.Printf("encryp_data : [%v] sign_data : [%v] \n", encryp_data, sign_data)
	v := url.Values{}
	v.Set("encryp_data", encryp_data)
	v.Set("extends_data", extends_data)
	v.Set("game_area", game_area)
	v.Set("game_group", game_group)
	v.Set("game_orderid", game_orderid)
	v.Set("game_price", game_price)
	v.Set("game_role_id", game_role_id)
	v.Set("subject", subject)
	v.Set("user_id", user_id)
	v.Set("xiao7_goid", xiao7_goid)
	v.Set("sign_data", sign_data)
	str4 := MakeStr(v)
	result := HttpPost(tarurl, str4)
	fmt.Println(result)
}


func TestHttpPost22(t *testing.T) {
	tarurl := "http://act1.lianwifi.com/h5/order/create"

	a := "game_id=xjy&open_id=873c5e9943c83d5a75478706a3520bc0"
	b := "&out_reserved="+""+"&out_trade_no=5c7f9ee148d88d041588e5f8"
	c := "&reserved=146637578130570603710020190306181959"
	d := "&sign=c39b3afd1a7de19e7cae2794e10ab9f1"+"&sign_type=md5"
	f := "&subject=60元宝&total_fee=600&_input_charset=UTF-8"


	H := a+b+c+d+f


	result := HttpPost(tarurl, H)
	fmt.Println(result)
}
