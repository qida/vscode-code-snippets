package str

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

//全角转半角
func DBCtoSBC(s string) string {
	retstr := ""
	for _, i := range s {
		inside_code := i
		if inside_code == 12288 {
			inside_code = 32
		} else {
			inside_code -= 65248
		}
		if inside_code < 32 || inside_code > 126 {
			retstr += string(i)
		} else {
			retstr += string(inside_code)
		}
	}
	return retstr
}

func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

/**
 *
 */
func Substr2(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < 0 || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}

// 按字节截取字符串 utf-8不乱码
func SubstrByByte(str string, length int) string {
	bs := []byte(str)[:length]
	bl := 0
	for i := len(bs) - 1; i >= 0; i-- {
		switch {
		case bs[i] >= 0 && bs[i] <= 127:
			return string(bs[:i+1])
		case bs[i] >= 128 && bs[i] <= 191:
			bl++
		case bs[i] >= 192 && bs[i] <= 253:
			cl := 0
			switch {
			case bs[i]&252 == 252:
				cl = 6
			case bs[i]&248 == 248:
				cl = 5
			case bs[i]&240 == 240:
				cl = 4
			case bs[i]&224 == 224:
				cl = 3
			default:
				cl = 2
			}
			if bl+1 == cl {
				return string(bs[:i+cl])
			}
			return string(bs[:i])
		}
	}
	return ""
}

func SubString(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	// 返回子串
	return string(rs[begin:end])
}
func Show_substr(s string, l int) string {
	if len(s) <= l {
		return s
	}
	ss, sl, rl, rs := "", 0, 0, []rune(s)
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			rl = 1
		} else {
			rl = 2
		}
		if sl+rl > l {
			break
		}
		sl += rl
		ss += string(r)
	}
	return ss
}

// 根据字符串显示获取显示长度
// 复制代码 代码如下:
func show_strlen(s string) int {
	sl := 0
	rs := []rune(s)
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			sl++
		} else {
			sl += 2
		}
	}
	return sl
}
func Sub(str string, start, length int) string {

	rs := []rune(str)

	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

//#用途：把用汉字表示的数字转换为阿拉伯数字

//=========================
func NumReg(han string) int {
	var num map[string]int
	num = make(map[string]int)
	num["零"] = 0
	num["〇"] = 0
	num["一"] = 1
	num["要"] = 1
	num["么"] = 1
	num["二"] = 2
	num["啊"] = 2
	num["啦"] = 2
	num["三"] = 3
	num["四"] = 4
	num["五"] = 5
	num["哦"] = 5
	num["六"] = 6
	num["七"] = 7
	num["切"] = 7
	num["其"] = 7
	num["八"] = 8
	num["九"] = 9
	num["十"] = 10
	num["是"] = 10
	num["实"] = 10
	num["百"] = 100
	num["白"] = 100
	num["千"] = 1000
	num["万"] = 10000
	var res int
	if strings.Contains(han, "十") || strings.Contains(han, "实") || strings.Contains(han, "是") {
		han = strings.Replace(han, "实", "十", -1)
		han = strings.Replace(han, "是", "十", -1)
		if shi := UnicodeIndex(han, "十"); shi > 0 {
			s, ok := num[SubString(han, shi-1, 1)]
			if ok {
				res += s * 10
			}

		} else {
			if shi == 0 {
				res += 10
			}
		}

	}
	if strings.Contains(han, "百") || strings.Contains(han, "白") {
		han = strings.Replace(han, "白", "百", -1)
		if bai := UnicodeIndex(han, "百"); bai > 0 {
			s, ok := num[SubString(han, bai-1, 1)]
			if ok {
				res += s * 100
			}

		}
	}
	s, ok := num[SubString(han, len(han)-1, 1)]
	if ok {
		res += s
	}
	// 	if shi := UnicodeIndex(han, "十"); shi > 0 {

	// 		g, ok1 := num[SubString(han, len(han)/3-1, 1)]
	// 		if ok1 {
	// 			if g != 10 {
	// 				res += g
	// 			}
	// 		}

	// 	}
	// } else {
	// 	for i := 0; i < len(han)/3; i++ {
	// 		b, ok := num[SubString(han, i, 1)]
	// 		if ok {
	// 			res += b * int(math.Pow(10, float64(len(han)/3-i-1)))
	// 		}
	// 	}
	// }
	return res
}

func UnicodeIndex(str, substr string) int {
	// 子串在字符串的字节位置
	result := strings.Index(str, substr)
	if result >= 0 {
		// 获得子串之前的字符串并转换成[]byte
		prefix := []byte(str)[0:result]
		// 将子串之前的字符串转换成[]rune
		rs := []rune(string(prefix))
		// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
		result = len(rs)
	}

	return result
}

func AmountConvert(p_money float64, p_Round bool) string {
	var NumberUpper = []string{"壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖", "零"}
	var Unit = []string{"分", "角", "圆", "拾", "佰", "仟", "万", "拾", "佰", "仟", "亿", "拾", "佰", "仟"}
	var regex = [][]string{
		{"零拾", "零"}, {"零佰", "零"}, {"零仟", "零"}, {"零零零", "零"}, {"零零", "零"},
		{"零角零分", "整"}, {"零分", "整"}, {"零角", "零"}, {"零亿零万零元", "亿元"},
		{"亿零万零元", "亿元"}, {"零亿零万", "亿"}, {"零万零元", "万元"}, {"万零元", "万元"},
		{"零亿", "亿"}, {"零万", "万"}, {"拾零圆", "拾元"}, {"零圆", "元"}, {"零零", "零"}}
	Str, DigitUpper, UnitLen, Round := "", "", 0, 0

	if p_money == 0 {
		return "零"
	}
	if p_money < 0 {
		Str = "负"
		p_money = math.Abs(p_money)
	}
	if p_Round {
		Round = 1
	} else {
		Round = 2
	}

	Digit_byte := []byte(strconv.FormatFloat(p_money, 'f', Round+1, 64)) //注意币种四舍五入
	UnitLen = len(Digit_byte) - Round

	for _, v := range Digit_byte {
		if UnitLen >= 1 && v != 46 {
			s, _ := strconv.ParseInt(string(v), 10, 0)
			if s != 0 {
				DigitUpper = NumberUpper[s-1]

			} else {
				DigitUpper = "零"
			}
			Str = Str + DigitUpper + Unit[UnitLen-1]
			UnitLen = UnitLen - 1
		}
	}

	for i := range regex {
		reg := regexp.MustCompile(regex[i][0])
		Str = reg.ReplaceAllString(Str, regex[i][1])
	}

	if string(Str[0:3]) == "元" {
		Str = string(Str[3:])
	}

	if string(Str[0:3]) == "零" {
		Str = string(Str[3:])
	}
	return Str
}

func GetKeysString(key_str string) (number int, py string, han string) {
	key_str = strings.TrimSpace(key_str)
	if len(key_str) == 0 {
		return
	}
	if unicode.Is(unicode.Latin, []rune(key_str)[0]) {
		py = strings.ToUpper(key_str)
	} else if unicode.Is(unicode.Number, []rune(key_str)[0]) {
		number, _ = strconv.Atoi(key_str)
	} else if unicode.Is(unicode.Han, []rune(key_str)[0]) {
		han = key_str
	} else {
		han = key_str
	}
	return
}
