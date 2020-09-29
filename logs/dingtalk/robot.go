package dingtalk

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
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
	var url string = r.Webhook
	if r.Secret != "" {
		timestamp := time.Now().Unix()
		string_to_sign := fmt.Sprintf("%d\n%s", timestamp, r.Secret)
		h := hmac.New(sha256.New, []byte(r.Secret))
		h.Write([]byte(string_to_sign))
		url = fmt.Sprintf("%s&timestamp=%d&sign=%s", url, timestamp, base64.StdEncoding.EncodeToString([]byte(hex.EncodeToString(h.Sum(nil)))))
	}
	fmt.Println(url)
	resp, err := httpClient.Post(url, "application/json", bytes.NewReader(bs))
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

type serverResponse struct {
	ErrorCode int `json:"errcode"` // just check error code
}
