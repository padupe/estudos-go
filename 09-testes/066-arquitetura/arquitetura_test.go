package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	// Para que os testes rodem em Paralelo, basta adicionar a linha abaixo:
	t.Parallel()
	if runtime.GOARCH == "amd64" {
		// Se a arquitetura for amd64, o teste será PULADO
		t.Skip("Não funciona em arquitetura amd64")
	}

	t.Fail()
}
