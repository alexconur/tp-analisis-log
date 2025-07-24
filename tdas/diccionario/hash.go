package diccionario

//LINK DE LA PAGINA:     https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function

//-------------------------------------- ALGORITMO DE WIKIPEDIA (ORIGINAL) ------------------------------------

//algorithm fnv-1a
//    hash := FNV_offset_basis
//    for each byte_of_data to be hashed do
//        hash := hash XOR byte_of_data
//        hash := hash × FNV_prime
//    return hash

//---------------------------------------- EXPLICACION DE ALGORITMO ORIGINAL ----------------------------------

//    All variables, except for byte_of_data, are 64-bit unsigned integers.
//    The variable, byte_of_data, is an 8-bit unsigned integer.
//    The FNV_offset_basis is the 64-bit value: 14695981039346656037 (in hex, 0xcbf29ce484222325).
//    The FNV_prime is the 64-bit value 1099511628211 (in hex, 0x100000001b3).
//    The multiply returns the lower 64 bits of the product.
//    The XOR is an 8-bit operation that modifies only the lower 8-bits of the hash value.
//    The hash value returned is a 64-bit unsigned integer.

import "fmt"

const (
	capacidadInicial        = 16
	factorDeCarga           = 0.70
	FNVOffsetBasis   uint64 = 14695981039346656037
	FNVPrime         uint64 = 1099511628211
	VACIO            int    = 0
	OCUPADO          int    = 1
	BORRADO          int    = 2
	FACTOR_EXPANSION int    = 2
)

type hashCerrado[K comparable, V any] struct {
	tabla    []elemento[K, V]
	cantidad int
	borrados int
}

type elemento[K comparable, V any] struct {
	clave  K
	valor  V
	estado int
}

type iteradorHash[K comparable, V any] struct {
	hash     *hashCerrado[K, V]
	posicion int
}

func (h *hashCerrado[K, V]) inicializarTabla(capacidad int) {
	h.tabla = make([]elemento[K, V], capacidad)
	h.cantidad = 0
	h.borrados = 0
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	h := &hashCerrado[K, V]{}
	h.inicializarTabla(capacidadInicial)
	return h
}

func (h *hashCerrado[K, V]) calcularHash(clave K) int {
	hash := fnv1aHash(convertirABytes(clave))
	return int(hash % uint64(len(h.tabla)))
}

func fnv1aHash(data []byte) uint64 {
	hash := FNVOffsetBasis

	for _, byteOfData := range data {
		hash ^= uint64(byteOfData)
		hash *= FNVPrime
	}

	return hash
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func (h *hashCerrado[K, V]) obtenerElemento(clave K) (int, *elemento[K, V]) {
	indice := h.calcularHash(clave)
	elem := &h.tabla[indice]

	for elem.estado == OCUPADO && elem.clave != clave {
		indice, elem = h.sondeoLineal(indice)
	}

	return indice, elem
}

func (h *hashCerrado[K, V]) verificarRedimension() bool {
	return float64(h.cantidad+h.borrados)/float64(len(h.tabla)) > factorDeCarga
}

func (h *hashCerrado[K, V]) sondeoLineal(indice int) (int, *elemento[K, V]) {
	return (indice + 1) % len(h.tabla), &h.tabla[indice]
}

func mensajePanic(tipo string) string {
	if tipo == "diccionario" {
		return "La clave no pertenece al diccionario"
	}
	if tipo == "iterador" {
		return "El iterador termino de iterar"
	}
	return "Error desconocido"
}

func (h *hashCerrado[K, V]) Guardar(clave K, valor V) {
	if h.verificarRedimension() {
		h.rehash()
	}

	_, elem := h.obtenerElemento(clave)

	if elem.estado == OCUPADO && elem.clave == clave {
		elem.valor = valor
	} else {
		if elem.estado == BORRADO {
			h.borrados--
		}
		h.cantidad++
		elem.clave = clave
		elem.valor = valor
		elem.estado = OCUPADO
	}
}

func (h *hashCerrado[K, V]) Pertenece(clave K) bool {
	_, elem := h.obtenerElemento(clave)
	return elem.estado == OCUPADO && elem.clave == clave
}

func (h *hashCerrado[K, V]) Obtener(clave K) V {
	_, elem := h.obtenerElemento(clave)

	if elem.estado == OCUPADO && elem.clave == clave {
		return elem.valor
	}
	panic(mensajePanic("diccionario"))
}

func (h *hashCerrado[K, V]) Borrar(clave K) V {
	if h.verificarRedimension() {
		h.rehash()
	}
	indice, elem := h.obtenerElemento(clave)

	for elem.estado != VACIO {
		if elem.estado == OCUPADO && elem.clave == clave {
			valor := elem.valor
			elem.estado = BORRADO
			h.cantidad--
			h.borrados++
			return valor
		}
		indice, elem = h.sondeoLineal(indice)
	}
	panic(mensajePanic("diccionario"))
}

func (h *hashCerrado[K, V]) Cantidad() int {
	return h.cantidad
}

func (h *hashCerrado[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for i := 0; i < len(h.tabla); i++ {
		elem := h.tabla[i]
		if elem.estado == OCUPADO {
			if !visitar(elem.clave, elem.valor) {
				return
			}
		}
	}
}

func (h *hashCerrado[K, V]) rehash() {
	viejaTabla := h.tabla
	h.cantidad, h.borrados = 0, 0
	h.inicializarTabla(len(viejaTabla) * FACTOR_EXPANSION)

	for i := 0; i < len(viejaTabla); i++ {
		elem := viejaTabla[i]
		if elem.estado == OCUPADO {
			h.Guardar(elem.clave, elem.valor)
		}
	}
}

func (h *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	return &iteradorHash[K, V]{hash: h}
}

func (it *iteradorHash[K, V]) HaySiguiente() bool {
	for it.posicion < len(it.hash.tabla) {
		if it.hash.tabla[it.posicion].estado == OCUPADO {
			return true
		}
		it.posicion++
	}
	return false
}

func (it *iteradorHash[K, V]) VerActual() (K, V) {
	if !it.HaySiguiente() {
		panic(mensajePanic("iterador"))
	}
	elem := it.hash.tabla[it.posicion]
	return elem.clave, elem.valor
}

func (it *iteradorHash[K, V]) Siguiente() {
	if !it.HaySiguiente() {
		panic(mensajePanic("iterador"))
	}
	it.posicion++
}
