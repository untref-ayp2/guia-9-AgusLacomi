package ejercicios

type Objeto struct {
	Nombre string
	Peso   int
	Valor  int
}

func Ejercicio3(objetos []Objeto, capacidad int) map[Objeto]float32 {
	mochila := map[Objeto]float32{}

	for capacidad != 0 && len(objetos) != 0 {

		aux := objetos[0]
		var pos int

		for i := 1; i < len(objetos); i++ {
			if aux.Valor/aux.Peso < objetos[i].Valor/objetos[i].Peso {
				aux = objetos[i]
				pos = i
			}
		}

		if capacidad >= aux.Peso {
			mochila[aux] = 1
			capacidad -= aux.Peso
			objetos = append(objetos[:pos], objetos[pos+1:]...)
		} else {
			mochila[aux] = float32(capacidad) / float32(aux.Peso)
			capacidad = 0
		}
	}
	return mochila
}
