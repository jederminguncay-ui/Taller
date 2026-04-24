package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var texto string
	fmt.Println("Ingresa una frase que mas te interesa:")
	reader := bufio.NewReader(os.Stdin)
	texto, _ = reader.ReadString('\n')

	texto = strings.TrimSpace(texto)

	texto = strings.ToLower(texto)

	listaPalabras := strings.Fields(texto)

	palabras := make(map[string]int)

	for _, p := range listaPalabras {
		palabras[p]++
	}

	fmt.Println("\nConteo de las palabras de la frase:")
	for p, c := range palabras {
		fmt.Println(p, ":", c)
	}
	max := 0
	maxPalabra := ""

	for p, c := range palabras {
		if c > max {
			max = c
			maxPalabra = p
		}
	}

	fmt.Println("\nPalabra más repetida es :", maxPalabra, "(", max, ")")
}

package main

import "fmt"

type Producto struct {
	nombre   string
	precio   float64
	cantidad int
}

func MostrarProducto(p Producto) {
	fmt.Println("Nombre:", p.nombre)
	fmt.Println("Precio:", p.precio)
	fmt.Println("Cantidad:", p.cantidad)
}

func CalcularTotal(p Producto) float64 {
	return p.precio * float64(p.cantidad)
}

func TieneStock(p Producto) string {
	if p.cantidad > 0 {
		return "Disponible"
	}
	return "Se acabó el stock"
}

func main() {
	p1 := Producto{"Telefono", 400, 5}
	p2 := Producto{"Licuadora", 120, 4}
	p3 := Producto{"Parlante", 320, 0}

	productos := []Producto{p1, p2, p3}

	for i := 0; i < len(productos); i++ {
		fmt.Println("\nProducto", i+1)

		MostrarProducto(productos[i])

		total := CalcularTotal(productos[i])
		estadoProducto := TieneStock(productos[i])

		fmt.Println("Total en stock:", total)
		fmt.Println("Estado:", estadoProducto)
	}
}