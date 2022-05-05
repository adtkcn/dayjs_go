## Installation, with go 1.18

```
go get github.com/adtkcn/dayjs_go
```

## 实现的函数 Implemented functions

### 方法

1. Format()
2. Add()
3. Subtract()
4. IsBefore()
5. IsAfter()
6. IsSame()
7. IsBetween()
8. IsSameOrBefore()
9. IsSameOrAfter()
10. IsLeapYear()
11. DaysInMonth()
12. Max()
13. Min()
14. Set()
15. Get()
16. Weekday() // 星期几，0 是星期日
17. ToArray()
18. Diff()

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
	fmt.Println(dayjs.Dayjs("2022年02月28").Add(2, "day").Add(2, "year").Add(2, "month").Format("YYYY年MM月DD HH时mm分ss秒"))

	fmt.Println(dayjs.Dayjs("2022年02月28").Subtract(-1, "month").Subtract(2, "hour").Format("YYYY年MM月DD HH时mm分ss秒"))

	fmt.Println(dayjs.Dayjs("2022年02月28").Subtract(-2, "month").Subtract(2, "month").DaysInMonth())
}

```
