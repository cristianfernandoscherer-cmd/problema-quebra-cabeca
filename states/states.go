package states

import "puzzle/types"

// Estado objetivo fixo
var EstadoObjetivo = types.State{
	1, 2, 3, 4,
	5, 6, 7, 8,
	9, 10, 11, 12,
	13, 14, 15, 0,
}

// Estados iniciais pré-definidos
var EstadosIniciais = map[string]types.State{
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
	"expert": {
        0, 2, 3, 4,
        1, 5, 7, 8,
        9, 6, 11, 12,
        13, 10, 14, 15,
    },
	"genius": {
        1, 3, 6, 4,
		5, 2, 7, 8,
		9, 10, 0, 11,
		13, 14, 15, 12,
    },
    "nightmare": {
     	3, 2, 6, 4,     
		1, 10, 7, 8, 
		5, 13, 11, 12,
		9, 14, 0, 15,   
	},		
}

// EscolherEstadoInicial retorna um estado baseado no nome
func EscolherEstadoInicial(nome string) (types.State, bool) {
	estado, ok := EstadosIniciais[nome]
	return estado, ok
}