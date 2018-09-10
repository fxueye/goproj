package server

import (
	"encoding/json"
	"fmt"
	"game/common/utils"
	"math/rand"
	"regexp"
	"strconv"

	log "github.com/cihub/seelog"
)

var (
	couponUrl = "http://api.test.php9.cn/api/coupon"
	message   = "%s\n【原价】: %s元\n【内部优惠券】: %s元\n【券后价】: %s元\n【淘口令下单】: 复制这条信息，打开→手机淘宝领取优惠券%s"
)

func GetCoupon(w string) map[string]interface{} {
	data := make(map[string]interface{})
	if w == "" {
		w = config.KeyWords[rand.Intn(len(config.KeyWords)-1)]
	}
	data["w"] = w
	data["app_key"] = config.AppKey
	data["app_secret"] = config.AppSecret
	data["app_pid"] = config.AppPid
	ret, err := utils.HttpPost(couponUrl, data)
	// log.Infof("ret:%v", ret)
	if err != nil {
		log.Errorf("http error %v", err)
		return nil
	}
	retMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(ret), &retMap)
	if err != nil {
		log.Errorf("json unmarshal err:%v", err)
		return nil
	}
	return retMap
}
func MakeCouponStr(retMap map[string]interface{}) string {
	coupon := retMap["data"].(map[string]interface{})
	title := coupon["title"].(string)
	couponInfo := coupon["coupon_info"].(string)
	zkFinalPrice := coupon["zk_final_price"].(string)

	// commissionRate := coupon["commission_rate"].(string)
	// couponClickUrl := coupon["coupon_click_url"].(string)
	tpwd := coupon["tpwd"].(string)
	reg := regexp.MustCompile(`\d+`)
	arr := reg.FindAllString(couponInfo, -1)

	log.Infof("%v", arr)
	zkp, err := strconv.ParseFloat(zkFinalPrice, 64)
	couponPrice, err := strconv.ParseFloat(arr[1], 64)
	if err != nil {
		log.Errorf("%v", err)
		return ""
	}
	originaPrice := fmt.Sprintf("%.2f", zkp-couponPrice)
	couponStr := fmt.Sprintf(message, title, zkFinalPrice, arr[1], originaPrice, tpwd)
	return couponStr
}
