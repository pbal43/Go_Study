package lessons

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	validator2 "github.com/go-playground/validator/v10"
	"io"
	"log"
	"math/rand"
	"net/http"
)

// enum
//const (
//	One   = iota // iota - счетчик строк. 0 - константный инт
//	Two          // 1
//	Three        // 2
//)

const (
	One = (iota + 1) * 2
	Two
	Three
)

type Status int

const (
	ServerRunning = iota
	ServerStopped
)

var statusNames = []string{
	"Running",
	"Stopped",
}

func (s Status) String() string {
	return statusNames[s]
}

func Lesson5() {
	// Константы - неизменяемы
	//fmt.Println("Const")
	//const intik = 14
	//intik += 1 // не даст

	// enum
	fmt.Println("Enum")
	fmt.Println(One, Two, Three)

	status := Status(ServerRunning)
	fmt.Printf("status id = %d, name = %s", status, status.String())

	// Работа с JSON
	// Сериализация - перевод структуры данных в байты
	fmt.Println("Работа с JSON")

	//fmt.Println("Используем Мараш/Анмаршал") // но подустарело, лучше не юзать
	//http.HandleFunc("/users-unmarshal", addUser)
	//http.HandleFunc("/users-marshal", getUsers)
	//http.HandleFunc("/users-marshal-random", getRandomUser)
	//
	//fmt.Println("Используем Decoder/Encoder")
	//http.HandleFunc("/users-add-decoder", addUsersDecoded)
	//http.HandleFunc("/users-get-encoder", getUsersEncoded)
	//
	//log.Println("Starting server")
	//log.Fatal(http.ListenAndServe(":8080", nil))

	// Делаем на Gin
	fmt.Println("Делаем на Gin")
	router := gin.Default()
	validator := validator2.New() // добавляем валидатр и помечаем поля в структуре
	router.GET("/get-users-gin", func(c *gin.Context) {
		c.JSON(200, users)
	})
	router.POST("/add-user-gin", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindBodyWithJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//if len(user.Password) < 8 { // валидация вручную
		//	c.JSON(http.StatusBadRequest, gin.H{"error": "password must be at least 8 characters"})
		//	return
		//}
		if err := validator.Struct(user); err != nil { // валидируем структуру валидатором
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		users = append(users, user)
	})
	router.Run(":8080")
}

type User struct {
	ID       int    `json:"id,omitempty" validate:"required"` // при omitempty не попадёт в JSON, если оно имеет нулевое значение для своего типа (см. getRandomUser)
	UserName string `json:"name" validate:"required"`         // установка соответствия поля структуры к полю json
	Email    string `json:"email" validate:"required,email"`  // проверка на то, что значение в поле - email
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
	Age      int    `json:"age" validate:"required,gte=18"` // делаем поля обязательными, при отсутствии поля - 400 ошибка (но палит поля структуры)
}

var users = []User{
	//{1, "John", "JohnDoe@gmail.com"},
	//{1, "Alice", "AlicePae@yahoo.com"},
}

func addUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	var user User
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println(string(body))
	log.Println(user)
	users = append(users, user)
	w.WriteHeader(http.StatusCreated)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	jsonData, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func getRandomUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	randomIndex := rand.Intn(len(users))

	user := users[randomIndex]
	user.Email = ""

	jsonData, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func addUsersDecoded(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var user User
	body := r.Body
	defer body.Close()
	if err := json.NewDecoder(body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err != nil {

		}
	}
	users = append(users, user)
}

func getUsersEncoded(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
