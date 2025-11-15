package lessons

import "fmt"

//func dataGenerator(done chan struct{}) <-chan int {
//	out := make(chan int)
//	go func() {
//		for i := 0; i < 100; i++ {
//			out <- i
//			time.Sleep(time.Millisecond * 500)
//		}
//		close(out)
//		done <- struct{}{}
//	}()
//	return out
//}

//type StopTheWorld struct {
//	stop chan struct{}
//	wg   sync.WaitGroup
//}
//
//func NewStopTheWorld() *StopTheWorld {
//	return &StopTheWorld{
//		stop: make(chan struct{}),
//	}
//}
//
//func (s *StopTheWorld) StartTask(name string, delay time.Duration) {
//	s.wg.Add(1)
//	go func() {
//		defer s.wg.Done()
//		for {
//			select {
//			case <-s.stop:
//				fmt.Println(name, "stopped")
//				return
//			case <-time.After(delay):
//				fmt.Println(name, "running")
//			}
//		}
//	}()
//}
//
//func (s *StopTheWorld) StopAll() {
//	close(s.stop)
//	s.wg.Wait()
//}

//func divideByZero(a, b int, wg *sync.WaitGroup, resChan chan int, errChan chan error) {
//	defer wg.Done()
//	if b == 0 {
//		errChan <- fmt.Errorf("divide by zero")
//		return
//	}
//	resChan <- a / b
//}

//func divideByZero1(a, b int) (int, error) {
//	if b == 0 {
//		return -1, fmt.Errorf("Divide by zero")
//	}
//	return a / b, nil
//}

//func stage1(input <-chan int) <-chan int {
//	out := make(chan int)
//	go func() {
//		for n := range input {
//			out <- n * 2
//			time.Sleep(time.Second)
//		}
//		close(out)
//	}()
//	return out
//}
//
//func stage2(input <-chan int) <-chan int {
//	out := make(chan int)
//	go func() {
//		for n := range input {
//			out <- n * 1000
//			time.Sleep(time.Millisecond * 500)
//		}
//		close(out)
//	}()
//	return out
//}
//
//func stage3(input <-chan int) <-chan int {
//	out := make(chan int)
//	go func() {
//		for n := range input {
//			out <- n * 11
//			time.Sleep(time.Second * 2)
//		}
//		close(out)
//	}()
//	return out
//}

//type WorkResult struct {
//	producerID string
//	workName   string
//	status     string
//}
//
//func (w WorkResult) String() string {
//	return fmt.Sprintf("ProducerID:%s, Work name: %s, Status: %s", w.producerID, w.workName, w.status)
//}
//
//func funIn(inputs ...<-chan WorkResult) <-chan string {
//	out := make(chan string)
//	for _, input := range inputs {
//		go func(in <-chan WorkResult) {
//			for res := range in {
//				out <- res.String()
//			}
//		}(input)
//	}
//	return out
//}
//
//func producer(pID string, out chan<- WorkResult, delay time.Duration) {
//	for i := 0; i < 10; i++ {
//		wName := fmt.Sprintf("work-%d", i)
//		out <- WorkResult{producerID: pID, workName: wName, status: "done"}
//		time.Sleep(delay)
//	}
//	close(out)
//}
//
//func funOut(input <-chan int, numWorkers int) []<-chan int {
//	channels := make([]<-chan int, numWorkers)
//	for i := 0; i < numWorkers; i++ {
//		out := make(chan int)
//		channels[i] = out
//		go func(c int) {
//			for v := range input {
//				out <- v * c
//				time.Sleep(300 * time.Millisecond)
//			}
//			close(out)
//		}(i + 1)
//	}
//	return channels
//}

//type Semaphore struct {
//	ch chan struct{}
//}
//
//func NewSemaphore(maxConcurrency int) *Semaphore {
//	return &Semaphore{
//		make(chan struct{}, maxConcurrency),
//	}
//}
//
//func (s *Semaphore) Acquire() {
//	s.ch <- struct{}{}
//}
//
//func (s *Semaphore) Release() {
//	<-s.ch
//}
//
//func workerSem(id int, sem *Semaphore, wg *sync.WaitGroup) {
//	defer wg.Done()
//	log.Printf("worker %d try to Acquire", id)
//	sem.Acquire()
//	log.Printf("worker %d acquired", id)
//	time.Sleep(2 * time.Second)
//	sem.Release()
//	log.Printf("worker %d released", id)
//}

//type Singleton struct {
//	id string
//}
//
//var instance *Singleton
//
//var once sync.Once
//
//func GetInstance() *Singleton {
//	once.Do(func() { // код в анонимной функции выполнится только 1 раз за все время работы программы
//		log.Println("Creating new instance of Singleton")
//		instance = &Singleton{id: "123"}
//		log.Println("Created new instance of Singleton")
//	})
//	return instance // возврат указателя будет каждый раз при вызове функции
//}

//type House struct {
//	Owner   string
//	Doors   int
//	Windows int
//	Floors  int
//}
//
//func (h *House) String() string {
//	return fmt.Sprintf("Owner: %s, Doors: %d, Windows: %d, Floors: %d", h.Owner, h.Doors, h.Windows, h.Floors)
//}
//
//type HouseBuilder struct {
//	house *House
//}
//
//func NewHouseBuilder() *HouseBuilder {
//	return &HouseBuilder{
//		house: &House{},
//	}
//}
//
//func (builder *HouseBuilder) SetOwner(owner string) *HouseBuilder {
//	builder.house.Owner = owner
//	return builder
//}
//
//func (builder *HouseBuilder) SetDoors(doors int) *HouseBuilder {
//	builder.house.Doors = doors
//	return builder
//}
//
//func (builder *HouseBuilder) SetWindows(windows int) *HouseBuilder {
//	builder.house.Windows = windows
//	return builder
//}
//
//func (builder *HouseBuilder) SetFloors(floors int) *HouseBuilder {
//	builder.house.Floors = floors
//	return builder
//}
//
//func (builder *HouseBuilder) Build() *House {
//	return builder.house
//}

//type Server struct {
//	Host string
//	Port string
//}
//
//type Option func(*Server)
//
//func WithHost(host string) Option {
//	return func(s *Server) {
//		s.Host = host
//	}
//}
//
//func WithPort(port string) Option {
//	return func(s *Server) {
//		s.Port = port
//	}
//}
//
//func NewServer(opts ...Option) *Server {
//	s := &Server{ // с дефолтными значениями
//		Host: "localhost",
//		Port: "8080",
//	}
//	for _, opt := range opts {
//		opt(s)
//	}
//	return s
//}

//type Shape interface {
//	Draw()
//}
//
//type Circle struct {
//}
//
//func (circle Circle) Draw() {
//	fmt.Println("Circle Draw")
//}
//
//type Rectangle struct {
//}
//
//func (rectangle Rectangle) Draw() {
//	fmt.Println("Rectangle Draw")
//}
//
//func ShapeFactory(shapeType string) Shape {
//	switch shapeType {
//	case "circle":
//		return &Circle{}
//	case "rectangle":
//		return &Rectangle{}
//	default:
//		return &Circle{}
//	}
//}

type Notifier interface {
	Notify(string)
}

type EmailNotifier struct {
}

func (e *EmailNotifier) Notify(msg string) {
	fmt.Printf("Email notifier: %s\n", msg)
}

type NotifierDecorator struct {
	Notifier Notifier
}

func (n *NotifierDecorator) Notify(msg string) {
	n.Notifier.Notify(msg)
}

type SmsNotifier struct {
	NotifierDecorator
}

func NewSmsNotifier(notifier Notifier) *SmsNotifier {
	return &SmsNotifier{
		NotifierDecorator{
			Notifier: notifier,
		},
	}
}

func (s *SmsNotifier) Notify(msg string) {
	s.Notifier.Notify(msg)
	fmt.Printf("SMS notified: %s\n", msg)
}

func Lesson16() {

	// Паттерны многопоточности

	// Генератор
	//done := make(chan struct{})
	//out := dataGenerator(done)
	//for v := range out {
	//	fmt.Println(v)
	//}
	//<-done

	// Стоп-кран (stop the world)
	//stw := NewStopTheWorld()
	//stw.StartTask("A", time.Second)
	//stw.StartTask("B", time.Second)
	//stw.StartTask("C", time.Second)
	//
	//time.Sleep(10 * time.Second)
	//fmt.Println("Stopping...")
	//stw.StopAll()

	// Обработка ошибок

	// через WG
	//wg := &sync.WaitGroup{}
	//resChan := make(chan int)
	//errChan := make(chan error)
	//
	//wg.Add(1)
	//go divideByZero(1, 0, wg, resChan, errChan)
	//
	//select {
	//case res := <-resChan:
	//	fmt.Println(res)
	//case err := <-errChan:
	//	fmt.Println(err)
	//}
	//wg.Wait()

	// через ErrGroup
	//errGroup, _ := errgroup.WithContext(context.Background())
	//
	//var result int
	//var err error
	//
	//errGroup.Go(func() error {
	//	result, err = divideByZero1(1, 0)
	//	return err
	//})
	//
	//err = errGroup.Wait() // Если хотя бы 1 горутина вернула ошибку - вернет ошибку, иначе - nil
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(result)

	// Конвеер (Pipeline)
	//input := make(chan int)
	//
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		input <- i
	//	}
	//	close(input)
	//}()
	//
	//result := stage3(stage2(stage1(input)))
	//
	//fmt.Println("Done")
	//
	//for n := range result {
	//	fmt.Println(n)
	//}

	// FunIn и FunOut
	// FunIn
	//ch1 := make(chan WorkResult)
	//ch2 := make(chan WorkResult)
	//ch3 := make(chan WorkResult)
	//ch4 := make(chan WorkResult)
	//ch5 := make(chan WorkResult)
	//ch6 := make(chan WorkResult)
	//
	//go producer("1", ch1, 100*time.Millisecond)
	//go producer("2", ch2, 200*time.Millisecond)
	//go producer("3", ch3, 300*time.Millisecond)
	//go producer("4", ch4, 400*time.Millisecond)
	//go producer("5", ch5, 500*time.Millisecond)
	//go producer("6", ch6, 600*time.Millisecond)
	//
	//result := funIn(ch1, ch2, ch3, ch4, ch5, ch6)
	//for i := 0; i < 60; i++ {
	//	fmt.Printf("result:\n\t%s\n", <-result)
	//}

	//	FunOut
	//input := make(chan int)
	//go func() {
	//	for i := 0; i < 5000; i++ {
	//		input <- i
	//	}
	//	close(input)
	//}()
	//
	//start := time.Now()
	//
	//workersChannels := funOut(input, 100)
	//
	//var wg sync.WaitGroup
	//for i, ch := range workersChannels {
	//	wg.Add(1)
	//	go func(id int, ch <-chan int) {
	//		defer wg.Done()
	//		for v := range ch {
	//			fmt.Printf("Worker: %d: %d\n", id, v)
	//		}
	//	}(i+1, ch)
	//}
	//wg.Wait()
	//elapsed := time.Since(start)
	//fmt.Printf("Done, time: %s", elapsed)

	// Семафор
	// Позволяет установить кол-во горутин, одновременно работающих с ресурсом
	//sem := NewSemaphore(2)
	//wg := sync.WaitGroup{}
	//start := time.Now()
	//for i := 0; i < 5; i++ {
	//	wg.Add(1)
	//	go workerSem(i, sem, &wg)
	//}
	//wg.Wait()
	//finish := time.Since(start)
	//log.Printf("All workers done in %s", finish)

	// Паттерны программирования
	// Порождающие паттерны
	// Синглтон
	//fmt.Printf("Instance is nill: %v\n", instance == nil) // true
	//instance = GetInstance()
	//fmt.Printf("Instance is nill: %v\n", instance == nil) // false
	//instance = GetInstance()                              // не вызовется создание экземплеяра еще раз, просто возврат указателя (можно глянуть по логам)
	//fmt.Printf("Instance is nill: %v\n", instance == nil) // false

	// Builder
	// Разделяет создание объекта на отдельные этапы
	//builder := NewHouseBuilder()
	//house := builder.SetDoors(5).SetFloors(3).SetWindows(15).SetOwner("Pasha").Build()
	//fmt.Println(house.String())

	// Functional options pattern
	// Использует функциональные опции для настройки объекта
	//srv := NewServer(
	//	WithHost("127.0.0.1"),
	//	WithPort("80"),
	//)
	//fmt.Println(srv.Host, srv.Port)

	// Factory
	// Делегирует создание объектов подклассам и функциям
	//shape1 := ShapeFactory("circle")
	//shape1.Draw()
	//shape2 := ShapeFactory("rectangle")
	//shape2.Draw()
	//shape3 := ShapeFactory("lil")
	//shape3.Draw()

	// Decorator
	// Динамически добавляет новую функциональность объекту
	email := &EmailNotifier{}
	email.Notify("Only Email")
	sms := NewSmsNotifier(email)
	sms.Notify("Hello World")
}
