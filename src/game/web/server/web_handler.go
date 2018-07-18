package server

import (
	"encoding/base64"
	"fmt"
	"game/common/server/web"
	"game/common/utils"
	"net/http"

	log "github.com/cihub/seelog"
)

type WebHandler struct{}

var (
	appid             = "wxbfdac7331dafd481"
	secret            = "c96533730072f3a9be92900b5f453f95"
	apiUrl            = "https://api.weixin.qq.com/sns/jscode2session"
	openid2SessionKey = make(map[string]string)
)

// func (*WebHandler) Api(val string) string {
// 	log.Infof("ctx : %v", ctx)
// 	return "hello " + val + "\n"
// }
func (*WebHandler) Api(ctx *web.Context, val string) string {
	log.Infof("ctx : %v", ctx)
	if val == "aes" {
		iv := ctx.Params["iv"]
		openid := ctx.Params["openid"]
		encryptedData := ctx.Params["encryptedData"]
		aesKey := base64.StdEncoding.EncodeToString([]byte(openid2SessionKey[openid]))
		aseIv := base64.StdEncoding.EncodeToString([]byte(iv))
		ret, err := utils.DesDecrypt([]byte(encryptedData), []byte(aesKey), []byte(aseIv))
		if err != nil {
			log.Errorf("%v", err)
			return ""
		}
		return string(ret)
	} else if val == "openid" {
		log.Infof("params : %v", ctx.Params)
		code := ctx.Params["code"]
		log.Infof("params : %v", code)
		params := make(map[string]interface{})
		params["appid"] = appid
		params["secret"] = secret
		params["js_code"] = code
		params["grant_type"] = "authorization_code"
		ret, err := utils.HttpGet(apiUrl, params)
		retMap := make(map[string]interface{})
		if err != nil {
			log.Errorf("%+v", err)
			return ""
		}
		err = utils.JsonDecode(ret, &retMap)
		if err != nil {
			log.Errorf("%+v", err)
			return ""
		}
		if _, ok := retMap["errcode"]; ok {
			return ""
		}
		if _, ok := retMap["session_key"]; !ok {
			return ""
		}
		if _, ok := retMap["openid"]; !ok {
			return ""
		}
		sessionKey := retMap["session_key"].(string)
		openid := retMap["openid"].(string)
		openid2SessionKey[openid] = sessionKey

		log.Infof("ret:%v", ret)
		return fmt.Sprintf(`{"openid":"%s"}`, openid)
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
