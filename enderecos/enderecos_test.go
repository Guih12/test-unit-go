package enderecos

import "testing"

type CenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTypeAddress(t *testing.T) {
	cenerioDeTest := []CenarioDeTeste{
		{"rua ABC", "rua"},
		{"avenida", "avenida"},
		{"rodovia do imigrantes", "rodovia"},
		{"praca das rosas", "Endereco invalido"},
	}

	for _, cenario := range cenerioDeTest {
		tipoDeEnceredoRecebido := TypeAddress(cenario.enderecoInserido)
		if tipoDeEnceredoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo de recebido %s é diferente do esperado %s", tipoDeEnceredoRecebido, cenario.retornoEsperado)
		}
	}
}
