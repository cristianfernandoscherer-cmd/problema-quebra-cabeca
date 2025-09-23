# Problema do Quebra-Cabeça (8-Puzzle)

Este projeto implementa um resolvedor do clássico **Puzzle** em Go, utilizando algoritmos de **Busca em Largura (BFS)**, **Heuristica de distância de Manhattan** e **Heurísticas de peças fora do lugar**.

---

## 🧩 Sobre o problema

O problema consiste em um tabuleiro **4x4** com 15 peças numeradas e um espaço vazio.  
O objetivo é reorganizar as peças a partir de um estado inicial até o **estado objetivo**.

---

## ⚙️ Funcionalidades

- Escolha entre diferentes **estados iniciais** (`easy`, `medium`, `hard`)
- Suporte a **BFS (Breadth-First Search)**  
- Suporte a **Heurísticas**:
  - `manhattan`
  - `foraDoLugar` (número de peças fora da posição correta)
- Impressão do **passo a passo da solução**

---

## 🚀 Como executar

### Pré-requisitos
- [Go](https://golang.org/) 1.20+

### Executar a aplicação
```bash
make build
make run