package server

import (
	"encoding/base64"
	"game/common/server/web"
	"game/common/utils"
	"net/http"

	log "github.com/cihub/seelog"
)

type WebHandler struct{}

var (
	appid      = "wx2f7a41ab7dac20d0"
	secret     = "fe7676c20068470a5b0074490876c13f"
	sessionKey = ""
	apiUrl     = "https://api.weixin.qq.com/sns/jscode2session"
)

// func (*WebHandler) Api(val string) string {
// 	log.Infof("ctx : %v", ctx)
// 	return "hello " + val + "\n"
// }
func (*WebHandler) Api(ctx *web.Context, val string) string {
	log.Infof("ctx : %v", ctx)
	if val == "aes" {
		iv := ctx.Params["iv"]
		encryptedData := ctx.Params["encryptedData"]

		aesKey := base64.StdEncoding.EncodeToString([]byte(sessionKey))
		aseIv := base64.StdEncoding.EncodeToString([]byte(iv))
		ret, err := utils.DesDecrypt([]byte(encryptedData), []byte(aesKey), []byte(aseIv))
		if err != nil {
			return ""
		}
		return string(ret)
	} else if val == "login" {
		log.Infof("params : %v", ctx.Params)
		code := ctx.Params["code"]
		log.Infof("params : %v", code)
		params := make(map[string]interface{})
		params["appid"] = appid
		params["secret"] = secret
		params["js_code"] = code
		params["grant_type"] = "authorization_code"
		ret, err := utils.HttpGet(apiUrl, params)
		if err != nil {
			log.Errorf("%+v", err)
			return ""
		}
		log.Infof("ret:%v", ret)
		// getUrl := "version_id=1.0.0&partner_id=1&timestamp=123456&sign=qwe+qwe"
		// u, err := url.Parse(getUrl)
		// if err != nil {
		// 	fmt.Print(err)
		// 	return ""
		// }
		// ms, _ := url.ParseQuery(u.RawQuery)
		// for k, v := range ms {
		// 	log.Infof("k:%v  v:%v", k, v)
		// }
		// log.Infof("%v", makeGetString(getUrl))
		// num := 1000000
		// for i := 0; i < 99999; i++ {
		// 	num = num + 1
		// 	fmt.Printf("before:%d\n", num)
		// 	fmt.Printf("after:%e\n", float64(num))
		// }

	}
	return ""
}

// func (*WebHandler) Api(ctx *web.Context) {
// 	log.Infof("ctx : %v", ctx)
// 	ctx.Write([]byte("hello 2"))
// }
func (*WebHandler) Test(w http.ResponseWriter, r *http.Request) {
	log.Info("text")
	w.Write([]byte("hello"))
}
