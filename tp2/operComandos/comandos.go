package operComandos

import (
	"fmt"
	"os"
	"strconv"
	TDAHEAP "tdas/cola_prioridad"
	TDADICC "tdas/diccionario"
)

// PRE: 'archivo' debe de ser una ruta valida a un archivo que se pueda abrir en modo lectura
// POST: devuelve un puntero al archivo abierto exitosamente. Si ocurre un error al intentar abrir el archivo escribe un mensaje de error en stderr, indicando el comando relacionado y devuelve nil
func abrirArchivo(archivo string, comando string) *os.File {
	file, err := os.Open(archivo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error en comando %s\n", comando)
		return nil
	}
	return file
}

// PRE: el arbol y el hash deben de estar vacios, el archivo debe de estar abierto en modo lectura y se debe de ingresar un comando existente
// POST: ejecuta el comando correspondiente, realizando la tarea del ejecutado. Si ocurre un error en algun caso, escribe el mensaje correspondiente en stderr y termina la ejecucion del comando.
func ProcesarEntrada(comando, parametro1, parametro2 string, hash TDADICC.Diccionario[string, int], arbol TDADICC.DiccionarioOrdenado[uint32, bool]) {
	switch comando {
	case "agregar_archivo":
		file := abrirArchivo(parametro1, comando)
		if file == nil {
			return
		}
		defer file.Close()
		actualizarIpsYRecursos(arbol, hash, file)
		sospechososDoS(file)

	case "ver_visitantes":
		if errorVerVisitantes(parametro2, comando) {
			return
		}
		verVisitantes(arbol, parametro1, parametro2)
	case "ver_mas_visitados":
		n, err := strconv.Atoi(parametro1)
		if errorVerMasVisitados(err, comando) {
			return
		}
		verMasVisitados(n, hash)
	default:
		fmt.Println("Comando no reconocido")
	}
}

// PRE: el arbol debe de existir, con las IPs inicializadas y ordenadas
// POST: itera el ABB y muestra las IPs dentro del rango especificado por parametro
func verVisitantes(arbol TDADICC.DiccionarioOrdenado[uint32, bool], desdeStr, hastaStr string) {
	desde := ipStringANumero(desdeStr)
	hasta := ipStringANumero(hastaStr)
	fmt.Println("Visitantes:")
	arbol.IterarRango(&desde, &hasta, func(clave uint32, dato bool) bool {
		fmt.Printf("\t%s\n", ipAString(clave))
		return true
	})
	fmt.Println("OK")
}

// PRE: se debe de pasar como parametro un comando valido
// POST: Devuelve `true` si `parametro2` esta vacio y escribe un mensaje de error en stderr. Devuelve `false` en caso contrario.
func errorVerVisitantes(parametro2, comando string) bool {
	if parametro2 == "" {
		fmt.Fprintf(os.Stderr, "Error en comando %s\n", comando)
		return true
	}
	return false
}

// PRE: debe de existir el hash con la información inicializada.
// POST: muestra los N recursos más solicitados en el log.
func verMasVisitados(n int, recursos TDADICC.Diccionario[string, int]) {
	heap := TDAHEAP.CrearHeap[recursoConConteo](compararRecursos)

	recursos.Iterar(func(recurso string, conteo int) bool {
		if heap.Cantidad() < n {
			heap.Encolar(recursoConConteo{recurso: recurso, conteo: conteo})
		} else if conteo > heap.VerMax().conteo {
			heap.Desencolar()
			heap.Encolar(recursoConConteo{recurso: recurso, conteo: conteo})
		}
		return true
	})

	// Mostramos los resultados en el orden en que fueron extraídos.
	fmt.Println("Sitios más visitados:")
	for i := 0; i < n && !heap.EstaVacia(); i++ {
		masVisitado := heap.Desencolar()
		fmt.Printf("\t%s - %d\n", masVisitado.recurso, masVisitado.conteo)
	}
	fmt.Println("OK")
}

// PRE: se debe de pasar como parametro un comando valido
// POST: Devuelve `true` si `err` no es nil y escribe un mensaje de error en stderr. Devuelve `false` en caso contrario.
func errorVerMasVisitados(err error, comando string) bool {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error en comando %s\n", comando)
		return true
	}
	return false
}
