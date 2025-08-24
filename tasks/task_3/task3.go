package task_3

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

type Task struct {
	TaskText string `json:"task"`
}

var tasks []string

func getTasks() []string {
	return tasks
}

func getAllTasks(w http.ResponseWriter, r *http.Request) {
	allTasks := getTasks()
	w.Header().Set("Content-Type", "text/plain")
	if len(allTasks) != 0 {
		for _, line := range tasks {
			if _, err := w.Write([]byte(line + "\n")); err != nil {
				http.Error(w, "Error writing response", http.StatusInternalServerError)
				return
			}
		}
	} else {
		fmt.Fprintf(w, "TaskForToDo list is empty!")
	}
}

func addTask(task string) {
	tasks = append(tasks, task)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if task.TaskText == "" {
		http.Error(w, "Missing required field 'task'", http.StatusBadRequest)
		return
	}
	addTask(task.TaskText)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

// chi
// habr: Пакет go-chi/chi поддерживает маршрутизацию на основе методов, переменные в URL-путях и шаблоны маршрутов на основе регулярных выражений.
// Как и httprouter, он также позволяет вам устанавливать собственные обработчики ответов 404 и 405.
// Можно создавать “группы” маршрутов, которые используют определенное middleware. Это очень полезно в больших приложениях.

func Task3Chi() {
	router := chi.NewRouter()
	router.Use(middleware.Logger) // логирование информации о каждом HTTP-запросе
	router.Route("/tasks", func(r chi.Router) {
		// обрабатываем get
		r.Get("/", getAllTasks)
		// обрабатываем post
		r.Post("/", createTask)
	})
	http.ListenAndServe(":8080", router)
}

// mux - выбрал по совету коллеги BE ("самый продакшн вариант")
// google: хорош в Go благодаря своей гибкости и расширенным функциям маршрутизации,
// которые выходят за рамки стандартной библиотеки Go Packages.
// Он поддерживает динамические URL-адреса, маршруты на основе регулярных выражений, маршрутизацию на основе хоста и методы запросов,
// что делает его мощным инструментом для создания сложных веб-приложений и API, недоступных в стандартном ServeMux

func Task3Mux() {
	router := mux.NewRouter()
	tasksRouter := router.PathPrefix("/tasks").Subrouter()
	// обрабатываем get
	tasksRouter.HandleFunc("", getAllTasks).Methods(http.MethodGet)
	// обрабатываем post
	tasksRouter.HandleFunc("", createTask).Methods(http.MethodPost)
	http.ListenAndServe(":8080", router)
}
