package dayjs

import (
	"testing"
)

func TestAdd(t *testing.T) {
	timeStr := Dayjs("2022年02月28").Add(-2, "day").Subtract(2, "month").Format("YYYY年MM月DD HH时mm分ss秒")
	gotStr := "2021年12月26 00时00分00秒"
	if timeStr != gotStr {
		t.Errorf("DaysInMonth() failed. Got %s, expected %s", timeStr, gotStr)
	}
}
func TestDaysInMonth(t *testing.T) {
	day := Dayjs("2022年02月28").Subtract(-2, "month").Subtract(2, "month").DaysInMonth()
	if day != 28 {
		t.Errorf("DaysInMonth() failed. Got %d, expected 28", day)
	}
}
