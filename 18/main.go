package main

/*Реализовать структуру-счетчик, которая будет
инкрементироваться в конкурентной среде.
 По завершению программа должна выводить
 итоговое значение счетчика */

import (
	"fmt"
	"sync"
)

func main() {
	counter := NewCounter(0)
	var wg sync.WaitGroup //создаем группу

	for i := 0; i < 2048; i++ {
		wg.Add(1)                 //каждую итерацию пишем + 1 в счетчик группы
		go counter.Increment(&wg) // запускаем 2048 горутин
	}

	wg.Wait()                   //ждет пока значение в ваит груп не будет равно 0
	fmt.Println(counter.Read()) //просто выводим
}

type Counter struct {
	sync.RWMutex //мьютекс для того чтобы можно было лочить наш коунтер, для чего rw написан снизу в Read
	count        int
}

func NewCounter(count int) *Counter { //ну вообще по хорошему прятать от всех реализацию + тут может быть логика какая то
	return &Counter{count: count}
}

func (c *Counter) Increment(wg *sync.WaitGroup) {
	c.Lock()         //лочим через мьютекс переменную
	defer c.Unlock() //разлочим когда закончится выполнение функции
	c.count += 1     //инекрементируем значение
	wg.Done()        //уменьшаем значение в группе на 1
}

func (c *Counter) Read() int {
	c.RLock()         //лочим переменеую для записи но не для чтения
	defer c.RUnlock() //разлочим когда функия выполнится
	return c.count
}
