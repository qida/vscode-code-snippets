package ding2

import (
	"fmt"

	"github.com/astaxie/beego/httplib"
	dingtalk "github.com/icepy/go-dingtalk/src"
)

type PackData struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type Token struct {
	PackData
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

//打卡记录
type WorkRecord struct {
	PackData
	HasMore      bool `json:"hasMore"`
	Recordresult []struct {
		CheckType      string `json:"checkType"`
		CorpId         string `json:"corpId"`
		LocationResult string `json:"locationResult"`
		BaseCheckTime  int64  `json:"baseCheckTime"`
		GroupId        int    `json:"groupId"`
		TimeResult     string `json:"timeResult"`
		UserId         string `json:"userId"`
		RecordId       int64  `json:"recordId"`
		WorkDate       int64  `json:"workDate"`
		SourceType     string `json:"sourceType"`
		UserCheckTime  int64  `json:"userCheckTime"`
		PlanId         int64  `json:"planId"`
		Id             int64  `json:"id"`
	} `json:"recordresult"`
}

type Deptment struct {
	PackData
	SubDeptIDList []int `json:"sub_dept_id_list"`
}

var c *dingtalk.DingTalkClient

func init() {
	NewClient("dingx52cfur2rzplmip7", "BFwrZ3ZtJKHc1sYGTzz78F74yxqD_C-48msYC3yOVR7i2Ew6Glk7UV4LxfvVUII-")
}
func NewClient(corp_id string, corp_secret string) {
	config := &dingtalk.DTConfig{
		CorpID:     corp_id,
		CorpSecret: corp_secret,
	}
	c = dingtalk.NewDingTalkCompanyClient(config)
}

func GetAccessToken() string {
	c.RefreshCompanyAccessToken()
	return c.AccessToken
}
func GetDeptment() (deptIds []int, err error) {
	var root Deptment
	req := httplib.Get("https://oapi.dingtalk.com/department/list_ids")
	req.Param("access_token", GetAccessToken())
	req.Param("fetch_child", "true")
	req.Param("id", "1")
	err = req.ToJSON(&root)
	if err != nil {
		return
	}
	fmt.Println(root)
	for i := 0; i < len(root.SubDeptIDList); i++ {
		var dept Deptment
		req := httplib.Get("https://oapi.dingtalk.com/department/list_ids")
		req.Param("access_token", GetAccessToken())
		req.Param("id", fmt.Sprintf("%d", root.SubDeptIDList[i]))
		err = req.ToJSON(&dept)
		if err != nil {
			return
		}
		deptIds = append(deptIds, dept.SubDeptIDList...)
	}
	deptIds = append(deptIds, root.SubDeptIDList...) //父级
	deptIds = append(deptIds, 1)                     //根级
	//父级
	fmt.Println("=======deptIds=========")
	fmt.Println(deptIds)
	return
}
