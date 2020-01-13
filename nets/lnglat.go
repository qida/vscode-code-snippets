package nets

import (
	"github.com/astaxie/beego/orm"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("china", "sqlite3", "go/nets/china.db")
	orm.RegisterModel(new(China))
	// orm.RunSyncdb("default", false, true)

}

type China struct {
	Id   int     `orm:"column(id);auto"`
	Fid  int     `orm:"column(fid);size(11);null"`
	Name string  `orm:"column(name);size(255);null"`
	Deep int8    `orm:"column(deep);null"`
	Lng  float64 `orm:"column(lng);null"`
	Lat  float64 `orm:"column(lat);null"`
	Tag  string  `orm:"column(tag);size(255);null"`
}
type Tx struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  struct {
		Title    string `json:"title"`
		Location struct {
			Lng float64 `json:"lng"`
			Lat float64 `json:"lat"`
		} `json:"location"`
		AdInfo struct {
			Adcode string `json:"adcode"`
		} `json:"ad_info"`
		AddressComponents struct {
			Province     string `json:"province"`
			City         string `json:"city"`
			District     string `json:"district"`
			Street       string `json:"street"`
			StreetNumber string `json:"street_number"`
		} `json:"address_components"`
		Similarity  float64 `json:"similarity"`
		Deviation   int     `json:"deviation"`
		Reliability int     `json:"reliability"`
		Level       int     `json:"level"`
	} `json:"result"`
}
