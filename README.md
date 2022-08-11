Esta es la implementacion del ensayo presentado por Peter Norvig acerca de la solucion de cada juego de sudoku

# La regla del sudoku 

Un juego de sudoku es solucionado si los cuadrados de cada unidad son llenados con una permitacion de los digitos del 1 al 9

# Representacion del sudoku
La siguiente cadena es un ejemplo de como se acepta en el programa el sudoku entrante:

```
"4.....8.5.3..........7......2.....6.....8.4......1... ....6.3.7.5..2.....1.4......"
```
Los primeros 9 caracteres forman la primera fila, los segundos 9 forman la segunda y asi sucesivamente. Todos los puntos se interpretan como cuadrados vacios 

## Extraccion de los valores

Para extraer los datos de la entrada se define la siguiente función , en donde cada dato es convertido es añadido a un map, donde la llave corresponde a una representacion de su fila y columna como se explica mas adelante 
```go
func gridValues(grid string) (map[string]string, error) {
	values := make(map[string]string, len(grid))
	chars := make([]string, len(grid))

	// The number of clues given in the grid
	nbClues := 0
	var diffDigits []string

	// For each square
	for i := 0; i < len(grid); i++ {
		// Value of the square
		str := grid[i : i+1]

		// Valid that the square value is a digit from 1 to 9 ('0' or '.' for empties)
		// and add it to the sudoku list of values.
		if strings.Contains(digitos, str) || strings.Contains("0.", str) {
			chars[i] = str
		}
		if strings.Contains(digitos, str) {
			nbClues++
			if !contains(diffDigits, str) {
				diffDigits = append(diffDigits, str)
			}
		}
	}
  ```
# Definiciones iniciales

## Cuadrados

Un juego de sudoku es una grilla de 81 cuadrados , para el ejercicio actual se considera entonces las columnas numeradas del 1 al 9 mientras que las filas van de la A a la I 

Para lo que entonces se define en el codigo lo siguiente:

```go
const digitos string = "123456789"
const filas string = "ABCDEFGHI"
const columnas string = digitos
```

Cabe tener presente que Peter hacer referencia a digitos como la lista de digits , mientras que cuando se hace referencia a columnas se hace referencia a los identificadores del 1 al 9 

Cada cuadro entonces es estructurado por la funcion 

```go
func cross(A, B string) []string {
	res := make([]string, len(A)*len(B))

	i := 0
	for _, a := range A {
		for _, b := range B {
			res[i] = string(a) + string(b)
			i++
		}
	}

	return res
}
```

Que lo que hace es construir nuestro tablero con los posibles 1 cuadros

```go
[A1 A2 A3 A4 A5 A6 A7 A8 A9 B1 B2 B3 B4 B5 B6 B7 B8 B9 C1 C2 C3 C4 C5 C6 C7 C8 C9 D1 D2 D3 D4 D5 D6 D7 D8 D9 E1 E2 E3 E4 E5 E6 E7 E8 E9 F1 F2 F3 F4 F5 F6 F7 F8 F9 G1 G2 G3 G4 G5 G6 G7 G8 G9 H1 H2 H3 H4 H5 H6 H7 H8 H9 I1 I2 I3 I4 I5 I6 I7 I8 I9]
```
## Lista de unidades

Cada cuadro tiene exactamente 3 unidades , esto hace referencia a las no repeticiones del numero dentro de 1. la fila 2. la columna y 3. su cuadro contenedor de 3x3 

Ahora se construyen las posible posibilidades para los cajones 3x3

```go
// CreateUnitList list the 20 peers of each sudoku squares
func createUnitList(rows, cols string) [][]string {
	res := make([][]string, len(rows)*3)

	i := 0
	for _, c := range cols {
		// A1 B1 C1 D1 E1 F1 G1 H1 I1...
		res[i] = cross(rows, string(c))
		i++
	}

	for _, r := range rows {
		res[i] = cross(string(r), cols)
		i++
	}

	rs := []string{`ABC`, `DEF`, `GHI`}
	cs := []string{`123`, `456`, `789`}
	for _, r := range rs {
		for _, c := range cs {
			res[i] = cross((r), string(c))
			i++
		}
	}

	return res
}

```

Ahora bien se conoce que hay 3 unidades para cada cuadrado esto puede ser expresado como un diccionario de la siguiente manera:
```go
func createUnits(squares []string, unitList [][]string) map[string][][]string {
	units := make(map[string][][]string, len(squares))

	for _, s := range squares {
		unit := make([][]string, 3)
		i := 0
		for _, u := range unitList {
			// For each squares of the unit
			for _, su := range u {
				if s == su {
					unit[i] = u
					i++
					break
				}
			}
		}
		units[s] = unit
	}

	return units
}
```
de tal manera que si se accede a las unidades de C2 por ejemplo se obtiene :
```
[[A2 B2 C2 D2 E2 F2 G2 H2 I2] [C1 C2 C3 C4 C5 C6 C7 C8 C9] [A1 A2 A3 B1 B2 B3 C1 C2 C3]]
```

## Pares
Los pares son analogos a la definicion de unidades con la diferencia que en este diccionarios se almacenan los registros sin repeticion
```go
func createPeers(units map[string][][]string) map[string]map[string]bool {
	peers := make(map[string]map[string]bool, len(units))

	for s, ul := range units {
		peer := make(map[string]bool, 20)
		for _, u := range ul {
			for _, su := range u {
				if s != su {
					peer[su] = true
				}
			}
		}
		peers[s] = peer
	}

	return peers
}
```

# Propagacion de restricciones
 Peter entonces enuncia dos reglas simples 

 1. Si un cuadrado tiene solo un valor posible, elimine ese valor de los pares del cuadrado.
 2. Si una unidad tiene solo un lugar posible para un valor, entonces coloque el valor allí

Para hacer entonces se define la siguiente funciona teniendo en cuenta lo siguiente , para todos los cuadros si tuviesen un valor nulo, entonces las posibilidades serian todos los digitos del 1 al 9 no obstante a medida que se encuentra un valor unico , lo que se hace entonces es propagar la restriccion asociada a sus pares:
```go
func parseGrid(grid string) (values map[string]string, err error) {
	values = make(map[string]string, len(cuadrados))
	for _, s := range cuadrados {
		values[s] = digitos
	}

	gr, err := gridValues(grid)
	for s, v := range gr {
		if strings.Contains(digitos, v) {
			values = assign(values, s, v)
			if values == nil {
				return nil, nil
			}
		}
	}
	return values, err
}
```

En la función anterior,

* los valores reciben dígitos como valor inicial en un diccionario donde cada cuadrado se usa como una clave.
* La cuadrícula dada contiene la representación inicial del rompecabezas Sudoku.
* Usamos la función grid_values ​​para extraer solo valores relevantes (dígitos, '0' y '.').
* Para cada cuadrado con valor de dígito inicial, llamamos a la función de asignación para asignar el valor al cuadrado mientras lo eliminamos de los pares.
* Si algo sale mal en la función de asignación, devuelve False para indicar la falla
* Si no hay errores, la función devuelve el diccionario de valores a la persona que llama

## Funcion de asignación

La función asignar actualiza los valores entrantes eliminando los otros valores que no sean v del cuadrado s llamando a la funcion eliminar.

Si alguna eliminacion devuelve falso, significa que asignacion devuelve falso , lo que indica que hay una contradiccion 

```go
func assign(values map[string]string, s string, v string) map[string]string {
	otherValues := strings.Replace(values[s], v, "", -1)
	for _, v := range otherValues {
		if _, ok := eliminate(values, s, string(v)); !ok {
			return nil
		}
	}
	return values
}
```


## Funcion eliminar
Que hace la funcion eliminar:
* Elimina el valor dado v de valores [s] que es una lista de valores potenciales para s.
* Si no quedan valores en valores[s] (es decir, no tenemos ningún valor potencial para ese cuadrado), devuelve Falso
* Cuando solo hay un valor potencial para s, elimina el valor de todos los pares de s (el valor es para s y los pares no pueden tenerlo para satisfacer la regla de Sudoku) <== estrategia (1)
* Asegurarse de que el valor dado v tenga un lugar en otro lugar (es decir, si ningún cuadrado tiene v como valor potencial, no podemos resolver el rompecabezas). Si esta prueba falla, devuelve Falso
* Donde solo hay un lugar para el valor d, eliminarlo de los pares <== estrategia (2)


```go
func eliminate(values map[string]string, s string, v string) (map[string]string, bool) {
	// The value is already eliminated
	if !strings.Contains(values[s], v) {
		return values, true
	}

	// Remove all occurrences of the value (v) from the square possible values
	values[s] = strings.Replace(values[s], v, "", -1)

	// If a square (s) is reduced to one value (v2), then eliminate the value from the peers.
	if len(values[s]) == 0 {
		return nil, false
	} else if len(values[s]) == 1 {
		v2 := values[s]

		for s2 := range pares[s] {
			if _, ok := eliminate(values, s2, v2); !ok {
				return nil, false
			}
		}
	}

	// If a unit (u) has only one possible place for a value (v), then put it there.
	for _, u := range unidades[s] {
		dplaces := []string{}
		for _, s := range u {
			if strings.Contains(values[s], v) {
				dplaces = append(dplaces, s)
			}
		}

		if len(dplaces) == 0 {
			return nil, false
		} else if len(dplaces) == 1 {
			if assign(values, dplaces[0], v) == nil {
				return nil, false
			}
		}
	}

	return values, true
}
```

## Visualizacion 

Esta funcion no es mas que mostrar el resultado obtenido

```go
func Display(values map[string]string) {
	for i, row := range filas {
		for j, col := range digitos {
			if j == 3 || j == 6 {
				fmt.Printf("| ")
			}
			fmt.Printf("%v ", values[string(row)+string(col)])
		}
		fmt.Println()
		if i == 2 || i == 5 {
			fmt.Println("------+-------+-------")
		}
	}
}
````

## Busqueda 

Si todo va bien hasta el momento aun contaremos con que cada cuadro tiene una cantidad < 9 de posibilidades de digito para lo cual el autor nos indica lo siguiente, podemos probrar sistematicamente todas las posibilidades hasta dar co una que funcione, no obstante corremos el riesgo de que tarde mucho tiempo en correr por ejemplo supongamos que el cuadro A2 tiene 4 posibilidades mientras que A3 tiene 5 posibilidades, juntos son 20 y si seguimos multiplicando obtenemos  4,62838344192 × 1038 para todo el puzzle.

Ahora bien el algoritmo de busqueda es claro en lo siguiente :
* primero asegúrese de que no hayamos encontrado una solución o una contradicción, y si no,
* elige un cuadrado vacío y considera todos sus valores posibles.
* Uno a la vez, intente asignar al cuadrado cada valor y busque desde la posición resultante.

Ahora bien se usa una heurística común llamada valores minimos restantes, lo que hace que evaluemos las probabilidades mas bajas de un cuadrado , evalando entonces primero aquellos que tienen solo 2 posilidades 


