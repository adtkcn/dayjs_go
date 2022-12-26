package dayjs

import (
	"fmt"
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

	Time time.Time `json:"-"`
}

/**
* 解析时间
* @param {string|*DayjsStruct|int64|int} timeStr 时间字符串
 */
func Dayjs(timeStr ...interface{}) *DayjsStruct {

	if len(timeStr) >= 1 {
		// 待区分时间戳和字符串时间
		return Parse(timeStr[0])
	} else {
		return Now()
	}

}

// TODO: 考虑换为私有
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
	dayTime := t.Clone()

	h1, _ := time.ParseDuration("1h")
	m1, _ := time.ParseDuration("1m")
	s1, _ := time.ParseDuration("1s")
	typeStr := strings.ToLower(Type)
	switch typeStr {
	case "year":
		dayTime.Time = dayTime.Time.AddDate(num, 0, 0)
	case "month":
		dayTime.Time = dayTime.Time.AddDate(0, num, 0)
	case "date":
		dayTime.Time = dayTime.Time.AddDate(0, 0, num)
	case "day":
		panic("Add 暂不支持 day")
	case "hour":
		dayTime.Time = dayTime.Time.Add(h1 * time.Duration(num))
	case "minute":
		dayTime.Time = dayTime.Time.Add(m1 * time.Duration(num))
	case "second":
		dayTime.Time = dayTime.Time.Add(s1 * time.Duration(num))
	}
	dayTime.SetTime()
	return dayTime
}

// 减去时间（传负数可以加）
func (t *DayjsStruct) Subtract(num int, Type string) *DayjsStruct {
	dayTime := t.Clone()

	h1, _ := time.ParseDuration("-1h")
	m1, _ := time.ParseDuration("-1m")
	s1, _ := time.ParseDuration("-1s")
	typeStr := strings.ToLower(Type)
	switch typeStr {
	case "year":
		dayTime.Time = dayTime.Time.AddDate(-num, 0, 0)
	case "month":
		dayTime.Time = dayTime.Time.AddDate(0, -num, 0)
	case "date":
		dayTime.Time = dayTime.Time.AddDate(0, 0, -num)
	case "day":
		panic("Add 暂不支持 day")
	case "hour":
		dayTime.Time = dayTime.Time.Add(h1 * time.Duration(num))
	case "minute":
		dayTime.Time = dayTime.Time.Add(m1 * time.Duration(num))
	case "second":
		dayTime.Time = dayTime.Time.Add(s1 * time.Duration(num))
	}
	dayTime.SetTime()
	return dayTime
}

// 设置年月日时分秒，
func (t *DayjsStruct) Set(Type string, value int) *DayjsStruct {
	t = t.Clone()
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
func (t *DayjsStruct) IsBefore(t2 *DayjsStruct, Type ...string) bool {
	typeStr := "millisecond"
	if len(Type) == 1 && Type[0] != "" {
		typeStr = strings.ToLower(Type[0])
	}
	diffVal := t.Diff(t2, typeStr)
	// fmt.Println("diffVal", diffVal)
	return diffVal < 0
	// return t.Time.Before(t2.Time)
}

// 如果 t 代表的时间点在 t2 之后，返回真
func (t *DayjsStruct) IsAfter(t2 *DayjsStruct, Type ...string) bool {
	typeStr := "millisecond"
	if len(Type) == 1 && Type[0] != "" {
		typeStr = strings.ToLower(Type[0])
	}
	diffVal := t.Diff(t2, typeStr)
	// fmt.Println("diffVal", diffVal)
	return diffVal > 0
	// return t.Time.After(t2.Time)
}

// 比较时间是否相等，相等返回真
func (t *DayjsStruct) IsSame(t2 *DayjsStruct, Type ...string) bool {
	typeStr := "millisecond"
	if len(Type) == 1 && Type[0] != "" {
		typeStr = strings.ToLower(Type[0])
	}
	diffVal := t.Diff(t2, typeStr)
	// fmt.Println("diffVal", diffVal)
	return diffVal == 0
	// return t.Time.Equal(t2.Time)
}

// 是否在两个时间范围之间
func (t *DayjsStruct) IsBetween(t2 *DayjsStruct, t3 *DayjsStruct, Type ...string) bool {
	typeStr := "millisecond"
	if len(Type) == 1 && Type[0] != "" {
		typeStr = strings.ToLower(Type[0])
	}
	diffVal_2 := t.Diff(t2, typeStr)
	diffVal_3 := t.Diff(t3, typeStr)

	// fmt.Println("diffVal", diffVal_2, diffVal_3)
	return diffVal_2 >= 0 && diffVal_3 <= 0
	// return t.Time.After(t2.Time) && t.Time.Before(t3.Time)
}

// 相同或之前
func (t *DayjsStruct) IsSameOrBefore(t2 *DayjsStruct, Type ...string) bool {
	typeStr := "millisecond"
	if len(Type) == 1 && Type[0] != "" {
		typeStr = strings.ToLower(Type[0])
	}
	diffVal := t.Diff(t2, typeStr)
	// fmt.Println("diffVal", diffVal)
	return diffVal <= 0
	// return t.Time.Before(t2.Time) || t.Time.Equal(t2.Time)
}

// 相同或之后
func (t *DayjsStruct) IsSameOrAfter(t2 *DayjsStruct, Type ...string) bool {
	typeStr := "millisecond"
	if len(Type) == 1 && Type[0] != "" {
		typeStr = strings.ToLower(Type[0])
	}
	diffVal := t.Diff(t2, typeStr)
	// fmt.Println("diffVal", diffVal)
	return diffVal >= 0
	// return t.Time.After(t2.Time) || t.Time.Equal(t2.Time)
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
	time := &DayjsStruct{
		Time:   t.Time,
		Year:   t.Year,
		Month:  t.Month,
		Date:   t.Date,
		Hour:   t.Hour,
		Minute: t.Minute,
		Second: t.Second,
	}
	return time
}

// StartOf
// TODO: 调用了过多Set,每一个Set创建了一个示例
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
// TODO: 调用了过多Set,每一个Set创建了一个示例
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

// 获取某月天数
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
	if day31[month] {
		return int64(31)
	}
	// 有30天的月份
	day30 := map[int]bool{
		4:  true,
		6:  true,
		9:  true,
		11: true,
	}
	if day30[month] {
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
func (t *DayjsStruct) Quarter() int64 {
	month := t.Month
	if month <= 3 {
		return int64(1)
	} else if month <= 6 {
		return int64(2)
	} else if month <= 9 {
		return int64(3)
	} else {
		return int64(4)
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
	// fmt.Println(diffYear, diffMonth, diffDate, diffHour, diffMinute, diffSecond)
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

// fmt.Println自动调用
func (t *DayjsStruct) String() string {
	return t.Format("YYYY-MM-DD HH:mm:ss")
}
