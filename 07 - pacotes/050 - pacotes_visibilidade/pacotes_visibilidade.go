package main

import "math"

// Inicializando (varíavel, struct, interface) com letra maiúscula é PÚBLICO (visível dentro e fora do pacote)!
// Inicializando (variável, struct, interface) com letra minúscula é PRIVADO (visível apenas dentro do pacote)!

/*
	POR EXEMPLO:

	Variavel -> gerará algo público
	variavel OU _Variável -> gerará algo privado
*/

// Struct Ponto (inicia com letra maiúscula) é PÚBLICA (visível dentro e fora do Pacote)
type Ponto struct {
	// Ponto representa uma coordenada no Plano Cartesiano
	x float64
	y float64
}

// Função Catetos (inicia com letra minúsculas) é PRIVADA (visível apenas dentro do Pacote)
func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Função Distanica (inicia com letra maiúscula) é PÚBLICA (viísvel dentro e fora do Pacote)
func Distancia(a, b Ponto) float64 {
	// Distancia é respónsável por calcular a distância linear entre dois pontos
	cx, cy := catetos(a, b)
	// Retorno a Raiz Quadrada das SOMAS dos quadrados de X e Y
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
