package util

import (
	"errors"
	"strconv"
	"sync"
)

var (
	sliceUtilInstance *Slice
	sliceOnce sync.Once
)

// SliceUtil 切片工具
func SliceUtil() *Slice {
	sliceOnce.Do(func() {
		sliceUtilInstance = new(Slice)
	})
	return sliceUtilInstance
}

type Slice struct {

}

// SliceToMap 切片转Map
func (u Slice) SliceToMap(stringSlice []string) map[string]string {
	stringMap := make(map[string]string)
	for i := range stringSlice {
		stringMap[strconv.Itoa(i)] = stringSlice[i]
	}
	return stringMap
}

// SliceToMapWithKey 切片转Map
func (u Slice) SliceToMapWithKey(stringSlice []string, keys []string) (stringMap map[string]string, err error) {
	if len(stringSlice) != len(keys) {
		return nil, errors.New("keys slice length discord")
	}

	stringMap = make(map[string]string)
	for i := range stringSlice {
		stringMap[keys[i]] = stringSlice[i]
	}

	return stringMap, nil
}
