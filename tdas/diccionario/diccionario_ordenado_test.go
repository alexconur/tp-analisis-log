package diccionario_test

import (
	"testing"

	TDADiccionario "tdas/diccionario"

	"math/rand"

	"github.com/stretchr/testify/require"
)

func cmpInt(a, b int) int {
	return a - b
}

func cmpStr(a, b string) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

func TestAbbVacio(t *testing.T) {
	abb := TDADiccionario.CrearABB[string, string](cmpStr)
	require.NotNil(t, abb, "Se espera que el ABB no sea nil")
	require.Equal(t, 0, abb.Cantidad(), "La cantidad de nodos debe de ser 0 al crear un nuevo ABB")
	require.False(t, abb.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar("A") })
}

func TestGuardarYObtener(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)

	// Guardamos claves y valores, verificando la cantidad
	abb.Guardar(10, "diez")
	abb.Guardar(20, "veinte")
	abb.Guardar(5, "cinco")
	require.Equal(t, 3, abb.Cantidad(), "La cantidad de nodos debe ser 3 despues de guardar tres nodos")

	// Obtenemos los valores asociados a las claves
	valor := abb.Obtener(10)
	require.Equal(t, "diez", valor, "Se espera obtener 'diez' para la clave 10")
	valor = abb.Obtener(5)
	require.Equal(t, "cinco", valor, "Se espera obtener 'cinco' para la clave 5")
	valor = abb.Obtener(20)
	require.Equal(t, "veinte", valor, "Se espera obtener 'veinte' para la clave 20")

	// Intentar obtener una clave que no existe
	require.Panics(t, func() { abb.Obtener(2) }, "Se esperaba un pánico al obtener una clave que no existe")
}

func TestGuardarYReemplazar(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)

	// Guardamos y luego reemplazamos un valor
	abb.Guardar(15, "quince")
	require.Equal(t, "quince", abb.Obtener(15), "Se espera obtener 'quince' para la clave 15")

	abb.Guardar(15, "nuevoQuince")
	require.Equal(t, "nuevoQuince", abb.Obtener(15), "Se espera obtener 'nuevoQuince' luego de haber actualizado la clave 15")

	// Cantidad sigue siendo 1
	require.Equal(t, 1, abb.Cantidad(), "La cantidad de nodos luego de actualizar un nodo debe de ser 1")
}

func TestPertenece(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)

	abb.Guardar(30, "treinta")
	abb.Guardar(10, "diez")

	require.True(t, abb.Pertenece(30), "Se espera que la clave 30 pertenezca al ABB")
	require.True(t, abb.Pertenece(10), "Se espera que la clave 10 pertenezca al ABB")
	require.False(t, abb.Pertenece(50), "Se espera que la clave 50 no pertenezca al ABB") // Clave no insertada
}

func TestBorrar(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)

	// Guardamos valores y borramos un nodo
	abb.Guardar(1, "uno")
	abb.Guardar(2, "dos")
	abb.Guardar(3, "tres")

	valor := abb.Borrar(2)
	require.Equal(t, "dos", valor, "Se espera obtener 'dos' al borrar la clave 2")
	require.Equal(t, 2, abb.Cantidad(), "La cantidad de nodos debe de ser 2 luego de haber eliminado un nodo")
	require.False(t, abb.Pertenece(2), "Se espera que la clave 2 no pertenezca al ABB")

	// Borrar una clave que no existe debe provocar pánico
	require.Panics(t, func() { abb.Borrar(5) }, "Se esperaba un panico al borrar una clave que no existe")
}

func TestObtenerClaveInexistente(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)

	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(100) })
}

func TestCantidad(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)

	require.Equal(t, 0, abb.Cantidad(), "La cantidad inicial de nodos debe de ser 0")
	abb.Guardar(1, "uno")
	require.Equal(t, 1, abb.Cantidad(), "La cantidad debe de ser 1 despues de agregar un nodo")
	abb.Guardar(2, "dos")
	require.Equal(t, 2, abb.Cantidad(), "La cantidad debe de ser 2 despues de agregar un nodo")
	abb.Borrar(2)
	require.Equal(t, 1, abb.Cantidad(), "La cantidad debe de ser 1 despues de borrar un nodo")
}

func TestIteradorInOrder(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)

	// Guardamos valores y creamos un iterador
	abb.Guardar(20, "valor20")
	abb.Guardar(10, "valor10")
	abb.Guardar(30, "valor30")

	iter := abb.Iterador()

	// Recorremos con el iterador
	require.True(t, iter.HaySiguiente(), "Se espera que haya un siguiente elemento en el iterador")
	clave, valor := iter.VerActual()
	require.Equal(t, 10, clave, "Se espera que la clave sea 10")
	require.Equal(t, "valor10", valor, "Se espera que el valor de la clave sea 'valor10'")

	iter.Siguiente()
	require.True(t, iter.HaySiguiente(), "Se espera que haya un siguiente elemento en el iterador")
	clave, valor = iter.VerActual()
	require.Equal(t, 20, clave, "Se espera que la clave sea 20")
	require.Equal(t, "valor20", valor, "Se espera que el valor de la clave sea 'valor20'")

	iter.Siguiente()
	require.True(t, iter.HaySiguiente(), "Se espera que haya un siguiente elemento en el iterador")
	clave, valor = iter.VerActual()
	require.Equal(t, 30, clave, "Se espera que la clave sea 30")
	require.Equal(t, "valor30", valor, "Se espera que el valor de la clave sea 'valor30'")

	iter.Siguiente()
	require.False(t, iter.HaySiguiente(), "Se espera que no haya un siguiente elemento en el iterador")
}

func TestIteradorVacio(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)

	iter := abb.Iterador()
	require.False(t, iter.HaySiguiente())

	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}
func TestIterar(t *testing.T) {

	abb := TDADiccionario.CrearABB[int, string](cmpInt)

	abb.Guardar(2, "dos")
	abb.Guardar(1, "uno")
	abb.Guardar(3, "tres")

	var claves []int
	abb.Iterar(func(clave int, valor string) bool {
		claves = append(claves, clave)
		return true
	})

	require.ElementsMatch(t, []int{1, 2, 3}, claves, "Se esperaba que las claves se iteraran en orden")
}

func TestIterarRango(t *testing.T) {

	abb := TDADiccionario.CrearABB[int, string](cmpInt)
	abb.Guardar(1, "uno")
	abb.Guardar(2, "dos")
	abb.Guardar(3, "tres")
	abb.Guardar(4, "cuatro")
	abb.Guardar(5, "cinco")

	// Iteramos en un rango que incluye solo algunos nodos
	var claves []int
	abb.IterarRango(&[]int{2}[0], &[]int{4}[0], func(clave int, valor string) bool {
		claves = append(claves, clave)
		return true
	})

	require.ElementsMatch(t, []int{2, 3, 4}, claves, "Se esperaba que las claves en el rango [2, 4] sean iteradas")

	// Iteramos en un rango fuera de los nodos
	var clavesVacias []int
	abb.IterarRango(&[]int{6}[0], &[]int{8}[0], func(clave int, valor string) bool {
		clavesVacias = append(clavesVacias, clave)
		return true
	})
	require.ElementsMatch(t, []int{}, clavesVacias, "Se esperaba que no hubiera claves en el rango [6, 8]")
}

func TestIterarConInterrupcion(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)

	abb.Guardar(1, "uno")
	abb.Guardar(2, "dos")
	abb.Guardar(3, "tres")

	var claves []int
	abb.Iterar(func(clave int, valor string) bool {
		claves = append(claves, clave)
		// Interrumpimos al encontrar la clave 2
		return clave != 2
	})

	require.ElementsMatch(t, []int{1, 2}, claves, "Se esperaba que la iteración se detuviera al encontrar la clave 2")
}

func TestAbbClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un Abb vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")
	abbStr := TDADiccionario.CrearABB[string, string](cmpStr)
	require.False(t, abbStr.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abbStr.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abbStr.Borrar("") })

	abbInt := TDADiccionario.CrearABB[int, string](cmpInt)
	require.False(t, abbInt.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abbInt.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abbInt.Borrar(0) })
}

func TestUnElemento(t *testing.T) {
	t.Log("Comprueba que Abb tiene sólo un elemento")
	abb := TDADiccionario.CrearABB[string, int](cmpStr)
	abb.Guardar("Rana", 10)
	require.EqualValues(t, 1, abb.Cantidad())
	require.True(t, abb.Pertenece("Rana"))
	require.False(t, abb.Pertenece("Sapo"))
	require.EqualValues(t, 10, abb.Obtener("Rana"))
	require.Panics(t, func() { abb.Obtener("Sapo") }, "La clave no pertenece al Abb")
}

func TestVolumen(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](cmpInt)
	cantidadTotal := 10000
	cantidadGuardada := 0
	dic := TDADiccionario.CrearHash[int, int]()
	for i := 0; i < cantidadTotal; i++ {
		clave := rand.Intn(cantidadTotal)
		if !abb.Pertenece(clave) {
			cantidadGuardada++
		}
		dic.Guardar(clave, i)
		abb.Guardar(clave, i)
	}
	require.EqualValues(t, cantidadGuardada, abb.Cantidad(), "La cantidad de elementos es incorrecta")
	ok := true
	iter := dic.Iterador()
	for iter.HaySiguiente() {
		clave, valor := iter.VerActual()
		ok = abb.Pertenece(clave)
		if !ok {
			break
		}
		ok = abb.Obtener(clave) == valor
		if !ok {
			break
		}
		abb.Borrar(clave)
		iter.Siguiente()
	}
	require.True(t, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(t, 0, abb.Cantidad())
}

func TestIteradorRangoVacio(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)

	var desde, hasta *int
	iter := abb.IteradorRango(desde, hasta)

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorRangoTodoAbb(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)
	claves := make([]int, 7)
	clavesEsperadas := []int{1, 2, 3, 4, 5, 7, 9}

	abb.Guardar(5, "")
	abb.Guardar(2, "")
	abb.Guardar(1, "")
	abb.Guardar(3, "")
	abb.Guardar(4, "")
	abb.Guardar(7, "")
	abb.Guardar(9, "")

	var desde, hasta *int
	iter := abb.IteradorRango(desde, hasta)

	i := 0
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		claves[i] = clave
		i++
		iter.Siguiente()
	}

	require.Equal(t, clavesEsperadas, claves, "Se espera que iteradorRango recorra In-Order")

}

func TestIteradorRangoUnElemento(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)
	claves := make([]int, 1)
	clavesEsperadas := []int{5}

	abb.Guardar(5, "")
	abb.Guardar(2, "")
	abb.Guardar(1, "")
	abb.Guardar(3, "")
	abb.Guardar(4, "")
	abb.Guardar(7, "")
	abb.Guardar(9, "")

	desde := new(int)
	hasta := new(int)
	*desde = 5
	*hasta = 6
	iter := abb.IteradorRango(desde, hasta)

	i := 0
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		claves[i] = clave
		i++
		iter.Siguiente()
	}

	require.Equal(t, clavesEsperadas, claves, "Se espera que iteradorRango recorra un solo elemento")

}

func TestIteradorRango(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)

	// Insertamos varios elementos en el ABB
	abb.Guardar(1, "uno")
	abb.Guardar(3, "tres")
	abb.Guardar(5, "cinco")
	abb.Guardar(7, "siete")
	abb.Guardar(9, "nueve")

	desde := new(int)
	hasta := new(int)
	*desde = 3
	*hasta = 7
	// Caso: Rango [3, 7] - Debe iterar por las claves 3, 5, 7
	iter := abb.IteradorRango(desde, hasta)
	require.True(t, iter.HaySiguiente(), "El iterador debería tener elementos")

	clave, _ := iter.VerActual()
	require.Equal(t, 3, clave, "La primera clave debería ser 3")
	iter.Siguiente()

	require.True(t, iter.HaySiguiente(), "El iterador debería tener más elementos")
	clave, _ = iter.VerActual()
	require.Equal(t, 5, clave, "La segunda clave debería ser 5")
	iter.Siguiente()

	require.True(t, iter.HaySiguiente(), "El iterador debería tener más elementos")
	clave, _ = iter.VerActual()
	require.Equal(t, 7, clave, "La tercera clave debería ser 7")
	iter.Siguiente()

	require.False(t, iter.HaySiguiente(), "El iterador no debería tener más elementos")
}

func TestIteradorRangoSinDesde(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)
	claves := make([]int, 5)
	clavesEsperadas := []int{1, 2, 3, 4, 5}

	abb.Guardar(5, "")
	abb.Guardar(2, "")
	abb.Guardar(1, "")
	abb.Guardar(3, "")
	abb.Guardar(4, "")
	abb.Guardar(7, "")
	abb.Guardar(9, "")

	desde := new(int)
	hasta := new(int)
	*hasta = 6
	iter := abb.IteradorRango(desde, hasta)

	i := 0
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		claves[i] = clave
		i++
		iter.Siguiente()
	}

	require.Equal(t, clavesEsperadas, claves, "Se espera que iteradorRango recorra hasta 'hasta' brindado por parámetro")

}

func TestIteradorRangoFueraDeRango(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)
	claves := make([]int, 0)
	clavesEsperadas := []int{}

	abb.Guardar(5, "")
	abb.Guardar(2, "")
	abb.Guardar(1, "")
	abb.Guardar(3, "")
	abb.Guardar(4, "")
	abb.Guardar(16, "")

	desde := new(int)
	hasta := new(int)
	*desde = 7
	*hasta = 10
	iter := abb.IteradorRango(desde, hasta)

	i := 0
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		claves[i] = clave
		i++
		iter.Siguiente()
	}

	require.Equal(t, clavesEsperadas, claves, "Se espera que iteradorRango no recorra Abb por estar fuera de rango")

}

func TestIterarABBEnRango(t *testing.T) {
	// Crear un nuevo ABB y agregar elementos
	abb := TDADiccionario.CrearABB[int, string](cmpInt)
	abb.Guardar(5, "Cinco")
	abb.Guardar(3, "Tres")
	abb.Guardar(8, "Ocho")
	abb.Guardar(2, "Dos")
	abb.Guardar(4, "Cuatro")
	abb.Guardar(10, "Diez")

	// Definir el rango
	desde := 3
	hasta := 8

	// Crear un slice para almacenar los resultados
	var resultados []string

	// Función que será llamada para cada elemento en el rango
	visitar := func(clave int, dato string) bool {
		// Agregamos el dato al slice y continuamos
		resultados = append(resultados, dato)
		return true
	}

	// Llamar a IterarRango
	abb.IterarRango(&desde, &hasta, visitar)

	// Verificar los resultados esperados
	expected := []string{"Tres", "Cuatro", "Cinco", "Ocho"}
	require.ElementsMatch(t, expected, resultados)
}

func TestIterarEnRangoIncluyeUnoSolo(t *testing.T) {
	// Crear un nuevo ABB y agregar elementos
	abb := TDADiccionario.CrearABB[int, string](cmpInt)
	abb.Guardar(5, "Cinco")
	abb.Guardar(3, "Tres")
	abb.Guardar(8, "Ocho")
	abb.Guardar(2, "Dos")
	abb.Guardar(4, "Cuatro")
	abb.Guardar(10, "Diez")

	// Definir el rango que incluye solo un elemento
	desde := 7
	hasta := 8

	// Crear un slice para almacenar los resultados
	var resultados []string

	// Función que será llamada para cada elemento en el rango
	visitar := func(clave int, dato string) bool {
		// Agregar el dato al slice y continuamos
		resultados = append(resultados, dato)
		return true
	}

	// Llamar a IterarRango
	abb.IterarRango(&desde, &hasta, visitar)

	// Verificar los resultados esperados
	expected := []string{"Ocho"} // Solo se espera "Ocho" en el rango
	require.ElementsMatch(t, expected, resultados)
}

func TestIteradorVolumenCompleto(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](cmpInt)
	cantidad := 10000

	// Insertar gran cantidad de elementos en el ABB
	for i := 0; i < cantidad; i++ {
		abb.Guardar(i, i*2)
	}

	// Iterador completo (sin rango) sobre todos los elementos
	iter := abb.Iterador()
	for i := 0; i < cantidad; i++ {
		require.True(t, iter.HaySiguiente(), "Al iterar debe de haber un siguiente elemento")
		clave, valor := iter.VerActual()
		require.Equal(t, i, clave, "Se espera que el clave asociada corresponda con su clave")
		require.Equal(t, i*2, valor, "Se espera que el valor sea el doble de la clave asociada")
		iter.Siguiente()
	}
	require.False(t, iter.HaySiguiente(), "Luego de iterar todos los elementos, se espera que no haya un siguiente")

}

func TestIteradorVolumenRango(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](cmpInt)
	cantidad := 10000

	// Insertar gran cantidad de elementos en el ABB
	for i := 0; i < cantidad; i++ {
		abb.Guardar(i, i*2)
	}

	// Iterador con rango parcial
	desde := 3000
	hasta := 7000
	iter := abb.IteradorRango(&desde, &hasta)

	for i := desde; i <= hasta; i++ {
		require.True(t, iter.HaySiguiente(), "Al iterar debe de haber un siguiente elemento")
		clave, valor := iter.VerActual()
		require.Equal(t, i, clave, "Se espera que la clave corresponda a la clave asociada")
		require.Equal(t, i*2, valor, "Se espera que el valor sea el doble de la clave asociada")
		iter.Siguiente()
	}
	require.False(t, iter.HaySiguiente(), "Luego de iterar todos los elementos, se espera que no haya un siguiente")
}

func TestIteradorABBVacio(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](cmpInt)

	iter := abb.Iterador()
	require.False(t, iter.HaySiguiente(), "Se espera que no haya un siguiente dentro de un iterador vacio")

	desde, hasta := 0, 100
	iterRango := abb.IteradorRango(&desde, &hasta)
	require.False(t, iterRango.HaySiguiente(), "Se espera que no haya un siguiente dentro de un iterador vacio")
}
