# 🛡️ TP - Análisis de Logs para Detección de Servicios e IPs Sospechosas

Este proyecto permite analizar archivos de logs de acceso a servidores para detectar visitantes únicos (IPs), determinar los recursos más visitados y detectar posibles ataques de denegación de servicio (DoS) de forma automática.

## 📁 Estructura del Proyecto

- `analisisLog.go`: Punto de entrada del programa. Se encarga de leer comandos desde la entrada estándar e invocar el procesamiento.
- `comandos.go`: Contiene la lógica de ejecución de los comandos disponibles (`agregar_archivo`, `ver_visitantes`, `ver_mas_visitados`).
- `funcionesIPs.go`: Funciones auxiliares para conversión y comparación de direcciones IP, así como la carga de IPs en un ABB.
- `funcionesAuxiliares.go`: Implementa el procesamiento de recursos y detección de IPs sospechosas de realizar ataques DoS.
- `tdas/`: Implementaciones de estructuras como Hash, ABB y Heap utilizadas internamente.

## ⚙️ Tecnologías utilizadas

- Lenguaje: **Go (Golang)**
- Estructuras de datos: Diccionario (Hash), Árbol Binario de Búsqueda (ABB), Heap (Cola de Prioridad)

## 📌 Comandos disponibles

El programa se ejecuta leyendo comandos desde `stdin`. Los comandos disponibles son:

### `agregar_archivo <file>`

Carga un archivo de log y detecta si una direccion IP realiza 5 o mas peticiones en **menos de 2 segundos**, alertandolo por salida estandar como sospechosa de intento de DoS. Cada archivo se los considera independientemente entre si.

- **_Ejemplo_**: al ejecutar `agregar_archivo test05.log` se procesa el archivo log y detecta posibles casos de denegacion de servicio.

- **_Ejemplo de salida_**:
```bash
DoS: 83.149.10.216
OK
```

### `ver_visitantes <IP1> <IP2>`
Lista en orden todas las IPs que realizaron alguna peticion. Se mostraran las IPs unicamente dentro del rango que se ingreso, con los limites inclusive.

- **_Ejemplo_**: al ejecutar `ver_visitantes 83.149.9.0 110.136.166.0` mostrara todas las IPs que empiezan con **83** hasta los que empiezan con **110**

- **_Ejemplo de salida_**:
```bash
OK
Visitantes:
	83.149.9.216
	83.149.10.216
	93.114.45.13
OK
```

**⚠️ Advertencia**: se debe usar primero `agregar_archivo <ruta>` para poder analizar los rangos.

### `ver_mas_visitados <n>`
Muestra los **n** recursos mas solicitados. 

- **_Ejemplo_**: al ejecutar `ver_mas_visitados 3` mostrara los 3 recursos mas solicitados para todos los logs analizados. Se mostrara en orden descendente y, en caso de empate, se puede mostrar en cualquier orden.

- **_Ejemplo de salida_**:
```bash
Sitios más visitados:
	/album/movingpictures - 3
	/album/presto - 2
	/album/clockworkangels - 1
OK
```
## 📄 Compilacion

Antes que todo se debe compilar el archivo principal `analisisLog.go` de la siguiente manera:
```bash
go build analisisLog.go
```

### Pruebas Analogicas

Para poder ejecutar todas las pruebas dentro de la carpeta `pruebasAnalog` se debe ingresar a la carpeta y ejecutar el binario `pruebas.sh`

```bash
./pruebas.sh ../analisisLog
```

### Pruebas Unitarias

Para poder ejecutar un comando de forma unitaria se necesitara tener un archivo `.log` valido y un archivo `.txt` donde tendran los comandos a ejecutar, luego ejecutar:

```bash
./analisisLog < command.txt
```

**⚠️ Advertencia**: la primera linea del archivo `.txt` debe tener el comando `agregar_archivo <file>`, de lo contario no se podra ejecutar correctamente los otros comandos.