package main

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"vguangQR/qr"
)

func main() {
	content := "___VBAR_CONFIG_V1.1.0___{w_mode=1,ochannel=64,nochannel=64,owifi=1,s_mode=0,interval=3000,nfc=0,haddr=\"http://192.168.16.66:8082/qrlogin?line=0\",houttime=2,ip_mode=0}"

	mac := hmac.New(md5.New, []byte("1234567887654321"))
	mac.Write([]byte(content))

	hmacBytes := mac.Sum(nil)

	content = content + "--" + base64.StdEncoding.EncodeToString(hmacBytes[:16])

	qr.RenderString(content)
	fmt.Println(content)

}
