package main

import (
	"bufio"       // Para leer entrada del usuario con más control
	"crypto/rand" // Para generar números aleatorios seguros (criptográficamente)
	"fmt"         // Para imprimir en pantalla
	"math/big"    // Para trabajar con números grandes (necesario para rand.Int)
	"os"          // Para acceder a la entrada estándar (teclado)
	"strconv"     // Para convertir texto a números (por ejemplo, string -> int)
	"strings"     // Para manipular texto (quitar espacios, convertir a minúsculas, etc.)
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Generador de Contraseñas ===")

	// Pedir longitud
	fmt.Print("Ingresá la longitud de la contraseña: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	longitud, err := strconv.Atoi(input)
	if err != nil || longitud <= 0 {
		fmt.Println("⚠️ Longitud inválida.")
		return
	}

	// Selección de tipos de caracteres
	var caracteres string

	fmt.Print("¿Incluir letras minúsculas? (s/n): ")
	if leerOpcion(reader) == "s" {
		caracteres += "abcdefghijklmnopqrstuvwxyz"
	}

	fmt.Print("¿Incluir letras mayúsculas? (s/n): ")
	if leerOpcion(reader) == "s" {
		caracteres += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	fmt.Print("¿Incluir números? (s/n): ")
	if leerOpcion(reader) == "s" {
		caracteres += "0123456789"
	}

	fmt.Print("¿Incluir símbolos? (s/n): ")
	if leerOpcion(reader) == "s" {
		caracteres += "!@#$%^&*()_+"
	}

	if caracteres == "" {
		fmt.Println("⚠️ No seleccionaste ningún tipo de carácter. Abortando.")
		return
	}

	// Generar la contraseña
	contraseña := generarContraseña(longitud, caracteres)

	fmt.Println("\n🔐 Contraseña generada:", contraseña)

	// Esperar Enter para salir
	fmt.Println("\nPresioná Enter para salir...")
	reader.ReadString('\n')
}

// Función auxiliar para leer "s" o "n" en minúscula
func leerOpcion(reader *bufio.Reader) string {
	opcion, _ := reader.ReadString('\n')
	opcion = strings.ToLower(strings.TrimSpace(opcion))
	return opcion
}

// Genera la contraseña aleatoria
func generarContraseña(longitud int, caracteres string) string {
	var contraseña string

	for i := 0; i < longitud; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(caracteres))))
		if err != nil {
			panic(err)
		}
		contraseña += string(caracteres[num.Int64()])
	}

	return contraseña
}
