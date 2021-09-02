package main

import "fmt"

func main() {
	// 初期化
	var a [5]int // a[0] ~ a[4]の扱いになる
	a[2] = 3
	a[4] = 5
	fmt.Println(a[2])
	// まとめて複数の値を宣言する場合
	b := [...]int{1, 3, 5}

	fmt.Println(b)
	// 要素数の確認 len(n)
	fmt.Println(len(b))
}
