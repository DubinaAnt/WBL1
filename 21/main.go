package main

import "fmt"

//представим что у нас есть некий клиент который хочет внести деньги на счет чтобы что нибудь приобрести
//если товар из россии то нужно оплатить рублями
//если товар из-за границы то нужно оплатить долларами
//допустим что у клиента может быть два баланса: рублевый и долларовый

type client struct { //клиент который хочет пополнить счет в рублях
}

func (c *client) refill(rs rubleRefill, rub float64) { //метод который инициализирует пополнение счета
	fmt.Println("Клиент пополняет счет")
	rs.refill(rub) //тут мы вызываем именно метод у нашего сервиса который занимается пополнением
}

type rubleRefill interface { //интерфейс который имплементит наш клиент
	refill(rub float64)
}

type rubleService struct { //как раз таки рублевый сервис, именно он пополняет рублевый счет
}

func (m *rubleService) refill(rub float64) { //и rubleService тоже иплементит интерфейс rubleRefill
	fmt.Printf("Счет пополнен на %f рублей.", rub) //ну и он пополняет рублевый счет
}

type dollarService struct{} //а это сторонний сервис который пополняет долларовый счет

func (w *dollarService) refillDollar(dollars float64) {
	fmt.Printf("Счет пополнен на %f долларов.", dollars)
}

type dollarsAdapter struct { //чтобы пользователь мог пополнить долларовый счет рублями, мы создаем некий адаптер
	dollarService *dollarService //который хранит в себе структуру сервиса доларового
}

func (w *dollarsAdapter) refill(rub float64) { //но наш адаптер тоже имплементит интерфейс rubleRefill
	dollars := rub / 36                   // :(
	w.dollarService.refillDollar(dollars) //мы внутри этого метода адаптера вызываем метод доларового сервиса, но уже
	//передаем конвертированные рубли в доллары
	fmt.Printf("Рубли конвертированы в доллары и внесены на счет.")
}

func main() {

	newClient := &client{} //просто создаем клиента

	newRubleService := &rubleService{}     //создаем рублевый сервис
	newClient.refill(newRubleService, 150) // если клиент хочет внести деньги на рублевый счет то вызываем этот метод

	newDollarService := &dollarService{}  //создаем доларовый сервис
	newDollarsAdapter := &dollarsAdapter{ //создаем адаптер
		dollarService: newDollarService,
	}

	newClient.refill(newDollarsAdapter, 150) //если клиент хочет внести деньги на долларовый счет,
	// то вызываем этот метод
}
