package server

import (
	"sync"
	"encoding/json"
	"encoding/base64"
	"fmt"
	"game/common/server/web"
	"game/common/utils"

	log "github.com/cihub/seelog"
)

type WebHandler struct{}

var (
	// wx8844e8b0bc33183b
	//1fe0ca7dc36651f64fc7de3fbeaafadd
	mu sync.RWMutex
	openid2SessionKey = make(map[string]string)
)
func (*WebHandler) Session(ctx *web.Context,val string) string{
	sessionKey:= ctx.Params["session_key"]
	openId:= ctx.Params["openId"]
	sessionOldKey:=""
	if _,ok := openid2SessionKey[openId]; ok {
		sessionOldKey = openid2SessionKey[openId]
	}
	mu.Lock()
	defer mu.Unlock()
	openid2SessionKey[openId] = sessionKey

	return sessionOldKey
}
func (*WebHandler) Api(ctx *web.Context, val string) string {
	log.Infof("ctx : %v", ctx)
	
	if val == "login" {
		if len(ctx.Params) == 0{
			return ret(-1,"prams len is 0","")
		}
		code := ctx.Params["code"]
		openid := getOpenId(code)
		if openid == "" {
			return ret(-1,"not get opnid ","")
		}
		iv := ctx.Params["iv"]
		encryptedData := ctx.Params["encryptedData"]
		signature := ctx.Params["signature"]
		rawData := ctx.Params["rawData"]
		hstr := fmt.Sprintf("%s%s", rawData, openid2SessionKey[openid])
		signature2 := utils.Sha1(hstr)
		if signature != signature2 {
			log.Errorf("signature:%s != signature2:%s", signature, signature2)
			return ret(-1,"signature != signature2 ","")
		}
		aesKey, err := base64.StdEncoding.DecodeString(openid2SessionKey[openid])
		if err != nil {
			log.Errorf("%v", err)
			return ret(-1,"aesKey error","")
		}
		aesIv, err := base64.StdEncoding.DecodeString(iv)
		if err != nil {
			return ret(-1,"aesIv error","")
		}
		dataBytes, err := base64.StdEncoding.DecodeString(encryptedData)
		userInfoMap := make(map[string]interface{})
		userInfo, err := utils.AesDecrypt(dataBytes, aesKey, aesIv)
		if err != nil {
			log.Errorf("%v", err)
			return ret(-1,"AesDecrypt error","")
		}
		err = json.Unmarshal(userInfo,&userInfoMap)
		if err != nil{
			return	ret(-1,"userinfo json Unmarshal error","")
		}
		return ret(0,"",userInfoMap)

	}
	return ""
}
func getOpenId(code string) string{
		log.Infof("params : %v", code)
		params := make(map[string]interface{})
		params["appid"] = config.WeiAppid
		params["secret"] = config.WeiSecret
		params["js_code"] = code
		params["grant_type"] = "authorization_code"
		ret, err := utils.HttpGet(config.WeiApiUrl, params)
		log.Infof("ret:%v",ret)
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
		mu.Lock()
		defer mu.Unlock()
		openid2SessionKey[openid] = sessionKey
		return openid
}

func ret(code int,msg string,data interface{}) string{
	ret := make(map[string]interface{})
	ret["code"] = code;
	ret["msg"] = msg;
	ret["data"] = data;
	str,_ := json.Marshal(ret)
	return string(str)
}