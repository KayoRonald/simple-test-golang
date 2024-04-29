package enderecos

import (
	"testing"
)

func TestTipoDeEndereco(t *testing.T) {
	enderecoParaTest := "Avenida Paulista"
	tipoEsperado := "Avenida"
	tipoRecebido := TipoDeEndereco(enderecoParaTest)
	if tipoRecebido != tipoEsperado{
		t.Error("Falha no endereço")
	}
}
