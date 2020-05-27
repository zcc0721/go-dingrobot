package main

import "fmt"

// More details about ding talk doc: https://ding-doc.dingtalk.com/doc#/serverapi2/qf2nxq/9e91d73c
func Example() {
	webHook := "https://oapi.dingtalk.com/robot/send?access_token=5f6585f2ed6b9461edb4391ff9a8a128a66643ba082054e90a5d1cdef547f53a"
	atMobiles := []string{"18612345678"}
	isAtAll := false
	dingTalk := NewDingTalk(webHook, atMobiles, isAtAll)
	message := "test hello ding talk"
	if err := dingTalk.sendDingMessage(message); err != nil {
		fmt.Println(err)
	}
}
