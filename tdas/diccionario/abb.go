package diccionario

import TDAPila "tdas/pila"

type funcion_cmp[K comparable] func(K, K) int

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	valor     V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      funcion_cmp[K]
}

// PRE: tipo debe de ser una cadena valida
// POST: retorna el mensaje de error correspondiente al tipo proporcionado
func mensajesPanic(tipo string) string {
	if tipo == "diccionario" {
		return "La clave no pertenece al diccionario"
	}
	if tipo == "iterador" {
		return "El iterador termino de iterar"
	}
	return "Error desconocido"
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{raiz: nil, cantidad: 0, cmp: funcion_cmp}
}

func (a *abb[K, V]) Iterar(visitar func(clave K, valor V) bool) {
	a.iterarRecursivo(a.raiz, visitar)
}

// PRE: Nodo de tipo *nodoAbb[K, V] declarado y funcion visitar de tipo func(clave K, valor V) bool inicializada
// POST: Devuelve true o false
func (a *abb[K, V]) iterarRecursivo(nodo *nodoAbb[K, V], visitar func(clave K, valor V) bool) bool {
	if nodo == nil {
		return true
	}
	if !a.iterarRecursivo(nodo.izquierdo, visitar) {
		return false
	}
	if !visitar(nodo.clave, nodo.valor) {
		return false
	}
	return a.iterarRecursivo(nodo.derecho, visitar)
}

// PRE: Nodo de tipo *nodoAbb[K, V] declarado y variable clave de tipo comparativo inicializada
// POST: retorna el nodo que obtiene la clave dada y un booleano true, o nil si la clave no se encuentra y un booleano false
func (a *abb[K, V]) obtenerNodo(nodoActual *nodoAbb[K, V], clave K) (*nodoAbb[K, V], bool) {
	for nodoActual != nil {
		resultado := a.cmp(clave, nodoActual.clave)
		if resultado == 0 {
			return nodoActual, true
		} else if resultado > 0 {
			nodoActual = nodoActual.derecho
		} else {
			nodoActual = nodoActual.izquierdo
		}
	}
	return nil, false
}

func (a *abb[K, V]) Obtener(clave K) V {
	nodo, encontrado := a.obtenerNodo(a.raiz, clave)
	if !encontrado {
		panic(mensajesPanic("diccionario"))
	}
	return nodo.valor
}

func (a *abb[K, V]) Pertenece(clave K) bool {
	_, encontrado := a.obtenerNodo(a.raiz, clave)
	return encontrado
}

func (a *abb[K, V]) Cantidad() int {
	return a.cantidad
}

// PRE: Nodo de tipo *nodoAbb[K, V] declarado, variable clave de tipo comparativo y valor de tipo any inicializadas
// POST: retorna la referencia al nodo raiz del subarbol tras la insercion de la nueva clave y valor, si la clave ya existia actualiza el valor
func (a *abb[K, V]) insertarNodo(nodoActual *nodoAbb[K, V], clave K, valor V) *nodoAbb[K, V] {
	if nodoActual == nil {
		a.cantidad++
		return &nodoAbb[K, V]{clave: clave, valor: valor}
	}
	resultado := a.cmp(clave, nodoActual.clave)
	if resultado == 0 {
		nodoActual.valor = valor
	} else if resultado > 0 {
		nodoActual.derecho = a.insertarNodo(nodoActual.derecho, clave, valor)
	} else {
		nodoActual.izquierdo = a.insertarNodo(nodoActual.izquierdo, clave, valor)
	}
	return nodoActual
}

func (a *abb[K, V]) Guardar(clave K, valor V) {
	a.raiz = a.insertarNodo(a.raiz, clave, valor)
}

func (a *abb[K, V]) Borrar(clave K) V {
	var valorEliminado V
	a.raiz, valorEliminado = a.borrarNodo(a.raiz, clave)
	return valorEliminado
}

// PRE: Nodo de tipo *nodoAbb[K, V] declarado, variable clave de tipo comparativo y direccion de tipo string inicializadas
// POST: recorre el subarbol especificado (izquierdo o derecho) y retorna el nodo raiz modificado y el valor del nodo eliminado.
func (a *abb[K, V]) recorrerSubarbol(nodoActual *nodoAbb[K, V], clave K, direccion string) (*nodoAbb[K, V], V) {
	var valorEliminado V
	if direccion == "izquierda" {
		nodoActual.izquierdo, valorEliminado = a.borrarNodo(nodoActual.izquierdo, clave)
	} else {
		nodoActual.derecho, valorEliminado = a.borrarNodo(nodoActual.derecho, clave)
	}
	return nodoActual, valorEliminado
}

// PRE: el nodo no debe de ser nil
// POST: retorna true si el nodo no tiene hijos (si los punteros izquierdo y derecho son nil)
func (a *abb[K, V]) esHoja(nodo *nodoAbb[K, V]) bool {
	return nodo.izquierdo == nil && nodo.derecho == nil
}

// PRE: el nodo no debe de ser nil
// POST: retorna true si el nodo tiene exactamente un hijo (si el puntero izquierdo o derecho es nil)
func (a *abb[K, V]) nodoConHijoUnico(nodo *nodoAbb[K, V]) bool {
	return nodo.izquierdo == nil || nodo.derecho == nil
}

// PRE: los punteros izquierdo y derecho deben de ser nil
// POST: el nodo se elimina y retorna nil y el valor del nodo eliminado
func (a *abb[K, V]) eliminarHoja(nodo *nodoAbb[K, V]) (*nodoAbb[K, V], V) {
	a.cantidad--
	return nil, nodo.valor
}

// PRE: nodo debe de tener un solo hijo
// POST: se elimina el nodo y se retorna su unico hijo y el valor del nodo eliminado
func (a *abb[K, V]) eliminarNodoConUnHijo(nodo *nodoAbb[K, V]) (*nodoAbb[K, V], V) {
	a.cantidad--
	if nodo.izquierdo == nil {
		return nodo.derecho, nodo.valor
	}
	return nodo.izquierdo, nodo.valor

}

// PRE: el nodo debe de tener dos hijos
// POST: se reemplaza el nodo con su sucesor In-Order y se elimina dicho sucesor
func (a *abb[K, V]) eliminarNodoConDosHijos(nodo *nodoAbb[K, V]) (*nodoAbb[K, V], V) {
	sucesor := a.buscarMinimo(nodo.derecho)
	valorEliminado := nodo.valor
	nodo.clave, nodo.valor = sucesor.clave, sucesor.valor
	nodo.derecho, _ = a.borrarNodo(nodo.derecho, sucesor.clave)
	return nodo, valorEliminado
}

// PRE: nodoActual no debe de ser nil y la clave debe de ser valida
// POST: se elimina el nodo con la clave dada y se retorna el subarbol modificado junto con el valor del nodo eliminado
func (a *abb[K, V]) borrarNodo(nodoActual *nodoAbb[K, V], clave K) (*nodoAbb[K, V], V) {
	if nodoActual == nil {
		panic(mensajesPanic("diccionario"))
	}

	resultado := a.cmp(clave, nodoActual.clave)
	if resultado < 0 {
		return a.recorrerSubarbol(nodoActual, clave, "izquierda")
	} else if resultado > 0 {
		return a.recorrerSubarbol(nodoActual, clave, "derecha")
	}
	if a.esHoja(nodoActual) {
		return a.eliminarHoja(nodoActual)
	}
	if a.nodoConHijoUnico(nodoActual) {
		return a.eliminarNodoConUnHijo(nodoActual)
	}
	return a.eliminarNodoConDosHijos(nodoActual)
}

// PRE: nodoActual no debe de ser nil y tiene subarbol izquierdo
// POST: retorna el nodo con la clave minima en el subarbol enraizado en nodoActual
func (a *abb[K, V]) buscarMinimo(nodoActual *nodoAbb[K, V]) *nodoAbb[K, V] {
	actual := nodoActual
	for actual.izquierdo != nil {
		actual = actual.izquierdo
	}
	return actual
}

func (a *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, valor V) bool) {
	a.iterarRangoRecursivo(a.raiz, desde, hasta, visitar)
}

// PRE: Nodo de tipo *nodoAbb[K, V] declarado, variables desde y hasta de tipo comparativo inicializadas
//
//	y visitar de tipo funcion func(clave K, valor V) bool) bool inicializada
//
// POST: Devuelve true o false
func (a *abb[K, V]) iterarRangoRecursivo(nodo *nodoAbb[K, V], desde *K, hasta *K, visitar func(clave K, valor V) bool) bool {
	if nodo == nil {
		return true
	}
	if desde != nil && a.cmp(nodo.clave, *desde) < 0 {
		return a.iterarRangoRecursivo(nodo.derecho, desde, hasta, visitar)
	}
	if !a.iterarRangoRecursivo(nodo.izquierdo, desde, hasta, visitar) {
		return false
	}
	if hasta != nil && a.cmp(nodo.clave, *hasta) > 0 {
		return true
	}
	if !visitar(nodo.clave, nodo.valor) {
		return false
	}
	return a.iterarRangoRecursivo(nodo.derecho, desde, hasta, visitar)
}

type iteradorABB[K comparable, V any] struct {
	pila  TDAPila.Pila[*nodoAbb[K, V]]
	desde *K
	hasta *K
	cmp   funcion_cmp[K]
}

// Reutiliza IteradorRango con límites `nil` para iterar todo el ABB
func (a *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return a.IteradorRango(nil, nil)
}

func (iter *iteradorABB[K, V]) HaySiguiente() bool {
	if iter.pila.EstaVacia() {
		return false
	}
	clave := iter.pila.VerTope().clave
	if iter.hasta != nil && iter.cmp(clave, *iter.hasta) > 0 {
		return false
	}
	return true
}

func (iter *iteradorABB[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic(mensajesPanic("iterador"))
	}
	nodo := iter.pila.VerTope()
	return nodo.clave, nodo.valor
}

func (iter *iteradorABB[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic(mensajesPanic("iterador"))
	}
	nodo := iter.pila.Desapilar()
	iter.apilarRango(nodo.derecho)
}

// Reemplazamos el uso del iterador estándar y aplicamos el rango
func (a *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := &iteradorABB[K, V]{
		pila:  TDAPila.CrearPilaDinamica[*nodoAbb[K, V]](),
		desde: desde,
		hasta: hasta,
		cmp:   a.cmp,
	}

	// Apilamos los nodos comenzando desde 'desde' (si es necesario)
	iter.apilarRango(a.raiz)

	return iter
}

// PRE: Nodo de tipo *nodoAbb[K, V] declarado
// POST: Modifica la Pila utilizada para iterar.
func (iter *iteradorABB[K, V]) apilarRango(nodoActual *nodoAbb[K, V]) {
	for nodoActual != nil {
		if iter.desde == nil || iter.cmp(nodoActual.clave, *iter.desde) >= 0 {
			iter.pila.Apilar(nodoActual)
			nodoActual = nodoActual.izquierdo
		} else {
			nodoActual = nodoActual.derecho
		}
	}
}
