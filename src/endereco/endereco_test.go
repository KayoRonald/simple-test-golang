package enderecos_test

import (
	enderecos "simple-test/src/endereco"
	"testing"
)

type cenarioDeTest struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()
	cenarioDeTest := []cenarioDeTest{
		{"Rua ABC", "Rua"},
		{"aua ABC", "Tipo Inválido"},
		{"Avenida ABC", "Avenida"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenarioDeTest {
		tipoRecebido := enderecos.TipoDeEndereco(cenario.enderecoInserido)
		if tipoRecebido != cenario.retornoEsperado {
			t.Error("Falha no tipo de enderço")
		}
	}
}
