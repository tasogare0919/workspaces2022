// range
// スライスやマップを一つずつ反復処理する
// for文の1つ目の変数はインデックス、2つ目のはインデックスの場所の要素のこぴー
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 18, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
