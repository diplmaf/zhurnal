# Журнал Николаева ИС-323 ✅
---
- package main
<br> oбъявление основного пакета программы

**import (
<br>"bufio"
<br>"fmt"
<br>"os"
<br>"strconv"
<br>"strings"
<br>)**
- далее импорт необходимых компонентов Go:
1. bufio - буферизированный ввод/вывод
2. fmt - ввод/вывод
3. os - доступ к аргументам командной строки
4. strconv - преобразование строк
5. strings - работа со строками
---
**type StudentInfo struct {
<br>	FullName    string
<br>	Grades []int
<br>	AvgGrade    float64
<br>}**
- создание структуры данных для хранения информации:
1. FullName - ФИО студента. тип данных - строка
2. Grades - список оценок. тип данных - массив целых чисел
3. AvgGrade - средний балл. тип данных - число с плавающей точкой
---
**func (s *StudentInfo) calculateAverage(){***
- объявление функции с указателем на StudentInfo для вычислениия среднего балла
---
**if len(s.Grades) == 0 {
<br>s.AvgGrade = 0.0
<br>return**
- проверка наличия оценок у студента, если оценок нет, устанавливается средний балл 0.0, и выходит из метода
---
**sum := 0
<br>for _, grade := range s.Grades {
<br>		sum += grade
<br>	}**
- инициализация переменной для суммы оценок, а так же цикл во всем оценкам студента и добавление каждой оценки к сумме
---
**s.AvgGrade= float64(sum) / float64(len(s.Grades))**
- вычисление среднего балла, как суммы оценок и добавление каждой оценки к сумме, после чего идет закрытие метода
---
**func addStudent(students map[string]StudentInfo, reader*bufio.Reader)**
- объявление функции добавления студента
---
**fmt.Print("Введите фамилию и имя студента: ")**
- вывод приглашения для ввода ФИО студента
---
**fullName, _ := reader.ReadString('\n')**
- чтение ввода пользователя до символа новой строки
---
**fullName = strings.TrimSpace(fullName)**
- удаление лишних пробелов с начала и конца введенной строки
**if _, exists := students[fullName]; exists {
<br>		fmt.Println("Студент с таким ФИО уже существует.")
<br>	return**
- проверка существования студента с введенным фио, если существует - вывод сообщения об ошибке и после выход из фукнции
--- 
**var grades []int**
- объявление переменной для хранения оценок студента
<br> fmt.Println("Введите оценки студента через пробел. Нажмите Enter еще раз для завершения.")
- вывод инструкции для ввода оценок
---
**for {
<br> fmt.Print(">")
<br> input, _ := reader.ReadString('\n')
<br>input = strings.TrimSpace(input)**
- цикл для ввода оценок и вывод приглашения для ввода. далее чтение ввода пользователя и очистка лишних пробелов
---
**if input == "" {
<br>if len(grades) == 0 {
<br>fmt.Println("Оценки не были введены. Введите оценки или нажмите Enter еще раз для пропуска.")
<br>continue
<br>}
<br>break
<br>}**
- проверка на пустой ввод, если пустой и оценки еще не введены - выводит предупреждение. если оценки есть - выходит из цикла
---
**gradesStrSlice := strings.Fields(input)**
- разделение введенной строки на раздельные оценки
---
**for _, gradeStr := range gradesStrSlice {
<br>grade, err := strconv.Atoi(gradeStr)**
- цикл по всем введенным оценкам и преобразование строки в число
---
**if err != nil {
<br>fmt.Printf("Некорректный ввод оценки '%s'. <br>Пожалуйста, вводите только числа.\n", gradeStr)
<br>continue
	<br>		}**
- проверка ошибки преобразования, если ошибка - сообщение и пропускает некорректную оценку 
---
**if grade < 1 || grade > 5 {
<br>fmt.Printf("Оценка %d вне допустимого диапазона (1-5)\n", grade)
<br>continue
<br>}**
- проверка диапозона, а так же сообщение о неверном вводе, соотвественно ее пропуск
---
**grades = append(grades, grade)
<br>}
<br>}**
- добавленные корректной оценки в массив и закрытие циклов
---
**student := StudentInfo{
<br>FullName : FullName,
<br>Grades: grades,
<br>}**
- создание структуры студента с полями ФИО и оценки
---
**student.calculateAverage()**
- вычисление среднего балла для созданного студента
---
**students[FullName] = student
<br>fmt.Println("Студент", FullName, "успешно добавлен!")**
- сохранение студента в map и сообщение об этом
---
**func filterStudentsByAvg(students map[string]StudentInfo, threshold float64) []StudentInfo {**
- объявление функции фильтрации студентов по среднему баллу
---
**filteredStudents := []StudentInfo{}
<br>for _, student := range students {
<br>if student.AvgGrade < threshold {
<br>filteredStudents = append(filteredStudents, student)
<br>}
<br>}**
- цикл по всем студентам, а так же проверка среднего балла. далее добавление студента в результат, если условие выполняется
---
**return filteredStudents
<br>}**
- возвращение результатов фильтрации и закрытие цикла
---
**func printStudentInfo(student Student) {**
- объявление функции для вывода информации о студенте
---
**fmt.Printf(" ФИО: %s, Оценки: %v, Средний балл: %.2f\n", student.FullName, student.Grades, student.AvgGrade)
}**
- вывод информации о студенте и закрытие функции
---
**func printAllStudents(students map[string]StudentInfo) {**
- объявление фукнции вывода всех студентов
---
**fmt.Println("Список всех студентов:")**
- вывод заголовка списка студентов
---
**for _, student := range students {
<br>printStudentInfo(student)
<br>}
<br>}**
- цикл по всем студентам, а так же вывод информации о каждом студенте. далее закрытия цикла и функции
---
**func main()**
- объявление главной функции программы
---
**students := make(map[string]StudentInfo)
<br>reader := bufio.NewReader(os.Stdin)**
- создание map для хранения студентов и буферного читателя для ввода c клавиатуры
---
**fmt.Println("Добро пожаловать в журнал!")**
- вывод приветствия
---
**for {**
- бесконечный цикл для работы с командами 
---
**fmt.Print("\nВведите команду (help - для вывода всех команд): ")
<br>command, _ := reader.ReadString('\n')
<br>command = strings.TrimSpace(command)**
- вывод инструкции с подсказкой, следом чтение команды и ее очистка от пробелов
---
**switch command {**
- оператор выбора команд
---
**case "add":
<br>addStudent(students, reader)**
- обработка команды add для добавления студента и вызов соответствующей функции
---
**case "list":
<br>if len(students) > 0 {
<br>printAllStudents(students)
<br>} else {
<br>fmt.Println("В базе данных пока нет студентов.")**
- обработка команды list для вывода студентов. выводит студентов, если они есть, в противном случае выводит ошибку
---
**case "filter":
<br>fmt.Print("Введите максимальный средний балл: ")
<br>thresholdStr, _ := reader.ReadString('\n')
<br>thresholdStr = strings.TrimSpace(thresholdStr)**
- обработка команды filter, которая запрашивает пороговое значение, затем последующее его чтение и очистка ввода
---
**threshold, err := strconv.ParseFloat(thresholdStr, 64)
<br>if err != nil {
<br>fmt.Println("Некорректный ввод среднего балла. Пожалуйста, введите число.")
<br>continue
<br>}**
- преобразование строки в число и проверка ошибок преобразоания. если ошибка, то следует ее вывод. далее продолжение цикла
---
**filteredStudents := filterStudentsByAvg(students, threshold)**
- фильтрация студентов
---
**if len(filteredStudents) > 0 {
<br>fmt.Printf("Студенты со средним баллом ниже %.2f:\n", threshold)
<br>for _, student := range filteredStudents {
<br>printStudentInfo(student)
<br>}
<br>} else {
<br>fmt.Printf("Нет студентов со средним баллом ниже %.2f.\n", threshold)
<br>}**
- проверка наличия отфильтрованных студентов. вывод заголовка с порогом. далее следует цикл по студентамЮ а так же вывод информации о каждом студенте. в противном случае сообщение об их отсутствии
**case "help":
<br>fmt.Println("Доступные команды:")
<br>fmt.Println("  add    - Добавить нового студента.")
<br>fmt.Println("  list   - Вывести информацию о всех студентах.")
<br>fmt.Println("  filter - Отфильтровать студентов по среднему баллу (ниже заданного порога).")
<br>fmt.Println("  help   - Показать список доступных команд.")
<br>fmt.Println("  exit   - Выйти из программы.")**
- обработка команды help и вывод всех существующих команд
---
**case "exit":
<br>fmt.Println("Выход из программы. До свидания!")
<br>return**
- обработка команды exit, прощальное сообщение и выход из программы
---
**default:
<br>fmt.Println("Неизвестная команда. Введите 'help' для просмотра списка команд.")**
- если неизвестная команда, то выводится сообщение об ошибке
- далее конец блока switch, конец цикла и фукнции main


---


# Спасибо за внимание, поставьте 5 пожалуйста



