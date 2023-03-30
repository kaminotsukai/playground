package main

import "fmt"

type Device struct {
	ID        int
	Published bool
}

// 値のスライスをイテレーターで回して、アドレスを出力すると全て同じアドレスになる
var ds = []Device{
	{ID: 1, Published: false},
	{ID: 2, Published: true},
	{ID: 3, Published: false},
	{ID: 4, Published: true},
}

func main() {
	// 1. Goの配列は参照ではなく、値
	array1 := [3]string{"aaa", "bbb", "ccc"}
	array2 := array1
	array2[0] = "ddd"
	fmt.Println(array1) // [aaa bbb ccc]
	fmt.Println(array2) // [ddd bbb ccc]

	// 2. スライスは、配列へのポインター、セグメントの長さ、容量を保持している
	var dd []*Device
	for _, d := range ds {

		// for range ループの value は常に同じアドレス値を取るため、再代入することで回避する
		// see: https://github.com/golang/go/wiki/CommonMistakes#using-reference-to-loop-iterator-variable
		d := d
		dd = append(dd, &d)
	}
	fmt.Println(dd)
}
