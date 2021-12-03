// Por padrão, os arquivos de testes devem terminar em "_test.go"

package matematica

import "testing"

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v."

// Por convenção, é necessário informar um ponteiro (t *testing.T) como parâmetro
// Ainda por convenção, os testes devem iniciar com "Test"
func TestMedia(t *testing.T) {
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}

}

// Comando para rodar os Testes via Terminal:
// $ go test       -> Este comando roda os testes no diretório em que se encontra
// $ go test ./... -> Este comando roda todos os testes, independente do diretório
