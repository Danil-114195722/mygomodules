package arrays


func StringInArray(arr []string, value string) (bool) {
	for _, arrValue := range arr {
		if arrValue == value {
			return true
		}
	}
	return false
}

func IntInArray(arr []int, value int) (bool) {
	for _, arrValue := range arr {
        if arrValue == value {
            return true
        }
    }
    return false
}

func FloatInArray(arr []float64, value float64) (bool) {
    for _, arrValue := range arr {
        if arrValue == value {
            return true
        }
    }
    return false
}
