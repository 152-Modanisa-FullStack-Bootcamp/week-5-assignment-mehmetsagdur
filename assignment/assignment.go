package assignment

import (
	"math"
	"sort"
	"strings"
)

func AddUint32(x, y uint32) (uint32, bool) {
	sum := x + y
	overflow := true
	if y > 0 && x > math.MaxUint32-y ||
		y < 0 && x < math.MaxUint32-y {
		overflow = true
	} else {
		overflow = false
	}
	return sum, overflow
}

func CeilNumber(f float64) float64 {
	myCeilNumber := math.Ceil(f*4) / 4

	return myCeilNumber
}

func AlphabetSoup(s string) string {
	myStringSplitArray := strings.Split(s, "")
	sort.Strings(myStringSplitArray)
	justString := strings.Join(myStringSplitArray, "")

	return justString
}

func StringMask(s string, n uint) string {
	myStringSplitArray := strings.Split(s, "")
	i := 0
	i = int(n)
	if len(myStringSplitArray) == 0 {
		myStringSplitArray = append(myStringSplitArray, "*")
	} else {
		if n == 0 || n > uint(len(myStringSplitArray)) || n == uint(len(myStringSplitArray)) {
			i = 0
		}
		for i < len(myStringSplitArray) {
			myStringSplitArray[i] = "*"
			i++
		}
	}

	justString := strings.Join(myStringSplitArray, "")

	return justString
}

func WordSplit(arr [2]string) string {
	wordToCompare := arr[0]
	stringDictionary := arr[1]
	singleStrings := strings.Split(stringDictionary, ",")
	var s []string
	for i := 0; i < len(singleStrings); i++ {
		splitMainWordArray := strings.ReplaceAll(wordToCompare, singleStrings[i], "")
		if splitMainWordArray != wordToCompare {
			s = append(s, splitMainWordArray)
		}
	}
	lastString := ""
	if len(s) > 1 {
		lastString = s[0] + "," + s[1]
	} else {
		lastString = "not possible"
	}

	return lastString
}

func VariadicSet(i ...interface{}) []interface{} {
	var myInterface []interface{}

	for a := 0; a < len(i); a++ {
		if !contains(myInterface, i[a]) {
			myInterface = append(myInterface, i[a])
		}
	}

	return myInterface
}

func contains(s []interface{}, str interface{}) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
