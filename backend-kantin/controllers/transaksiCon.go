package controllers

import (
	"backend_golang/models"
	"backend_golang/middlewares"
	"backend_golang/setup"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTransaksi(c *gin.Context) {
	var Transaksi models.Transaksi

	// Cek role admin
	middlewares.Admin(c)
	if c.IsAborted() {
		return
	}

	if err := setup.DB.Find(&Transaksi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get stan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Transaksi})
}

func GetTransaksiById(c *gin.Context) {
	id := c.Param("id")
	var Transaksi models.Transaksi
	if err := setup.DB.First(&Transaksi, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stan not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Transaksi})
}

func CreateTransaksi(c *gin.Context) {

	var input struct {
		StanId   int64 `json:"stan_id" binding:"required`
		SiswaId  int64 `json:"siswa_id" binding:"required`
		StatusId int64 `json:"status"  binding:"required`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if (input.StatusId){
	// 	input.StatusId = "1"
	// }

	// StandIdInt, err := strconv.ParseInt(input.StanId, 10, 64)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stand_id format"})
	// 	return
	// }
	// SiswaIdInt, err := strconv.ParseInt(input.SiswaId, 10, 64)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid siswa_id format"})
	// 	return
	// }
	// StatusIdInt, err := strconv.ParseInt(input.StatusId, 10, 64)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status_id format"})
	// 	return
	// }

	// Tran := models.Transaksi{
	// 	StanId: StandIdInt,
	// 	SiswaId: SiswaIdInt,
	// 	StatusId: StatusIdInt,
	// }

	Tran := models.Transaksi{
		StanId:   input.StanId,
		SiswaId:  input.SiswaId,
		StatusId: input.StatusId,
	}

	if err := setup.DB.Create(&Tran).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaksi created successfully"})
}

func UpdateStatusTransaksi(c *gin.Context) {

	id := c.Param("id")
	var TransCheck models.Transaksi

	// Cek role admin
	middlewares.Admin(c)
	if c.IsAborted() {
		return
	}

	if err := setup.DB.First(&TransCheck, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var input struct {
		StatusId string `json:"stat" binding:"required"`
	}

	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	StatusIdInt, err := strconv.ParseInt(input.StatusId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id format"})
		return
	}

	Transaksi := models.Transaksi{
		StatusId: StatusIdInt,
	}

	if err := setup.DB.Save(&Transaksi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status changed successfully"})
}
