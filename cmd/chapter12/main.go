package main

func main() {
	var s int8 = 127
	s += 1
	println(s) // -128

	u := uint8(1)
	u -= 2
	println(u) // 255
}
