package iteracao

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepetir(t *testing.T) {
	repeticoes := Repetir("a", 5)
	esperado := "aaaaa"

	if repeticoes != esperado {
		t.Errorf("esperado '%s' mas obteve '%s'", esperado, repeticoes)
	}
}
func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a", 5)
	}
}

func ExampleRepetir() {
	repetir := Repetir("a", 3)
	fmt.Println(repetir + strings.Repeat("a", 2))
	//Output: aaaaa
}
