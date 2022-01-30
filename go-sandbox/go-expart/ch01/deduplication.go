package main

import "fmt"

func main() {
	followers := []string{"John", "Richard", "John", "Jane", "Jane", "Alan"}

	// ユニークであることを考慮して大正のスライスの長さ分の要領を確保する
	// 繰り返し処理が実行される中でスライスの容量拡張を抑える
	unique := make([]string, 0, len(followers))

	// 存在有無チェック機構でマップを利用
	// 値を空の構造体にすることでメモリのアロケーションを０にできる
	m := make(map[string]struct{})
	for _, v := range followers {
		if _, ok := m[v]; ok {
			continue
		}
		unique = append(unique, v)
		m[v] = struct{}{}
	}

	fmt.Println(unique)
}