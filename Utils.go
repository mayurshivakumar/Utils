package main

import (
	"sort"
	"strconv"
)

type IUtils interface {
	IntArrayToInterface(array []int) []interface{}
	IntArrayToString(array []int) string
	GetIntKeyArray(mymap map[int]int) []int
	SortedKeys(m map[string]float64) []string
}
type Utils struct {
	IUtils
}

func GetUtils() *Utils {
	var utils = Utils{}
	return &utils
}

// Gets a int array and returns
// interface
func (utils *Utils) IntArrayToInterface(array []int) []interface{} {
	res := make([]interface{}, len(array))
	for i := range array {
		res[i] = array[i]
	}
	return res
}

// Gets a int array and returns
// string
func (utils *Utils) IntArrayToString(array []int) string {
	var str string
	for _, value := range array {
		if str == "" {
			str += strconv.Itoa(value)
		} else {
			str += "," + strconv.Itoa(value)
		}
	}
	return str
}

// Gets a map[int]int{} and returns
// int array of keys.
func (utils *Utils) GetIntKeyArray(mymap map[int]int) []int {
	keys := make([]int, len(mymap))
	i := 0
	for k := range mymap {
		keys[i] = k
		i += 1
	}
	return keys
}


type sortedMap struct {
	m map[string]float64
	s []string
}

func (sm *sortedMap) Len() int {
	return len(sm.m)
}

func (sm *sortedMap) Less(i, j int) bool {
	return sm.m[sm.s[i]] > sm.m[sm.s[j]]
}

func (sm *sortedMap) Swap(i, j int) {
	sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
}

// Sorts a given map by value in desc and returns the key
func SortedKeys(m map[string]float64) []string {
	sm := new(sortedMap)
	sm.m = m
	sm.s = make([]string, len(m))
	i := 0
	for key, _ := range m {
		sm.s[i] = key
		i++
	}
	sort.Sort(sm)
	return sm.s
}
