package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"strings"
)

type TianRobot struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	Newslist []newslist
}

type newslist struct {
	news []string
}

type reply struct {
	Reply    string `json:"reply"`
	Datatype string `json:"datatype"`
}

func TianRebootApi(question string) string {
	url := tianApiUlr(question)
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return err.Error()
	}
	//fmt.Println("respone newslist:", data["newslist"])
	newslist := data["newslist"].([]interface{})[0]
	replyMsg := newslist.(map[string]interface{})["reply"]

	strNews := fmt.Sprintf("%v", replyMsg)
	multLineNews := strings.Replace(strNews, "<br>","\n", -1)
	fmt.Println(multLineNews)
	return multLineNews
}

func tianApiUlr(question string) string {
	baseUrl, err := url2.Parse("http://api.tianapi.com/txapi/robot/index?")
	if err != nil {
		fmt.Println("Malformed URL: ", err.Error())
		return err.Error()
	}
	// Add a Path Segment (Path segment is automatically escaped)
	//baseUrl.Path += "path with?reserved characters"

	// 增加参数
	params := url2.Values{}
	params.Add("key", TIANAPIKEY)
	params.Add("question", question)

	// 向URL增加 Query 参数
	baseUrl.RawQuery = params.Encode()
	fmt.Println("Encoded URL is: ", baseUrl.String())
	return baseUrl.String()
}


func TianReboot(msg string) (replayMsg string) {
	msg = strings.TrimLeft(msg, WECHATNAME)
	msg = strings.Trim(msg, " ")
	replayMsg = TianRebootApi(msg)
	return replayMsg
}
