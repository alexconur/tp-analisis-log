package pila

const (
	TAMANIO_INICIAL   int = 4
	FACTOR_EXPANSION  int = 2
	DIVISOR_REDUCCION int = 2
	FACTOR_REDUCCION  int = 4
	CAPACIDAD_MINIMA  int = 1
)

/* Definición del struct pila proporcionado por la cátedra. */
type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

// PRE:
// POST: retorna un mensaje de pila vacia
func mensajePanic() string {
	return "La pila esta vacia"
}

// PRE: la capacidad nueva debe ser un numero entero
// POST: retorna una pila con una nueva capacidad que la anterior
func redimensionarPila[T any](datos []T, capacidadNueva int) []T {
	nuevosDatos := make([]T, capacidadNueva)
	copy(nuevosDatos, datos)
	return nuevosDatos
}

// PRE:
// POST: crea una pila dinamica de tamanio 4
func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{
		datos:    make([]T, TAMANIO_INICIAL),
		cantidad: 0,
	}
}

// PRE:
// POST: retorna true si la cantidad de la pila es 0, si no es asi retorna false
func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

// PRE:
// POST: retorna el valor que esta en el tope de la pila, si no retorna un panic
func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic(mensajePanic())
	}
	return p.datos[p.cantidad-1]
}

// PRE:
// POST: apila el elemento deseado si hay espacio suficiente, sino se redimensiona y luego lo apila
func (p *pilaDinamica[T]) Apilar(elemento T) {
	if p.cantidad == cap(p.datos) {
		nuevaCapacidad := cap(p.datos) * FACTOR_EXPANSION
		p.datos = redimensionarPila(p.datos, nuevaCapacidad)
	}
	p.datos[p.cantidad] = elemento
	p.cantidad++
}

// PRE:
// POST: retorna el elemento que fue desapilado, si la pila esta vacia retorna un panic
func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic(mensajePanic())
	}
	tope := p.datos[p.cantidad-1]
	p.cantidad--

	if p.cantidad*FACTOR_REDUCCION <= cap(p.datos) && cap(p.datos) > CAPACIDAD_MINIMA {
		nuevaCapacidad := cap(p.datos) / DIVISOR_REDUCCION
		p.datos = redimensionarPila(p.datos, nuevaCapacidad)
	}
	return tope
}
