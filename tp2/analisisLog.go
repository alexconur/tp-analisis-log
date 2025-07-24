package main

import (
	"bufio"
	"os"
	"strings"
	TDADICC "tdas/diccionario"
	operacionesComandos "tp2/operComandos"
)

func main() {
	hash := TDADICC.CrearHash[string, int]()
	arbol := TDADICC.CrearABB[uint32, bool](operacionesComandos.CompararIPs)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		partes := strings.Fields(line)
		comando := partes[0]
		parametro1 := ""
		parametro2 := ""
		if len(partes) > 1 {
			parametro1 = partes[1]
		}
		if len(partes) > 2 {
			parametro2 = partes[2]
		}
		operacionesComandos.ProcesarEntrada(comando, parametro1, parametro2, hash, arbol)
	}
}
