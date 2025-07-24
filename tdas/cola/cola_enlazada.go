package cola

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

// PRE:
// POST: retorna un mensaje de cola vacia
func mensajePanic() string {
	return "La cola esta vacia"
}

// PRE:
// POST: crea un nuevo nodo de la cola
func nuevoNodoCola[T any](valor T) *nodoCola[T] {
	return &nodoCola[T]{dato: valor}
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{}
}

// PRE: la cola debe de existir
// POST: retorna true si no hay elementos en la cola, en caso contrario retorna false
func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.primero == nil
}

// PRE: la cola no debe de estar vacia
// POST: retorna el primer valor de la cola
func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic(mensajePanic())
	}
	return c.primero.dato
}

// PRE: la cola debe de existir
// POST: el elemento es agregado al final de la cola
func (c *colaEnlazada[T]) Encolar(valor T) {
	nuevoNodo := nuevoNodoCola(valor)
	if c.EstaVacia() {
		c.primero = nuevoNodo
	} else {
		c.ultimo.prox = nuevoNodo
	}
	c.ultimo = nuevoNodo
}

// PRE: debe de haber elementos en la cola
// POST: el primer elemento es removido y retornado, actualizando la cola
func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic(mensajePanic())
	}
	valor := c.primero.dato
	c.primero = c.primero.prox

	if c.primero == nil {
		c.ultimo = nil
	}
	return valor
}
