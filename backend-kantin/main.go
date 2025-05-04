package main

import (
	"backend_golang/controllers"
	"backend_golang/middlewares"
	"backend_golang/setup"

	"github.com/gin-gonic/gin"
)

func main() {
	//Declare New Gin Route System
	router := gin.New()
	router.Use(middlewares.CORSMiddleware())
	setup.ConnectDatabase()

	router.POST("/login", controllers.Login)
	router.POST("/register", controllers.Register)

	protected := router.Group("/api")
	protected.Use(middlewares.AuthMiddleware())

	protected.GET("/user", controllers.GetCurrentUser)
	protected.POST("/logout", controllers.Logout)
	protected.PUT("/user-role/:id", controllers.UpdateRole)

	// Siswa
	protected.GET("/siswa", controllers.GetSiswa)
	protected.GET("/siswa/:id", controllers.GetSiswaById)
	protected.GET("/siswa-by-user-id/:user_id", controllers.GetSiswaByUserId)
	protected.POST("/siswa", controllers.CreateSiswa)
	protected.PUT("/siswa/:id", controllers.UpdateSiswa)
	protected.DELETE("/siswa/:id", controllers.DeleteSiswa)

	//Stan
	protected.GET("/stan", controllers.GetStan)
	protected.GET("/stan/:id", controllers.GetStanById)
	protected.POST("/stan", controllers.CreateStan)
	protected.PUT("/stan/:id", controllers.UpdateStan)
	protected.DELETE("/stan/:id", controllers.DeleteStan)

	//Menu
	protected.GET("/menu", controllers.GetMenu)
	protected.GET("/menu/:id", controllers.GetMenuById)
	protected.POST("/menu", controllers.AddMenu)
	protected.PUT("/menu/:id", controllers.UpdateMenu)
	protected.DELETE("/menu/:id", controllers.DeleteMenu)

	//Transaksi
	protected.GET("/transaksi", controllers.GetTransaksi)
	protected.GET("/transaksi/:id", controllers.GetTransaksiById)
	protected.POST("/transaksi", controllers.CreateTransaksi)
	protected.PUT("/transaksi/statupdate", controllers.UpdateStatusTransaksi)

	//Detail
	protected.GET("/detail", controllers.GetDetail)
	protected.GET("/detail/:id", controllers.GetDetailById)
	protected.POST("/detail", controllers.CreateDetail)

	//Diskon
	protected.GET("/diskon", controllers.GetAllDiskon)
	protected.GET("/diskon/:id", controllers.GetDiskonById)
	protected.POST("/diskon", controllers.AddDiskon)
	protected.PUT("/diskon/:id", controllers.UpdateDiskon)
	protected.DELETE("/diskon/:id", controllers.DeleteDiskon)





	router.Static("/uploads", "./uploads")



	router.Run(":6969")
}