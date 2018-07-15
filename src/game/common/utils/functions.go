package utils

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

func determinant(v1 float32, v2 float32, v3 float32, v4 float32) float32 {
	return (v1*v4 - v2*v3)
}

//校验两线段是否相交
func CheckIntersect(aa *Vector2, bb *Vector2, cc *Vector2, dd *Vector2) bool {
	delta := determinant(bb.X-aa.X, cc.X-dd.X, bb.Y-aa.Y, cc.Y-dd.Y)
	if delta <= (1E-5) && delta >= -(1E-5) { // delta=0，表示两线段重合或平行
		return false
	}
	namenda := determinant(cc.X-aa.X, cc.X-dd.X, cc.Y-aa.Y, cc.Y-dd.Y) / delta
	if namenda > 1 || namenda < 0 {
		return false
	}
	miu := determinant(bb.X-aa.X, cc.X-aa.X, bb.Y-aa.Y, cc.Y-aa.Y) / delta
	if miu > 1 || miu < 0 {
		return false
	}
	return true
}

func EscapeSqlString(src string) string {
	s := strings.Replace(src, "\\", "\\\\", -1)
	s = strings.Replace(s, "\"", "\\\"", -1)
	s = strings.Replace(s, "'", "\\'", -1)
	s = strings.Replace(s, "\r", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "\x00", "", -1)
	s = strings.Replace(s, "\x1a", "", -1)
	return s
}

func GetIPFromAddr(addr string) string {
	idxLeft := strings.Index(addr, "[")
	if idxLeft > -1 {
		idxRight := strings.Index(addr, "]")
		if idxRight > -1 {
			return addr[idxLeft+1 : idxRight]
		} else {
			return ""
		}
	} else {
		return strings.Split(addr, ":")[0]
	}
}

func CheckWhiteIPs(whites map[string]bool, addr string) bool {
	nums := strings.Split(addr, ".")
	ipformats := make([]string, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		ip := ""
		for j := 0; j < len(nums); j++ {
			if j < i {
				ip = fmt.Sprintf("%s%v.", ip, nums[j])
			} else {
				ip = fmt.Sprintf("%s*.", ip)
			}
		}
		ip = ip[:len(ip)-1]
		ipformats[i] = ip
	}
	ipformats[len(nums)] = addr
	for _, v := range ipformats {
		if val, ok := whites[v]; ok && val {
			return true
		}
	}
	return false
}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetTripleValue(cond bool, tValue interface{}, fValue interface{}) interface{} {
	if cond {
		return tValue
	} else {
		return fValue
	}
}

func GetTripleIntValue(cond bool, tValue int, fValue int) int {
	if cond {
		return tValue
	} else {
		return fValue
	}
}

//参数：weights 权重，randCount 随机数量 repeat是否可以重复
//返回：选中的数组下标的集合
func GetRandomValueByWeight(weights []int, randCount int, repeat bool) (ret []int) {
	ret = make([]int, 0, randCount)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	//总权重
	totalWeight := 0
	for _, w := range weights {
		totalWeight += w
	}
	//生成临时序号
	tmpIndexList := make([]int, len(weights))
	for i := 0; i < len(weights); i++ {
		tmpIndexList[i] = i
	}
	//不可重复，并且要求数量大于权重数量，直接返回所有下标
	if !repeat && randCount >= len(weights) {
		return tmpIndexList
	}
	for i := 0; i < randCount; i++ {
		rdValue := r.Intn(totalWeight)
		//			log.Debugf("rdValue[%v]/totalWeight[%v]", rdValue, totalWeight)
		gotIdx := 0 //对应tmpIndexList下标
		for j := 0; j < len(tmpIndexList); j++ {
			if rdValue < weights[tmpIndexList[j]] {
				gotIdx = j
				break
			}
			rdValue -= weights[j]
		}

		ret = append(ret, tmpIndexList[gotIdx])
		if !repeat {
			totalWeight -= weights[tmpIndexList[gotIdx]]
			tmpIndexList = append(tmpIndexList[:gotIdx], tmpIndexList[gotIdx+1:]...)
		}
		//			log.Debug(tmpIndexList)

	}
	return
}

func TimeToZeroTime(sec int64) int64 {
	t := time.Unix(sec, 0)
	nt := t.Format("2006-01-02 00:00:00")
	the_time, _ := time.ParseInLocation("2006-01-02 15:04:05", nt, time.Local)
	return the_time.Unix()
}

func DateZeroTime() int64 {
	the_time, _ := time.ParseInLocation("2006-01-02 15:04:05", time.Now().Format("2006-01-02 00:00:00"), time.Local)
	return the_time.Unix()
}

//获取游戏中每日零点(非自然时间零点)
func GameDateZeroTime(freshClock int) int64 {
	now := time.Now()
	if now.Hour() >= freshClock {
		//如果当前小时大于刷新钟点，则返回今日时间
		return DateZeroTime() + int64(freshClock*3600)
	} else {
		//如果当前小时小于刷新钟点，则返回昨日时间
		return DateZeroTime() + int64(freshClock*3600) - 3600*24
	}
}

func TimeDistanceDay(old int64, cur int64) int {
	curTime := time.Unix(cur, 0)
	the_time, _ := time.ParseInLocation("2006-01-02 15:04:05", curTime.Format("2006-01-02 00:00:00"), time.Local)
	curStp := the_time.Unix()
	dis := curStp - old
	if dis < 0 {
		return 0
	} else {
		return int(math.Ceil(float64(dis) / (3600 * 24)))
	}
}

func TimeToWeekday(st int64) int {
	week := time.Unix(st, 0).Weekday().String()
	switch week {
	case "Monday":
		return 1
	case "Tuesday":
		return 2
	case "Wednesday":
		return 3
	case "Thursday":
		return 4
	case "Friday":
		return 5
	case "Saturday":
		return 6
	case "Sunday":
		return 7
	}
	return 1
}

func GetRandomData(count int) (idx int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	idx = r.Intn(count)
	return idx
}

//功能尚不完善，勿使用
//func ShowValue(prefix string, name string, data interface{}) {
//	t := reflect.TypeOf(data)
//	v := reflect.ValueOf(data)
//
//	switch v.Kind() {
//	case reflect.Struct:
//		{
//			log.Debugf("%v{", prefix)
//			for j := 0; j < t.NumField(); j++ {
//				ShowValue(fmt.Sprintf("%v\t", prefix), t.Field(j).Name, v.Field(j).Interface())
//			}
//			log.Debugf("%v}", prefix)
//		}
//	case reflect.Slice:
//		{
//			log.Debugf("%v%v:", prefix, name)
//			log.Debugf("%v[", prefix)
//			for k := 0; k < v.Len(); k++ {
//				ShowValue(fmt.Sprintf("%v\t", prefix), fmt.Sprintf("%v", k), v.Index(k).Interface())
//			}
//			log.Debugf("%v]", prefix)
//		}
//	case reflect.Ptr:
//		{
//			log.Debug(v)
//			log.Debug(v.Elem())
//			//			ShowValue(prefix, name, v.Elem())
//		}
//	default:
//		log.Debugf("%v%v : %v", prefix, name, v.Interface())
//	}
//
//}

func GetStringFromNum(num interface{}) string {
	switch num.(type) {

	case int16:
		return strconv.FormatInt(int64(num.(int16)), 10)
	case int:
		return strconv.FormatInt(int64(num.(int)), 10)
	case int32:
		return strconv.FormatInt(int64(num.(int32)), 10)
	case int64:
		return strconv.FormatInt(int64(num.(int64)), 10)
	case uint:
		return strconv.FormatUint(uint64(num.(uint)), 10)
	case uint16:
		return strconv.FormatUint(uint64(num.(uint16)), 10)
	case uint32:
		return strconv.FormatUint(uint64(num.(uint32)), 10)
	case uint64:
		return strconv.FormatUint(uint64(num.(uint64)), 10)
	case float32:
		return strconv.FormatFloat(float64(num.(float32)), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(float64(num.(float64)), 'f', -1, 64)
	default:
		return ""
	}
}

//获取字符串长度，ascII小于128算1，其他算2
func GetStringLength(str string) int {
	cnt := 0
	for _, r := range str {
		if r < 128 {
			cnt++
		} else {
			cnt += 2
		}
	}
	return cnt

}

var filterChar = []byte{'#', ',', '\n', ';', '|', ':'}

func ContainFilterWords(str string) bool {
	for i := 0; i < len(filterChar); i++ {
		idx := strings.IndexByte(str, filterChar[i])
		if idx != -1 {
			return true
		}
	}
	return false
}

func GenerateLogXml(filename string) (error, string) {
	bs, err := ioutil.ReadFile("etc/log.xml")
	if err != nil {
		return err, ""
	}
	logString := strings.Replace(string(bs), "#FILE", filename, -1)
	return nil, logString
}

func IsInArray(nums []int, num int) bool {
	for _, n := range nums {
		if n == num {
			return true
		}
	}
	return false
}

func IsRepeat(ids []int) bool {
	sort.Ints(ids)
	for i, id := range ids {
		if i == len(ids)-1 {
			return false
		} else {
			if id == ids[i+1] {
				return true
			}
		}
	}
	return false
}

func GetTelNum(s string) []string {
	reg := "(((\\+\\d{2}-)?0\\d{2,3}-\\d{7,8})|((\\+\\d{2}-)?(\\d{2,3}-)?([1][3,4,5,7,8][0-9]\\d{8})))"
	rgx := regexp.MustCompile(reg)
	return rgx.FindAllString(s, -1)
}

func RemoveDuplicatesAndEmpty(a []string) (ret []string) {
	a_len := len(a)
	for i := 0; i < a_len; i++ {
		if (i > 0 && a[i-1] == a[i]) || len(a[i]) == 0 {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}
