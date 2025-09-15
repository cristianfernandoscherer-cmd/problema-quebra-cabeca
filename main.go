package main

import (
	"container/heap"
	"container/list"
	"fmt"
	"math"
)

// Tipo para representar o estado
type State [16]int

// Estados iniciais fixos
var estadosIniciais = map[string]State{
	"easy": {
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 0, 15,
	},
	"medium": {
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 0, 14, 15,
	},
	"hard": {
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 0, 12,
		13, 14, 11, 15,
	},
}

// Estado objetivo
var estadoObjetivo = State{
	1, 2, 3, 4,
	5, 6, 7, 8,
	9, 10, 11, 12,
	13, 14, 15, 0,
}

// Movimentos possíveis no grid 4x4
var movimentos = map[string]int{
	"cima":     -4,
	"baixo":    4,
	"esquerda": -1,
	"direita":  1,
}

// Verifica se o movimento é válido
func movimentoValido(indiceZero int, direcao string) bool {
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

// Gera sucessores
func gerarSucessores(estado State) []State {
	var sucessores []State
	var indiceZero int

	for i, v := range estado {
		if v == 0 {
			indiceZero = i
			break
		}
	}

	for direcao, desloc := range movimentos {
		if movimentoValido(indiceZero, direcao) {
			novoIndice := indiceZero + desloc
			novoEstado := estado
			novoEstado[indiceZero], novoEstado[novoIndice] = novoEstado[novoIndice], novoEstado[indiceZero]
			sucessores = append(sucessores, novoEstado)
		}
	}

	return sucessores
}

// Verifica se dois estados são iguais
func estadosIguais(a, b State) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// -------- BFS --------
func bfs(estadoInicial, estadoObjetivo State) []State {
	type Node struct {
		estado  State
		caminho []State
	}

	visitados := make(map[State]bool)
	fila := list.New()
	fila.PushBack(Node{estadoInicial, []State{}})

	for fila.Len() > 0 {
		elemento := fila.Front()
		fila.Remove(elemento)

		no := elemento.Value.(Node)
		estadoAtual := no.estado

		if visitados[estadoAtual] {
			continue
		}
		visitados[estadoAtual] = true

		if estadosIguais(estadoAtual, estadoObjetivo) {
			return append(no.caminho, estadoAtual)
		}

		for _, sucessor := range gerarSucessores(estadoAtual) {
			novaCaminho := append([]State{}, no.caminho...)
			novaCaminho = append(novaCaminho, estadoAtual)
			fila.PushBack(Node{sucessor, novaCaminho})
		}
	}

	return nil
}

// -------- A* --------

// Distância de Manhattan como heurística
func heuristica(estado State) int {
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

type Node struct {
	estado  State
	custo   int
	f       int // custo + heurística
	caminho []State
	index   int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].f < pq[j].f
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(*Node)
	node.index = n
	*pq = append(*pq, node)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	node.index = -1
	*pq = old[0 : n-1]
	return node
}

func aStar(estadoInicial, estadoObjetivo State) []State {
	pq := &PriorityQueue{}
	heap.Init(pq)

	heap.Push(pq, &Node{estado: estadoInicial, custo: 0, f: heuristica(estadoInicial), caminho: []State{}})

	visitados := make(map[State]bool)

	for pq.Len() > 0 {
		no := heap.Pop(pq).(*Node)
		estadoAtual := no.estado

		if visitados[estadoAtual] {
			continue
		}
		visitados[estadoAtual] = true

		if estadosIguais(estadoAtual, estadoObjetivo) {
			return append(no.caminho, estadoAtual)
		}

		for _, sucessor := range gerarSucessores(estadoAtual) {
			if !visitados[sucessor] {
				novaCaminho := append([]State{}, no.caminho...)
				novaCaminho = append(novaCaminho, estadoAtual)
				custo := no.custo + 1
				f := custo + heuristica(sucessor)
				heap.Push(pq, &Node{estado: sucessor, custo: custo, f: f, caminho: novaCaminho})
			}
		}
	}

	return nil
}

// -------- Utilidades --------

// Imprime um estado
func imprimirEstado(s State) {
	for i := 0; i < 16; i += 4 {
		fmt.Printf("%2d %2d %2d %2d\n", s[i], s[i+1], s[i+2], s[i+3])
	}
	fmt.Println()
}

// Escolhe estado baseado em entrada do usuário
func escolherEstadoInicial(nome string) (State, bool) {
	estado, ok := estadosIniciais[nome]
	return estado, ok
}

func main() {
	var escolha, tipoBusca string
	fmt.Println("Escolha um estado inicial: easy, medium ou hard")
	fmt.Print("> ")
	fmt.Scanln(&escolha)

	estadoInicial, ok := escolherEstadoInicial(escolha)
	if !ok {
		fmt.Println("Escolha inválida.")
		return
	}

	fmt.Println("Escolha o tipo de busca: bfs ou astar")
	fmt.Print("> ")
	fmt.Scanln(&tipoBusca)

	fmt.Println("Estado inicial selecionado:")
	imprimirEstado(estadoInicial)

	var caminho []State
	if tipoBusca == "bfs" {
		caminho = bfs(estadoInicial, estadoObjetivo)
	} else if tipoBusca == "astar" {
		caminho = aStar(estadoInicial, estadoObjetivo)
	} else {
		fmt.Println("Tipo de busca inválido.")
		return
	}

	if caminho != nil {
		fmt.Printf("Solução encontrada em %d movimentos:\n\n", len(caminho)-1)
		for _, estado := range caminho {
			imprimirEstado(estado)
		}
	} else {
		fmt.Println("Nenhuma solução encontrada.")
	}
}
