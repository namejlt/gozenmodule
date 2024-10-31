package mtools

func InUint32Slice(input uint32, list []uint32) bool {
	for _, item := range list {
		if input == item {
			return true
		}
	}
	return false
}

func InUint8Slice(input uint8, list []uint8) bool {
	for _, item := range list {
		if input == item {
			return true
		}
	}
	return false
}

func InStringSlice(input string, list []string) bool {
	if len(list) == 0 {
		return false
	}

	for _, item := range list {
		if input == item {
			return true
		}
	}

	return false
}

func FilterStringSlice(inputs []string) (outputs []string) {
	if len(inputs) == 0 {
		return outputs
	}

	m := make(map[string]struct{}, len(inputs))
	outputs = make([]string, 0, len(inputs))
	for _, item := range inputs {
		if _, exist := m[item]; exist {
			continue
		} else {
			m[item] = struct{}{}
			outputs = append(outputs, item)
		}
	}
	return outputs
}

// StringTenSlice uid每10个一组
func StringTenSlice(inputs []string) [][]string {
	var (
		inputList     [][]string
		inputListTemp = make([]string, 0, 10)
	)
	for _, uid := range inputs {
		inputListTemp = append(inputListTemp, uid)
		if len(inputListTemp) >= 10 {
			inputList = append(inputList, inputListTemp)
			inputListTemp = make([]string, 0, 10)
		}
	}
	if len(inputListTemp) > 0 {
		inputList = append(inputList, inputListTemp)
	}
	return inputList
}
