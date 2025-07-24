package operComandos

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	TDADICC "tdas/diccionario"
	"time"
)

const (
	LAYOUT               = "2006-01-02T15:04:05-07:00"
	CAPMINARRSOSPECHOSOS = 10
	FACTOR_EXPANSION     = 2
)

type recursoConConteo struct {
	recurso string
	conteo  int
}

type timestamps struct {
	times    [5]time.Time
	index    int
	contador int
}

// PRE: r1 y r2 son estructuras de tipo recursoConConteo inicializadas.
// POST: compara recursos por conteo en orden descendente
func compararRecursos(r1, r2 recursoConConteo) int {
	if r1.conteo > r2.conteo {
		return 1
	} else if r1.conteo < r2.conteo {
		return -1
	}
	return 0
}

// PRE: el hash debe de existir
// POST: ordena los recursos en un hash con un contador de repeticiones
func actualizarRecursos(hash TDADICC.Diccionario[string, int], file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linea := scanner.Text()
		campo := strings.Split(linea, "\t")
		recurso := campo[3]
		if hash.Pertenece(recurso) {
			conteoActual := hash.Obtener(recurso)
			hash.Guardar(recurso, conteoActual+1)
		} else {
			hash.Guardar(recurso, 1)
		}
	}
}

// PRE: 'logHash' y 'detectedDoS' son diccionarios válidos
// POST: Procesa el log, actualiza 'logHash' y registra IPs sospechosas de DoS en 'detectedDoS'.
func inicializarSospechososDoS(logHash TDADICC.Diccionario[string, timestamps], detectedDoS TDADICC.Diccionario[string, bool], file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		ip := parts[0]
		timestamp := parts[1]

		t, err := time.Parse(LAYOUT, timestamp)
		if err != nil {
			fmt.Println("Error al parsear la fecha:", err)
			continue
		}

		if !logHash.Pertenece(ip) {
			logHash.Guardar(ip, timestamps{index: 0, contador: 0})
		}

		timestamps := logHash.Obtener(ip)

		timestamps.times[timestamps.index] = t

		if timestamps.contador < 5 {
			timestamps.contador++
		}

		timestamps.index = (timestamps.index + 1) % 5

		logHash.Guardar(ip, timestamps)

		var diferencia time.Duration
		if timestamps.index == 0 {
			diferencia = timestamps.times[4].Sub(timestamps.times[0])
		} else {
			diferencia = timestamps.times[timestamps.index-1].Sub(timestamps.times[timestamps.index])
		}

		if timestamps.contador == 5 && diferencia < 2*time.Second {
			if !detectedDoS.Pertenece(ip) {
				detectedDoS.Guardar(ip, true)
			}
		}
	}
}

// PRE: 'detectedDoS' debe de estar inicializado con las IPs sospechosas de DoS
// POST: Crea un arreglo de IPs sospechosas desde el diccionario 'detectedDoS' con la cantidad de IPs encontradas.
func arrayDeSospechososDoS(detectedDoS TDADICC.Diccionario[string, bool]) []string {
	// Creo arreglo de tipo string de ips
	sospechosos := make([]string, CAPMINARRSOSPECHOSOS)
	i := 0
	//Recorro hash y guardo sospechosos en vector sospechosos
	for iter := detectedDoS.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		if i == cap(sospechosos) {
			sospechosos = redimensionarArregloDeSospechosos(sospechosos, cap(sospechosos)*FACTOR_EXPANSION, cap(sospechosos)*FACTOR_EXPANSION)
		}
		ip, _ := iter.VerActual()
		sospechosos[i] = ip
		i++
	} // ESto es O(m) donde m es mucho menor que n (en el peor caso, m es igual a n, pero esto es raro)
	return redimensionarArregloDeSospechosos(sospechosos, i, cap(sospechosos)) // Redimensionar el slice para que tenga solo la cantidad de elementos válidos (IPs sospechosas)
}

// PRE: sospechosos debe ser un slice de strings inicializado y nuevaCapacidad debe ser mayor que len(sospechosos).
// POST: retorna un nuevo slice con capacidad igual a nuevaCapacidad y copia los elementos de sospechosos.
func redimensionarArregloDeSospechosos(sospechosos []string, nuevoLargo, nuevaCapacidad int) []string {
	nuevosSospechosos := make([]string, nuevoLargo, nuevaCapacidad)
	copy(nuevosSospechosos, sospechosos)
	return nuevosSospechosos
}

// PRE: arr debe ser un slice de strings que representan direcciones IP validas.
// POST: retorna un slice con las IPs ordenadas en orden ascendente.
func radixSort(arr []string) []string {
	nums := make([]uint32, len(arr))
	for i, ip := range arr {
		num := ipStringANumero(ip)

		nums[i] = num
	}

	nums = countingSort(nums, 24)

	sortedIPs := make([]string, len(arr))
	for i, num := range nums {
		sortedIPs[i] = fmt.Sprintf("%d.%d.%d.%d", (num>>24)&0xFF, (num>>16)&0xFF, (num>>8)&0xFF, num&0xFF)
	}

	return sortedIPs
}

// PRE: arr debe ser un slice de numeros uint32 inicializado y exp debe ser un multiplo de 8 entre 0 y 24.
// POST: retorna un slice con los numeros ordenados segun el dígito correspondiente al exp.
func countingSort(arr []uint32, exp int) []uint32 {
	count := make([]int, 256)
	output := make([]uint32, len(arr))
	for _, num := range arr {
		index := (num >> exp) & 0xFF
		count[index]++
	}

	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	for i := len(arr) - 1; i >= 0; i-- {
		num := arr[i]
		index := (num >> exp) & 0xFF
		output[count[index]-1] = num
		count[index]--
	}

	copy(arr, output)
	return arr
}

// PRE: el arbol y el hash deben de estar vacios y el archivo debe de estar abierto en modo lectura
// POST: actualiza el ABB con direcciones IP del archivo y actualiza el hash con los recursos del archivo
func actualizarIpsYRecursos(arbol TDADICC.DiccionarioOrdenado[uint32, bool], hash TDADICC.Diccionario[string, int], file *os.File) {
	actualizarIPS(arbol, file)
	file.Seek(0, 0)
	actualizarRecursos(hash, file)
}

// PRE:
// POST: imprime los sospechosos DoS dentro del arreglo almacenado
func imprimirSospechosos(sospechosos []string) {
	for i := 0; i < len(sospechosos); i++ {
		fmt.Printf("DoS: %s\n", sospechosos[i])
	}
	fmt.Println("OK")
}

// PRE: el archivo debe de abrir sin error
// POST: procesa un archivo de log y detecta los DoS almacenandolos en un hash auxiliar
func sospechososDoS(file *os.File) {
	file.Seek(0, 0)
	logHash := TDADICC.CrearHash[string, timestamps]()
	detectedDoS := TDADICC.CrearHash[string, bool]()
	inicializarSospechososDoS(logHash, detectedDoS, file)

	sospechosos := arrayDeSospechososDoS(detectedDoS)
	sospechosos = radixSort(sospechosos)

	imprimirSospechosos(sospechosos)
}
