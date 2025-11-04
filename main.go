package main

import (
	"fmt"
)

func main() {
	fmt.Println("Выберите операцию:")
	fmt.Println("1. Сложение матриц")
	fmt.Println("2. Вычитание матриц")
	fmt.Println("3. Умножение матрицы на скаляр")
	fmt.Println("4. Умножение двух матриц")

	var userChoice int
	fmt.Print("Ваш выбор: ")
	fmt.Scan(&userChoice)

	switch userChoice {
	case 1:
		executeMatrixOperation(addMatrices, "сложения")
	case 2:
		executeMatrixOperation(subtractMatrices, "вычитания")
	case 3:
		multiplyByScalar()
	case 4:
		multiplyMatrices()
	default:
		fmt.Println("Неверный ввод. Пожалуйста, выберите число от 1 до 4.")
	}
}

func readMatrix(label string) [][]float64 {
	var rows, cols int
	if label != "" {
		fmt.Printf("Введите размерность матрицы %s (строки столбцы): ", label)
	} else {
		fmt.Print("Введите размерность матрицы (строки столбцы): ")
	}
	fmt.Scan(&rows, &cols)

	matrix := make([][]float64, rows)
	fmt.Println("Введите элементы матрицы построчно (через пробел):")
	for i := 0; i < rows; i++ {
		matrix[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}
	return matrix
}

func printMatrix(matrix [][]float64) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%.2f\t", val)
		}
		fmt.Println()
	}
}

func addMatrices(a, b [][]float64) [][]float64 {
	return elementWiseOperation(a, b, func(x, y float64) float64 { return x + y })
}

func subtractMatrices(a, b [][]float64) [][]float64 {
	return elementWiseOperation(a, b, func(x, y float64) float64 { return x - y })
}

func elementWiseOperation(a, b [][]float64, op func(float64, float64) float64) [][]float64 {
	rows := len(a)
	cols := len(a[0])
	result := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		result[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			result[i][j] = op(a[i][j], b[i][j])
		}
	}
	return result
}

func multiplyByScalar() {
	matrix := readMatrix("")
	var scalar float64
	fmt.Print("Введите скаляр для умножения: ")
	fmt.Scan(&scalar)

	result := make([][]float64, len(matrix))
	for i := range matrix {
		result[i] = make([]float64, len(matrix[i]))
		for j := range matrix[i] {
			result[i][j] = matrix[i][j] * scalar
		}
	}

	fmt.Println("Результат умножения матрицы на скаляр:")
	printMatrix(result)
}

func multiplyMatrices() {
	a := readMatrix("A")
	b := readMatrix("B")

	if len(a[0]) != len(b) {
		fmt.Println("Ошибка: количество столбцов матрицы A должно совпадать с количеством строк матрицы B")
		return
	}

	result := make([][]float64, len(a))
	for i := range result {
		result[i] = make([]float64, len(b[0]))
		for j := range result[i] {
			for k := 0; k < len(b); k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	fmt.Println("Результат умножения матриц:")
	printMatrix(result)
}

type matrixOperation func([][]float64, [][]float64) [][]float64

func executeMatrixOperation(op matrixOperation, operationName string) {
	fmt.Printf("Введите матрицы для операции %s\n", operationName)
	a := readMatrix("A")
	b := readMatrix("B")

	if len(a) != len(b) || len(a[0]) != len(b[0]) {
		fmt.Println("Ошибка: матрицы должны иметь одинаковую размерность")
		return
	}

	result := op(a, b)
	fmt.Printf("Результат %s:\n", operationName)
	printMatrix(result)
}
