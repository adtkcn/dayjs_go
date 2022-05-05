## 实现的函数

1. Format
2. Add
3. Subtract
4. IsBefore
5. IsAfter
6. IsSame
7. IsBetween
8. IsSameOrBefore
9. IsSameOrAfter
10. IsLeapYear
11. DaysInMonth

## 文档参考 dayjs

https://dayjs.fenxianglu.cn/

## 调用示例

```go
import (
	"dayjs_go/dayjs"
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
