package routes

import (
	"backend_golang/controllers"
	"backend_golang/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Group untuk transaksi
	transaksi := r.Group("/transaksi")
	{
		transaksi.POST("/", controllers.CreateTransaksi)
		transaksi.POST("/with-detail", controllers.CreateTransaksiWithDetail) // Endpoint baru untuk membuat transaksi dengan detail
		transaksi.GET("/", controllers.GetTransaksi)
		transaksi.GET("/:id", controllers.GetTransaksiById)
		transaksi.GET("/:id/nota", controllers.PrintNota)                      // Endpoint untuk print nota
		transaksi.GET("/bulan/:tahun/:bulan", controllers.GetTransaksiByBulan) // Endpoint untuk melihat transaksi per bulan
		transaksi.GET("/rekap/:tahun", controllers.GetRekapBulanan)            // Endpoint untuk melihat rekap per bulan
	}

	// Group untuk detail transaksi
	detail := r.Group("/detail")
	{
		detail.POST("/", controllers.CreateDetail)
		detail.GET("/", controllers.GetDetail)
		detail.GET("/:id", controllers.GetDetailById)
	}

	// Group untuk menu (dengan middleware admin)
	menu := r.Group("/menu")
	menu.Use(middlewares.Admin)
	{
		menu.POST("/", controllers.AddMenu)
		menu.PUT("/:id", controllers.UpdateMenu)
		menu.DELETE("/:id", controllers.DeleteMenu)
		menu.GET("/", controllers.GetMenu)
		menu.GET("/:id", controllers.GetMenuById)
		menu.GET("/stand/:stand_id", controllers.GetMenuByStand) // Endpoint untuk mendapatkan menu berdasarkan stand
	}

	// Group untuk siswa
	siswa := r.Group("/siswa")
	{
		siswa.POST("/", controllers.CreateSiswa)
		siswa.GET("/", controllers.GetSiswa)
		siswa.GET("/:id", controllers.GetSiswaById)
		siswa.PUT("/:id", controllers.UpdateSiswa)
		siswa.DELETE("/:id", controllers.DeleteSiswa)
		siswa.GET("/stand/:stand_id", controllers.GetSiswaByStand) // Endpoint untuk mendapatkan siswa berdasarkan stand
	}
}
