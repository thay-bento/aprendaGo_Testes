package main

func Soma(numeros []int) int {
	soma := 0
	// for i := 0; i < 5; i++ {
	// 	soma += numeros[i]
	// }
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}
func main() {

}
