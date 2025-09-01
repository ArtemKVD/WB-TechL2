package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	// создаём новый  канал
	go func() {
		for {
			select { // случайным образом выбираем, из какого канала мы передадим элемент в новый канал
			case v, ok := <-a:
				if ok {
					c <- v
				} else {
					a = nil
				}
			case v, ok := <-b:
				if ok {
					c <- v
				} else {
					b = nil
				}
			}
			if a == nil && b == nil {
				close(c)
				// когда элементы в обоих каналах закончились, закрываем новый канал
				return
			}
		}
	}()
	return c
}

func main() {
	rand.Seed(time.Now().Unix())
	a := asChan(1, 3, 5, 7) // Создаём канал a со значениями 1 3 5 7
	b := asChan(2, 4, 6, 8) // Создаём канал b со значениями 2 4 6 8
	c := merge(a, b)        // Объединяем эти каналы
	for v := range c {
		//выводим значения из объединения двух каналов
		fmt.Print(v)
		//в качестве результата получим набор из 8 элементов расположенных в случайном порядке
	}
}
