package GoCalculator

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

type Calculator struct {}

func (Calculator) Operate(number1 string, number2 string, operador string) int {
	response := 0

	operador1 := parsear(number1) 
	operador2 := parsear(number2) 

	switch operador {
		case "+":
			response = operador1 + operador2
			break
		case "-":
			response = operador1 - operador2
			break
		case "*":
			response = operador1 * operador2
			break
		case "/":
			response = operador1 / operador2
			break
		default:
			response = 0
	}

	return response
}


func parsear (entrada string) int {
	number, error := strconv.Atoi(entrada)

	if error != nil {
		fmt.Printf("Ocurrio un error al intentar parsear la entrada: %s, valor de salida: %d", error, number)
	}

	return number
}


func LeerEntrada(mensaje string) string {
	fmt.Println(mensaje)
	
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	entrada := scanner.Text()

	return entrada
}