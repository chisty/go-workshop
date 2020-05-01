package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Group Anagrams")

	fmt.Println("Result= ", groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println("Result= ", groupAnagrams([]string{"chisty", "hicsty", "chisty", "cchisty", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {
	var result [][]string
	store := make(map[string]int)
	resIndex := 0
	for _, val := range strs {
		key := generateKey(val)
		if index, ok := store[key]; ok {
			result[index] = append(result[index], val)
			continue
		}
		result = append(result, []string{val})
		store[key] = resIndex
		resIndex++
	}

	return result
}

func generateKey(temp string) string {
	small := strings.Split(temp, "")
	sort.Strings(small)
	return strings.Join(small, "")
}
