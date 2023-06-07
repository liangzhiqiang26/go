package main

import (
	"github.com/gin-gonic/gin"
	"go-parent/database"
	"go-parent/handlers"
	"go-parent/repositories"
	"go-parent/services"
)

func main() {
	// 初始化数据库连接
	err := database.InitDB()
	if err != nil {
		panic(err)
	}

	// 创建仓库和服务
	studentRepo := repositories.NewStudentRepository(database.DB)
	studentService := services.NewStudentService(studentRepo)

	// 创建处理函数和路由
	studentHandler := handlers.NewStudentHandler(studentService)

	router := gin.Default()
	router.GET("/students", studentHandler.GetStudents)

	// 启动服务器
	router.Run(":8080")
}
