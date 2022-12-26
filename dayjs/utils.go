package dayjs

import "fmt"

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
