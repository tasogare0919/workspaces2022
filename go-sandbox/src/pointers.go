// ポインタ
// *T でポインタを表現
// &オペレータは、ポイントを引きだす
// Goにおけるポインタを理解する

package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)

	*p = 21

	fmt.Println(i)

	p = &j
	*p = *p

	fmt.Println(j)
}
