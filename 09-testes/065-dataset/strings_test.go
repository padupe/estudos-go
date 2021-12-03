package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parte: %s) - índices: esperado (%d) <> encontrado (%d)."

func TestIndex(t *testing.T) {
	testes := []struct {
		texto    string
		parte    string
		esperado int
	}{
		// Estou procurando dentro da string o valor "Cod3r" e espero que sua posição seja a 0
		{"Cod3r é show", "Cod3r", 0},
		// Estou procurando uma string vazia e espero que sua posição seja 0
		{"", "", 0},
		// Estou procurando por 'opa', com O minúsculo, e espero que o resultado seja que NÃO existe. Ou seja, -1
		{"Opa", "opa", -1},
		// Estou procurando pela letra 'o', e no nome a letra ocupe a posição 2
		{"Leonardo", "o", 2},
	}

	for _, teste := range testes {
		t.Logf("Massa: %v", teste)

		atual := strings.Index(teste.texto, teste.parte)

		if atual != teste.esperado {
			t.Errorf(msgIndex, teste.texto, teste.parte, teste.esperado, atual)
		}
	}
}

// Comando para rodar os Testes via Terminal:
// $ go test       -> Este comando roda os testes no diretório em que se encontra
// $ go test ./... -> Este comando roda todos os testes, independente do diretório
// $ go test -v    -> Este comando roda os testes no modo "verboso"
