package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)        // <nil>
	fmt.Println(err == nil) // false
}

//Значение ошибки = nil, так как мы присвоили значение nil для ошибки в нашей функции
//Но тип интерфейса не равен nil, тип интерфейса = *os.PathError
//Чтобы интерфейс был равен nil, нужно чтобы и тип, и значение были равны nil (
