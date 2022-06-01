package iteracao

//const quantidadeRepeticoes = 5

func Repetir(caractere string, repetir int) string {
	var repeticoes string
	for i := 0; i < repetir; i++ {
		repeticoes += caractere

	}
	return repeticoes
}
