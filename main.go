package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type StudentInfo struct {
	FullName string
	Grades   []int
	AvgGrade float64
}

func (s *StudentInfo) calculateAverage() {
	if len(s.Grades) == 0 {
		s.AvgGrade = 0.0
		return
	}
	sum := 0
	for _, grade := range s.Grades {
		sum += grade
	}
	s.AvgGrade = float64(sum) / float64(len(s.Grades))
}

func addStudent(students map[string]StudentInfo, reader *bufio.Reader) {
	fmt.Print("Введите фамилию и имя студента: ")
	fullName, _ := reader.ReadString('\n')
	fullName = strings.TrimSpace(fullName)

	if _, exists := students[fullName]; exists {
		fmt.Println("Студент с таким ФИО уже существует.")
		return
	}

	var grades []int
	fmt.Println("Введите оценки студента через пробел. Нажмите Enter еще раз для завершения.")
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			if len(grades) == 0 {
				fmt.Println("Оценки не были введены. Введите оценки или нажмите Enter еще раз для пропуска.")
				continue
			}
			break
		}

		gradesStrSlice := strings.Fields(input)
		for _, gradeStr := range gradesStrSlice {
			grade, err := strconv.Atoi(gradeStr)
			if err != nil {
				fmt.Printf("Некорректный ввод оценки '%s'. Пожалуйста, вводите только числа.\n", gradeStr)
				continue
			}
			if grade < 1 || grade > 5 {
				fmt.Printf("Оценка %d вне допустимого диапазона (1-5)\n", grade)
				continue
			}
			grades = append(grades, grade)
		}
	}

	student := StudentInfo{
		FullName: fullName,
		Grades:   grades,
	}
	student.calculateAverage()

	students[fullName] = student
	fmt.Println("Студент", fullName, "успешно добавлен!")
}

func filterStudentsByAvg(students map[string]StudentInfo, threshold float64) []StudentInfo {
	filteredStudents := []StudentInfo{}
	for _, student := range students {
		if student.AvgGrade < threshold {
			filteredStudents = append(filteredStudents, student)
		}
	}
	return filteredStudents
}

func printStudentInfo(student StudentInfo) {
	fmt.Printf("  ФИО: %s, Оценки: %v, Средний балл: %.2f\n", student.FullName, student.Grades, student.AvgGrade)
}

func printAllStudents(students map[string]StudentInfo) {
	fmt.Println("Список всех студентов:")
	for _, student := range students {
		printStudentInfo(student)
	}
}

func main() {
	students := make(map[string]StudentInfo)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Добро пожаловать в журнал!")

	for {
		fmt.Print("\nВведите команду (help - для вывода всех команд): ")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		switch command {
		case "add":
			addStudent(students, reader)
		case "list":
			if len(students) > 0 {
				printAllStudents(students)
			} else {
				fmt.Println("В базе данных пока нет студентов.")
			}
		case "filter":
			fmt.Print("Введите максимальный средний балл: ")
			thresholdStr, _ := reader.ReadString('\n')
			thresholdStr = strings.TrimSpace(thresholdStr)
			threshold, err := strconv.ParseFloat(thresholdStr, 64)
			if err != nil {
				fmt.Println("Некорректный ввод среднего балла. Пожалуйста, введите число.")
				continue
			}

			filteredStudents := filterStudentsByAvg(students, threshold)
			if len(filteredStudents) > 0 {
				fmt.Println("Студенты с средним баллом ниже", threshold)
				for _, student := range filteredStudents {
					printStudentInfo(student)
				}
			} else {
				fmt.Println("Нет студентов с средним баллом ниже", threshold)
			}
		case "help":
			fmt.Println("Доступные команды:")
			fmt.Println("  add - Добавить студента")
			fmt.Println("  list - Вывести список всех студентов")
			fmt.Println("  filter - Вывести студентов с средним баллом ниже указанного")
			fmt.Println("  help - Вывести список команд")
			fmt.Println("  exit - Выйти из программы")
		case "exit":
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Неизвестная команда. Введите 'help' для списка команд.")
		}
	}
}
