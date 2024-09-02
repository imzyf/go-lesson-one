package main

import "fmt"

func main() {
	m1 := map[int][]string{
		1: []string{"var1-1", "var1-2"},
		2: []string{"var2-1", "var2-2"},
		3: {"3-1", "3-2"},
	}

	fmt.Printf("%T\n", m1)
	fmt.Println(m1)

	// 操作一:插入新键值对
	m2 := make(map[int]string)
	m2[1] = "value1"
	m2[2] = "value2"
	fmt.Println(m2)
	m2[2] = "value3"
	fmt.Println(m2)

	// 操作二:获取键值对数量
	fmt.Println(len(m2))
	m2[3] = "value3"
	fmt.Println(len(m2))

	test1()

}

// 稳定 map key 输出
func test1() {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	for k, v := range m {
		fmt.Printf("键: %s, 值: %d\n", k, v)
	}
}
