package main

import (
	"fmt"
)

func main() {
	var data1, data2 float64
	var operator string
	fmt.Print("silahkan ketik angka pertama")
	fmt.Scan(&data1)
	fmt.Print("silahkan ketik operator")
	fmt.Scan(&operator)
	fmt.Print("silahkan ketik angka kedua")
	fmt.Scan(&data2)

	var hasil float64

	switch operator {
	case "+":
		hasil = data1 + data2
	case "-":
		hasil = data1 - data2
	case "*":
		hasil = data1 * data2
	case "/":
		if data2 != 0 {
			hasil = data1 / data2
		} else {
			fmt.Println("Eror : pastikan bilangan tidak 0")
			return
		}
	default:
		fmt.Println("Invalid salah operasi")
		return

	}
	fmt.Println("Result: %f\n", hasil)

}
