package lessons

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Lesson4() {
	// исп. пакет netHTTP
	//fmt.Println("Пакет netHTTP")
	// стучимся по http://127.0.0.1:8080/hello
	//http.HandleFunc("/hello", helloHandler) // аргументы в helloHandler подставятся, как только придет запрос по маршруту
	//http.ListenAndServe(":8080", nil)       //  указываем nil, так как до этого создали хэндлер глобально
	//http.ListenAndServeTLS() - для https, но нужны подписанные сертификаты

	// фреймворк GIN
	//fmt.Println("Пакет gin")
	//router := gin.Default()
	//router.GET("/hello", newHelloHandler)
	//router.POST("/hello", func(c *gin.Context) {
	//	c.String(http.StatusOK, "Hello Post!")
	//})
	//router.GET("/view", viewHelloHandler)

	// группировки хендлеров
	//fmt.Println("Группировка хендлеров")
	//hello := router.Group("/hello") // при Group прикол - после пишем фигурные скобки, так принято, для читабельности, типа замыкание
	//{
	//	hello.GET("/en", func(c *gin.Context) { // en - дочернее от hello (/hello/en)
	//		c.String(http.StatusOK, "Hello")
	//	})
	//	hello.GET("/ru", func(c *gin.Context) {
	//		c.String(http.StatusOK, "Привет")
	//	})
	//}
	//router.GET("/user/:name", func(c *gin.Context) { // создаем запрос с квери параметром, все что вобьет юзер - name
	//	name := c.Param("name") // создали переменную, подвязав ее к квери-параметру
	//	c.String(http.StatusOK, "Hello %s", name)
	//})
	//router.GET("/nekto", func(c *gin.Context) { // считываем только если совпали параметры
	//	firstName := c.Query("firstname")                                        // совпало, если юзер вбил /nekto?firstname=
	//	lastName := c.Query("lastname")                                          //  /nekto?firstname=&lastname=
	//	newParam := c.DefaultQuery("newParam", "parapam")                        // установка дефолтного параметра
	//	c.String(http.StatusOK, "Hello %s %s %s", firstName, lastName, newParam) // если запрос будет без параметров, то выведется просто "Hello"
	//	// как сделать так, чтобы не играл ролял регистр?
	//})
	//router.Run(":8080")

	//uuid := uuid.New().String() // пример создания uuid пользователя
	//fmt.Print(uuid)
	// фреймворк Chi

	// самостоятельная
	router := gin.Default()
	fmt.Println("Самостоятельная")
	toDoList := router.Group("/todolist")
	{
		toDoList.GET("/tasks", func(c *gin.Context) { // получение списка всех тасок
			c.JSON(http.StatusOK, tasks)
		})
		toDoList.POST("/tasks", func(c *gin.Context) { // запись таски
			var req TaskRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			writeTask(req.Task, &id)
			c.JSON(http.StatusOK, gin.H{
				"message": "task created",
				"id":      id,
			})
		})
		toDoList.PUT("tasks/:id", func(c *gin.Context) { // изменить таску по id
			taskIDStr := c.Param("id")                         // записали id
			taskID, err := strconv.ParseInt(taskIDStr, 10, 64) // преобразуем id в int
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task id"})
				return
			}
			taskIDNewInt := int(taskID)
			var req TaskRequest
			if err := c.ShouldBindJSON(&req); err != nil { // получили новый текст
				c.JSON(http.StatusBadRequest, gin.H{"error": "неверно указана задача"})
				return
			}
			foundedTask := FindTaskByID(tasks, taskIDNewInt)
			if foundedTask != nil {
				oldTask := foundedTask.ToDo
				foundedTask.ToDo = req.Task
				c.JSON(http.StatusOK, gin.H{
					"message":  "task updated",
					"old task": oldTask,
					"new task": req.Task,
					"id":       taskIDNewInt,
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "task not found"})
			}
		})
		toDoList.GET("tasks/:id", func(c *gin.Context) {
			taskIDStr := c.Param("id")
			taskID, err := strconv.ParseInt(taskIDStr, 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task id"})
				return
			}
			taskIDNewInt := int(taskID)
			foundedTask := FindTaskByID(tasks, taskIDNewInt)
			if foundedTask != nil {
				c.JSON(http.StatusOK, gin.H{
					"message": foundedTask,
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "task not found"})
			}

		})
	}
	router.Run(":8080")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// r - запрос, который к нам пришел
	// w - для того, чтобы возвращать ответ
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Hello World!") // ожидает на вход объект, реализующий интерфейс io writer
}

func newHelloHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello World!")
}

func viewHelloHandler(ctx *gin.Context) {
	helloes := []string{"hello", "wassap", "cockorel"}
	ctx.JSON(http.StatusOK, helloes)
}

// для самостоятельной

type Task struct {
	ID   int    `json:"id"`
	ToDo string `json:"name"`
}

type TaskRequest struct {
	Task string `json:"task"`
}

// мини-BD
var tasks []Task // список всех тасков
var id int

func writeTask(todo string, id *int) {
	*id += 1
	newTask := Task{
		ID:   *id,
		ToDo: todo,
	}
	tasks = append(tasks, newTask)
}

func FindTaskByID(tasks []Task, searchID int) *Task {
	var foundTask *Task
	for i, _ := range tasks {
		if tasks[i].ID == searchID {
			foundTask = &tasks[i]
			break
		}
	}
	return foundTask
}
