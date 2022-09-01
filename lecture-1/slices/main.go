package main

import (
	"fmt"
	"math"
	"sort"
)

func appendAtEnd(s *[]int, n int) []int {

	if s == nil {
		return nil
	}

	slc := append(*s, n)

	return slc
}

func appendAtStart(s *[]int, n int) []int {

	if s == nil {
		return nil
	}

	slc := append(*s, (*s)[len(*s)-1])
	for i := len(slc) - 2; i >= 0; i = i - 1 {
		slc[i+1] = slc[i]
	}
	slc[0] = n
	return slc
}

func removeLastElem(s *[]int) []int {

	if s == nil {
		return nil
	}

	fmt.Print("Last element", (*s)[len(*s)-1])
	return (*s)[:len(*s)-1]
}

func removeFirstElem(s *[]int) []int {

	if s == nil {
		return nil
	}

	fmt.Print("Last element", (*s)[0])
	for i := 0; i < len(*s)-1; i = i + 1 {
		(*s)[i] = (*s)[i+1]
	}
	return (*s)[:len(*s)-1]
}

func mergeWithDuplicate(slc1 []int, slc2 []int) []int {

	mp := make(map[int]bool)

	for _, elem := range slc1 {
		mp[elem] = true
	}

	for _, elem := range slc2 {

		_, ok := mp[elem]
		if !ok {
			slc1 = append(slc1, elem)
		}

	}

	return slc1

}

func removeFromSlc(slc1 []int, slc2 []int) []int {

	mp := make(map[int]bool)

	for _, elem := range slc2 {
		mp[elem] = true
	}

	var i int = 0
	var j int
	for ; i < len(slc1); i = i + 1 {
		_, ok := mp[slc1[i]]
		if ok {
			j = i
			i = i + 1
			break
		}
	}

	var count int
	for ; i < len(slc1); i = i + 1 {

		_, ok := mp[slc1[i]]
		if ok {
			slc1[j] = slc1[i]
			j = j + 1
			count = count + 1
		}

	}

	slc1 = slc1[:len(slc1)-count]
	return slc1

}

func leftShiftBy1(s *[]int) {

	if s == nil || len(*s) == 1 {
		return
	}

	for i := 1; i < len(*s); i = i + 1 {
		(*s)[i-1], (*s)[i] = (*s)[i], (*s)[i-1]
	}

}

func leftShiftByI(s *[]int) {

	if s == nil || len(*s) == 1 {
		return
	}

	fmt.Print("Enter the positive number for the left shift of slice")
	var input int
	fmt.Scan(&input)

	for {
		fmt.Scan(&input)

		if input < 0 {
			fmt.Println("Enter the valid number")
			continue
		}

		break

	}

	input = int(math.Mod(float64(input), float64(len(*s))))

	for ; input != 0; input = input - 1 {
		leftShiftBy1(s)
	}

}

func rightShiftBy1(s *[]int) {

	if s == nil || len(*s) == 1 {
		return
	}

	for i := len(*s) - 1; i > 0; i = i - 1 {
		(*s)[i], (*s)[i-1] = (*s)[i-1], (*s)[i]
	}

}

func rightShiftByI(s *[]int) {

	if s == nil || len(*s) == 1 {
		return
	}

	fmt.Print("Enter the positive number for the right shift of slice")
	var input int
	fmt.Scan(&input)

	for {
		fmt.Scan(&input)

		if input < 0 {
			fmt.Println("Enter the valid number")
			continue
		}

		break

	}

	input = int(math.Mod(float64(input), float64(len(*s))))

	for ; input != 0; input = input - 1 {
		rightShiftBy1(s)
	}

}

func copyOfSlice(s *[]int) []int {

	if s == nil || len(*s) == 1 {
		return nil
	}

	copySlc := make([]int, len(*s), len(*s))

	for i, elem := range *s {
		copySlc[i] = elem
	}

	return copySlc

}

func swapOddEven(s *[]int) {

	if s == nil || len(*s) == 1 {
		return
	}

	for i := 1; i < len(*s); i = i + 2 {

		(*s)[i], (*s)[i-1] = (*s)[i-1], (*s)[i]

	}

}

func deleteIthElement(s *[]int) *[]int {
	fmt.Printf("Address %p", s)

	if s == nil || len(*s) == 0 {
		return s
	}

	var input int

	slc := *s

	for {
		fmt.Println("Please enter the index to get the element")

		fmt.Scanln(&input)

		if input >= len(slc) || input < 0 {
			fmt.Println("Please enter the valid index in range", 0, "to", len(slc))
			continue
		}

		break

	}

	result := slc[input]
	fmt.Println("Number:", result)

	for i := input; i < len(slc)-1; i = i + 1 {
		slc[i] = slc[i+1]
	}

	slc = (*s)[:len(slc)-1]

	return &slc
}

func sortDirect(slc *[]int) {

	sort.Sort(sort.IntSlice(*slc))

}

func sortReverse(slc *[]int) {

	sort.Sort(sort.Reverse(sort.IntSlice(*slc)))

}

func sortString(slc *[]string) {

	sort.Strings(*slc)

}

func main() {

}
