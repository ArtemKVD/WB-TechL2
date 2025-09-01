package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	for n := range ch {
		println(n)
	}
}

//Вывод:
// 0
// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
// fatal error: all goroutines are asleep - deadlock!
//Выводятся числа от 0 до 9, которые были записаны в канал, после чего ошибка: fatal error: all goroutines are asleep - deadlock!
//Ошибка выводится из-за того, что главная горутина пытается читать из канала, но в него уже ничего не записывается, при этом канал не закрыт
