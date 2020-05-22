package times

import (
	"time"
)

type Time time.Time

const (
	timeFormart = `"2006-01-02 15:04:05"`
)

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(timeFormart, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart))
	b = time.Time(t).AppendFormat(b, timeFormart)
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormart)
}

//获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

//获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

//获取传入的时间所在年第一天
func GetFirstDateOfYear(d time.Time) time.Time {
	return time.Date(d.Year(), 1, 1, 0, 0, 0, 0, d.Location())
}

//获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

func GetFirstDateOfWeek(d time.Time) time.Time {
	timeFirst := d.AddDate(0, 0, -1*(int(d.Weekday())-1))
	timeFirst = GetZeroTime(timeFirst)
	return timeFirst
}
