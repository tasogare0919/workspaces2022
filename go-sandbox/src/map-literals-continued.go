// 要素の操作
// m[key] = elem 挿入
// elem = m[key]
// 削除 delete(m,key)
// キーに対する要素が存在するかどうかの確認 elem,ok = m[key]

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.4202, -122.08408},
}

func main() {
	fmt.Println(m)
}
