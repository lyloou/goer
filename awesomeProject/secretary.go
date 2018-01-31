package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	// sign := ApiSign("201708111010007", "{\"MerchantNo\":\"201708111010007\",\"vcRobotSerialNo\":\"AB424E9092CB32DC368A7DA686F0D99E\"}")
	sign := ApiSign("201708111010007", `{MerchantNo:"201708111010007",vcRobotSerialNo:"",vcChatRoomSerialNo:"",nType:10,NAddMinute:10,nCodeCount:1,nAddMinute:10,nIsNotify:0,vcNotifyContent:"wowo"}`)
	fmt.Println(sign)

}

func ApiSign(secret string, strContext string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(strContext + secret))
	result := md5Ctx.Sum(nil)
	return hex.EncodeToString(result)
}
