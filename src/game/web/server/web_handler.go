package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"game/common/server/web"
	"game/common/utils"
	"net/http"
	"strconv"
	"strings"

	log "github.com/cihub/seelog"
)

type WebHandler struct{}

var (
	appid      = "wxcb74c9581beb6739"
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
		sessionKey = ctx.Params["code"]
		log.Infof("params : %v", sessionKey)
		params := make(map[string]interface{})
		// params["appid"] =
		// utils.HttpGet(apiUrl)
		jsMap := make(map[string]interface{})
		str := `{"errno":0,"msg":"成功","data":{"user_id":1109261,"user_name":"F143001040223","uuid":"ffffffff-cfe9-f796-ffff-ffffef05ac4a","mobile":"","nickname":"","avatar":"","gid":211}}`
		dec := json.NewDecoder(strings.NewReader(str))
		dec.UseNumber()
		err := dec.Decode(&jsMap)
		log.Infof("jsMap : %v", jsMap)
		if err != nil {
			log.Error(err)
			return ""
		}
		errno, err := strconv.Atoi(jsMap["errno"].(json.Number).String())
		if err != nil {
			return ""
		}
		if errno > 0 {
			return ""
		}
		data := jsMap["data"].(map[string]interface{})
		params["openid"] = fmt.Sprintf("%v", data["user_id"])
		log.Infof("%v", params)

		ret := make(map[string]interface{})

		ret["return_code"] = 1
		ret["body"] = params

		byte, err := json.Marshal(ret)
		if err != nil {
			log.Error(err)
			return ""
		}
		log.Infof("json: %v", string(byte))

		getUrl := "version_id=1.0.0&partner_id=1&timestamp=123456&sign=qwe+qwe"
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
