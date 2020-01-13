package net

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/astaxie/beego/orm"
)

//初始化TAG
func (this *QQwry) UpdateTag() (err error) {
	o := orm.NewOrm()
	fmt.Println("===========")
	var addresses []China
	var num int64
	num, err = o.QueryTable(new(China)).Filter("Deep__gte", 2).OrderBy("Deep").Limit(-1).All(&addresses)
	if num == 0 {
		fmt.Println("NO DATA")
	}
	if err != nil {
		return
	}
	fmt.Println("===========")

	for i := 0; i < len(addresses); i++ {
		var address = addresses[i]
		if address.Tag == "" {
			var address_f China
			err = o.QueryTable(new(China)).Filter("Id", address.Fid).One(&address_f)
			if err != nil {
				return
			}
			if address_f.Tag != "" {
				address.Tag = address_f.Tag + address.Name
				o.Update(&address, "Tag")
			}
		}
	}
	fmt.Println("===========")

	return
}
func (this *QQwry) UpdateLngLat() (err error) {
	o := orm.NewOrm()
	fmt.Println("===========")
	var addresses []China
	var num int64
	num, err = o.QueryTable(new(China)).Filter("Deep__gte", 2).OrderBy("Deep").Limit(-1).All(&addresses)
	if num == 0 {
		fmt.Println("NO DATA")
	}
	if err != nil {
		return
	}
	fmt.Println("===========")
	for i := 0; i < len(addresses); i++ {
		var address = addresses[i]
		if address.Lng == 0 && address.Tag != "" {
			lng, lat, err1 := TxLngLat(address.Tag)
			if err1 != nil {
				err = err1
				return
			}
			address.Lng = lng
			address.Lat = lat
			o.Update(&address, "Lng", "Lat")
		}
	}
	fmt.Println("===========")
	return
}
func TxLngLat(address string) (lng, lat float64, err error) {
	params := url.Values{}
	Url, err := url.Parse("https://apis.map.qq.com/ws/geocoder/v1/")
	if err != nil {
		return
	}
	params.Set("address", address)
	params.Set("key", "ROPBZ-BCIKX-CA446-ZGFNQ-24FUF-GFFP7")
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := http.Get(urlPath)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	var tx Tx
	err = json.Unmarshal(body, &tx)
	if err != nil {
		return
	}
	fmt.Println(address, tx.Result.Location.Lng, tx.Result.Location.Lat)
	lng = tx.Result.Location.Lng
	lat = tx.Result.Location.Lat
	return
}
