package main

import "testing"

// func TestOla(t *testing.T) {
// 	resultado := Ola("Chris")
// 	esperado := "Olá, Chris"

// 	if resultado != esperado {
// 		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
// 	}
// }
func TestOla(t *testing.T) {
	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Chris", "")
		esperado := "Olá, Chris"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz olá, mundo quando uma string vazia é passada como parâmetro", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá, Mundo"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em espanhol", func(t *testing.T) {
		resultado := Ola("Elodie", "espanhol")
		esperado := "Hola, Elodie"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em francês", func(t *testing.T) {
		resultado := Ola("Jean", "francês")
		esperado := "Bonjour, Jean"
		verificaMensagemCorreta(t, resultado, esperado)
	})
	t.Run("em italiano", func(t *testing.T) {
		resultado := Ola("Paolo", "italiano")
		esperado := "Ciao, Paolo"
		verificaMensagemCorreta(t, resultado, esperado)
	})
}
