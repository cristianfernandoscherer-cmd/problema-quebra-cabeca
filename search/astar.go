package search

import (
	"container/heap"
	"fmt"
	"puzzle/moves"
	"puzzle/types"
	"puzzle/utils"
	"strings"
)

// AStar implementa a busca A*
func AStar(estadoInicial, estadoObjetivo types.State, tipoHeuristica string) []types.State {
	pq := &types.PriorityQueue{}
	heap.Init(pq)

	iteracao := 0
	estadosGerados := 0

	calcularHeuristica := func(estado types.State) int {
		switch tipoHeuristica {
		case "foraDoLugar":
			return HeuristicaForaDoLugar(estado)
		default:
			return HeuristicaManhattan(estado)
		}
	}

	imprimirHeap := func(pq *types.PriorityQueue, iteracao int) {
		fmt.Printf("\n📦 HEAP (Iteração %d) - %d elementos:\n", iteracao, pq.Len())
		fmt.Println(strings.Repeat("-", 50))
		for i, node := range *pq {
			fmt.Printf("[%d] f=%d (g=%d + h=%d)\n", i, node.F, node.Custo, node.F-node.Custo)
			utils.ImprimirEstado(node.Estado)
			fmt.Println()
		}
		fmt.Println(strings.Repeat("-", 50))
	}

	hInicial := calcularHeuristica(estadoInicial)
	fmt.Println("🎯 Estado objetivo:")
	utils.ImprimirEstado(estadoObjetivo)
	fmt.Println("🚀 Estado inicial adicionado ao heap:")
	fmt.Printf("f=%d (g=0 + h=%d)\n", hInicial, hInicial)
	fmt.Printf("📊 Heurística utilizada: %s\n", tipoHeuristica)
	utils.ImprimirEstado(estadoInicial)
	fmt.Println(strings.Repeat("=", 60))

	heap.Push(pq, &types.Node{
		Estado:  estadoInicial,
		Custo:   0,
		F:       hInicial,
		Caminho: []types.State{},
	})

	visitados := make(map[types.State]bool)

	for pq.Len() > 0 {
		iteracao++
		
		fmt.Printf("\n\n🔷🔷🔷 ITERAÇÃO %d 🔷🔷🔷\n", iteracao)
		imprimirHeap(pq, iteracao)

		no := heap.Pop(pq).(*types.Node)
		estadoAtual := no.Estado

		fmt.Printf("🎯 POP: Elemento com menor f=%d selecionado:\n", no.F)
		fmt.Printf("   g=%d (custo acumulado)\n", no.Custo)
		fmt.Printf("   h=%d (heurística %s)\n", no.F-no.Custo, tipoHeuristica)
		utils.ImprimirEstado(estadoAtual)

		if visitados[estadoAtual] {
			fmt.Printf("⏩ ESTADO JÁ VISITADO! Pulando...\n")
			continue
		}

		visitados[estadoAtual] = true
		
		if EstadosIguais(estadoAtual, estadoObjetivo) {
			fmt.Printf("🎉🎉🎉 OBJETIVO ENCONTRADO! 🎉🎉🎉\n")
			fmt.Printf("📊 Estatísticas finais:\n")
			fmt.Printf("   - Iterações: %d\n", iteracao)
			fmt.Printf("   - Estados visitados: %d\n", len(visitados))
			fmt.Printf("   - Estados gerados: %d\n", estadosGerados)
			fmt.Printf("   - Tamanho do caminho: %d movimentos\n", len(no.Caminho)+1)
			fmt.Printf("   - Heurística utilizada: %s\n", tipoHeuristica)
			return append(no.Caminho, estadoAtual)
		}

		sucessores := moves.GerarSucessores(estadoAtual)
		estadosGerados += len(sucessores)
		
		fmt.Printf("\n🌱 GERANDO SUCESSORES (%d):\n", len(sucessores))
		for i, sucessor := range sucessores {
			if !visitados[sucessor] {
				h := calcularHeuristica(sucessor)
				g := no.Custo + 1
				f := g + h
				movimento := moves.ObterMovimento(estadoAtual, sucessor)
				
				fmt.Printf("   %d. %s → f=%d (g=%d + h=%d)\n", i+1, movimento, f, g, h)
				utils.ImprimirEstado(sucessor)
				
				novaCaminho := append([]types.State{}, no.Caminho...)
				novaCaminho = append(novaCaminho, estadoAtual)
				heap.Push(pq, &types.Node{
					Estado:  sucessor,
					Custo:   g,
					F:       f,
					Caminho: novaCaminho,
				})
			} else {
				fmt.Printf("   %d. ⏩ Já visitado\n", i+1)
			}
		}

		fmt.Printf("\n📊 STATUS APÓS EXPANSÃO:\n")
		fmt.Printf("   - Heap agora tem: %d elementos\n", pq.Len())
		fmt.Printf("   - Total visitados: %d estados\n", len(visitados))
		fmt.Printf("   - Total gerados: %d estados\n", estadosGerados)
		
		fmt.Println(strings.Repeat("=", 60))
	}

	return nil
}