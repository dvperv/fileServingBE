package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	//router.Static("/assets", "./assets")
	//router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/more_static/1", "./more_static/1.txt")
	router.StaticFile("/more_static/1", "./more_static/2.txt")
	router.StaticFile("/more_static/1", "./more_static/3.txt")

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8085")
}
