package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
    "strings"
)

func main()  {
	
	//declaracion de variables
	var min, max float64
	var values [] float64
	var slide [] float64

	//entrada de min & max
	fmt.Println("Escribir un valores en una sola linea separada por espacios")

	fmt.Scan(&min)
	fmt.Scan(&max)

	//entrada de values
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	linea := scanner.Text()

	partes := strings.Fields(linea)

	for _, str := range partes{
		num, err := strconv.ParseFloat(str, 64)
		if err == nil {
			values = append(values, num)	
		}
	}

	for _, s := range values{

		if s >= min && s < max{
			slide = append(slide, s)
		}
	}
	
	//salida de datos 

	fmt.Println("valor minimo", min, "valor maximo", max, "valores", values)
	fmt.Println("EL rango es", slide)

	
}