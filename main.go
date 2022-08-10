package main

import "fmt"

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

	fmt.Println(len(definePeers()["A2"]))
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
