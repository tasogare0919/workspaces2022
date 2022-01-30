// Goの埋め込みは継承ではなく移譲

package main

import "fmt"

type Chip struct {
	Number int
}

type Card struct {
	string
	Chip

	Number int
}

func (c *Chip) Scan() {
	fmt.Println(c.Number)
}

func main() {
	c := Card{
		string: "Credit",
		Chip: Chip{
			Number: 4242424242424242,
		},

		Number: 545454545454545454,
	}
	// CardにはScanメソッドがないため、c.Chip.Scan()が実行される
	c.Scan()
	// ScanメソッドのレシーバはCardではなくChipであることがわかる
	// Output:4242424242424242
}