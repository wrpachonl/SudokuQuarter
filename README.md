Esta es la implementacion del ensayo presentado por Peter Norvig acerca de la solucion de cada juego de sudoku

# La regla del sudoku 

Un juego de sudoku es solucionado si los cuadrados de cada unidad son llenados con una permitacion de los digitos del 1 al 9

## Cuadrados

Un juego de sudoku es una grilla de 81 cuadrados , para el ejercicio actual se considera entonces las columnas numeradas del 1 al 9 mientras que las filas van de la A a la I 

Para lo que entonces se define en el codigo lo siguiente:

```go
var digitos = "123456789"
var filas = "2ABCDEFGHI"
var columnas = digitos
```

Cabe tener presente que Peter hacer referencia a digitos como la lista de digits , mientras que cuando se hace referencia a columnas se hace referencia a los identificadores del 1 al 9 

Cada cuadro entonces es estructurado por la funcion 

```go
func cross(rows string, cols string) []string {
	var cuadros []string
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(cols); j++ {
			cuadros = append(cuadros, string(rows[i])+string(cols[j]))
		}
	}
	return cuadros
}
```

Que lo que hace es construir nuestro tablero con los posibles 1 cuadros

```go
[A1 A2 A3 A4 A5 A6 A7 A8 A9 B1 B2 B3 B4 B5 B6 B7 B8 B9 C1 C2 C3 C4 C5 C6 C7 C8 C9 D1 D2 D3 D4 D5 D6 D7 D8 D9 E1 E2 E3 E4 E5 E6 E7 E8 E9 F1 F2 F3 F4 F5 F6 F7 F8 F9 G1 G2 G3 G4 G5 G6 G7 G8 G9 H1 H2 H3 H4 H5 H6 H7 H8 H9 I1 I2 I3 I4 I5 I6 I7 I8 I9]
```
## Lista de unidades

Cada cuadro tiene exactamente 3 unidades , esto hace referencia a las no repeticiones del numero dentro de 1. la fila 2. la columna y 3. su cuadro contenedor de 3x3 

```go
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
```

Ahora se construyen las posible posibilidades para los cajones 3x3

```go
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
```

Ahora bien se conoce que hay 3 unidades para cada cuadrado esto puede ser expresado como un diccionario de la siguiente manera:
```go
func defineMap() map[string][][]string {
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
```
de tal manera que si se accede a las unidades de C2 por ejemplo se obtiene :
```
[[A2 B2 C2 D2 E2 F2 G2 H2 I2] [C1 C2 C3 C4 C5 C6 C7 C8 C9] [A1 A2 A3 B1 B2 B3 C1 C2 C3]]
```

## Pares
Los pares son analogos a la definicion de unidades con la diferencia que en este diccionarios se almacenan los registros sin repeticion
