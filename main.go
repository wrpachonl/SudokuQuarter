package main

import (
	"fmt"
	"strings"
)

var digitos = "123456789"
var filas = "ABCDEFGHI"
var columnas = digitos
var cuadros = cross(filas, columnas)
var unidadColumnas = crossLines(columnas, filas, false)
var unidadFilas = crossLines(filas, columnas, true)
var squares = squaresConstruction()
var listaUnidades = joinArrays()
var unidades = defineUnits()
var pares = definePeers()

func main() {
	// var matriz string
	// fmt.Println("Ingrese el sudoku a resolver")
	// fmt.Scanf("%s", &matriz)
	matriz := "4.....8.5.3..........7......2.....6.....8.4......1.......6.3.7.5..2.....1.4......"
	matrizArray := []string{}
	for i := 0; i < len(matriz); i++ {
		matrizArray = append(matrizArray, string(matriz[i]))
	}
	result, _ := parse_grid(matrizArray)
	display(result)
}
func cross(rows string, cols string) []string {
	var cuadros []string
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(cols); j++ {
			cuadros = append(cuadros, string(rows[i])+string(cols[j]))
		}
	}
	return cuadros
}
func crossLines(row string, col string, order bool) [][]string {
	var cuadroTotal [][]string
	var cuadroLongitudinal []string
	for i := 0; i < len(row); i++ {
		for j := 0; j < len(col); j++ {
			if order {
				cuadroLongitudinal = append(cuadroLongitudinal, string(row[i])+string(col[j]))
			} else {
				cuadroLongitudinal = append(cuadroLongitudinal, string(col[j])+string(row[i]))
			}
		}
		cuadroTotal = append(cuadroTotal, cuadroLongitudinal)
		cuadroLongitudinal = nil
	}
	return cuadroTotal
}
func squaresConstruction() [][]string {
	var cuadroTotal [][]string
	var cuadroLongitudinal []string
	rowsPossible := [...]string{"ABC", "DEF", "GHI"}
	colsPossible := [...]string{"123", "456", "789"}
	for i := 0; i < len(rowsPossible); i++ {
		for j := 0; j < len(colsPossible); j++ {
			cuadroLongitudinal = append(cuadroLongitudinal, cross(rowsPossible[i], colsPossible[j])...)
			cuadroTotal = append(cuadroTotal, cuadroLongitudinal)
			cuadroLongitudinal = nil
		}

	}
	return cuadroTotal
}
func joinArrays() [][]string {
	result := append(unidadColumnas, unidadFilas...)
	result = append(result, squares...)
	return result
}

func contains(element []string, value string) bool {
	for _, s := range element {
		if value == s {
			return true
		}
	}
	return false
}
func containsString(baseElement string, searchElement string) bool {
	if position := strings.Index(baseElement, searchElement); position != -1 {
		return true
	}
	return false
}
func defineUnits() map[string][][]string {
	unidades := make(map[string][][]string)
	for s := 0; s < len(cuadros); s++ {
		for u := 0; u < len(listaUnidades); u++ {
			if contains(listaUnidades[u], cuadros[s]) {
				unidades[cuadros[s]] = append(unidades[cuadros[s]], listaUnidades[u])
			}
		}
	}
	return unidades
}
func unique(array []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range array {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
func peerUnit(key string) []string {
	result := []string{}
	for u := 0; u < len(unidades[key]); u++ {
		for s := 0; s < len(unidades[key][u]); s++ {
			if unidades[key][u][s] != key {
				result = append(result, unidades[key][u][s])
			}
		}
	}
	return unique(result)
}

func definePeers() map[string][]string {
	pares := make(map[string][]string)
	for s := 0; s < len(cuadros); s++ {
		pares[cuadros[s]] = append(pares[cuadros[s]], peerUnit(cuadros[s])...)
	}
	return pares
}
func grid_values(matriz []string) map[string]string {
	result := make(map[string]string)
	for i := 0; i < len(matriz); i++ {
		result[cuadros[i]] = matriz[i]
	}
	return result
}

func parse_grid(grid []string) (map[string]string, bool) {
	values := make(map[string]string)
	grid_v := grid_values(grid)
	for i := 0; i < len(grid); i++ {
		values[cuadros[i]] = digitos
	}
	for square, digits := range grid_v {
		_, error := assign(values, square, digits)
		if containsString(digitos, digits) && !error {
			return nil, false
		}
	}
	return values, true
}

func assign(values map[string]string, square string, deleteItem string) (map[string]string, bool) {
	other_values := strings.Replace(values[square], deleteItem, "", -1)
	for _, deleteItem2 := range other_values {
		valueResult, resultError := eliminate(values, square, string(deleteItem2))
		if !resultError {
			return nil, false
		}
		return valueResult, true
	}
	return values, true
}

func eliminate(values map[string]string, square string, deleteItem string) (map[string]string, bool) {
	if !containsString(values[square], deleteItem) {
		return values, true
	}
	values[square] = strings.Replace(values[square], deleteItem, "", -1)
	if len(values[square]) == 0 {
		return nil, false
	} else if len(values[square]) == 1 {
		deleteItem2 := values[square]
		for i := 0; i < len(pares[square]); i++ {
			_, resultEliminate := eliminate(values, pares[square][i], deleteItem2)
			if !resultEliminate {
				return nil, false
			}
		}
	}
	for i := 0; i < len(unidades[square]); i++ {
		dplaces := []string{}
		for s := 0; s < len(unidades[square][i]); i++ {
			if containsString(values[square], deleteItem) {
				dplaces = append(dplaces, square)
			}
		}
		if len(dplaces) == 0 {
			return nil, false
		} else if len(dplaces) == 1 {
			_, resultAssign := assign(values, dplaces[0], deleteItem)
			if !resultAssign {
				return nil, false
			}
		}
	}
	return values, true
}

func display(values map[string]string) {
	fmt.Println("+-------+-------+-------+")
	for row := 0; row < 9; row++ {
		fmt.Print("| ")
		for col := 0; col < 9; col++ {
			if col == 3 || col == 6 {
				fmt.Print("| ")
			}
			fmt.Printf("%s ", values[string(filas[row])+string(columnas[col])])
			if col == 8 {
				fmt.Print("|")
			}
		}
		if row == 2 || row == 5 || row == 8 {
			fmt.Println("\n+-------+-------+-------+")
		} else {
			fmt.Println()
		}
	}
}
