package file

import (
	"os"
	"testing"
)

func TestGetMP4Duration(t *testing.T) {
	t.Run("tt.name", func(t *testing.T) {
		file, err := os.Open("http://wechat.cdn.zxjy.work/audio/rec_customer_follow/1608373085_老潘语音.m4a")
		if err != nil {
			t.Errorf("Err = %v ", err)
			return
		}
		gotLengthOfTime, err := GetMP4Duration(file)
		if err != nil {
			t.Errorf("GetMP4Duration() error = %v", err)
			return
		}
		t.Errorf("GetMP4Duration() = %v ", gotLengthOfTime)
	})
}

// func main() {
// 	file, err := os.Open(os.Args[1])
// 	if err != nil {
// 		panic(err)
// 	}
// 	duration, err := GetMP4Duration(file)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(filepath.Base(os.Args[1]), duration)
// }
