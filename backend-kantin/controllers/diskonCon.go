package controllers

import (
	"backend_golang/models"
	"backend_golang/middlewares"
	"backend_golang/setup"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllDiskon(c *gin.Context) {
	var diskon []models.Diskon

	// Cek role admin
	middlewares.Admin(c)
	if c.IsAborted() {
		return
	}


	if err := setup.DB.Find(&diskon).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": diskon})
}

func GetDiskonById(c *gin.Context) {
	var diskon models.Diskon
	id := c.Param("id")
	// Cek role admin
	middlewares.Admin(c)
	if c.IsAborted() {
		return
	}
	if err := setup.DB.First(&diskon, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": diskon})
}

func AddDiskon(c *gin.Context) {

	// Cek role admin
	middlewares.Admin(c)
	if c.IsAborted() {
		return
	}

	var input struct {
		NamaDiskon       string `json:"nama_diskon" binding:"required"`
		PresentaseDiskon string `json:"presentase_diskon" binding:"required"`
		TanggalAwal      time.Time `json:"tanggal_awal" binding:"required"`
		TanggalAkhir     time.Time `json:"tanggal_akhir" binding:"required"`
	}

	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	diskon := models.Diskon{
		NamaDiskon: input.NamaDiskon,
		PresentaseDiskon: input.PresentaseDiskon,
		TanggalAwal: input.TanggalAwal,
		TanggalAkhir: input.TanggalAkhir,
	}

	if err := setup.DB.Create(&diskon).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Diskon created successfully"})
}

func UpdateDiskon(c *gin.Context) {
	var diskonn models.Diskon
	id := c.Param("id")
	// Cek role admin
	middlewares.Admin(c)
	if c.IsAborted() {
		return
	}

	if err := setup.DB.First(&diskonn, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var input struct {
		NamaDiskon       string `json:"nama_diskon" binding:"required"`
		PresentaseDiskon string `json:"presentase_diskon" binding:"required"`
		TanggalAwal      time.Time `json:"tanggal_awal" binding:"required"`
		TanggalAkhir     time.Time `json:"tanggal_akhir" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	diskon := models.Diskon{
		NamaDiskon: input.NamaDiskon,
		PresentaseDiskon: input.PresentaseDiskon,
		TanggalAwal: input.TanggalAwal,
		TanggalAkhir: input.TanggalAkhir,
	}

	if err := setup.DB.Save(&diskon).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
}

func DeleteDiskon(c *gin.Context) {
	var diskon models.Diskon
	id := c.Param("id")

	// Cek role admin
	middlewares.Admin(c)
	if c.IsAborted() {
		return
	}
	
	if err := setup.DB.Delete(&diskon, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Diskon deleted successfully"})
}
