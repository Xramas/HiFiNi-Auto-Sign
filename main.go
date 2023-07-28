package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	client := &http.Client{}
	success := SignIn(client)
	if success {
		fmt.Println("签到成功")
	} else {
		fmt.Println("签到失败")
		os.Exit(3)
	}
}

// SignIn 签到
func SignIn(client *http.Client) bool {
	//生成要访问的url
	url := "https://www.hifini.com/sg_sign.htm"
	cookie := os.Getenv("COOKIE")
	if cookie == "" {
		fmt.Println("COOKIE不存在，请检查是否添加")
		return false
	}
	//提交请求
	reqest, err := http.NewRequest("POST", url, nil)
	reqest.Header.Add("Cookie", cookie)
	reqest.Header.Add("x-requested-with", "XMLHttpRequest")
	//处理返回结果
	response, err := client.Do(reqest)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	buf, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(buf))
	dingding()
	return strings.Contains(string(buf), "成功")
}

func dingding(){
	// 构造要发送的消息
	message := DingTalkMessage{
		MsgType: "text",
		Text: struct {
			Content string `json:"content"`
		}{
			Content: "Hello, DingTalk!",
		},
	}

	// 将消息转换为JSON格式
	messageJson, _ := json.Marshal(message)
	DINGDING_WEBHOOK := os.Getenv("DINGDING_WEBHOOK")
	// 发送HTTP POST请求
	resp, err := http.Post(DINGDING_WEBHOOK,
		"application/json", bytes.NewBuffer(messageJson))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
