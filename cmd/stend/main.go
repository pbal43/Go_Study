package stend

import (
	"fmt"
	storage "go_study/internal/repository"
	service "go_study/internal/service"
)

func Main() {
	storageData := storage.NewStorage()
	serviceData := service.NewService("stend", *storageData)

	serviceData.SaveData("Lollipop")

	fmt.Println(service.Test)

	//fmt.Print(service.neTest) - не достанет, так как с маленькой буквы. Можно исп. только в том пакете, где объявлена
	service.PrintNeTest() // тут достанет, так как через доступный метод (с большой буквы)
}
