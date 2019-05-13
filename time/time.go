package time

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
