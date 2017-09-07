package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang/glog"
)

const AccessTokenFilePath = "/root/PropertyManagement/server/dist/access_token.json"

func WriteFile(path string, content string) {
	outputFile, outputError := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if outputError != nil {
		glog.Error(outputError)
		return
	}
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	outputWriter.WriteString(content)
	outputWriter.Flush()
}

func GetAccessToken() string {
	access_token, _ := ReadFile(AccessTokenFilePath)
	if access_token != nil { //存在则要判断时间
		curTime := time.Now().Unix()
		saveTime, _ := strconv.ParseInt(access_token["time"], 10, 64)
		expiresTime, _ := strconv.ParseInt(access_token["expires_in"], 10, 64)
		if curTime >= (saveTime + expiresTime) { //access_token失效
			return FetchNewAccessToken()
		}
		return access_token["access_token"]
	}
	// access_token不存在
	return FetchNewAccessToken()
}

func FetchNewAccessToken() string {
	wxinfo, _ := ReadFile("/root/PropertyManagement/server/dist/wx.json")
	resp := GetJson(
		"https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" +
			wxinfo["appid"] + "&secret=" + wxinfo["appsecret"])
	resp["time"] = time.Now().Unix()
	jsonResp, _ := json.Marshal(resp)
	fmt.Println(jsonResp)
	os.Remove(AccessTokenFilePath)
	WriteFile(AccessTokenFilePath, string(jsonResp))
	return resp["access_token"].(string)
}
