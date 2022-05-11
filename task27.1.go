package main

import (
	"bufio"
	"errors"
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
	arr := strings.Split(text, " ") //строку делим на срез по пробелу
	s.Name = arr[0]
	s.Age, _ = strconv.Atoi(arr[1])
	s.Grade, _ = strconv.Atoi(arr[2])
	return
}

func (p ptu) putMap(s Student) { // метод добавление в мапу
	_, ok := p.studentMap[s.Name]
	if ok {
		fmt.Println(errors.New("Имя уже существует"))
		return
	} else {
		p.studentMap[s.Name] = s
	}
}

func (p ptu) getMap(k string) Student { // метод считывание с мапы
	return p.studentMap[k]
}

func main() {
	fmt.Println("27.7 Практическая работа")
	fmt.Println("")
	fmt.Println("Задание 1. Научиться работать с композитными типами данных: структурами и картами.")
	fmt.Println("------------")

	m := ptu{}
	m.studentMap = make(map[string]Student)
	sliceName := make([]string, 0)
	var input = bufio.NewScanner(os.Stdin)

	fmt.Println("Введите данные студентов:")
	for input.Scan() { //Цикл который прерывается при EOF
		studentData := input.Text()                                   //Получаем бесконечные строки
		m.putMap(newStudent(studentData))                             //Строки кидаем в фукцию newStudent, вывод, который летит в метод добавления в мапу
		sliceName = append(sliceName, strings.Fields(studentData)[0]) // Добавляем имя в разрез
	}
	if err := input.Err(); err != nil {
		return
	}

	fmt.Println()
	fmt.Println("Студенты из хранилища:")
	fmt.Println("------------")
	for _, k := range sliceName {
		fmt.Println(m.getMap(k).Name, m.getMap(k).Age, m.getMap(k).Grade)
	}
}
