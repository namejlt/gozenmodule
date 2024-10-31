package mtools

import "strconv"

// FenToYuan 金额分转换成元字符串
func FenToYuan(money float64) string {
	money = money / 100
	return strconv.FormatFloat(money, 'f', 2, 64)
}

// MinInt64 取最小值
func MinInt64(inputs ...int64) int64 {
	if len(inputs) == 0 {
		return 0
	}
	min := inputs[0]
	for _, item := range inputs {
		if item < min {
			min = item
		}
	}
	return min
}

// MinUint16 取最小值
func MinUint16(inputs ...uint16) uint16 {
	if len(inputs) == 0 {
		return 0
	}
	min := inputs[0]
	for _, item := range inputs {
		if item < min {
			min = item
		}
	}
	return min
}

// SafeReduceUint16 安全的减法
func SafeReduceUint16(bigVal, smallVal uint16) uint16 {
	if bigVal > smallVal {
		return bigVal - smallVal
	}
	return 0
}

// SafeReduceUint32 安全的减法
func SafeReduceUint32(bigVal, smallVal uint32) uint32 {
	if bigVal > smallVal {
		return bigVal - smallVal
	}
	return 0
}

// SafeReduceFloat64 安全的减法
func SafeReduceFloat64(bigVal, smallVal float64) float64 {
	if bigVal > smallVal {
		return bigVal - smallVal
	}
	return 0
}

// SafeReduceInt64 安全的减法
func SafeReduceInt64(bigVal, smallVal int64) int64 {
	if bigVal > smallVal {
		return bigVal - smallVal
	}
	return 0
}

// SafeReduceTime 安全的减法
func SafeReduceTime(bigVal, smallVal uint32) uint32 {
	if bigVal > smallVal {
		return bigVal - smallVal
	}
	return 0
}
