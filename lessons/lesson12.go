package lessons

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//func workerPrinter(wg *sync.WaitGroup) {
//	defer wg.Done() // указать на то, что горутинка выполнилась (счетчик -= 1)
//	for i := range 15 {
//		fmt.Printf("print %d\n", i)
//		time.Sleep(1 * time.Second)
//	}
//}
//
//func workerScaner(wg *sync.WaitGroup) {
//	defer wg.Done()
//	for i := range 15 {
//		fmt.Printf("scan %d\n", i)
//		time.Sleep(1 * time.Second)
//	}
//}

var counter int32
var mutex sync.Mutex
var rwMmutex sync.RWMutex

func Lesson12() {
	// вейт группа - 1 из способов синхронизации горутин
	//wg := sync.WaitGroup{}
	//start := time.Now()
	//wg.Add(2) // указание на запуск 2 горутин, будет ожидать их выполнения (счетчик = 2)
	//go workerPrinter(&wg)
	//go workerScaner(&wg)
	//wg.Wait() // делаем ожидание завершения горутин (обе должны быть done) (счетчик = 0)
	//fmt.Println("time: ", time.Since(start))

	//wg := sync.WaitGroup{}
	//
	//wg.Add(1) // пишем перед каждой горутиной, а не сразу на все
	//go func() {
	//	defer wg.Done()
	//	for i := 0; i < 1000; i++ {
	//		counter = counter + 2
	//	}
	//}()
	//
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	for i := 0; i < 1000; i++ {
	//		counter--
	//	}
	//}()
	//wg.Wait()
	//fmt.Println("Counter:", counter)
	// будет непонятное число, так как берется какое-то состояние каунтера, а он уже мог измениться (гонка данных)
	// необходимо синхронизировать потоки (mutex, атомарные операции; семафоры - их нет в го)
	// mutex позволяет работать с общим ресурсом только 1 потоку одновременно

	//wg := sync.WaitGroup{}
	//
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	for i := 0; i < 1000; i++ {
	//		mutex.Lock() // блокировка
	//		counter = counter + 2
	//		mutex.Unlock()
	//	}
	//}()
	//
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	for i := 0; i < 1000; i++ {
	//		mutex.Lock() // блокировка, но не сработает, если уже будет залочено 1-ой горутиной
	//		counter--
	//		mutex.Unlock()
	//	}
	//}()
	//wg.Wait()
	//fmt.Println("Counter:", counter)

	// rwMutex - 1 пишет (блок на запись), а остальные могут читать

	// атомарная операция - либо выполняется, либо нет (как транзакция в БД)
	// даёт низкоуровневые примитивы для работы с общими переменными в многопоточном (concurrent) коде, когда несколько горутин читают и пишут один и тот же ресурс.
	// Без синхронизации такие операции небезопасны: процессор может "разорвать" запись или выполнить их неатомарно,
	// а горутины могут видеть "грязное" состояние. sync/atomic решает это, гарантируя атомарность (операция выполняется полностью или не выполняется вообще).
	// То есть два потока физически не могут "перебить" друг друга посередине операции.
	// только простые операции и с простыми типами данных

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			atomic.AddInt32(&counter, 2)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			atomic.AddInt32(&counter, -1)
		}
	}()
	wg.Wait()
	fmt.Println("Counter:", counter)
}
