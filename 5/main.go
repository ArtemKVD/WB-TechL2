package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	// ... do something
	return nil
}

func main() {
	var err error
	err = test()
	//error это интерфейс, чтобы интерфейс был равен nil, необходимо чтобы оба его поля были равны nil
	//У интерфейса 2 поля (тип и значение), в нашел случае тип = customError, а значение = nil
	if err != nil {
		//err != nil, так как тип != nil, вывод: error
		println("error")
		return
	}
	println("ok")
}
