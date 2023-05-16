package ejercicios

type Tarea struct {
	Nombre string
	Tiempo int
}

func Ejercicio2(tareas []Tarea) {

	for i := 0; i < len(tareas)-1; i++ {
		for tareas[i].Tiempo > tareas[i+1].Tiempo {
			tareaMayor := tareas[i]
			tareas[i] = tareas[i+1]
			tareas[i+1] = tareaMayor
			if i != 0 {
				i--
			}
		}
	}
}
