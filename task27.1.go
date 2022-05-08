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

var studentMap = map[string]Student{} //глобальная мапа(походу сразу инициализированная:))

func newStudent(text string) { // функция заполняет поля структуры студент
	s := Student{}
	arr := strings.Split(text, " ") //строку делим на срез по пробелу
	s.Name = arr[0]
	s.Age, _ = strconv.Atoi(arr[1])
	s.Grade, _ = strconv.Atoi(arr[2])
	s.putMap() //структура уходит в мапу через метод put
}

func (s Student) putMap() { // метод добавление в мапу
	studentMap[s.Name] = s
}

func (s Student) getMap() { // метод считывание с мапы
	fmt.Println(s.Name, s.Age, s.Grade)
}

func main() {
	fmt.Println("27.7 Практическая работа")
	fmt.Println("")
	fmt.Println("Задание 1. Научиться работать с композитными типами данных: структурами и картами.")
	fmt.Println("------------")

	var input = bufio.NewScanner(os.Stdin)

	fmt.Println("Введите данные студентов:")
	for input.Scan() { //Цикл который прерывается при EOF
		studentData := input.Text() //Получаем бесконечные строки
		newStudent(studentData)     //Строки кидаем в функцию newStudent, вывод который летит в метод добавленияв мапу

	}
	if err := input.Err(); err != nil {
		return

	}

	fmt.Println()
	fmt.Println("Студенты из хранилища:")
	fmt.Println("------------")
	for _, v := range studentMap {
		v.getMap()
	}
}
