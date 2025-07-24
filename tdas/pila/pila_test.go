package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	// mas pruebas para este caso...

	// Verificar que la pila esta vacia al inicio y que la cantidad de elementos es 0
	require.True(t, pila.EstaVacia(), "La pila deberia estar vacia al inicializarla")

	// Intentar desapilar y ver el tope en una pila vacia y verificar el panico
	require.Panics(t, func() { pila.Desapilar() }, "Se esperaba un panico al intentar desapilar en una pila vacia")
	require.Panics(t, func() { pila.VerTope() }, "Se esperaba un panico al intentar ver el tope de una pila vacia")
}

func TestPilaVaciada(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	// Apilar y desapilar un elemento, verificar que la pila vuelva a estar vacía
	pila.Apilar(1)
	require.False(t, pila.EstaVacia(), "La pila no debería estar vacía después de apilar un elemento")
	require.Equal(t, 1, pila.VerTope(), "El tope de la pila debería ser el elemento apilado")

	require.Equal(t, 1, pila.Desapilar(), "El valor desapilado debería ser igual al valor apilado")
	require.True(t, pila.EstaVacia(), "La pila debería estar vacía después de desapilar el único elemento")

	// Intentar desapilar y ver el tope en una pila vaciada y verificar el pánico
	require.Panics(t, func() { pila.Desapilar() }, "Se esperaba un pánico al intentar desapilar en una pila vacía")
	require.Panics(t, func() { pila.VerTope() }, "Se esperaba un pánico al intentar ver el tope de una pila vacía")
}

func TestApilarYDesapilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(3)

	// Verificar el orden LIFO
	require.Equal(t, 3, pila.VerTope(), "El tope deberia ser 3")
	require.Equal(t, 3, pila.Desapilar(), "El elemento desapilado deberia ser 3")
	require.Equal(t, 2, pila.VerTope(), "El tope deberia ser 2")
	require.Equal(t, 2, pila.Desapilar(), "El elemento desapilado deberia ser 2")
	require.Equal(t, 1, pila.VerTope(), "El tope deberia ser 1")
	require.Equal(t, 1, pila.Desapilar(), "El elemento desapilado deberia ser 1")

	require.True(t, pila.EstaVacia(), "La pila deberia estar vacía después de desapilar todos los elementos")
}

func TestCondicionesBordes(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	// Comprobar que al desapilar hasta que esta vacia hace que la pila se comporte como recien creada.
	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(3)

	pila.Desapilar()
	pila.Desapilar()
	pila.Desapilar()

	require.True(t, pila.EstaVacia(), "La pila deberia de estar vacia despues de desapilar todos los elementos")

	// Las  acciones de desapilar y ver_tope en una pila recien creada son inválidas.
	require.Panics(t, func() { pila.Desapilar() }, "Se esperaba un panico al intentar desapilar en una pila vacia")
	require.Panics(t, func() { pila.VerTope() }, "Se esperaba un panico al intentar ver el tope de una pila vacia")

	// La accion de EstaVacia en una pila recien creada es verdadero.
	require.True(t, pila.EstaVacia())

	// Las acciones de desapilar y VerTope en una pila a la que se le apilo y desapilo hasta estar vacia son invalidas.
	pila.Apilar(1)
	pila.Desapilar()

	require.Panics(t, func() { pila.Desapilar() }, "Se esperaba un panico al intentar desapilar en una pila vacia")
	require.Panics(t, func() { pila.VerTope() }, "Se esperaba un panico al intentar ver el tope de una pila vacia")

}

func TestApilarDiferentesTipos(t *testing.T) {
	// Prueba con pila de enteros
	pilaInt := TDAPila.CrearPilaDinamica[int]()
	pilaInt.Apilar(1)
	pilaInt.Apilar(2)
	require.Equal(t, 2, pilaInt.VerTope(), "El tope de la pila de enteros deberia ser 2")
	require.False(t, pilaInt.EstaVacia(), "La pila no deberia de estar vacia despues de apilar los elementos")

	require.Equal(t, 2, pilaInt.Desapilar(), "El elemento desapilado deberia ser 2")
	require.Equal(t, 1, pilaInt.Desapilar(), "El elemento desapilado deberia ser 1")
	require.True(t, pilaInt.EstaVacia(), "La pila deberia estar vacia despues de desapilar todos los elementos")

	// Prueba con pila de cadenas
	pilaStr := TDAPila.CrearPilaDinamica[string]()
	pilaStr.Apilar("a")
	pilaStr.Apilar("b")
	require.Equal(t, "b", pilaStr.VerTope(), "El tope de la pila de cadenas deberia ser 'b'")
	require.False(t, pilaStr.EstaVacia(), "La pila no deberia de estar vacia despues de apilar los elementos")

	require.Equal(t, "b", pilaStr.Desapilar(), "El elemento desapilado deberia ser 'b'")
	require.Equal(t, "a", pilaStr.Desapilar(), "El elemento desapilado deberia ser 'a'")
	require.True(t, pilaStr.EstaVacia(), "La pila deberia estar vacia despues de desapilar todos los elementos")

	// Prueba con pila de strings
	pilaFloat64 := TDAPila.CrearPilaDinamica[float64]()
	pilaFloat64.Apilar(2.34)
	pilaFloat64.Apilar(60.82)
	require.Equal(t, 60.82, pilaFloat64.VerTope(), "El tope de la pila de cadenas deberia ser '60.82'")
	require.False(t, pilaFloat64.EstaVacia(), "La pila no deberia de estar vacia despues de apilar los elementos")

	require.Equal(t, 60.82, pilaFloat64.Desapilar(), "El elemento desapilado deberia ser '60.82'")
	require.Equal(t, 2.34, pilaFloat64.Desapilar(), "El elemento desapilado deberia ser '2.34'")
	require.True(t, pilaFloat64.EstaVacia(), "La pila deberia estar vacia despues de desapilar todos los elementos")
}

func TestVolumenPila(t *testing.T) {
	// Prueba con 1000 elementos
	const numElementos1000 = 1000
	pila := TDAPila.CrearPilaDinamica[int]()

	// Apilar elementos
	for i := 1; i <= numElementos1000; i++ {
		pila.Apilar(i)
	}

	// Verificar el orden LIFO
	for i := numElementos1000; i > 0; i-- {
		require.Equal(t, i, pila.VerTope(), "El tope deberoa ser %d", i)
		require.Equal(t, i, pila.Desapilar(), "El elemento desapilado deberia ser %d", i)
	}

	require.True(t, pila.EstaVacia(), "La pila deberia estar vacia despues de desapilar todos los elementos")
}
