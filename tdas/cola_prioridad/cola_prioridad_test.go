package cola_prioridad_test

import (
	TDAHeap "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

func cmpInt(a, b int) int {
	return a - b
}

func TestHeapVacio(t *testing.T) {
	// Creamos un heap vacio
	heap := TDAHeap.CrearHeap[int](cmpInt)

	// Verificamos que esta vacia
	require.True(t, heap.EstaVacia(), "El heap debe de estar vacio al ser creado")
	require.Equal(t, 0, heap.Cantidad(), "Se espera que la cantidad de un heap vacio sea 0")
}

func TestHeapUnElemento(t *testing.T) {
	// Creamos un heap vacio y encolamos el elemento 10
	heap := TDAHeap.CrearHeap[int](cmpInt)
	heap.Encolar(10)

	// Verificamos que solo haya un elemento
	require.False(t, heap.EstaVacia(), "Se espera que al encolar un solo elemento el heap no este vacio")
	require.Equal(t, 1, heap.Cantidad(), "Se espera que haya solo un elemento dentro del heap")
	require.Equal(t, 10, heap.VerMax(), "El valor maximo del unico elemento encolado debe de ser '10'")
}

func TestDesencolarUnicoElemento(t *testing.T) {
	// Creamos un heap vacio y encolamos el elemento 10
	heap := TDAHeap.CrearHeap[int](cmpInt)
	heap.Encolar(10)

	// Desencolamos el unico elemento y verificamos que el heap vuelve a estar vacio
	require.Equal(t, 10, heap.Desencolar(), "Se espera que el elemento desencolado sea '10'")
	require.True(t, heap.EstaVacia(), "El heap debe de estar vacio")
}

func TestDesencolar(t *testing.T) {
	// Creamos un heap vacio y encolamos elementos
	heap := TDAHeap.CrearHeap[int](cmpInt)
	heap.Encolar(15)
	heap.Encolar(20)
	heap.Encolar(10)

	// Desencolamos y verificamos que se realice de la manera correcta
	require.Equal(t, 20, heap.Desencolar(), "El elemento desencolado debe de ser '20'")
	require.Equal(t, 15, heap.VerMax(), "El elemento maximo del heap actual debe de ser '15'")
	require.Equal(t, 15, heap.Desencolar(), "El elemento desencolado debe de ser '15'")
	require.Equal(t, 10, heap.Desencolar(), "El elemento desencolado debe de ser '10'")
	require.True(t, heap.EstaVacia(), "El heap vuelve a estar vacio una vez que se desencolaron todos los elementos")
}

func TestEncolarYVerMax(t *testing.T) {
	// Creamos un heap vacio y encolamos elementos
	heap := TDAHeap.CrearHeap[int](cmpInt)

	heap.Encolar(10)
	require.False(t, heap.EstaVacia(), "El heap no deberia de estar vacio luego de haber encolado un elemento")
	require.Equal(t, 10, heap.VerMax(), "El elemento maximo del heap debe ser '10'")
	heap.Encolar(20)
	require.Equal(t, 20, heap.VerMax(), "El elemento maximo del heap debe ser '20'")
	heap.Encolar(5)
	require.Equal(t, 20, heap.VerMax(), "El elemento maximo del heap debe ser '20'")
}

func TestHeapConValoresMixtos(t *testing.T) {
	// Inicializamos un array y le damos forma de heap
	heap := TDAHeap.CrearHeap[int](cmpInt)
	valores := []int{10, -1, 0, 5, -10, 3}

	// Encolamos los valores del array y verificamos sus operaciones
	for _, numeros := range valores {
		heap.Encolar(numeros)
	}

	require.Equal(t, 10, heap.VerMax(), "El elemento maximo deberia ser '10'")

	require.Equal(t, 10, heap.Desencolar(), "Desencolar deberia devolver '10'")
	require.Equal(t, 5, heap.VerMax(), "El siguiente maximo deberia ser '5'")
}

func TestCrearHeapArr(t *testing.T) {
	// Inicializamos un array y le damos forma de heap
	array := []int{3, 5, 1, 8, 7}
	heap := TDAHeap.CrearHeapArr(array, cmpInt)

	// Comprobamos sus valores
	require.Equal(t, 8, heap.VerMax(), "El elemento maximo del array debe de ser '8'")
	require.Equal(t, len(array), heap.Cantidad(), "La cantidad de elementos del array debe de ser correcta")
}

func TestCrearHeapArrVacio(t *testing.T) {
	// Inicializamos un array y le damos forma de heap
	array := []int{}
	heap := TDAHeap.CrearHeapArr(array, cmpInt)

	// Comprobamos sus valores
	require.True(t, heap.EstaVacia(), "El heap deberia estar vacio al ser creado desde un array vacio")
	require.Equal(t, 0, heap.Cantidad(), "La cantidad de un heap creado desde un array vacio deberia ser 0")
}

func TestHeapVolumen(t *testing.T) {
	// Creamos un heap vacio
	elementos := 10000
	heap := TDAHeap.CrearHeap[int](cmpInt)

	// Encolamos todos los elementos en el heap y verificamos su cantidad y su maximo
	for i := 0; i < elementos; i++ {
		heap.Encolar(i)
	}
	require.Equal(t, elementos, heap.Cantidad(), "La cantidad de elementos en el heap deberia ser %d", elementos)
	require.Equal(t, 9999, heap.VerMax(), "El elemento maximo en el heap deberia ser '9999'")

	// Verificamos que se desencolen correctamente
	for i := 9999; i >= 0; i-- {
		require.Equal(t, i, heap.Desencolar(), "Se esperaba desencolar el elemento %d", i)
	}

	// Verificamos que el heap este vacia
	require.True(t, heap.EstaVacia(), "El heap deberia estar vacio despues de desencolar todos los elementos")
}

func TestHeapArrVolumen(t *testing.T) {
	// Creamos un arreglo vacio
	elementos := 10000
	array := make([]int, elementos)

	// Guardamos todos los elementos en el arreglo de menor a mayor
	for i := 0; i < elementos; i++ {
		array[i] = i
	}

	// Pasamos arreglo por parametro para crear heap en base al mismo
	// Tendra que hacer el respectivo heapify al crear el heap
	heap := TDAHeap.CrearHeapArr(array, cmpInt)

	// Verifico si se agregaron correctamente los elementos al heap
	require.Equal(t, elementos, heap.Cantidad(), "La cantidad de elementos en el heap deberia ser %d", elementos)
	require.Equal(t, 9999, heap.VerMax(), "El elemento maximo en el heap deberia ser '9999'")

	// Verificamos que se desencolen correctamente
	for i := 9999; i >= 0; i-- {
		require.Equal(t, i, heap.Desencolar(), "Se esperaba desencolar el elemento %d", i)
	}

	// Verificamos que el heap este vacia
	require.True(t, heap.EstaVacia(), "El heap deberia estar vacio despues de desencolar todos los elementos")
}

func TestHeapSortArregloVacio(t *testing.T) {
	// Creamos un arreglo vacio y verificamos que al aplicar Heapsort este vacio
	elementos := []int{}
	TDAHeap.HeapSort(elementos, cmpInt)
	require.Empty(t, elementos, "El arreglo debe de estar vacio")
}

func TestHeapSortUnElemento(t *testing.T) {
	// Creamos un arreglo con un solo elemento y verificamos que al aplicar Heapsort sea el elemento correcto
	elementos := []int{5}
	TDAHeap.HeapSort(elementos, cmpInt)
	require.Equal(t, []int{5}, elementos, "El arreglo debe de contener el elemento '5'")

}

func TestHeapSortArregloOrdenadoAsc(t *testing.T) {
	// Creamos un arreglo ordenado ascendentemente y verificamos que al aplicar Heapsort se mantenga igual
	elementos := []int{1, 2, 3, 4, 5}
	TDAHeap.HeapSort(elementos, cmpInt)
	require.Equal(t, []int{1, 2, 3, 4, 5}, elementos, "El arreglo debe de mantener su orden (ascendente)")
}

func TestHeapSortArregloOrdenadoDesc(t *testing.T) {
	// Creamos un arreglo desordenado descendentemente y verificamos que al aplicar Heapsort se ordene ascendentemente
	elementos := []int{5, 4, 3, 2, 1}
	TDAHeap.HeapSort(elementos, cmpInt)
	require.Equal(t, []int{1, 2, 3, 4, 5}, elementos, "El arreglo debe de estar ordenado ascendentemente")
}

func TestHeapSortDesordenado(t *testing.T) {
	// Creamos un arreglo desordenado y verificamos que al aplicar Heapsort se ordene ascendentemente
	elementos := []int{3, 1, 4, 5, 2}
	TDAHeap.HeapSort(elementos, cmpInt)
	require.Equal(t, []int{1, 2, 3, 4, 5}, elementos, "El arreglo debe de estar ordenado ascendentemente")
}

func TestHeapSortVolumen(t *testing.T) {
	// Creamos un arreglo de volumen descendiente y verificamos que al aplicar Heapsort se ordene ascendentemente
	elementos := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		elementos[i] = 10000 - i
	}
	TDAHeap.HeapSort(elementos, cmpInt)

	// Verificamos que el arreglo este en orden ascendente y verificamos que cada elemento sea menor o igual al siguiente
	for i := 1; i < len(elementos); i++ {
		require.True(t, elementos[i-1] <= elementos[i], "El arreglo no esta en orden ascendente")
	}
}

func TestHeapCrearArrNoModificaArregloOriginal(t *testing.T) {
	// Dado un arreglo original
	original := []int{2, 6, 5, 19, 13, 18, 15, 10, 17, 7}

	// Cuando se crea un heap desde el arreglo
	heap := TDAHeap.CrearHeapArr(original, cmpInt)

	// Entonces el arreglo original no debe modificarse
	require.Equal(t, []int{2, 6, 5, 19, 13, 18, 15, 10, 17, 7}, original, "El arreglo original no debe de ser modificado")
	require.Equal(t, 19, heap.VerMax(), "El maximo deberia de ser '19'") // Verificamos que el máximo del heap sea correcto
}

func TestEncolar2DesencolarEncolar(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](cmpInt)

	// Encolar dos elementos
	heap.Encolar(200)
	heap.Encolar(300)

	// Verificar que el máximo sea el correcto
	require.Equal(t, 300, heap.VerMax(), "El elemento maximo deberia de ser '300'")

	// Desencolar el elemento máximo
	desencolado := heap.Desencolar()
	require.Equal(t, 300, desencolado, "Se esperaba desencolar '300'")

	// Encolar de nuevo un elemento
	heap.Encolar(250)

	// Verificar que el nuevo máximo sea correcto
	require.Equal(t, 250, heap.VerMax(), "El nuevo elemento maximo deberia de ser '250'")
	require.Equal(t, 2, heap.Cantidad(), "Se esperaba una cantidad total de '2'")
}

func TestHeapCrearArrGeneral(t *testing.T) {
	// Dado un arreglo de prueba
	array := []int{10, 20, 5, 30, 15}

	// Cuando se crea un heap desde el arreglo
	heap := TDAHeap.CrearHeapArr(array, cmpInt)

	// Verificar que el maximo inicial sea correcto
	require.Equal(t, 30, heap.VerMax(), "El maximo del arreglo deberia de ser '30'")
	require.Equal(t, 5, heap.Cantidad(), "La cantidad del arreglo debe de ser '5'")

	// Encolar un nuevo elemento
	heap.Encolar(25)

	// Verificar que el nuevo maximo no se haya actualizado
	require.Equal(t, 30, heap.VerMax(), "El maximo del arreglo deberia de ser '30'")
	require.Equal(t, 6, heap.Cantidad(), "La cantidad del arreglo debe de ser '6'")

	// Desencolar el máximo
	_ = heap.Desencolar()

	// Verificar el nuevo máximo
	require.Equal(t, 25, heap.VerMax(), "El nuevo maximo del arreglo deberia de ser '25'")
	require.Equal(t, 5, heap.Cantidad(), "La cantidad del arreglo debe de ser '5'")
}
