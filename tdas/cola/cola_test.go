package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	// Verificar que la cola esta vacia al inicio y que la cantidad de elementos es 0
	require.True(t, cola.EstaVacia(), "La cola deberia estar vacía al inicializarla")

	// Intentar desencolar y ver el primero en una cola vacia y verificar el panico
	require.Panics(t, func() { cola.Desencolar() }, "Se esperaba un pánico al intentar desencolar en una cola vacía")
	require.Panics(t, func() { cola.VerPrimero() }, "Se esperaba un pánico al intentar desencolar en una cola vacía")
}

func TestEncolarYDesencolar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	// Encolamos elementos
	cola.Encolar(5)
	cola.Encolar(10)
	cola.Encolar(15)

	// Verificamos el orden FIFO
	require.Equal(t, 5, cola.VerPrimero(), "El primero deberia de ser 5")
	require.Equal(t, 5, cola.Desencolar(), "El elemento desencolado deberia ser 5")
	require.Equal(t, 10, cola.VerPrimero(), "El primero deberia de ser 10")
	require.Equal(t, 10, cola.Desencolar(), "El elemento desencolado deberia ser 10")
	require.Equal(t, 15, cola.VerPrimero(), "El primero deberia de ser 15")
	require.Equal(t, 15, cola.Desencolar(), "El elemento desencolado deberia ser 15")
	require.True(t, cola.EstaVacia(), "La cola deberia de estar vacia despues de desencolar todos los elementos")
}

func TestCondicionesBordes(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	// Comprobar que al desencolar hasta que este vacia hace que la cola se comporte como recien creada.
	cola.Encolar(1)
	cola.Encolar(2)
	cola.Encolar(3)

	cola.Desencolar()
	cola.Desencolar()
	cola.Desencolar()

	// Comprobamos que la cola esta vacia
	require.True(t, cola.EstaVacia(), "La cola deberia de estar vacia despues de desencolar todos los elementos")

	// Las acciones de desencolar y ver primero en una cola recien creada son invalidas.
	require.Panics(t, func() { cola.Desencolar() }, "Se esperaba un pánico al intentar desencolar en una cola vacía")
	require.Panics(t, func() { cola.VerPrimero() }, "Se esperaba un pánico al intentar desencolar en una cola vacía")
}

func TestEncolarDiferentesTipos(t *testing.T) {
	// Prueba con cola de enteros
	colaInt := TDACola.CrearColaEnlazada[int]()
	colaInt.Encolar(1)
	colaInt.Encolar(2)
	require.Equal(t, 1, colaInt.VerPrimero(), "El primero de la cola de enteros deberia ser 1")
	require.False(t, colaInt.EstaVacia(), "La cola no deberia de estar vacia despues de encolar los elementos")

	require.Equal(t, 1, colaInt.Desencolar(), "El elemento desencolado deberia ser 1")
	require.Equal(t, 2, colaInt.Desencolar(), "El elemento desencolado deberia ser 2")
	require.True(t, colaInt.EstaVacia(), "La cola deberia estar vacia despues de desencolar todos los elementos")

	// Prueba con cola de cadenas
	colaStr := TDACola.CrearColaEnlazada[string]()
	colaStr.Encolar("a")
	colaStr.Encolar("b")
	require.Equal(t, "a", colaStr.VerPrimero(), "El primero de la cola de enteros deberia ser 'a'")
	require.False(t, colaStr.EstaVacia(), "La cola no deberia de estar vacia despues de encolar los elementos")

	require.Equal(t, "a", colaStr.Desencolar(), "El elemento desencolado deberia ser 'a'")
	require.Equal(t, "b", colaStr.Desencolar(), "El elemento desencolado deberia ser 'b'")
	require.True(t, colaStr.EstaVacia(), "La cola deberia estar vacia despues de desencolar todos los elementos")

	// Prueba con cola de float64
	colaFloat64 := TDACola.CrearColaEnlazada[float64]()
	colaFloat64.Encolar(2.34)
	colaFloat64.Encolar(60.82)
	require.Equal(t, 2.34, colaFloat64.VerPrimero(), "El primero de la cola de enteros deberia ser 2.34")
	require.False(t, colaFloat64.EstaVacia(), "La cola no deberia de estar vacia despues de encolar los elementos")

	require.Equal(t, 2.34, colaFloat64.Desencolar(), "El elemento desencolado deberia ser 2.34")
	require.Equal(t, 60.82, colaFloat64.Desencolar(), "El elemento desencolado deberia ser 60.82")
	require.True(t, colaFloat64.EstaVacia(), "La cola deberia estar vacia despues de desencolar todos los elementos")
}

func TestVolumenCola(t *testing.T) {
	// Prueba con 1000 elementos
	const numElementos1000 = 1000
	cola := TDACola.CrearColaEnlazada[int]()

	// Encolamos elementos
	for i := 0; i < numElementos1000; i++ {
		cola.Encolar(i)
	}

	// Verificamos orden FIFO
	for i := 0; i < numElementos1000; i++ {
		require.Equal(t, i, cola.VerPrimero(), "El primero deberia ser %d", i)
		require.Equal(t, i, cola.Desencolar(), "El elemento desencolado deberia ser %d", i)
	}

	// Verificamos que la cola haya quedado vacia
	require.True(t, cola.EstaVacia(), "La cola deberia estar vacía despues de desencolar todos los elementos")
}
