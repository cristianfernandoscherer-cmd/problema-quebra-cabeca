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
	var escolha, tipoBusca, tipoHeuristica string
	
	fmt.Println("Escolha um estado inicial: easy, medium ou hard")
	fmt.Print("> ")
	fmt.Scanln(&escolha)

	estadoInicial, ok := states.EscolherEstadoInicial(escolha)
	if !ok {
		fmt.Println("Escolha inválida.")
		return
	}

	fmt.Println("Escolha o tipo de busca: bfs ou astar")
	fmt.Print("> ")
	fmt.Scanln(&tipoBusca)

	if tipoBusca == "astar" {
		fmt.Println("Escolha a heurística: manhattan ou foraDoLugar")
		fmt.Print("> ")
		fmt.Scanln(&tipoHeuristica)
		
		if tipoHeuristica != "manhattan" && tipoHeuristica != "foraDoLugar" {
			fmt.Println("Heurística inválida. Usando Manhattan como padrão.")
			tipoHeuristica = "manhattan"
		}
	}

	fmt.Println("Estado inicial selecionado:")
	utils.ImprimirEstado(estadoInicial)
	fmt.Println()

	var caminho []types.State

	if tipoBusca == "bfs" {
		caminho = search.BFS(estadoInicial, states.EstadoObjetivo)
	} else if tipoBusca == "astar" {
		caminho = search.AStar(estadoInicial, states.EstadoObjetivo, tipoHeuristica)
	} else {
		fmt.Println("Tipo de busca inválido.")
		return
	}

	utils.ImprimirCaminho(caminho, moves.ObterMovimento)
}