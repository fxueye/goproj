package utils

import (
	"encoding/json"
	"fmt"
	"game/common/utils"
	"regexp"
	"strconv"

	log "github.com/cihub/seelog"
)

var (
	couponUrl = "http://api.test.php9.cn/api/coupon"
	message   = "%s\n【原价】: %s元\n【内部优惠券】: %s元\n【券后价】: %s元\n【淘口令下单】: 复制这条信息，打开→手机淘宝领取优惠券%s"
)

func get_coupon() string {
	data := make(map[string]interface{})
	data["w"] = ""
	ret, err := utils.HttpPost(couponUrl, data)
	if err != nil {
		log.Errorf("http error %v", err)
		return ""
	}
	retMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(ret), &retMap)
	if err != nil {
		log.Errorf("json unmarshal err:%v", err)
		return ""
	}
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
	originaPrice := fmt.Sprintf("%v", zkp+couponPrice)
	couponStr := fmt.Sprintf(message, title, originaPrice, arr[1], zkFinalPrice, tpwd)
	return couponStr
}
