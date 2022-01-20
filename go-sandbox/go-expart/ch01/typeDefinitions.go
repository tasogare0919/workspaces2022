package main

import (
	"fmt"
	"time"
)

func main() {
	// 新しい型MyDurationをtime.Durationを基底として定義
	type MyDuration time.Duration
	d := MyDuration(100)

	// %Tを使うことで変数に代入されている値の型情報を出力する
	fmt.Printf("%T", d)

	// MyDuration型で基底型として定義しているtime.Durationへのキャスト
	td := time.Duration(d)

	// 型の定義がされている定数（１００）に対して明示的なキャストなしでの演算
	md := 100 * d

	fmt.Printf("td: %T, md; %T", td, md)
}
