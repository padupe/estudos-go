package main

import (
	"fmt"
	"runtime"
)

func main() {
	// O método asseguir apresenta quantos processadores físicos há disponíveis no seu ambiente
	fmt.Println(runtime.NumCPU())
}
