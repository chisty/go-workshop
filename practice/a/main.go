package main

import (
	"fmt"
)

type model struct {
	value string
}

type models []model

func (m *model) reverse() {
	sz := len(m.value)
	result := make([]byte, sz)

	for i := 0; i < sz; i++ {
		result[i] = m.value[sz-1-i]
	}

	m.value = string(result)
}

func (m *model) String() string {
	return m.value
}

func (ms models) Len() int {
	return len(ms)
}

func (ms models) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

func (ms models) Less(i, j int) bool {
	return ms[i].value < ms[j].value
}

func main() {
	fmt.Println("Hello")

	item := model{
		value: "Chisty",
	}

	item.reverse()
	fmt.Println(item)

	// items := []model{
	// 	model{
	// 		value: "Tamjid",
	// 	},
	// 	model{
	// 		value: "Chisty",
	// 	},
	// 	model{
	// 		value: "Ahmed",
	// 	},
	// }

	// _ = items

	// sort.Sort(models(items))
	// fmt.Println(items)
}
