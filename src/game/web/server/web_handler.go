package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"game/common/server/web"
	"game/common/utils"
	"strings"
	"sync"

	log "github.com/cihub/seelog"
)

type WebHandler struct{}

var (
	// wx8844e8b0bc33183b
	//1fe0ca7dc36651f64fc7de3fbeaafadd
	aesKey            = []byte("abcdefghi@qq.com")
	aesIv             = []byte("123456789@qq.com")
	mu                sync.RWMutex
	openid2SessionKey = make(map[string]string)
)

func (*WebHandler) Session(ctx *web.Context, val string) string {
	sessionKey := ctx.Params["sessionKey"]
	openId := ctx.Params["openId"]

	sessionKeyEncrypt := read(openId)
	if sessionKeyEncrypt != "" {
		sessionOldKey, _ := utils.AesDecrypt([]byte(sessionKeyEncrypt), aesKey, aesIv)
		return string(sessionOldKey)
	}
	sessionKeyEn, _ := utils.AesEncrypt([]byte(sessionKey), aesKey, aesIv)
	write(openId, string(sessionKeyEn))
	return ""
}
func read(openId string) string {
	defer mu.Unlock()
	mu.Lock()
	if _, ok := openid2SessionKey[openId]; ok {
		return openid2SessionKey[openId]
	}
	return ""
}
func write(openId, sessionKey string) {
	defer mu.Unlock()
	mu.Lock()
	openid2SessionKey[openId] = sessionKey
}
func (*WebHandler) Api(ctx *web.Context, val string) string {
	log.Infof("ctx : %v", ctx)

	if val == "login" {
		if len(ctx.Params) == 0 {
			return ret(-1, "prams len is 0", "")
		}
		token := ctx.Params["token"]
		params := make(map[string]interface{})
		err := JsonDecode(token, &params)
		if err != nil {
			return ret(-1, "token error ", "")
		}
		if _, ok := params["code"]; !ok {
			return ret(-1, "code error ", "")
		}
		code := params["code"].(string)
		openId := getOpenId(code)
		if openId == "" {
			return ret(-1, "not get opnid ", "")
		}
		if _, ok := params["iv"]; !ok {
			return ret(-1, "iv error ", "")
		}
		iv := params["iv"].(string)
		if _, ok := params["encryptedData"]; !ok {
			return ret(-1, "encryptedData error ", "")
		}
		encryptedData := params["encryptedData"].(string)
		if _, ok := params["signature"]; !ok {
			return ret(-1, "signature error ", "")
		}
		signature := params["signature"].(string)
		if _, ok := params["rawData"]; !ok {
			return ret(-1, "rawData error ", "")
		}
		rawData := params["rawData"].(string)
		sessionKey := read(openId)
		hstr := fmt.Sprintf("%s%s", rawData, sessionKey)
		signature2 := utils.Sha1(hstr)
		if signature != signature2 {
			log.Errorf("signature:%s != signature2:%s", signature, signature2)
			return ret(-1, "signature != signature2 ", "")
		}
		aesKey, err := base64.StdEncoding.DecodeString(sessionKey)
		if err != nil {
			log.Errorf("%v", err)
			return ret(-1, "aesKey error", "")
		}
		aesIv, err := base64.StdEncoding.DecodeString(iv)
		if err != nil {
			return ret(-1, "aesIv error", "")
		}
		dataBytes, err := base64.StdEncoding.DecodeString(encryptedData)
		userInfoMap := make(map[string]interface{})
		userInfo, err := utils.AesDecrypt(dataBytes, aesKey, aesIv)
		if err != nil {
			log.Errorf("%v", err)
			return ret(-1, "AesDecrypt error", "")
		}
		err = json.Unmarshal(userInfo, &userInfoMap)
		if err != nil {
			return ret(-1, "userinfo json Unmarshal error", "")
		}
		return ret(0, "", userInfoMap)

	}
	return ""
}
func JsonDecode(str string, v interface{}) error {
	dec := json.NewDecoder(strings.NewReader(str))
	dec.UseNumber()
	err := dec.Decode(v)
	if err != nil {
		return err
	}
	return nil
}
func getOpenId(code string) string {
	log.Infof("params : %v", code)
	params := make(map[string]interface{})
	params["appid"] = config.WeiAppid
	params["secret"] = config.WeiSecret
	params["js_code"] = code
	params["grant_type"] = "authorization_code"
	ret, err := utils.HttpGet(config.WeiApiUrl, params)
	log.Infof("ret:%v", ret)
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
	openId := retMap["openid"].(string)
	write(openId, sessionKey)
	return openId
}

func ret(code int, msg string, data interface{}) string {
	ret := make(map[string]interface{})
	ret["code"] = code
	ret["msg"] = msg
	ret["data"] = data
	str, _ := json.Marshal(ret)
	log.Debugf("ret:%v", string(str))
	return string(str)
}
