package main

import "fmt"

type line struct {
	value []int
}

func main() {

	var item line
	for i := 1; i < 10; i += 2 {
		item.value = append(item.value, i)
	}

	newItem := item.GetNewSlice()
	fmt.Println(item)
	fmt.Println(newItem)
	item.Update()
	fmt.Println(item)
	fmt.Println(newItem)

	// item := line{
	// 	value: "Chisty",
	// }
	// newItem := item.reverse()
	// fmt.Println(item)
	// fmt.Println(newItem)
	// item.reverseSelf()
	// fmt.Println(item)
}

func (data line) GetNewSlice() line {
	newList := make([]int, len(data.value))
	for i, value := range data.value {
		newList[i] = value
	}
	return line{value: newList}
}

func (data line) Update() {
	for i := 0; i < len(data.value); i++ {
		data.value[i] += 10
	}
}

// func (data line) reverse() line {
// 	size := len(data.value)
// 	text := make([]byte, size)

// 	for i := 0; i < size; i++ {
// 		text[i] = data.value[size-1-i]
// 	}

// 	return line{
// 		value: string(text),
// 	}
// }

// func (data *line) reverseSelf() {
// 	size := len(data.value)
// 	text := make([]byte, size)

// 	for i := 0; i < size; i++ {
// 		text[i] = data.value[size-1-i]
// 	}
// 	data.value = string(text)
// }
