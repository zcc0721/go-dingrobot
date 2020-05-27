package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

var (
	errPostDingFail = errors.New("post ding talk robot fail")
)

// ContentInfo is ding talk request content info
type ContentInfo struct {
	Content string `json:"content"`
}

// AtInfo is ding talk request at info
type AtInfo struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

// Contents is request of ding talk robot
type Contents struct {
	MsgType string       `json:"msgtype"`
	Text    *ContentInfo `json:"text"`
	At      *AtInfo      `json:"at"`
}

// DingTalkResp is response of ding talk robot
type DingTalkResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// DingTalk implements ding talk client
type DingTalk struct {
	webHook   string
	atMobiles []string
	isAtAll   bool
}

// NewDingTalk create new ding talk client
func NewDingTalk(webHook string, atMobiles []string, isAtAll bool) *DingTalk {
	return &DingTalk{
		webHook:   webHook,
		atMobiles: atMobiles,
		isAtAll:   isAtAll,
	}
}

func (ding *DingTalk) sendDingMessage(message string) error {
	contents := &Contents{
		MsgType: "text",
		Text: &ContentInfo{
			Content: message,
		},
		At: &AtInfo{
			AtMobiles: ding.atMobiles,
			IsAtAll:   ding.isAtAll,
		},
	}

	payload, err := json.Marshal(contents)
	if err != nil {
		return errors.Wrap(err, "json marshal contents")
	}

	dingResp := new(DingTalkResp)
	if err := Post(ding.webHook, payload, dingResp); err != nil {
		return errors.Wrap(err, "post ding talk robot message")
	}

	if dingResp.ErrCode != 0 {
		return errors.Wrapf(errPostDingFail, "err: %s", dingResp.ErrMsg)
	}

	return nil
}

func Post(url string, payload []byte, result interface{}) error {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if result == nil {
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, result)
}
