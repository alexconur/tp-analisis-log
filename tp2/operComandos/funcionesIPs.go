package operComandos

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	TDADICC "tdas/diccionario"
)

// PRE: ipStr debe ser una dirección IP válida en formato string
// POST: convierte una direccion IP de tipo string a un numero comparable (uint32)
func ipStringANumero(ipStr string) uint32 {
	ip := net.ParseIP(ipStr).To4()
	if ip == nil {
		return 0
	}
	return uint32(ip[0])<<24 + uint32(ip[1])<<16 + uint32(ip[2])<<8 + uint32(ip[3])
}

// PRE: ip debe ser una dirección IP válida representada como uint32.
// POST: convierte un numero IP a formato de tipo string
func ipAString(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

// PRE: ip1 e ip2 deben ser direcciones IP válidas representadas como uint32.
// POST: compara dos direcciones IP representadas como uint32
func CompararIPs(ip1, ip2 uint32) int {
	if ip1 < ip2 {
		return -1
	} else if ip1 > ip2 {
		return 1
	}
	return 0
}

// PRE: el arbol debe de existir
// POST: ordena las IPs en un ABB
func actualizarIPS(arbol TDADICC.DiccionarioOrdenado[uint32, bool], file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linea := scanner.Text()
		parte := strings.Split(linea, "\t")
		ip := ipStringANumero(parte[0])
		if !arbol.Pertenece(ip) {
			arbol.Guardar(ip, true)
		}
	}
}
