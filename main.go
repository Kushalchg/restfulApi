package main

import (
	"fmt"
	"log"
	"os"
	"practice/restfulApi/global"
	"practice/restfulApi/initializers"
	"practice/restfulApi/routes"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func init() {

	global.Validate = validator.New(validator.WithRequiredStructEnabled())
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func loggerInitializer(File *os.File) *log.Logger {
	infoLog := log.New(File, "Info:", log.Ltime|log.Lshortfile)
	return infoLog

}

func main() {

	fileName := fmt.Sprintf("logfiles/log_%s.txt", time.Now().Format("2006-01-02"))
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0700)

	if err != nil {
		fmt.Printf("error occurred while opening file: %v", err)
		os.Exit(1)
	}
	defer file.Close()

	global.Logger = loggerInitializer(file)
	// test.Test()
	r := gin.Default()
	// Registering routers
	routes.PostRoutes(r)
	routes.UserRoutes(r)

	r.Run()
}
