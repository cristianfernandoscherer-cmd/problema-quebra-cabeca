package main

import (
	"fmt"
	"puzzle/moves"
	"puzzle/search"
	"puzzle/states"
	"puzzle/utils"
	"puzzle/types"
)

func main() {
	var escolhaNum, tipoBuscaNum, tipoHeuristicaNum int

	// Mapear opções para estados iniciais
	opcoesEstados := []string{"easy", "medium", "hard", "expert", "genius", "nightmare"}

	fmt.Println("Escolha um estado inicial:")
	for i, nome := range opcoesEstados {
		fmt.Printf("%d) %s\n", i+1, nome)
	}
	fmt.Print("> ")
	fmt.Scanln(&escolhaNum)

	if escolhaNum < 1 || escolhaNum > len(opcoesEstados) {
		fmt.Println("Escolha inválida.")
		return
	}
	estadoInicial := states.EstadosIniciais[opcoesEstados[escolhaNum-1]]

	// Escolher tipo de busca
	fmt.Println("Escolha o tipo de busca:")
	fmt.Println("1) Busca em largura")
	fmt.Println("2) Busca por Heurística")
	fmt.Print("> ")
	fmt.Scanln(&tipoBuscaNum)

	var tipoBusca, tipoHeuristica string
	if tipoBuscaNum == 1 {
		tipoBusca = "bfs"
	} else if tipoBuscaNum == 2 {
		tipoBusca = "astar"

		// Escolher heurística
		fmt.Println("Escolha a heurística:")
		fmt.Println("1) Manhattan")
		fmt.Println("2) OutOfPlace")
		fmt.Print("> ")
		fmt.Scanln(&tipoHeuristicaNum)

		if tipoHeuristicaNum == 1 {
			tipoHeuristica = "manhattan"
		} else if tipoHeuristicaNum == 2 {
			tipoHeuristica = "foraDoLugar"
		} else {
			fmt.Println("Heurística inválida. Usando Manhattan como padrão.")
			tipoHeuristica = "manhattan"
		}

	} else {
		fmt.Println("Tipo de busca inválida.")
		return
	}

	fmt.Println("\nEstado inicial selecionado:")
	utils.ImprimirEstado(estadoInicial)
	fmt.Println()

	var caminho []types.State
	if tipoBusca == "bfs" {
		caminho = search.BFS(estadoInicial, states.EstadoObjetivo)
	} else if tipoBusca == "astar" {
		caminho = search.AStar(estadoInicial, states.EstadoObjetivo, tipoHeuristica)
	}

	utils.ImprimirCaminho(caminho, moves.ObterMovimento)
}
