package tools

// 查询list中是否程存在元素
func InList[T int | uint | string](arr []T, target T) bool {
	for _, item := range arr {
		if item == target {
			return true
		}
	}
	return false
}

// list排重
func UniqueList[T uint | string](list []T) []T {
	keys := make(map[T]bool)
	var uniqueList []T

	for _, entry := range list {
		if _, found := keys[entry]; !found {
			keys[entry] = true
			uniqueList = append(uniqueList, entry)
		}
	}

	return uniqueList
}
