package main

import (
	"fmt"
	"io"
)

func main() {
	Contagem()
}
func Contagem(saida io.Writer) {
	fmt.Fprint(saida, "3")
}
