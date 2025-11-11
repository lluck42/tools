package array

func InArray[T comparable](target T, array []T) bool {
	for _, item := range array {
		if item == target {
			return true
		}
	}
	return false
}

// Difference 泛型版本，适用于任何可比较的类型
// 差集计算
func Difference[T comparable](a, b []T) []T {
	setB := make(map[T]bool, len(b))
	for _, item := range b {
		setB[item] = true
	}

	var result []T
	for _, item := range a {
		if !setB[item] {
			result = append(result, item)
		}
	}
	return result
}

// 使用泛型对可比较（comparable）类型的切片去重
func RemoveDuplicates[T comparable](slice []T) []T {
	encountered := make(map[T]bool)
	result := []T{}

	for _, v := range slice {
		if !encountered[v] {
			encountered[v] = true
			result = append(result, v)
		}
	}
	return result
}
