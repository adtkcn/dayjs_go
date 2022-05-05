package dayjs

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type DayjsStruct struct {
	Year   int //年
	Month  int //月
	Day    int //日
	Hour   int //小时
	Minute int //分钟
	Second int //秒

	Time time.Time
}

func Dayjs(args ...string) *DayjsStruct {
	dayTime := &DayjsStruct{}

	if len(args) == 1 {
		dayTime.Parse(args[0])
	} else {
		dayTime.Now()
	}

	return dayTime
}

// 前置补零
func ZeroFillByStr[T any](str T, resultLen int) string {
	newStr := fmt.Sprintf("%v", str)
	if len(newStr) > resultLen || resultLen <= 0 {
		return newStr
	}
	result := newStr
	for i := 0; i < resultLen-len(newStr); i++ {
		result = "0" + result
	}
	return result
}

// 解析时间，每个时间需要任意字符分开； YYYY年MM月DD HH时mm分ss秒
func (t *DayjsStruct) Parse(str string) *DayjsStruct {
	re := regexp.MustCompile("[0-9]+")
	timeArr := re.FindAllString(str, -1) //-1以表明您想要全部
	fmt.Println(timeArr)
	if len(timeArr) > 6 || len(timeArr) == 0 {
		panic("时间格式最少需要一个时间")
	}
	year := fmt.Sprint(time.Now().Year()) //年
	month := "01"                         //月
	day := "01"                           //日
	hour := "00"                          //小时
	minute := "00"                        //分钟
	second := "00"                        //秒

	for key, val := range timeArr {
		if key == 0 {
			year = ZeroFillByStr(val, 4)
		} else if key == 1 {
			month = ZeroFillByStr(val, 2)
		} else if key == 2 {
			day = ZeroFillByStr(val, 2)
		} else if key == 3 {
			hour = ZeroFillByStr(val, 2)
		} else if key == 4 {
			minute = ZeroFillByStr(val, 2)
		} else if key == 5 {
			second = ZeroFillByStr(val, 2)
		}
	}
	timeStr := year + "-" + month + "-" + day + " " + hour + ":" + minute + ":" + second

	strTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		panic(err)
	}
	t.Time = strTime
	t.SetTime()
	return t
}

// 设置为当前时间
func (t *DayjsStruct) Now() *DayjsStruct {
	t.Time = time.Now() //获取当前时间
	t.SetTime()
	return t
}
func (t *DayjsStruct) SetTime() {
	t.Year = t.Time.Year()        //年
	t.Month = int(t.Time.Month()) //月
	t.Day = t.Time.Day()          //日
	t.Hour = t.Time.Hour()        //小时
	t.Minute = t.Time.Minute()    //分钟
	t.Second = t.Time.Second()    //秒
}

// 格式化 YYYY-MM-DD HH:mm:ss
func (t *DayjsStruct) Format(format ...string) string {
	formatStr := "YYYY-MM-DD HH:mm:ss"
	if len(format) == 1 && format[0] != "" {
		formatStr = format[0]
	}
	timeStr := strings.ReplaceAll(formatStr, "YYYY", fmt.Sprint(t.Year))
	timeStr = strings.ReplaceAll(timeStr, "MM", ZeroFillByStr(t.Month, 2))
	timeStr = strings.ReplaceAll(timeStr, "DD", ZeroFillByStr(t.Day, 2))
	timeStr = strings.ReplaceAll(timeStr, "HH", ZeroFillByStr(t.Hour, 2))
	timeStr = strings.ReplaceAll(timeStr, "mm", ZeroFillByStr(t.Minute, 2))
	timeStr = strings.ReplaceAll(timeStr, "ss", ZeroFillByStr(t.Second, 2))

	return timeStr
}

// 加上时间（传负数可以减）
func (t *DayjsStruct) Add(num int, Type string) *DayjsStruct {
	h1, _ := time.ParseDuration("1h")
	m1, _ := time.ParseDuration("1m")
	s1, _ := time.ParseDuration("1s")
	switch Type {
	case "year":
		t.Time = t.Time.AddDate(num, 0, 0)
	case "month":
		t.Time = t.Time.AddDate(0, num, 0)
	case "day":
		t.Time = t.Time.AddDate(0, 0, num)
	case "hour":
		t.Time = t.Time.Add(h1 * time.Duration(num))
	case "minute":
		t.Time = t.Time.Add(m1 * time.Duration(num))
	case "second":
		t.Time = t.Time.Add(s1 * time.Duration(num))
	}
	// t.Time.Add(year)
	t.SetTime()
	return t
}

// 减去时间（传负数可以加）
func (t *DayjsStruct) Subtract(num int, Type string) *DayjsStruct {
	h1, _ := time.ParseDuration("-1h")
	m1, _ := time.ParseDuration("-1m")
	s1, _ := time.ParseDuration("-1s")
	switch Type {
	case "year":
		t.Time = t.Time.AddDate(-num, 0, 0)
	case "month":
		t.Time = t.Time.AddDate(0, -num, 0)
	case "day":
		t.Time = t.Time.AddDate(0, 0, -num)
	case "hour":
		t.Time = t.Time.Add(h1 * time.Duration(num))
	case "minute":
		t.Time = t.Time.Add(m1 * time.Duration(num))
	case "second":
		t.Time = t.Time.Add(s1 * time.Duration(num))
	}
	// t.Time.Add(year)
	t.SetTime()
	return t
}

// // 设置年月日时分秒
// func (t *DayjsStruct) Set(year int, month int, day int, hour int, minute int, second int) *DayjsStruct {
// 	t.Time = time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local)
// 	t.SetTime()
// 	return t
// }

// Unix纪元以来的毫秒数
func (t *DayjsStruct) ValueOf() int64 {
	return t.Time.UnixMilli()
}

// Unix纪元以来的秒数
func (t *DayjsStruct) Unix() int64 {
	return t.Time.Unix()
}

// 如果 t 代表的时间点在 t2 之前，返回真；
func (t *DayjsStruct) IsBefore(t2 *DayjsStruct) bool {
	return t.Time.Before(t2.Time)
}

// 如果 t 代表的时间点在 t2 之后，返回真
func (t *DayjsStruct) IsAfter(t2 *DayjsStruct) bool {
	return t.Time.After(t2.Time)
}

// 比较时间是否相等，相等返回真
func (t *DayjsStruct) IsSame(t2 *DayjsStruct) bool {
	return t.Time.Equal(t2.Time)
}

// 是否在两个时间范围之间
func (t *DayjsStruct) IsBetween(t2 *DayjsStruct, t3 *DayjsStruct) bool {
	return t.Time.After(t2.Time) && t.Time.Before(t3.Time)
}

// 相同或之前
func (t *DayjsStruct) IsSameOrBefore(t2 *DayjsStruct) bool {
	return t.Time.Before(t2.Time) || t.Time.Equal(t2.Time)
}

// 相同或之后
func (t *DayjsStruct) IsSameOrAfter(t2 *DayjsStruct) bool {
	return t.Time.After(t2.Time) || t.Time.Equal(t2.Time)
}

// 是否闰年
func (t *DayjsStruct) IsLeapYear() bool {
	if t.Year%4 == 0 && t.Year%100 != 0 || t.Year%400 == 0 {
		return true
	}
	return false
}

//获取某月天数
func (t *DayjsStruct) DaysInMonth() int {
	year := t.Year
	month := t.Month
	// 有31天的月份
	day31 := map[int]bool{
		1:  true,
		3:  true,
		5:  true,
		7:  true,
		8:  true,
		10: true,
		12: true,
	}
	if day31[month] == true {
		return 31
	}
	// 有30天的月份
	day30 := map[int]bool{
		4:  true,
		6:  true,
		9:  true,
		11: true,
	}
	if day30[month] == true {
		return 30
	}
	// 计算是平年还是闰年
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		// 得出2月的天数
		return 29
	}
	// 得出2月的天数
	return 28
}

// func main() {
// 	dayTime := Dayjs()
// 	fmt.Println(dayTime.format("YYYY-MM-DD HH:mm:ss"))
// 	fmt.Println(Dayjs("2022年02月28").Add(2, "day").Add(2, "year").Add(2, "month").format("YYYY年MM月DD HH时mm分ss秒"))

// 	fmt.Println(Dayjs("2022年02月28").Subtract(-1, "month").Subtract(2, "hour").format("YYYY年MM月DD HH时mm分ss秒"))

// 	fmt.Println(Dayjs("2022年02月28").Subtract(-2, "month").Subtract(2, "month").DaysInMonth())
// }
