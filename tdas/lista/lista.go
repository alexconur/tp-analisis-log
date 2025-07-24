package lista

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si la lista no tiene elementos, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero agrega un nuevo elemento al principio de la lista.
	InsertarPrimero(T)

	// InsertarUltimo agrega un nuevo elemento al final de la lista.
	InsertarUltimo(T)

	// BorrarPrimero saca el primer elemento de la lista y lo devuelve.
	//Si la lista está vacía, entra en pánico con un mensaje "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero obtiene el valor del primero de la lista. Si está vacía, entra en pánico con un mensaje
	// "La lista esta vacia".
	VerPrimero() T

	// VerUltimo obtiene el valor del ultimo de la lista. Si está vacía, entra en pánico con un mensaje
	// "La lista esta vacia".
	VerUltimo() T

	// Largo devuelve la cantidad de elementos en la lista.
	Largo() int

	// Iterar recibe una funcion "visitar" que devuelve un bool recibiendo previamente un dato.
	Iterar(visitar func(T) bool)

	// Iterador recibe un tipo de dato correspondiente a la lista.
	// Devuelve un iterador de lista con elementos del tipo de dato recibido.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// VerActual devuelve el elemento actual en el que está posicionado el iterador.
	VerActual() T

	// HaySiguiente devuelve verdadero si hay un elemento siguiente en la lista.
	HaySiguiente() bool

	// Siguiente avanza el iterador al siguiente elemento de la lista.
	Siguiente()

	// Insertar inserta un nuevo elemento en la posición actual del iterador.
	Insertar(T)

	// Borrar elimina el elemento actual del iterador y devuelve su valor.
	Borrar() T
}
