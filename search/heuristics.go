package search

import (
	"math"
	"puzzle/states"
	"puzzle/types"
)

// HeuristicaManhattan calcula a distância de Manhattan
func HeuristicaManhattan(estado types.State) int {
	dist := 0
	for i, v := range estado {
		if v != 0 {
			linhaAtual, colAtual := i/4, i%4
			linhaObj, colObj := (v-1)/4, (v-1)%4
			dist += int(math.Abs(float64(linhaAtual-linhaObj)) + math.Abs(float64(colAtual-colObj)))
		}
	}
	return dist
}

// HeuristicaForaDoLugar conta peças fora do lugar
func HeuristicaForaDoLugar(estado types.State) int {
	count := 0
	for i := 0; i < 16; i++ {
		if estado[i] != 0 && estado[i] != states.EstadoObjetivo[i] {
			count++
		}
	}
	return count
}

// EstadosIguais verifica se dois estados são iguais
func EstadosIguais(a, b types.State) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}