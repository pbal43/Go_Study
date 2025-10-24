package lessons

import (
	"fmt"
	"sync"
	"time"
)

func reader(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("reader is waiting")
	fmt.Println(time.Now())
	time.Sleep(time.Second * 5)
	fmt.Println(<-ch)
	fmt.Println("reader done")
	fmt.Println(time.Now())
}

func writer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- 443
	fmt.Println("writer done")
}

func generator() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch) // сигнал конца данных, аналогичный EOF (end of file) при чтении файла все текущие получатели (<-ch) могут дочитать оставшиеся значения;
		// новые попытки отправки (ch <- x) вызовут panic;
		// новые попытки чтения из уже пустого канала вернут нулевое значение типа (например, 0 для int) и флаг ok=false:
	}()
	return ch
}

func sum(ch chan int) int {
	sum := 0
	value, ok := <-ch // эта проверка есть под капотом for range, указано для демонстрации
	if !ok {          // если канал будет закрыт
		fmt.Println("OK: ", ok)
		return value // 0
	} else {
		fmt.Println("OK: ", ok)
	}
	for v := range ch { // если канал не закрыть → range зависнет навсегда, потому что он будет ждать ещё данные, которых уже не придёт.
		sum += v
	}
	return sum
}

func generatorForReading() <-chan int { // возвращаем канал, в который нельзя будет писать
	ch := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch) // сигнал конца данных, аналогичный EOF (end of file) при чтении файла все текущие получатели (<-ch) могут дочитать оставшиеся значения;
		// новые попытки отправки (ch <- x) вызовут panic;
		// новые попытки чтения из уже пустого канала вернут нулевое значение типа (например, 0 для int) и флаг ok=false:
	}()
	return ch
}

func sum1(ch <-chan int) int {
	sum := 0
	value, ok := <-ch // эта проверка есть под капотом for range, указано для демонстрации
	if !ok {          // если канал будет закрыт
		fmt.Println("OK: ", ok)
		return value // 0
	} else {
		fmt.Println("OK: ", ok)
	}
	for v := range ch { // если канал не закрыть → range зависнет навсегда, потому что он будет ждать ещё данные, которых уже не придёт.
		sum += v
	}
	return sum
}

func generatorNew(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		results <- job * 2
	}
}

func Lesson13() {
	//ch := make(chan int) // канал и тип данных, которые будут черен него передаваться, можно передавать любые данные
	// канал - ссылочный тип, можно передавать без указателя

	//ch <- 12    // запись числа в канал
	//res := <-ch // чтение из канала
	// если у канала нет потребителя - блокировка, пока кто-то что-то не прочтет и со чтением аналогично (ждем, пока появится пишущий)

	//fmt.Println(res)

	//ch := make(chan int)
	//wg := sync.WaitGroup{}
	//
	//wg.Add(1)
	//go reader(ch, &wg) // заблокируется на чтение до записи, поэтому даже если запуск раньше, то ждет записи в канал
	//
	//wg.Add(1)
	//go writer(ch, &wg)
	//
	//wg.Wait()

	// аналогично и в другую сторону, writer будет ждать, пока не будет reader'а

	// Буферизированные каналы - можно класть данные и без читателя (емкость > 0, просто каналы - емкость == 0)
	//ch := make(chan int, 5) // где 5 - емкость
	//wg := sync.WaitGroup{}
	//
	//wg.Add(1)
	//go reader(ch, &wg)
	//
	//wg.Add(1)
	//go writer(ch, &wg) // не заблокируется на запись до чтения, так как есть емкость
	//
	//wg.Wait()

	// при перегрузке буфера - блокируемся до вычитывания (освободится место)
	// если буфер не пустой - не заблокируемся на чтение
	//ch2 := make(chan int, 2)
	//ch2 <- 1
	//ch2 <- 2
	//go func() {
	//	fmt.Println("try to write to channel")
	//	ch2 <- 3
	//	fmt.Println("data is written to channel") // выведет после done, так как сначала необходимо вычитать (освобождение буфера)
	//}()
	//time.Sleep(time.Second * 2)
	//fmt.Println(<-ch2) // 1
	//fmt.Println(<-ch2) // 2

	//fmt.Println(<-ch2) // fatal error: all goroutines are asleep - deadlock! - читаем из пустого канала

	//fmt.Println("done")
	//time.Sleep(time.Second * 5)

	// запись - чтение - FiFo

	//g := generator()
	//fmt.Println(sum(g))

	// Направленные каналы - только на чтение и только на запись - для защиты
	//directedCh := generatorForReading()
	//directedCh <- 1 // Ошибка: Invalid operation: directedCh <- 1 (send to the receive-only type <-chan int)
	//fmt.Println(sum1(directedCh))

	//ch := make(chan int)
	//go generatorNew(ch)   // только пишем
	//fmt.Println(sum1(ch)) // только читаем

	// select - похоже на switch, но с каналами
	//ch1 := make(chan int)
	//ch2 := make(chan string)
	//ch3 := make(chan float32)
	//
	//go func() {
	//	for i := 0; i < 1000; i++ {
	//		time.Sleep(time.Second)
	//		ch1 <- i
	//	}
	//	close(ch1)
	//}()
	//
	//go func() {
	//	for i := 0; i < 1000; i++ {
	//		time.Sleep(time.Millisecond * 700)
	//		ch2 <- fmt.Sprintf("str %d", i)
	//	}
	//	close(ch2)
	//}()
	//
	//go func() {
	//	for i := 0; i < 1000; i++ {
	//		time.Sleep(time.Second * 2)
	//		res := float32(i) * 5.5
	//		ch3 <- res
	//	}
	//	close(ch3)
	//}()

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	//defer cancel()

	//stopCh := make(chan struct{})
	//
	//go func() {
	//	time.Sleep(time.Second * 5)
	//	close(stopCh) // ① Отправка (done <- struct{}{})  Отправляет одно значение в канал. Только одна горутина сможет это получить (если канал не буферизирован).
	//	// Остальные горутины останутся ждать, и не получат сигнал.
	//	// ② Закрытие (close(done)) Говорит Go: «канал больше не будет иметь новых значений». Все, кто делают <-done, сразу пробуждаются и получают нулевое значение типа.
	//	// Закрытие — широковещательный сигнал (broadcast), который видят все горутины.
	//}()

	//for {
	//	select { // выберет канал случайно, а не в порядке тут указанном
	//	//case <-ctx.Done(): // отработает, когда контекст Done
	//	//	fmt.Println("timeout")
	//	//	return
	//	case <-stopCh: // отработает, когда канал закроется
	//		fmt.Println("timeout")
	//		return
	//	case v := <-ch1:
	//		fmt.Println("ch1: ", v)
	//	case v := <-ch2:
	//		fmt.Println("ch2: ", v)
	//	case v := <-ch3:
	//		fmt.Println("ch3: ", v)
	//		//default: // позводит избежать блокировки, если нет ни 1 значения
	//		//	fmt.Println("nothing") // отработает, если нет ни 1 case, чтобы избежать блокировок
	//	}
	//}

	jobs := make(chan int, 5)
	result := make(chan int, 5)
	var wg sync.WaitGroup

	for w := 0; w < 10; w++ {
		wg.Add(1)
		go worker(w, jobs, result, &wg)
	}

	go func() {
		for i := 0; i < 100; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(result)
	}()

	for res := range result {
		println(res)
	}
}
