package dingtalk

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	httpClient http.Client
)

type Robot struct {
	// Name    string `json:"name"`
	Webhook string `json:"webhook"`
	Secret  string `json:"secret"`
}

func New(webhook string, secret string) *Robot {
	return &Robot{
		// Name:    name,
		Webhook: webhook,
		Secret:  secret,
	}
}

// func (r *Robot) String() string {
// 	return fmt.Sprintf("DingtalkRobot#[%s](%s)", r.Name, r.Webhook)
// }

func (r *Robot) NewTextMessage() *textMessage {
	return &textMessage{robot: r}
}

func (r *Robot) NewLinkMessage() *linkMessage {
	return &linkMessage{robot: r}
}

func (r *Robot) NewMarkdownMessage() *markdownMessage {
	return &markdownMessage{robot: r}
}

func (r *Robot) sendMessagePayload(payload *messagePayload) error {
	bs, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	var urlWeb string = "https://oapi.dingtalk.com/robot/send?access_token="
	urlWeb += r.Webhook
	if r.Secret != "" {
		timestamp := time.Now().UnixNano() / 1e6
		urlWeb = fmt.Sprintf("%s&timestamp=%d&sign=%s", urlWeb, timestamp, sign(timestamp, r.Secret))
	}
	fmt.Println(urlWeb)
	resp, err := httpClient.Post(urlWeb, "application/json", bytes.NewReader(bs))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var sr serverResponse
	responseTextBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(responseTextBytes, &sr); err != nil {
		return errors.New(err.Error() + " :" + string(responseTextBytes))
	}
	if sr.ErrorCode != 0 {
		return errors.New(string(responseTextBytes))
	}
	return nil
}
func sign(t int64, secret string) string {
	strToHash := fmt.Sprintf("%d\n%s", t, secret)
	hmac256 := hmac.New(sha256.New, []byte(secret))
	hmac256.Write([]byte(strToHash))
	return base64.StdEncoding.EncodeToString(hmac256.Sum(nil))
}

type serverResponse struct {
	ErrorCode int `json:"errcode"` // just check error code
}
