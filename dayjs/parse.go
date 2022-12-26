package dayjs

import (
	"fmt"
	"regexp"
	"time"
)

// 解析时间
// @param {string|*DayjsStruct|int64|int} str 时间字符串

func Parse(str interface{}) *DayjsStruct {
	var dayjsStruct *DayjsStruct

	switch str.(type) {
	case string:
		dayjsStruct = ParseString(fmt.Sprint(str))
	case int64:
		dayjsStruct = ParseUnix(str.(int64))
	case int:
		dayjsStruct = ParseUnix(int64(str.(int)))
	case *DayjsStruct:
		dayTimetemplate := str.(*DayjsStruct)
		dayjsStruct = dayTimetemplate.Clone()
	default:
		// TODO: 可能需要支持time.Time格式
		panic("时间格式有误")
	}
	return dayjsStruct
}

// 解析时间，每个时间需要任意字符分开； YYYY年MM月DD HH时mm分ss秒
func ParseString(str string) *DayjsStruct {
	dayjsStruct := &DayjsStruct{}

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
	dayjsStruct.Time = strTime
	dayjsStruct.SetTime()
	return dayjsStruct
}

// 解析秒级时间戳
func ParseUnix(unix int64) *DayjsStruct {
	dayjsStruct := &DayjsStruct{}
	dayjsStruct.Time = time.Unix(unix, 0) // 一个参数是时间戳（秒），第二个参数是纳秒，设置0即可
	dayjsStruct.SetTime()
	return dayjsStruct
}

// 解析毫秒级时间戳
func ParseUnixMilli(unix int64) *DayjsStruct {
	dayjsStruct := &DayjsStruct{}
	dayjsStruct.Time = time.Unix(unix/1000, 0) // 一个参数是时间戳（秒），第二个参数是纳秒，设置0即可
	dayjsStruct.SetTime()
	return dayjsStruct
}

// 设置为当前时间
func Now() *DayjsStruct {
	dayjsStruct := &DayjsStruct{}
	dayjsStruct.Time = time.Now() // 获取当前时间
	dayjsStruct.SetTime()
	return dayjsStruct
}
