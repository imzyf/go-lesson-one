package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var s = "中国人"

	// 字节视角
	// 一个可空的【字节】序列构成
	// 一个个的字节只是孤立数据，不表意。
	fmt.Printf("the length of s = %d\n", len(s)) // 9
	for i := 0; i < len(s); i++ {
		fmt.Printf("0x%x ", s[i])
		// Unicode 字符串的 UTF-8 编码值
		// 0xe4 0xb8 0xad 0xe5 0x9b 0xbd 0xe4 0xba 0xba
	}
	fmt.Printf("\n")

	// 一个可空的【字符】序列构成
	fmt.Printf("the character count in s is %d\n", utf8.RuneCountInString(s)) // 3
	for _, c := range s {
		fmt.Printf("0x%x ", c)
		// Unicode 字符中的码点(Code Point)
		// 0x4e2d 0x56fd 0x4eba
	}
	fmt.Printf("\n")

	fmt.Printf("%d\n", 'a')

	fmt.Printf("%d\n", '中')
	fmt.Printf("%c\n", '\u4e2d')
	fmt.Printf("%c\n", '\U00004e2d')
	fmt.Printf("%c\n", '\U00004e2d')
	fmt.Printf("%s\n", "\U00004e2d")

	fmt.Printf("%d\n", '\'')
	fmt.Printf("%d\n", '\u0027')
	fmt.Printf("%d\n", '\x27')

	// 中
	// 0x4e2d
	// 0xe4 0xb8 0xad
	encodeRune(0x4e2d)
	decodeRune([]byte{0xe4, 0xb8, 0xad})

	s = "中国人"
	// string -> []rune
	fmt.Printf("%x\n", []rune(s))
	// string -> []byte
	fmt.Printf("%x\n", []byte(s))
	// []rune -> string
	fmt.Println(string([]rune(s)))
	// []byte -> string
	fmt.Println(string([]byte(s)))
	// [4e2d 56fd 4eba]
	// e4b8ade59bbde4baba
	// 中国人
	// 中国人
}

// rune -> []byte
func encodeRune(r rune) {
	fmt.Printf("the unicode charactor is %c\n", r)
	buf := make([]byte, 3)
	_ = utf8.EncodeRune(buf, r) // 对rune进行utf-8编码
	fmt.Printf("utf-8 representation is 0x%X\n", buf)
}

func decodeRune(p []byte) {
	r, _ := utf8.DecodeLastRune(p)
	fmt.Printf("the unicode charactor after decoding is %s\n", string(r))
}
