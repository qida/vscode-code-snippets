package times

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type Person struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Birthday Time   `json:"birthday"`
}

func TestTimeJson(t *testing.T) {
	now := Time(time.Now())
	t.Log(now)
	src := `{"id":5,"name":"xiaoming","birthday":"2016-06-30 16:09:51"}`
	p := new(Person)
	err := json.Unmarshal([]byte(src), p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(p)
	t.Log(time.Time(p.Birthday))
	js, _ := json.Marshal(p)
	t.Log(string(js))
}

func TestTimeWeek(t *testing.T) {

	l, _ := time.LoadLocation("Asia/Shanghai")
	startTime, _ := time.ParseInLocation("2006-01-02", "2018-12-22", l)
	endTime, _ := time.ParseInLocation("2006-01-02", "2019-05-17", l)

	datas := GroupByWeekDate(startTime, endTime)
	for _, d := range datas {
		fmt.Println(d)
	}

}
