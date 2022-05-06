package dayjs

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type DayjsStruct struct {
	Year   int `json:"year"`   //年
	Month  int `json:"month"`  //月
	Date   int `json:"date"`   //日
	Day    int `json:"day"`    //周几
	Hour   int `json:"hour"`   //小时
	Minute int `json:"minute"` //分钟
	Second int `json:"second"` //秒

	Time time.Time
}

func Dayjs(timeStr ...interface{}) *DayjsStruct {
	dayTime := &DayjsStruct{}

	if len(timeStr) == 1 {
		// 待区分时间戳和字符串时间

		switch timeStr[0].(type) {
		case string:
			dayTime.Parse(fmt.Sprint(timeStr[0]))
		case int64:
			dayTime.ParseUnix(timeStr[0].(int64))
		case int:
			dayTime.ParseUnix(int64(timeStr[0].(int)))
		default:
			panic("时间格式有误")
		}
	} else {
		dayTime.Now()
	}

	return dayTime
}

// 前置补零
func ZeroFill(str interface{}, resultLen int) string {
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
			year = ZeroFill(val, 4)
		} else if key == 1 {
			month = ZeroFill(val, 2)
		} else if key == 2 {
			day = ZeroFill(val, 2)
		} else if key == 3 {
			hour = ZeroFill(val, 2)
		} else if key == 4 {
			minute = ZeroFill(val, 2)
		} else if key == 5 {
			second = ZeroFill(val, 2)
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

// 解析秒级时间戳
func (t *DayjsStruct) ParseUnix(unix int64) *DayjsStruct {
	t.Time = time.Unix(unix, 0) // 一个参数是时间戳（秒），第二个参数是纳秒，设置0即可
	fmt.Println("unix", unix, t.Time)
	t.SetTime()
	return t
}

// 设置为当前时间
func (t *DayjsStruct) Now() *DayjsStruct {
	t.Time = time.Now() // 获取当前时间
	t.SetTime()
	return t
}
func (t *DayjsStruct) SetTime() {
	t.Year = t.Time.Year()        //年
	t.Month = int(t.Time.Month()) //月
	t.Date = t.Time.Day()         //日
	t.Hour = t.Time.Hour()        //小时
	t.Minute = t.Time.Minute()    //分钟
	t.Second = t.Time.Second()    //秒

	t.Day = int(t.Time.Weekday())
}

// 格式化 YYYY-MM-DD HH:mm:ss
func (t *DayjsStruct) Format(format ...string) string {
	formatStr := "YYYY-MM-DD HH:mm:ss"
	if len(format) == 1 && format[0] != "" {
		formatStr = format[0]
	}
	timeStr := strings.ReplaceAll(formatStr, "YYYY", fmt.Sprint(t.Year))
	timeStr = strings.ReplaceAll(timeStr, "MM", ZeroFill(t.Month, 2))
	timeStr = strings.ReplaceAll(timeStr, "DD", ZeroFill(t.Date, 2))
	timeStr = strings.ReplaceAll(timeStr, "HH", ZeroFill(t.Hour, 2))
	timeStr = strings.ReplaceAll(timeStr, "mm", ZeroFill(t.Minute, 2))
	timeStr = strings.ReplaceAll(timeStr, "ss", ZeroFill(t.Second, 2))

	return timeStr
}

// 加上时间（传负数可以减）
func (t *DayjsStruct) Add(num int, Type string) *DayjsStruct {
	h1, _ := time.ParseDuration("1h")
	m1, _ := time.ParseDuration("1m")
	s1, _ := time.ParseDuration("1s")
	typeStr := strings.ToLower(Type)
	switch typeStr {
	case "year":
		t.Time = t.Time.AddDate(num, 0, 0)
	case "month":
		t.Time = t.Time.AddDate(0, num, 0)
	case "date":
		t.Time = t.Time.AddDate(0, 0, num)
	case "day":
		panic("Add 暂不支持 day")
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
	typeStr := strings.ToLower(Type)
	switch typeStr {
	case "year":
		t.Time = t.Time.AddDate(-num, 0, 0)
	case "month":
		t.Time = t.Time.AddDate(0, -num, 0)
	case "date":
		t.Time = t.Time.AddDate(0, 0, -num)
	case "day":
		panic("Add 暂不支持 day")
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

// 设置年月日时分秒，
func (t *DayjsStruct) Set(Type string, value int) *DayjsStruct {
	typeStr := strings.ToLower(Type)
	switch typeStr {
	case "year":
		t.Time = time.Date(value, time.Month(t.Month), t.Date, t.Hour, t.Minute, t.Second, 0, time.Local)
	case "month":
		t.Time = time.Date(t.Year, time.Month(value), t.Date, t.Hour, t.Minute, t.Second, 0, time.Local)
	// case "day":
	// 	t.Time = time.Date(t.Year, time.Month(t.Month), value, t.Hour, t.Minute, t.Second, 0, time.Local)
	case "date":
		t.Time = time.Date(t.Year, time.Month(t.Month), value, t.Hour, t.Minute, t.Second, 0, time.Local)
	case "hour":
		t.Time = time.Date(t.Year, time.Month(t.Month), t.Date, value, t.Minute, t.Second, 0, time.Local)
	case "minute":
		t.Time = time.Date(t.Year, time.Month(t.Month), t.Date, t.Hour, value, t.Second, 0, time.Local)
	case "second":
		t.Time = time.Date(t.Year, time.Month(t.Month), t.Date, t.Hour, t.Minute, value, 0, time.Local)
	default:
		panic("Dayjs().Set(Type,value) Type Error: " + Type)
	}
	t.SetTime()
	return t
}

func (t *DayjsStruct) Get(Type string) int64 {
	typeStr := strings.ToLower(Type)
	switch typeStr {
	case "year":
		return int64(t.Year)
	case "month":
		return int64(t.Month)
	case "day":
		return int64(t.Weekday())
	case "date":
		return int64(t.Date)
	case "hour":
		return int64(t.Hour)
	case "minute":
		return int64(t.Minute)
	case "second":
		return int64(t.Second)
	}
	panic("Dayjs().Get(Type) Type Error ：" + Type)
	// return int64(0)
}

// 周几0-6，0是星期日
func (t *DayjsStruct) Weekday() int64 {
	return int64(t.Day)
}

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

// 差异
func (t *DayjsStruct) Diff(t2 *DayjsStruct, Type ...string) int64 {
	typeStr := "millisecond"
	if len(Type) == 1 && Type[0] != "" {
		typeStr = strings.ToLower(Type[0])
	}
	// diffTime :=t.ValueOf() - t2.ValueOf()
	switch typeStr {
	case "millisecond":
		return t.ValueOf() - t2.ValueOf()
	case "second":
		return t.Unix() - t2.Unix()
	case "minute":
		return (t.Unix() - t2.Unix()) / 60
	case "hour":
		return (t.Unix() - t2.Unix()) / 60 / 60
	case "date":
		return (t.Unix() - t2.Unix()) / 60 / 60 / 24
	case "day":
		panic("Diff 暂不支持 day")
	case "week":
		panic("Diff 暂不支持 week")
	case "quarter":
		panic("Diff 暂不支持 quarter")
	// 	return ( t.Unix() - t2.Unix()) / 60 / 60 / 24 / 7
	case "month":
		return (t.Unix() - t2.Unix()) / 60 / 60 / 24 / 30
	case "year":
		return (t.Unix() - t2.Unix()) / 60 / 60 / 24 / 365
	}

	return t.ValueOf() - t2.ValueOf()
}

// 是否闰年
func (t *DayjsStruct) IsLeapYear() bool {
	if t.Year%4 == 0 && t.Year%100 != 0 || t.Year%400 == 0 {
		return true
	}
	return false
}

// 转数组
func (t *DayjsStruct) ToArray() []int {
	return []int{t.Year, t.Month, t.Date, t.Hour, t.Minute, t.Second}
}

// 克隆
func (t *DayjsStruct) Clone() *DayjsStruct {
	return &DayjsStruct{
		Time:   t.Time,
		Year:   t.Year,
		Month:  t.Month,
		Date:   t.Date,
		Hour:   t.Hour,
		Minute: t.Minute,
		Second: t.Second,
	}
}

// StartOf
func (t *DayjsStruct) StartOf(Type string) *DayjsStruct {
	typeStr := strings.ToLower(Type)
	switch typeStr {
	case "year":
		return t.Set("month", 1).Set("date", 1).Set("hour", 0).Set("minute", 0).Set("second", 0)
	case "quarter":
		panic("StartOf 暂不支持 quarter")
		// return t.Set("month", 1).Set("date", 1).Set("hour", 0).Set("minute", 0).Set("second", 0)
	case "month":
		return t.Set("date", 1).Set("hour", 0).Set("minute", 0).Set("second", 0)
	case "week":
		panic("StartOf 暂不支持 week")
		// return t.Set("weekday", 0).Set("hour", 0).Set("minute", 0).Set("second", 0)
	case "date":
		return t.Set("hour", 0).Set("minute", 0).Set("second", 0)
	case "day":
		return t.Set("hour", 0).Set("minute", 0).Set("second", 0)
	case "hour":
		return t.Set("minute", 0).Set("second", 0)
	case "minute":
		return t.Set("second", 0)
	}
	panic("Dayjs().StartOf(Type) Type Error ：" + Type)
}

// EndOf
func (t *DayjsStruct) EndOf(Type string) *DayjsStruct {
	typeStr := strings.ToLower(Type)
	switch typeStr {
	case "year":
		return t.Set("month", 12).Set("date", 31).Set("hour", 23).Set("minute", 59).Set("second", 59)
	case "quarter":
		panic("EndOf 暂不支持 quarter")
		// return t.Set("month", 3).Set("date", 31).Set("hour", 23).Set("minute", 59).Set("second", 59)
	case "month":
		return t.Set("date", 31).Set("hour", 23).Set("minute", 59).Set("second", 59)
	case "week":
		panic("EndOf 暂不支持 week")
		// return t.Set("weekday", 6).Set("hour", 23).Set("minute", 59).Set("second", 59)
	case "date":
		return t.Set("hour", 23).Set("minute", 59).Set("second", 59)
	case "day":
		return t.Set("hour", 23).Set("minute", 59).Set("second", 59)
	case "hour":
		return t.Set("minute", 59).Set("second", 59)
	case "minute":
		return t.Set("second", 59)
	}
	panic("Dayjs().EndOf(Type) Type Error ：" + Type)
}

//获取某月天数
func (t *DayjsStruct) DaysInMonth() int64 {
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
		return int64(31)
	}
	// 有30天的月份
	day30 := map[int]bool{
		4:  true,
		6:  true,
		9:  true,
		11: true,
	}
	if day30[month] == true {
		return int64(30)
	}
	// 计算是平年还是闰年
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		// 得出2月的天数
		return int64(29)
	}
	// 得出2月的天数
	return int64(28)
}

// 获取季度
func (t *DayjsStruct) Quarter() int {
	month := t.Month
	if month <= 3 {
		return 1
	} else if month <= 6 {
		return 2
	} else if month <= 9 {
		return 3
	} else {
		return 4
	}
}

// fromNow
func (t *DayjsStruct) FromNow() string {
	now := Dayjs()
	diffYear := now.Year - t.Year
	diffMonth := now.Month - t.Month
	diffDate := now.Date - t.Date
	diffHour := now.Hour - t.Hour
	diffMinute := now.Minute - t.Minute
	diffSecond := now.Second - t.Second
	fmt.Println(diffYear, diffMonth, diffDate, diffHour, diffMinute, diffSecond)
	if diffYear > 0 {
		return fmt.Sprintf("%d年前", diffYear)
	} else if diffYear < 0 {
		return fmt.Sprintf("%d年后", -diffYear)
	} else if diffMonth > 0 {
		return fmt.Sprintf("%d个月前", diffMonth)
	} else if diffMonth < 0 {
		return fmt.Sprintf("%d个月后", -diffMonth)
	} else if diffDate > 0 {
		return fmt.Sprintf("%d天前", diffDate)
	} else if diffDate < 0 {
		return fmt.Sprintf("%d天后", -diffDate)
	} else if diffHour > 0 {
		return fmt.Sprintf("%d小时前", diffHour)
	} else if diffHour < 0 {
		return fmt.Sprintf("%d小时后", -diffHour)
	} else if diffMinute > 0 {
		return fmt.Sprintf("%d分钟前", diffMinute)
	} else if diffMinute < 0 {
		return fmt.Sprintf("%d分钟后", -diffMinute)
	} else if diffSecond > 0 {
		return fmt.Sprintf("%d秒前", diffSecond)
	} else if diffSecond < 0 {
		return fmt.Sprintf("%d秒后", -diffSecond)
	} else {
		return "刚刚"
	}
}

// 获取最大时间
func Max(dayjs ...*DayjsStruct) *DayjsStruct {
	var max DayjsStruct
	for k, v := range dayjs {
		if k == 0 {
			max = *v
		}
		if v.IsAfter(&max) {
			max = *v
		}
	}
	return &max
}

// 获取最小时间
func Min(dayjs ...*DayjsStruct) *DayjsStruct {
	var min DayjsStruct
	for k, v := range dayjs {
		if k == 0 {
			min = *v
		}
		if v.IsBefore(&min) {
			min = *v
		}
	}
	return &min
}
