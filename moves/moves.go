package moves

import "puzzle/types"

// Movimentos possíveis no grid 4x4
var Movimentos = map[string]int{
	"cima":     -4,
	"baixo":    4,
	"esquerda": -1,
	"direita":  1,
}

// MovimentoValido verifica se o movimento é válido
func MovimentoValido(indiceZero int, direcao string) bool {
	switch direcao {
	case "cima":
		return indiceZero >= 4
	case "baixo":
		return indiceZero <= 11
	case "esquerda":
		return indiceZero%4 != 0
	case "direita":
		return indiceZero%4 != 3
	default:
		return false
	}
}

// GerarSucessores gera todos os estados sucessores válidos
func GerarSucessores(estado types.State) []types.State {
	var sucessores []types.State
	var indiceZero int

	// Encontra a posição do zero
	for i, v := range estado {
		if v == 0 {
			indiceZero = i
			break
		}
	}

	for direcao, desloc := range Movimentos {
		if MovimentoValido(indiceZero, direcao) {
			novoIndice := indiceZero + desloc
			novoEstado := estado
			novoEstado[indiceZero], novoEstado[novoIndice] = novoEstado[novoIndice], novoEstado[indiceZero]
			sucessores = append(sucessores, novoEstado)
		}
	}

	return sucessores
}

// ObterMovimento retorna a direção do movimento entre dois estados
func ObterMovimento(estadoAtual, sucessor types.State) string {
	var indiceZeroAtual, indiceZeroSucessor int
	
	for i, v := range estadoAtual {
		if v == 0 {
			indiceZeroAtual = i
			break
		}
	}
	
	for i, v := range sucessor {
		if v == 0 {
			indiceZeroSucessor = i
			break
		}
	}
	
	diferenca := indiceZeroSucessor - indiceZeroAtual
	
	switch diferenca {
	case -4:
		return "↑ cima"
	case 4:
		return "↓ baixo"
	case -1:
		return "← esquerda"
	case 1:
		return "→ direita"
	default:
		return "? desconhecido"
	}
}