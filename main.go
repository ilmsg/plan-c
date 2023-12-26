package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hi.")

	app := gin.Default()
	app.Run(":3040")
}
