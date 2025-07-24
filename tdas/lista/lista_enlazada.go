package lista

type nodo[T any] struct {
	dato      T
	siguiente *nodo[T]
}

type listaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
	largo   int
}

type iteradorListaEnlazada[T any] struct {
	actual   *nodo[T]
	anterior *nodo[T]
	lista    *listaEnlazada[T]
}

func (lista *listaEnlazada[T]) crearNodo(elemento T) *nodo[T] {
	return &nodo[T]{dato: elemento}
}

/*
 *	Pre: Tipo de dato para elementos de la lista.
 *	Post: Devuelve una lista.
 */
func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{
		primero: nil,
		ultimo:  nil,
	}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(elemento T) {
	nuevoNodo := lista.crearNodo(elemento)
	if lista.EstaVacia() {
		lista.ultimo = nuevoNodo
	} else {
		nuevoNodo.siguiente = lista.primero
	}
	lista.primero = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(elemento T) {
	nuevoNodo := lista.crearNodo(elemento)
	if lista.EstaVacia() {
		lista.primero = nuevoNodo
	} else {
		lista.ultimo.siguiente = nuevoNodo
	}
	lista.ultimo = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	primerElemento := lista.VerPrimero()
	if lista.primero.siguiente == nil {
		lista.ultimo = nil
	}
	lista.primero = lista.primero.siguiente
	lista.largo--
	return primerElemento
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	for actual := lista.primero; actual != nil && visitar(actual.dato); actual = actual.siguiente {
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iteradorListaEnlazada[T]{
		actual: lista.primero,
		lista:  lista,
	}
}

func (it *iteradorListaEnlazada[T]) VerActual() T {
	if !it.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return it.actual.dato
}

func (iter *iteradorListaEnlazada[T]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iteradorListaEnlazada[T]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iter.anterior = iter.actual
	iter.actual = iter.actual.siguiente
}

func (iter *iteradorListaEnlazada[T]) Insertar(elemento T) {
	nuevoNodo := iter.lista.crearNodo(elemento)

	if iter.anterior == nil {
		nuevoNodo.siguiente = iter.actual
		iter.lista.primero = nuevoNodo
		if iter.lista.ultimo == nil {
			iter.lista.ultimo = nuevoNodo
		}
	} else if iter.actual == nil {
		iter.anterior.siguiente = nuevoNodo
		iter.lista.ultimo = nuevoNodo
	} else {
		nuevoNodo.siguiente = iter.actual
		iter.anterior.siguiente = nuevoNodo
	}
	iter.actual = nuevoNodo
	iter.lista.largo++

}

func (iter *iteradorListaEnlazada[T]) Borrar() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	dato := iter.VerActual()
	if iter.anterior == nil {
		iter.lista.primero = iter.actual.siguiente
	} else {
		iter.anterior.siguiente = iter.actual.siguiente
	}
	if iter.actual.siguiente == nil {
		iter.lista.ultimo = iter.anterior
	}
	iter.actual = iter.actual.siguiente
	iter.lista.largo--
	return dato
}
