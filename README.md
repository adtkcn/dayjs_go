## Installation

```
go get github.com/adtkcn/dayjs_go
```

💪 不可变数据 (Immutable)

🔥 支持链式操作 (Chainable)

🌐 不支持 I18n 国际化，只有 Format 函数格式化

📦 纯 go 实现的仅 十几 kb 大小的微型库

## 实现的功能 Implemented functions

### 方法

```go
t := dayjs.Dayjs()
```

```go
t.Format() 			// 格式化： YYYY-MM-DD HH:mm:ss
t.Add()		 		// 增加时间
t.Subtract() 		// 减少时间
t.IsBefore(t2) 		// t 是否在 t2 之前
t.IsAfter(t2) 		// t 是否在 t2 之后
t.IsSame(t2) 		// t 是否与 t2 相同
t.IsBetween(t2,t3)  // t 是否在 t2 和 t3 之间
t.IsSameOrBefore() 	// t 是否在 t2 之前或者与 t2 相同
t.IsSameOrAfter() 	// t 是否在 t2 之后或者与 t2 相同
t.IsLeapYear() 		// t 是否为闰年
t.DaysInMonth() 	// t 的月份的天数
t.Set(Type,value) 	// 设置时间,Type 可以是"year","month","date","hour","minute","second"
t.Get(Type) 		// 获取时间,Type 可以是"year","month","day","date","hour","minute","second"
t.Weekday() 		// 星期几，0 是星期日
t.ToArray() 		// 转换成数组 [year,month,day,hour,minute,second]
t.Diff(t2) 			// t2 与 t 的差值
t.Clone() 			// 克隆
t.StartOf(Type) 	// 开始时间,Type 可以是"year","month","date","hour","minute","second"
t.EndOf(Type) 		// 结束时间,Type 可以是"year","month","date","hour","minute","second"
t.Quarter() 		// 获取季度 1，2，3，4
t.FromNow() 		// 从现在开始返回相对时间的字符串。(2 小时前)

dayjs.Max(t2,t3,t4,...) // 最大值
dayjs.Min(t2,t3,t4,...) // 最小值


dayjs.Now() 			// 当前时间
dayjs.Parse(interface) 	// 解析时间,支持ParseString和ParseUnix的参数
dayjs.ParseString("2022年12月25日 23:59:59") 	// 解析字符串时间
dayjs.ParseUnix( int64 ) 						// 解析秒级时间戳
dayjs.ParseUnixMilli( int64 ) 					// 解析毫秒级时间戳

```

### 属性

1. Year
2. Month
3. Date
4. Hours
5. Minutes
6. Seconds
7. Day
8. Time

| Type   | 说明                     |
| ------ | ------------------------ |
| date   | 日期                     |
| day    | 星期(星期日 0，星期六 6) |
| month  | 月份(0-11)               |
| year   | 年                       |
| hour   | 小时                     |
| minute | 分钟                     |
| second | 秒                       |

## 文档参考 dayjs, Document reference dayjs

https://dayjs.fenxianglu.cn/

## 调用示例 example

```go
import (
	"github.com/adtkcn/dayjs_go/dayjs"
	"fmt"
)
func main() {
	dayTime := dayjs.Dayjs()
	fmt.Printf("%+v", dayTime)

	fmt.Println(dayTime.Year)

	fmt.Println(dayTime.Format())
	fmt.Println(dayjs.Dayjs("2022年02月28").Add(-2, "date").Add(2, "year").Add(2, "month").Format("YYYY年MM月DD HH时mm分ss秒"))

	fmt.Println(dayjs.Dayjs("2022年02月28").Subtract(-1, "month").Subtract(2, "hour").Format("YYYY-MM-DD HH-mm-ss"))

	fmt.Println(dayjs.Dayjs("2022年02月28").DaysInMonth())
}

```
