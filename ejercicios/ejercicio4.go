package ejercicios

type Farolas struct {
	Posicion  int
	Radio     int
	Encendida bool
}

// Farolas públicas
// Hay M farolas en las posiciones y1,…,yM de una recta y N puntos x1,…,xN.
// Cada farola tiene un radio de iluminación ri, tal que la i-ésima farola ilumina puntos en
// el intervalo [yi−ri,yi+ri]. Se quiere encender el mínimo número de farolas tales que cada
// uno de los N puntos x1,…xN esté iluminado por al menos una farola.
// Encuentra este mínimo número.
// Entrada: un arreglo de Farolas de tamaño M y
// un arreglo de enteros de tamaño N con los puntos a iluminar.
// Salida: un arreglo de farolas encendidas
// Pre condiciones:
// El arreglo de farolas está ordenado por posicion de menor a mayor.
// Paso greedy: Dado un punto x, encender la farola más a la derecha que pueda iluminarlo.

func ordenarPuntos(x []int) []int {

	for i := 0; i < len(x)-1; i++ {
		if x[i] > x[i+1] {
			mayor := x[i]
			x[i] = x[i+1]
			x[i+1] = mayor

			if i > 0 {
				i -= 2
			}
		}

	}

	return x
}

func EncenderFarolas(farolas []Farolas, x []int) ([]Farolas, error) {
	x = ordenarPuntos(x)

	var resultado []Farolas
	var luz int // hasta qué punto llega la luz, (inicialmente es 0)

	for _, punto := range x {

		if punto > luz {

			for j := len(farolas) - 1; j >= 0; j-- {
				farolaActual := &farolas[j]
				inicio := farolaActual.Posicion - farolaActual.Radio
				fin := farolaActual.Posicion + farolaActual.Radio

				if inicio <= punto && punto <= fin {

					farolaActual.Encendida = true
					resultado = append(resultado, *farolaActual)
					luz = fin

					break
				}

			}
		}

	}

	return resultado, nil
}
