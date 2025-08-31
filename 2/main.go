package main

import "fmt"

func test() (x int) {
	defer func() {
		x++ // 2. увеличиваем это значение на 1
	}()
	x = 1  // 1. Устанавлием значение 1 для переданной переменной
	return // 3. Возвращается значение x (2)
}

func anotherTest() int {
	var x int //1.Инициализируем x
	defer func() {
		x++ //3. Увеличиваем на 1, но то значение, которое было инициализировано (0)
	}()
	x = 1    //2. Присваиваем 1
	return x //Возвращаем x (1)
}

func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
