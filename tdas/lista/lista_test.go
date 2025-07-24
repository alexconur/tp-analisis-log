package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListaEnlazadaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	// Verificar que la lista este vacia al inicio y que la cantidad de elementos sea 0
	require.True(t, lista.EstaVacia(), "La lista debería estar vacia al ser creada.")
	require.Equal(t, 0, lista.Largo(), "La lista deberia tener 0 elementos al ser creada")

	// Verificar los panicos correspondientes de una lista vacia
	require.Panics(t, func() { lista.VerPrimero() }, "Se esperaba un panico al ver el primer elemento de una lista vacia")
	require.Panics(t, func() { lista.VerUltimo() }, "Se esperaba un panico al ver el ultimo elemento de una lista vacia")
	require.Panics(t, func() { lista.BorrarPrimero() }, "Se esperaba un panico al borrar el ultimo elemento de una lista vacia")
}

func TestInsertarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	// Verificar que el primer elemento insertado sea el correspondiente
	lista.InsertarPrimero(10)
	require.False(t, lista.EstaVacia(), "La lista no deberia de estar vacia despues de haber insertado el primer elemento")
	require.Equal(t, 10, lista.VerPrimero(), "El primer elemento deberia de ser 10")
}

func TestInsertarUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	// Verificar que el ultimo elemento insertado sea el correspondiente
	lista.InsertarPrimero(10)
	lista.InsertarUltimo(20)
	require.False(t, lista.EstaVacia(), "La lista no deberia de estar vacia despues de haber insertado el ultimo elemento")
	require.Equal(t, 20, lista.VerUltimo(), "El ultimo elemento deberia ser 20")
}

func TestInsertarYEliminar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	// Insertamos los elementos
	lista.InsertarPrimero(10)
	lista.InsertarUltimo(20)

	// Verificamos que los elementos borrados sea el correspondiente de acuerdo al orden insertado
	require.Equal(t, 10, lista.BorrarPrimero(), "Deberia de eliminarse el primer elemento 10")
	require.Equal(t, 20, lista.VerPrimero(), "El nuevo primer elemento deberia de ser 20")
	require.Equal(t, 20, lista.BorrarPrimero(), "Deberia de eliminarse el primer elemento 20")
	require.True(t, lista.EstaVacia(), "La lista deberia de estar vacia luego de haber eliminado todos los elementos de la lista")
}

func TestVerificarLargo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	// Verificamos el largo al insertar y borrar el elemento
	require.Equal(t, 0, lista.Largo(), "El largo de una lista recien creada deberia de ser 0")
	lista.InsertarPrimero(10)
	require.Equal(t, 1, lista.Largo(), "El largo de la lista deberia de ser 1 despues de haber insertado el primer elemento")
	lista.InsertarUltimo(20)
	require.Equal(t, 2, lista.Largo(), "El largo de una lista deberia de ser 2 despues de haber insertado el ultimo elemento")
	lista.BorrarPrimero()
	require.Equal(t, 1, lista.Largo(), "EL largo de una lista deberia de ser 1 luego de haber borrado el primer elemento")
}

func TestCondicionesBordes(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	// Comprobar que al borrar elementos hasta que este vacia hace que la lista se comporte como recien creada.
	lista.InsertarPrimero(10)
	lista.InsertarUltimo(20)
	lista.InsertarUltimo(30)
	lista.InsertarUltimo(40)

	lista.BorrarPrimero()
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	lista.BorrarPrimero()

	// Comprobamos que la lista esta vacia
	require.True(t, lista.EstaVacia(), "La lista deberia de estar vacia despues de eliminar todos los elementos")
	require.Equal(t, 0, lista.Largo(), "El largo de la lista deberia de ser 0 luego de haber eliminado todos los elementos contenidos")

	// Las acciones de ver primero, ver ultimo y borrar ultimo deberian de dar panic
	require.Panics(t, func() { lista.VerPrimero() }, "Se esperaba un panico al ver el primer elemento de una lista vacia")
	require.Panics(t, func() { lista.VerUltimo() }, "Se esperaba un panico al ver el ultimo elemento de una lista vacia")
	require.Panics(t, func() { lista.BorrarPrimero() }, "Se esperaba un panico al borrar el ultimo elemento de una lista vacia")
}

func TestInsertarDiferentesTipos(t *testing.T) {
	// Prueba con lista de enteros
	listaInt := TDALista.CrearListaEnlazada[int]()
	listaInt.InsertarPrimero(10)
	listaInt.InsertarUltimo(20)
	listaInt.InsertarUltimo(30)
	require.Equal(t, 3, listaInt.Largo(), "El largo de la lista deberia de ser 3 luego de haber insertado tres elementos")
	require.Equal(t, 10, listaInt.VerPrimero(), "El primer elemento de la lista deberia de ser 10")
	require.Equal(t, 30, listaInt.VerUltimo(), "El ultimo elemento de la lista deberia de ser 30")
	require.False(t, listaInt.EstaVacia(), "La lista no deberia de estar vacia luego de haber insertado los elementos")

	require.Equal(t, 10, listaInt.BorrarPrimero(), "El elemento borrado deberia de ser 10")
	require.Equal(t, 20, listaInt.BorrarPrimero(), "El elemento borrado deberia de ser 20")
	require.Equal(t, 30, listaInt.BorrarPrimero(), "El elemento borrado deberia de ser 30")

	require.True(t, listaInt.EstaVacia(), "La lista deberia de estar vacia luego de haber borrado todos sus elementos")

	// Prueba con lista de cadenas
	listaString := TDALista.CrearListaEnlazada[string]()
	listaString.InsertarPrimero("a")
	listaString.InsertarUltimo("b")
	listaString.InsertarUltimo("c")
	require.Equal(t, 3, listaString.Largo(), "El largo de la lista deberia de ser 3 luego de haber insertado tres elementos")
	require.Equal(t, "a", listaString.VerPrimero(), "El primer elemento de la lista deberia de ser 'a'")
	require.Equal(t, "c", listaString.VerUltimo(), "El ultimo elemento de la lista deberia de ser 'c'")
	require.False(t, listaString.EstaVacia(), "La lista no deberia de estar vacia luego de haber insertado los elementos")

	require.Equal(t, "a", listaString.BorrarPrimero(), "El elemento borrado deberia de ser 'a'")
	require.Equal(t, "b", listaString.BorrarPrimero(), "El elemento borrado deberia de ser 'b'")
	require.Equal(t, "c", listaString.BorrarPrimero(), "El elemento borrado deberia de ser 'c'")

	require.True(t, listaString.EstaVacia(), "La lista deberia de estar vacia luego de haber borrado todos sus elementos")

	// Prueba con lista de float64
	listaFloat64 := TDALista.CrearListaEnlazada[float64]()
	listaFloat64.InsertarPrimero(10.25)
	listaFloat64.InsertarUltimo(20.35)
	listaFloat64.InsertarUltimo(30.45)
	require.Equal(t, 3, listaFloat64.Largo(), "El largo de la lista deberia de ser 3 luego de haber insertado tres elementos")
	require.Equal(t, 10.25, listaFloat64.VerPrimero(), "El primer elemento de la lista deberia de ser 10.25")
	require.Equal(t, 30.45, listaFloat64.VerUltimo(), "El ultimo elemento de la lista deberia de ser 30.45")
	require.False(t, listaFloat64.EstaVacia(), "La lista no deberia de estar vacia luego de haber insertado los elementos")

	require.Equal(t, 10.25, listaFloat64.BorrarPrimero(), "El elemento borrado deberia de ser 10.25")
	require.Equal(t, 20.35, listaFloat64.BorrarPrimero(), "El elemento borrado deberia de ser 20.35")
	require.Equal(t, 30.45, listaFloat64.BorrarPrimero(), "El elemento borrado deberia de ser 30.45")

	require.True(t, listaFloat64.EstaVacia(), "La lista deberia de estar vacia luego de haber borrado todos sus elementos")
}

func TestVolumen1000Elementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	const elementos = 10000

	// Insertar 10000 elementos en la lista, verificando el largo y el orden
	for i := 1; i <= elementos; i++ {
		lista.InsertarUltimo(i)
	}
	require.Equal(t, elementos, lista.Largo(), "El largo de la lista deberia ser 1000 despues de las inserciones.")

	for i := 1; i <= elementos; i++ {
		require.Equal(t, i, lista.BorrarPrimero(), "El primer elemento deberia ser %d.", i)
	}

	// Verificar que la lista esta vacia despues de borrar todos los elementos
	require.True(t, lista.EstaVacia(), "La lista deberia estar vacia despues de eliminar todos los elementos.")
	require.Equal(t, 0, lista.Largo(), "El largo de la lista deberia ser 0 despues de eliminar todos los elementos.")
}

// ------------------------------------ TEST PARA ITERADOR EXTERNO ------------------------------------ //
func TestIteradorInsertarAlPrincipio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()

	// Insertar al principio
	iter.Insertar(10)
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 10, lista.VerPrimero())
}

func TestIteradorInsertarAlFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iter := lista.Iterador()
	for iter.HaySiguiente() {
		iter.Siguiente()
	}

	iter.Insertar(10)
	require.Equal(t, 4, lista.Largo(), "El largo de la lista deberia de ser 4 luego de haber insertado al final")
	require.Equal(t, 10, lista.VerUltimo(), "El ultimo valor deberia de ser 10")

}

func TestIteradorInsertarEnElMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(3)

	iter := lista.Iterador()
	iter.Siguiente()
	iter.Insertar(2)

	require.Equal(t, 3, lista.Largo(), "El largo de la lista deberia de ser 3")
	iteradorPrueba := lista.Iterador()
	require.Equal(t, 1, iteradorPrueba.VerActual(), "El valor actual deberia de ser 1")
	iteradorPrueba.Siguiente()
	require.Equal(t, 2, iteradorPrueba.VerActual(), "El valor actual deberia de ser 2")
	iteradorPrueba.Siguiente()
	require.Equal(t, 3, iteradorPrueba.VerActual(), "El valor actual deberia de ser 3")
}

func TestIteradorBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iter := lista.Iterador()
	iter.Borrar()

	require.Equal(t, 2, lista.Largo(), "El largo de la lista deberia de ser 2")
	require.Equal(t, 2, lista.VerPrimero(), "El primer elemento deberia de ser 2")
}

func TestIteradorBorrarElementoEnMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iter := lista.Iterador()
	iter.Siguiente()
	iter.Borrar()

	require.Equal(t, 2, lista.Largo(), "El largo de la lista deberia de ser 2")

	iteradorPrueba := lista.Iterador()
	require.Equal(t, 1, iteradorPrueba.VerActual(), "El valor actual deberia de ser 1")
	iteradorPrueba.Siguiente()
	require.Equal(t, 3, iteradorPrueba.VerActual(), "El valor actual deberia de ser 3")
}

func TestIteradorBorrarUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iter := lista.Iterador()
	require.Equal(t, true, iter.HaySiguiente(), "El iterador deberia de tener siguiente al inicio")
	iter.Siguiente()
	require.Equal(t, true, iter.HaySiguiente(), "El iterador deberia de tener siguiente despues del primero")
	iter.Siguiente()
	iter.Borrar()

	require.Equal(t, 2, lista.Largo(), "El largo de la lista deberia de ser 2 luego de haber eliminado el ultimo elemento")
	require.Equal(t, 2, lista.VerUltimo(), "El ultimo elemento de la lista deberia de ser 2")
	require.Panics(t, func() { iter.VerActual() }, "VerActual deberia de dar panic cuando el iterador esta en nil")
}

func TestListaVaciaConIteradores(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()

	//Verificar panic al realizar siguientes funciones en lista vacia.
	require.Panics(t, func() { iter.VerActual() }, "VerActual deberia de dar panic en una lista vacia")
	require.Panics(t, func() { iter.Siguiente() }, "Siguiente deberia de dar panic en una lista vacia")
	require.Panics(t, func() { iter.Borrar() }, "Borrar deberia de dar panic en una lista vacia")

	//Comportamiento de iterador con lista vacia
	require.Equal(t, false, iter.HaySiguiente(), "El iterador no deberia de tener siguiente en una lista vacia")
}

func TestIterarListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	llamadas := 0
	lista.Iterar(func(v int) bool {
		llamadas++
		return true
	})
	//Al iterar no entra en ningun momento a la funcion
	require.Equal(t, 0, llamadas, "No se deberian de haber realizado llamadas a la funcion en una lista vacia")
}

func TestIterarListaConUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarPrimero("w")
	llamadas := 0
	lista.Iterar(func(v string) bool {
		llamadas++
		return true
	})
	//Al iterar entra una vez en en la funcion
	require.Equal(t, 1, llamadas, "Se deberia de haber realizado una sola llamada a la funcion con un elemento")
}

func TestIterarListaConElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarUltimo("w")
	lista.InsertarUltimo("a")
	lista.InsertarUltimo("s")
	letras := make([]string, 3)
	llamadas := 0
	lista.Iterar(func(v string) bool {
		letras[llamadas] = v
		llamadas++
		return true
	})
	//Recorre todas la lista sin restricciones
	require.Equal(t, 3, llamadas, "La funcion deberia haber sido llamada 3 veces")
	//Recorre en el orden correcto al utilizar iterar
	require.Equal(t, []string{"w", "a", "s"}, letras, "El orden de los elementos iterados deberia ser 'w', 'a', 's'")
}

func SumarPares(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i <= 10; i++ {
		lista.InsertarUltimo(i)
	}
	restriccion := 3
	sumaAlgunosPares := 0
	lista.Iterar(func(v int) bool {
		if v == restriccion {
			return false
		}
		if v%2 == 0 {
			sumaAlgunosPares += v
		}
		return true
	})
	sumaPares := 0
	lista.Iterar(func(v int) bool {
		if v%2 == 0 {
			sumaPares += v
		}
		return true
	})
	//Suma los numeros pares de la lista hasta la restriccion indicada.
	require.Equal(t, 2, sumaAlgunosPares, "La suma de los pares hasta la restriccion deberia ser 2")
	//Suma todos los numeros pares de la lista sin restricciones
	require.Equal(t, 30, sumaPares, "La suma de todos los pares deberia ser 30")
}

func TestEncontrarElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i <= 10; i++ {
		lista.InsertarUltimo(i)
	}
	elementoBuscado := 9
	encontrado := false
	lista.Iterar(func(v int) bool {
		if v == elementoBuscado {
			encontrado = true
			return false
		}
		return true
	})
	require.Equal(t, true, encontrado, "El elemento buscado deberia haberse encontrado en la lista")
}

func TestSumarAlgunosElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i <= 10; i++ {
		lista.InsertarPrimero(i)
	}
	restriccion := 9
	contador := 0
	suma := 0
	lista.Iterar(func(v int) bool {
		suma += v
		contador++
		return contador != restriccion
	})
	//Sumo primeros 9 numeros
	require.Equal(t, 54, suma, "La suma de los primeros 9 elementos deberia ser 54")
}

func TestNoEncontrarElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i <= 10; i++ {
		lista.InsertarUltimo(i)
	}
	elementoBuscado := 11
	encontrado := false
	lista.Iterar(func(v int) bool {
		if v == elementoBuscado {
			encontrado = true
			return false
		}
		return true
	})
	require.Equal(t, false, encontrado, "El elemento buscado no deberia estar en la lista")
}

func TestCasosBordesIterador(t *testing.T) {
	// Crear una lista vacia e iterar sobre esta
	listaVacia := TDALista.CrearListaEnlazada[int]()
	iterVacia := listaVacia.Iterador()
	require.False(t, iterVacia.HaySiguiente(), "El iterador de una lista vacia no deberia de tener siguiente")

	// Crear una lista con un solo elemento
	listaConUnElemento := TDALista.CrearListaEnlazada[int]()
	listaConUnElemento.InsertarUltimo(1)

	iterUnElemento := listaConUnElemento.Iterador()
	require.True(t, iterUnElemento.HaySiguiente(), "El iterador debe encontrar un elemento")
	require.Equal(t, 1, iterUnElemento.VerActual(), "El iterador debe devolver el unico elemento")
	iterUnElemento.Siguiente()
	require.False(t, iterUnElemento.HaySiguiente(), "Despues de avanzar, no deberia de haber mas elementos")

	// Crear una lista con varios elementos y eliminar mientras se esta iterando
	listaVarios := TDALista.CrearListaEnlazada[int]()
	listaVarios.InsertarUltimo(1)
	listaVarios.InsertarUltimo(2)
	listaVarios.InsertarUltimo(3)
	listaVarios.InsertarUltimo(4)

	iterVarios := listaVarios.Iterador()
	require.True(t, iterVarios.HaySiguiente(), "Debe haber un siguiente elemento")
	require.Equal(t, 1, iterVarios.VerActual(), "El primer elemento debe ser 1")

	iterVarios.Borrar()
	require.True(t, iterVarios.HaySiguiente(), "Debe haber un siguiente elemento")
	require.Equal(t, 2, iterVarios.VerActual(), "El nuevo elemento actual debe ser 2")

	iterVarios.Borrar()
	require.True(t, iterVarios.HaySiguiente(), "Debe haber un siguiente elemento")
	require.Equal(t, 3, iterVarios.VerActual(), "El nuevo elemento actual debe ser 3")

	iterVarios.Borrar()
	require.True(t, iterVarios.HaySiguiente(), "Debe haber un siguiente elemento")
	require.Equal(t, 4, iterVarios.VerActual(), "El nuevo elemento actual debe ser 4")

	iterVarios.Borrar()
	require.False(t, iterVarios.HaySiguiente(), "No debe de haber un siguiente elemento")

}

func TestIterarDistintosTiposDeDatos(t *testing.T) {
	// Caso string
	listaString := TDALista.CrearListaEnlazada[string]()
	listaString.InsertarUltimo("a")
	listaString.InsertarUltimo("b")
	listaString.InsertarUltimo("c")
	listaString.InsertarUltimo("d")

	iterString := listaString.Iterador()
	require.True(t, iterString.HaySiguiente(), "Debe de haber un siguiente elemento")
	require.Equal(t, "a", iterString.VerActual(), "El elemento actual deberia se 'a'")

	iterString.Siguiente()
	require.True(t, iterString.HaySiguiente(), "Debe de haber un siguiente elemento")
	require.Equal(t, "b", iterString.VerActual(), "El elemento actual deberia ser 'b'")

	iterString.Siguiente()
	require.True(t, iterString.HaySiguiente(), "Debe de haber un siguiente elemento")
	require.Equal(t, "c", iterString.VerActual(), "El elemento actual deberia de ser 'c'")

	iterString.Siguiente()
	require.True(t, iterString.HaySiguiente(), "Debe de haber un siguiente elemento")
	require.Equal(t, "d", iterString.VerActual(), "El elemento actual deberia de ser 'd'")

	iterString.Siguiente()
	require.False(t, iterString.HaySiguiente(), "No deberia de haber un siguiente elemento ")

	// Caso float64
	listaFloat64 := TDALista.CrearListaEnlazada[float64]()
	listaFloat64.InsertarUltimo(10.20)
	listaFloat64.InsertarUltimo(20.30)
	listaFloat64.InsertarUltimo(30.40)
	listaFloat64.InsertarUltimo(40.50)

	iterFloat64 := listaFloat64.Iterador()
	require.True(t, iterFloat64.HaySiguiente(), "Debe de haber un siguiente elemento")
	require.Equal(t, 10.20, iterFloat64.VerActual(), "El elemento actual deberia se 10.20")

	iterFloat64.Siguiente()
	require.True(t, iterFloat64.HaySiguiente(), "Debe de haber un siguiente elemento")
	require.Equal(t, 20.30, iterFloat64.VerActual(), "El elemento actual deberia ser 20.30")

	iterFloat64.Siguiente()
	require.True(t, iterFloat64.HaySiguiente(), "Debe de haber un siguiente elemento")
	require.Equal(t, 30.40, iterFloat64.VerActual(), "El elemento actual deberia de ser 30.40")

	iterFloat64.Siguiente()
	require.True(t, iterFloat64.HaySiguiente(), "Debe de haber un siguiente elemento")
	require.Equal(t, 40.50, iterFloat64.VerActual(), "El elemento actual deberia de ser 40.50")

	iterFloat64.Siguiente()
	require.False(t, iterFloat64.HaySiguiente(), "No deberia de haber un siguiente elemento ")
}

func TestPruebaVolumen(t *testing.T) {
	const volumen = 10000

	lista := TDALista.CrearListaEnlazada[int]()

	for i := 0; i < volumen; i++ {
		lista.InsertarUltimo(i)
	}

	require.Equal(t, volumen, lista.Largo(), "El largo de la lista deberia de ser igual al volumen insertado")

	// Verificamos que los elementos esten en el orden correcto
	iter := lista.Iterador()
	for i := 0; i < volumen; i++ {
		require.True(t, iter.HaySiguiente(), "El iterador deberia de tener un siguiente elemento")
		require.Equal(t, i, iter.VerActual(), "El elemento actual deberia de coincidir con el valor esperado")
		iter.Siguiente()
	}
}
