package dayjs

import (
	"testing"
)

func TestDayjs(t *testing.T) {
	timeStr := Dayjs("2022年02月28").Format()
	gotStr := "2022-02-28 00:00:00"
	if timeStr != gotStr {
		t.Errorf("Dayjs('2022年02月28') failed. Got %s, expected %s", timeStr, gotStr)
	}
	timeStr = Dayjs(1651909230).Format()
	gotStr = "2022-05-07 15:40:30"
	if timeStr != gotStr {
		t.Errorf("Dayjs(1651909230) failed. Got %s, expected %s", timeStr, gotStr)
	}
	timeStr = Dayjs(int64(1651909230)).Format()
	gotStr = "2022-05-07 15:40:30"
	if timeStr != gotStr {
		t.Errorf("Dayjs(int64(1651909230)) failed. Got %s, expected %s", timeStr, gotStr)
	}

	timeStr = Dayjs(timeStr).Format()
	gotStr = "2022-05-07 15:40:30"
	if timeStr != gotStr {
		t.Errorf("Dayjs(*DayjsStruct) failed. Got %s, expected %s", timeStr, gotStr)
	}
}
func TestParse(t *testing.T) {
	timeStr := Dayjs().Parse("2022年02月28").Format()
	gotStr := "2022-02-28 00:00:00"
	if timeStr != gotStr {
		t.Errorf("Parse() failed. Got %s, expected %s", timeStr, gotStr)
	}
}
func TestAdd(t *testing.T) {
	timeStr := Dayjs("2022年02月28").Add(-2, "date").Subtract(2, "month").Format("YYYY年MM月DD HH时mm分ss秒")
	gotStr := "2021年12月26 00时00分00秒"
	if timeStr != gotStr {
		t.Errorf("DaysInMonth() failed. Got %s, expected %s", timeStr, gotStr)
	}
}

func TestDaysInMonth(t *testing.T) {
	day := Dayjs("2022年02月20").DaysInMonth()
	if day != 28 {
		t.Errorf("DaysInMonth() failed. Got %d, expected 28", day)
	}
	day = Dayjs("2020年02月20").DaysInMonth()
	if day != 29 {
		t.Errorf("DaysInMonth() failed. Got %d, expected 29", day)
	}
}

func TestZeroFill(t *testing.T) {
	timeStr := ZeroFill(0, 6)
	gotStr := "000000"
	if timeStr != gotStr {
		t.Errorf("ZeroFill(0,6) failed. Got %s, expected %s", timeStr, gotStr)
	}
	timeStr = ZeroFill("0", 6)
	// gotStr := "000000"
	if timeStr != gotStr {
		t.Errorf("ZeroFill('0',6) failed. Got %s, expected %s", timeStr, gotStr)
	}
}

func TestMax(t *testing.T) {
	timeStr := Max(Dayjs("2020年02月28"), Dayjs("2020年02月29")).Format()
	gotStr := "2020-02-29 00:00:00"
	if timeStr != gotStr {
		t.Errorf("Max() failed. Got %s, expected %s", timeStr, gotStr)
	}
}
func TestMin(t *testing.T) {
	timeStr := Min(Dayjs("2020年02月28"), Dayjs("2020年02月29")).Format()
	gotStr := "2020-02-28 00:00:00"
	if timeStr != gotStr {
		t.Errorf("Max() failed. Got %s, expected %s", timeStr, gotStr)
	}
}
