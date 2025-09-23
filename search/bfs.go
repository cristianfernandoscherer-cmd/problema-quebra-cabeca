package search

import (
	"container/list"
	"fmt"
	"puzzle/moves"
	"puzzle/types"
	"puzzle/utils"
	"strings"
)

// BFS implementa a busca em largura
func BFS(estadoInicial, estadoObjetivo types.State) []types.State {
	type NodeBFS struct {
		Estado  types.State
		Caminho []types.State
	}

	visitados := make(map[types.State]bool)
	fila := list.New()
	fila.PushBack(NodeBFS{estadoInicial, []types.State{}})

	fmt.Println("🌳 INICIANDO BUSCA BFS")
	fmt.Println("🎯 Estado objetivo:")
	utils.ImprimirEstado(estadoObjetivo)
	fmt.Println("🚀 Estado inicial:")
	utils.ImprimirEstado(estadoInicial)
	fmt.Println(strings.Repeat("=", 50))

	iteracao := 1
	estadosGerados := 0

	imprimirFila := func(fila *list.List, iteracao int) {
		fmt.Printf("\n📋 FILA (Iteração %d) - %d elementos:\n", iteracao, fila.Len())
		fmt.Println(strings.Repeat("-", 50))
		elemento := fila.Front()
		posicao := 0
		for elemento != nil {
			no := elemento.Value.(NodeBFS)
			fmt.Printf("[%d] Profundidade: %d\n", posicao, len(no.Caminho))
			utils.ImprimirEstado(no.Estado)
			fmt.Println()
			elemento = elemento.Next()
			posicao++
		}
		fmt.Println(strings.Repeat("-", 50))
	}

	for fila.Len() > 0 {
		imprimirFila(fila, iteracao)

		elemento := fila.Front()
		fila.Remove(elemento)
		no := elemento.Value.(NodeBFS)
		estadoAtual := no.Estado

		if visitados[estadoAtual] {
			fmt.Printf("\n⏩ ITERAÇÃO %d: Pulando estado já visitado\n", iteracao)
			iteracao++
			continue
		}

		visitados[estadoAtual] = true
		
		fmt.Printf("\n%d️⃣ ITERAÇÃO %d:\n", iteracao, iteracao)
		fmt.Printf("📦 Estado atual (profundidade: %d):\n", len(no.Caminho))
		utils.ImprimirEstado(estadoAtual)

		if EstadosIguais(estadoAtual, estadoObjetivo) {
			fmt.Println("🎯 OBJETIVO ENCONTRADO!")
			fmt.Printf("📈 Estatísticas finais:\n")
			fmt.Printf("   - Iterações: %d\n", iteracao)
			fmt.Printf("   - Estados visitados: %d\n", len(visitados))
			fmt.Printf("   - Estados gerados: %d\n", estadosGerados)
			fmt.Printf("   - Tamanho do caminho: %d movimentos\n", len(no.Caminho))
			return append(no.Caminho, estadoAtual)
		}

		sucessores := moves.GerarSucessores(estadoAtual)
		estadosGerados += len(sucessores)
		
		fmt.Printf("🌱 Gerou %d sucessores:\n", len(sucessores))
		for i, sucessor := range sucessores {
			movimento := moves.ObterMovimento(estadoAtual, sucessor)
			fmt.Printf("   %d. %s:\n", i+1, movimento)
			utils.ImprimirEstado(sucessor)
			
			if visitados[sucessor] {
				fmt.Printf("      (já visitado)\n")
			} else {
				fmt.Printf("      (novo - será adicionado à fila)\n")
			}
			fmt.Println()
		}

		novosEstados := 0
		for _, sucessor := range sucessores {
			if !visitados[sucessor] {
				novaCaminho := append([]types.State{}, no.Caminho...)
				novaCaminho = append(novaCaminho, estadoAtual)
				fila.PushBack(NodeBFS{sucessor, novaCaminho})
				novosEstados++
			}
		}

		fmt.Printf("📊 Status da busca:\n")
		fmt.Printf("   - Fila: %d elementos\n", fila.Len())
		fmt.Printf("   - Novos estados adicionados: %d\n", novosEstados)
		fmt.Printf("   - Total visitados: %d estados\n", len(visitados))
		fmt.Printf("   - Total gerados: %d estados\n", estadosGerados)
		fmt.Println(strings.Repeat("-", 50))
		
		iteracao++
	}

	fmt.Println("❌ Nenhuma solução encontrada")
	return nil
}