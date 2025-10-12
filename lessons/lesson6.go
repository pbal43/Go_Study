package lessons

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func Lesson6() {
	password := "123456789qwerty"

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hash))

	password = "123456789qwerty1"

	err = bcrypt.CompareHashAndPassword(hash, []byte(password))

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Pass is correct")
	}

}
