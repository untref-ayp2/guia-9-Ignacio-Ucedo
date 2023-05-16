package ejercicios

type Objeto struct {
	Nombre string
	Peso   int
	Valor  int
}

func Ejercicio3(objetos []Objeto, capacidad int) map[Objeto]float32 {

	// debería ordenar los objetos por su valor por unidad de peso
	for i := 0; i < len(objetos)-1; i++ {
		if objetos[i].Valor/objetos[i].Peso < objetos[i+1].Valor/objetos[i+1].Peso {
			objetoMayorValor := objetos[i]
			objetos[i] = objetos[i+1]
			objetos[i+1] = objetoMayorValor
			if i != 0 {
				i -= 2
			}
		}
	}

	//creo la mochila y meto los objetos

	mochila := make(map[Objeto]float32)

	for i := 0; i < len(objetos); i++ {
		if capacidad-objetos[i].Peso >= 0 {
			mochila[objetos[i]] = 1
			capacidad -= objetos[i].Peso
		} else {
			mochila[objetos[i]] = float32(capacidad) / float32(objetos[i].Peso)
			break
		}
	}

	return mochila
}
