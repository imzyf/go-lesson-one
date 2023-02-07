package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var arr = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("数组长度:", len(arr))           // 6
	fmt.Println("数组大小:", unsafe.Sizeof(arr)) // 48
	// 在 64 位平台上，int 类型的大小为 8，数组 arr 一共有 6 个元素，因此它的总大小为 6x8=48 个字节

	var arr3 = [...]int{
		22, 23, 24,
	}
	fmt.Printf("%T\n", arr3) // [3]int

	var arr4 = [...]int{
		99: 39, // 将第100个元素(下标值为99)的值赋值为39，其余元素值均为0
	}
	fmt.Printf("%T\n", arr4) // [100]int

	// 切片的声明
	var num = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(len(num)) // 6
	num = append(num, 7)
	fmt.Println(len(num)) // 7
	fmt.Println(cap(num)) // 12

	arr5 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// array[low : high : max] 数组的切片化
	sl := arr5[3:7:9]
	fmt.Println(sl)      // [4 5 6 7]
	fmt.Println(len(sl)) // 4 high-low
	fmt.Println(cap(sl)) // 6 max-low

	// 由于切片 sl 的底层数组就是数组 arr，对切片 sl 中元素的修改将直接影响数组 arr 变量
	sl[0] += 10
	fmt.Println("arr[3] =", arr5[3]) // 14

	// 关系解除
	u := [...]int{11, 12, 13, 14, 15}
	fmt.Println("array:", u) // [11, 12, 13, 14, 15]
	s := u[1:3]
	fmt.Printf("slice(len=%d, cap=%d): %v\n", len(s), cap(s), s) // [12, 13]
	//array: [11 12 13 14 15]
	//slice(len=2, cap=4): [12 13]

	s = append(s, 24)
	fmt.Println("after append 24, array:", u)
	fmt.Printf("after append 24, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)
	//after append 24, array: [11 12 13 24 15]
	//after append 24, slice(len=3, cap=4): [12 13 24]
	// 对数组进行了插入

	s = append(s, 25)
	s = append(s, 26)
	fmt.Println("after append 26, array:", u)
	fmt.Printf("after append 26, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)
	//after append 26, array: [11 12 13 24 25]
	//after append 26, slice(len=5, cap=8): [12 13 24 25 26]
	// 发生了关系解除
	// append 发现底层数组已经无法满足 append 的要求，于是新创建了一 个底层数组(数组长度为 cap(s) 的 2 倍，即 8)，并将 slice 的元素拷贝到新数组中了。

	s[0] = 22
	fmt.Println("after reassign 1st elem of slice, array:", u)
	fmt.Printf("after reassign 1st elem of slice, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)
	//after reassign 1st elem of slice, array: [11 12 13 24 25]
	//after reassign 1st elem of slice, slice(len=5, cap=8): [22 13 24 25 26]
	// 因为发生了关系解除，修改 slice 的第一个元素，array 的第一个元素没有发生变化
}
