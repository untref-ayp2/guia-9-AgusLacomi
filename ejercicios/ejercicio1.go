package ejercicios

type Actividad struct {
	Nombre string
	Inicio int
	Fin    int
}

// Selector de actividades recursivo
// Reescribir la función que resuelve el problema de seleccionar actividades en forma recursiva
func SelectorRecursivo(actividades []Actividad) []Actividad {
	if len(actividades) <= 1 {
		return actividades
	}
	panic("not implemented")
}
