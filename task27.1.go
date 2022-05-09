package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name  string
	Age   int
	Grade int
}

type ptu struct {
	studentMap map[string]Student
}

func newStudent(text string) (s Student) { // функция заполняет поля структуры студент
	// s = Student{}
	arr := strings.Split(text, " ") //строку делем на срез по пробелу
	s.Name = arr[0]
	s.Age, _ = strconv.Atoi(arr[1])
	s.Grade, _ = strconv.Atoi(arr[2])
	return
}

func (p ptu) putMap(s Student) { // метод добавление в мапу
	p.studentMap[s.Name] = s
}

func (p ptu) getMap(k string) { // метод считывание с мапы
	fmt.Println(p.studentMap[k].Name, p.studentMap[k].Age, p.studentMap[k].Grade)
}

func main() {
	fmt.Println("27.7 Практическая работа")
	fmt.Println("")
	fmt.Println("Задание 1. Научиться работать с композитными типами данных: структурами и картами.")
	fmt.Println("------------")
	m := ptu{}
	m.studentMap = make(map[string]Student)
	var input = bufio.NewScanner(os.Stdin)

	fmt.Println("Введите данные студентов:")
	for input.Scan() { //Цикл который прерывается при EOF
		studentData := input.Text()       //Получаем бесконечные строки
		m.putMap(newStudent(studentData)) //Строки кидаем в фукцию newStudent, вывод который летит в метод добавленияв мапу
	}
	if err := input.Err(); err != nil {
		return
	}

	fmt.Println()
	fmt.Println("Студенты из хранилища:")
	fmt.Println("------------")
	for k, _ := range m.studentMap {
		m.getMap(k)
	}
}
