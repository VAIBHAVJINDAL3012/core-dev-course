package main

import (
	"fmt"
	"sort"
	"strings"
)

func countWord(s *string) map[string]int {

	stringArray := strings.Fields(*s)

	mp := make(map[string]int)

	for _, word := range stringArray {
		mp[word] = mp[word] + 1
	}

	return mp
}

func mentionedAtleastOnce(s *[]int) []int {

	slc := make([]int, 0, len(*s))
	mp := make(map[int]bool)

	for _, ele := range *s {

		_, ok := mp[ele]
		if !ok {
			slc = append(slc, ele)
			mp[ele] = true
		}

	}

	return slc

}

func mentionedInBoth(s1 *[]int, s2 *[]int) []int {

	slc := make([]int, 0, len(*s1))
	mp := make(map[int]bool)

	for _, ele := range *s1 {
		mp[ele] = true
	}

	for _, ele := range *s2 {
		val, ok := mp[ele]
		if ok && val == true {
			mp[ele] = false
			slc = append(slc, ele)
		}

	}

	return slc

}

func Fibbonaci(n int, mp *map[int]int) int {

	if n == 1 {
		return 0

	} else if n == 2 {

		return 1

	} else if val, ok := (*mp)[n]; ok {
		return val
	}

	var res int
	res = Fibbonaci(n-1, mp) + Fibbonaci(n-2, mp)
	(*mp)[n] = res

	return res

}

func sorted(mp *map[int](map[int]int)) {

	keys := make([]int, 0, len(*mp))

	for key := range *mp {
		keys = append(keys, key)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	var mapBuff map[int]int

	for _, key := range keys {

		mapBuff = (*mp)[key]
		keys := make([]int, 0, len(*mp))

		for key := range mapBuff {
			keys = append(keys, key)
		}

		sort.Sort(sort.IntSlice(keys))

		for _, ele := range keys {
			fmt.Printf("Fee: %v Nounce %v: Tx:%v", key, ele, mapBuff[ele])
		}

	}

}

func main() {

}
