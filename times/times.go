package times

import (
	"time"
)

type Time time.Time

const (
	timeFormart = `"2006-01-02 15:04:05"`
)

var ShangHaiZone, _ = time.LoadLocation("Asia/Shanghai")

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(timeFormart, string(data), ShangHaiZone)
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
	return GetZeroTimeOfDay(d)
}

//获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

//获取传入的时间所在年第一天
func GetFirstDateOfYear(d time.Time) time.Time {
	return time.Date(d.Year(), 1, 1, 0, 0, 0, 0, ShangHaiZone)
}

//获取某一天的0点时间
func GetZeroTimeOfDay(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, ShangHaiZone)
}

func GetFirstDateOfWeek(d time.Time, start_sunday bool) time.Time {
	var timeFirst time.Time
	var n = int(d.Weekday())
	if !start_sunday {
		if d.Weekday() == 0 {
			n = 7
		}
	}
	timeFirst = d.AddDate(0, 0, -1*(n-1))
	timeFirst = GetZeroTimeOfDay(timeFirst)
	return timeFirst
}

func GetBetweenDates(date_start, date_end time.Time) (d []time.Time) {
	if date_end.Before(date_start) {
		// 如果结束时间小于开始时间，异常
		return
	}
	// 输出日期格式固定
	d = append(d, date_start)
	for {
		date_start = date_start.AddDate(0, 0, 1)
		if date_start == date_end {
			break
		}
		d = append(d, date_start)
	}
	return
}
