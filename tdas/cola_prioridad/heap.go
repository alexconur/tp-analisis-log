package cola_prioridad

const (
	TAMANIO_INICIAL   int = 4
	FACTOR_EXPANSION  int = 2
	DIVISOR_REDUCCION int = 2
	FACTOR_REDUCCION  int = 4
)

// Función auxiliar para redimensionar el arreglo `datos` del heap
func redimensionar[T any](h *colaPrioridad[T], nuevaCapacidad int) {
	nuevosDatos := make([]T, nuevaCapacidad)
	copy(nuevosDatos, h.datos[:h.cantidad])
	h.datos = nuevosDatos
}

type colaPrioridad[T any] struct {
	datos       []T
	capacidad   int
	cantidad    int
	funcion_cmp func(T, T) int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &colaPrioridad[T]{
		datos:       make([]T, TAMANIO_INICIAL),
		capacidad:   TAMANIO_INICIAL,
		cantidad:    0,
		funcion_cmp: funcion_cmp,
	}
}

func CrearHeapArr[T any](array []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	capacidad := len(array)
	if capacidad < TAMANIO_INICIAL {
		capacidad = TAMANIO_INICIAL
	}

	datos := make([]T, capacidad)
	copy(datos, array)

	heap := &colaPrioridad[T]{
		datos:       datos,
		cantidad:    len(array),
		capacidad:   capacidad,
		funcion_cmp: funcion_cmp,
	}
	heapify(heap)
	return heap
}

// PRE: el heap contiene los elementos sin orden.
// POST: el heap cumple con la propiedad de orden de un heap.
func heapify[T any](h *colaPrioridad[T]) {
	ultimoInterno := (h.cantidad / 2) - 1
	for i := ultimoInterno; i >= 0; i-- {
		downHeap(h, i)
	}
}

// PRE:
// POST: retorna un mensaje de panico
func mensajePanic() string {
	return "La cola esta vacia"
}

// PRE: i y j tienen que ser indices validos en el arreglo h.datos.
// POST: se intercambian los elementos en las posiciones i y j de h.datos
func swap[T any](h *colaPrioridad[T], i, j int) {
	h.datos[i], h.datos[j] = h.datos[j], h.datos[i]
}

// Implementacion Primitiva: EstaVacia
func (h *colaPrioridad[T]) EstaVacia() bool {
	return h.cantidad == 0
}

// Implementacion Primitiva: Encolar
// Función para insertar un elemento en el heap
func (h *colaPrioridad[T]) Encolar(elem T) {
	if h.cantidad == h.capacidad {
		redimensionar(h, h.capacidad*FACTOR_EXPANSION)
		h.capacidad *= FACTOR_EXPANSION
	}
	h.datos[h.cantidad] = elem
	h.cantidad++
	upheap(h, h.cantidad-1)
}

// PRE: el heap es valido, excepto por el elemento en la posicion indicada.
// POST: el heap cumple con la propiedad de orden.
func upheap[T any](h *colaPrioridad[T], elem int) {
	if elem == 0 {
		return
	}
	padre := (elem - 1) / 2
	if h.funcion_cmp(h.datos[elem], h.datos[padre]) > 0 {
		swap(h, elem, padre)
		upheap(h, padre)
	}
}

// Implementacion Primitiva: VerMax
func (h *colaPrioridad[T]) VerMax() T {
	if h.EstaVacia() {
		panic(mensajePanic())
	}
	return h.datos[0]
}

// Implementacion Primitiva: Desencolar
func (h *colaPrioridad[T]) Desencolar() T {
	if h.EstaVacia() {
		panic(mensajePanic())
	}
	elemesEncolado := h.VerMax()
	swap(h, 0, h.cantidad-1)
	h.cantidad--
	downHeap(h, 0)

	if h.cantidad*FACTOR_REDUCCION <= h.capacidad && h.capacidad > TAMANIO_INICIAL {
		nuevaCapacidad := h.capacidad / DIVISOR_REDUCCION
		if nuevaCapacidad < TAMANIO_INICIAL {
			nuevaCapacidad = TAMANIO_INICIAL
		}
		redimensionar(h, nuevaCapacidad)
		h.capacidad = nuevaCapacidad
	}

	return elemesEncolado
}

// Implementacion Primitiva: HeapSort
func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	n := len(elementos)
	heap := &colaPrioridad[T]{datos: elementos, cantidad: n, funcion_cmp: funcion_cmp}
	heapify(heap)

	for i := n - 1; i > 0; i-- {
		elementos[i] = heap.Desencolar()
	}
}

// PRE: i tiene que ser un valor >= 0.
// POST: Retorna el índice de un hijo de i en un heap completo (depende de su orientacion)
func hijo(i int, orientacion string) int {
	if orientacion == "derecho" {
		return 2*i + 2
	}
	return 2*i + 1
}

// PRE: el heap es valido, excepto por el elemento en la posicion indicada.
// POST: el heap cumple con la propiedad de orden.
func downHeap[T any](h *colaPrioridad[T], i int) {
	hijoIzq := hijo(i, "izquierdo")

	if hijoIzq >= h.cantidad {
		return
	}
	hijoDer := hijo(i, "derecho")
	hijoMayor := hijoIzq

	if hijoDer < h.cantidad && h.funcion_cmp(h.datos[hijoDer], h.datos[hijoIzq]) >= 0 {
		hijoMayor = hijoDer
	}

	if h.funcion_cmp(h.datos[i], h.datos[hijoMayor]) >= 0 {
		return
	}

	swap(h, i, hijoMayor)
	downHeap(h, hijoMayor)
}

// Implementacion Primitiva: Cantidad
func (h *colaPrioridad[T]) Cantidad() int {
	return h.cantidad
}
