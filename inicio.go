package main

import (
	"ejemplo/src/utils"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func inicio() {
	fmt.Println("Hola mundo")

	entorno := godotenv.Load()

	if entorno != nil {
		fmt.Println(entorno)
	}

	fmt.Println(os.Getenv("nombre"))
	fmt.Println(2 + 3)

	token := utils.TokenGenerator()
	fmt.Println(token)

	md5Encrypt := utils.Md5("Hola")
	fmt.Println(md5Encrypt)

}
