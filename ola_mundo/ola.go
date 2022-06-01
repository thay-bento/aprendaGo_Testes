package main

import "fmt"

const espanhol = "espanhol"
const frances = "francês"
const italiano = "italiano"
const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "
const prefixoOlaItaliano = "Ciao, "

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}
	return prefixodeSaudacao(idioma) + nome
}

func prefixodeSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case frances:
		prefixo = prefixoOlaFrances
	case espanhol:
		prefixo = prefixoOlaEspanhol
	case italiano:
		prefixo = prefixoOlaItaliano
	default:
		prefixo = prefixoOlaPortugues
	}

	return
}

func main() {
	fmt.Println(Ola("mundo", ""))
}
