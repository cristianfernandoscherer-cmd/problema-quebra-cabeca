package utils

import (
	"fmt"
	"puzzle/types"
)

// ImprimirEstado exibe o estado formatado
func ImprimirEstado(s types.State) {
	for i := 0; i < 16; i += 4 {
		fmt.Printf("%2d %2d %2d %2d\n", s[i], s[i+1], s[i+2], s[i+3])
	}
}

// ImprimirCaminho exibe a solução encontrada
func ImprimirCaminho(caminho []types.State, obterMovimento func(types.State, types.State) string) {
	if caminho != nil {
		fmt.Printf("\n✅ SOLUÇÃO ENCONTRADA em %d movimentos:\n\n", len(caminho)-1)
		
		for i, estado := range caminho {
			if i > 0 {
				movimento := obterMovimento(caminho[i-1], estado)
				fmt.Printf("🔹 Movimento %d: %s\n", i, movimento)
			} else {
				fmt.Printf("🔹 Estado inicial\n")
			}
			ImprimirEstado(estado)
			fmt.Println()
		}
	} else {
		fmt.Printf("❌ Nenhuma solução encontrada")
	}
}