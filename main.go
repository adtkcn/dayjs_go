package main

import (
	"fmt"

	"github.com/adtkcn/dayjs_go/dayjs"
)

func main() {
	dayTime := dayjs.Dayjs()
	dayTime.Now()
	dayTime.Parse("2022年02月28")
	dayjs.Dayjs("2022年02月28")

	dayjs.Dayjs("2022年02月28").ValueOf() //Unix纪元以来的毫秒数
	dayjs.Dayjs("2022年02月28").Unix()    //Unix纪元以来的秒数

	fmt.Printf("%+v", dayTime)

	fmt.Println(dayTime.Year)

	fmt.Println(dayTime.Format())
	fmt.Println(dayjs.Dayjs("2022年02月28").Add(2, "day").Add(2, "year").Add(2, "month").Format("YYYY年MM月DD HH时mm分ss秒"))

	fmt.Println(dayjs.Dayjs("2022年02月28").Subtract(-1, "month").Subtract(2, "hour").Format("YYYY年MM月DD HH时mm分ss秒"))

	fmt.Println(dayjs.Dayjs("2022年02月28").Subtract(-2, "month").Subtract(2, "month").DaysInMonth())
}
