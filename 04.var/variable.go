package main

import "fmt"

func main() {
	var iVal int = 10
	var sMsg string = "school"

	iVal = 20
	sMsg = "msg plus"

	fmt.Println(iVal, "-----", sMsg)

	test()
	ex4_4()

	var a int16 = 3456
	var b int8 = int8(a)
	fmt.Println(a, b)

}

func v_test() {
	// 모든 정수타입의 기본값은 0
	// 모든 실수타입의 기본값은 0.0
	// 불리언의 기본값은 false
	// 문자열의 기본값은 "" 빈문자열
	// 이외 기본값 ni (정의되지 않은 메모리 주소를 나타내는 go키워드

	// go에서 연산시 타입이 같아야 된다. int + float 안된다.
}

func ex4_4() {
	a := 3
	var b float64 = 3.5
	var c int = int(b)

	d := a * int(b)

	var e int64 = 7
	f := a * int(e)

	fmt.Println(a, b, c, d, e, f)

	//Output:

}

func test() {
	var a int = 3
	var b int
	var c = 4
	var e = "hello"

	f := 5.14

	d := 5

	fmt.Println(a, b, c, d, e, f)
}

func test2() {
	var a string = "1111"

	fmt.Println(a)
}

func basicVar() {

	// // usingned 부호없는 정수 즉 다 양수, 8 bit == 1byte
	// var vui8 uint8 = 0
	// var vui16 uint16 // 2byte
	// var vui32 uint32
	// var vui64 uint64

	// // 부호있는 정수
	// var vi8 int8
	// var vi4 int
	// var vi16 int16
	// var vi32 int32
	// var vi64 int64

	// // 소숫점이 있는 실수 3.1415
	// var vf32 float32
	// var vf64 float64

	// var vByte byte // uint8과 같다.
	// var vRune rune // rune은 문자하나를 말한다.
	// var vInt int   // 32bit 컴퓨터의 경우는 int32와 같고, 64bit컴퓨터의 경우 int64와 같다.
	// var vUint uint

}

func etc() {
	// var vBool bool
	// var vString string
	// var vArray[3] vArray
	// var vSlice vSlice
	// var vStruct
}
