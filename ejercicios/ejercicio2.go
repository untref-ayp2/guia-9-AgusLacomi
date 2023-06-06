package ejercicios

type Tarea struct {
	Nombre string
	Tiempo int
}

func Ejercicio2(tareas []Tarea) {
	for i := 0; i < len(tareas)-1; i++ {
		for j := 0; j < len(tareas)-1; j++ {
			if tareas[j].Tiempo > tareas[j+1].Tiempo {
				aux := tareas[j]
				tareas[j] = tareas[j+1]
				tareas[j+1] = aux
			}
		}
	}
}
