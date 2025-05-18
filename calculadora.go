package main

// Importamos el paquete "fmt" para entrada/salida en consola
import (
    "fmt"
)

func main() {
    // Declaramos las variables donde se guardarán los números ingresados
    var a, b float64

    // Variable para guardar la operación (+, -, *, /)
    var op string

    // Solicitamos el primer número al usuario
    fmt.Print("Ingresa el primer número: ")
    fmt.Scanln(&a) // Lee el número desde consola y lo guarda en 'a'

    // Solicitamos el segundo número
    fmt.Print("Ingresa el segundo número: ")
    fmt.Scanln(&b)

    // Pedimos al usuario que indique la operación
    fmt.Print("Elige operación (+, -, *, /): ")
    fmt.Scanln(&op)

    // Variable donde almacenaremos el resultado final
    var resultado float64

    // Usamos un switch para decidir qué operación realizar
    switch op {
    case "+":
        resultado = a + b
    case "-":
        resultado = a - b
    case "*":
        resultado = a * b
    case "/":
        // Validamos que el segundo número no sea cero antes de dividir
        if b != 0 {
            resultado = a / b
        } else {
            fmt.Println("Error: división por cero") // Mensaje de error
            return // Terminamos el programa si hay error
        }
    default:
        fmt.Println("Operación inválida") // Mensaje si el operador no es reconocido
        return
    }

    // Mostramos el resultado con 2 decimales
    fmt.Printf("Resultado: %.2f\n", resultado)
}
