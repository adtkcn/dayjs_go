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

	maxTime := dayjs.Max(dayjs.Dayjs().Now(), dayTime, dayjs.Dayjs().Parse("2022年02月28"))
	fmt.Println(maxTime)
	minTime := dayjs.Min(dayjs.Dayjs().Now(), dayTime, dayjs.Dayjs().Parse("2022年02月28"))
	fmt.Println(minTime.Set("Year", 2000))
	fmt.Println("Weekday", dayjs.Dayjs().Add(3, "date").Weekday())
	fmt.Println(minTime.Get("Year"))

	fmt.Println(dayjs.Dayjs().Add(1, "date").Diff(dayjs.Dayjs().Add(3, "date")))
	fmt.Println(dayjs.Dayjs().Add(1, "date").Diff(dayjs.Dayjs().Add(3, "date"), "year"))

	// dayjs.Dayjs("2022年02月28").ValueOf() //Unix纪元以来的毫秒数
	// dayjs.Dayjs("2022年02月28").Unix()    //Unix纪元以来的秒数

	// fmt.Printf("%+v", dayTime)

	// fmt.Println(dayTime.Year)

	// fmt.Println(dayTime.Format())
	// fmt.Println(dayjs.Dayjs("2022年02月28").Add(2, "date").Add(2, "year").Add(2, "month").Format("YYYY年MM月DD HH时mm分ss秒"))

	// fmt.Println(dayjs.Dayjs("2022年02月28").Subtract(-1, "month").Subtract(2, "hour").Format("YYYY年MM月DD HH时mm分ss秒"))

	// fmt.Println(dayjs.Dayjs("2022年02月28").Subtract(-2, "month").Subtract(2, "month").DaysInMonth())
}
