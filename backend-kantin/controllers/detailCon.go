package controllers

import (
	"backend_golang/models"
	"backend_golang/setup"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetDetail(c *gin.Context) {

	var detail []models.Detail

	if err := setup.DB.First(&detail).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": detail})
}

func GetDetailById(c *gin.Context) {

	id := c.Param("id")
	var detail []models.Detail

	if err := setup.DB.First(&detail, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": detail})
}

func CreateDetail(c *gin.Context) {
	var input struct {
		TransaksiId string `json:"transaksi_id" binding:"required"`
		MenuId      string `json:"menu_id" binding:"required"`
		Qty         string `json:"qty" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	TransaksiIdInt, err := strconv.ParseInt(input.TransaksiId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaksi_id format"})
		return
	}
	MenuIdInt, err := strconv.ParseInt(input.MenuId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu_id format"})
		return
	}
	QtyInt, err := strconv.ParseInt(input.Qty, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid qty format"})
		return
	}

	// Ambil data menu
	var Menu models.Menu
	if err := setup.DB.First(&Menu, input.MenuId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Menu not found"})
		return
	}

	// Ambil data diskon yang berlaku hari ini
	var diskon models.Diskon
	today := time.Now().Format("2006-01-02")
	if err := setup.DB.Where("tanggal_mulai <= ? AND tanggal_selesai >= ?", today, today).First(&diskon).Error; err != nil {
		// Jika tidak ada diskon, gunakan harga normal
		Final := Menu.Harga * float64(QtyInt)
		detail := models.Detail{
			TransaksiId: TransaksiIdInt,
			MenuId:      MenuIdInt,
			Qty:         QtyInt,
			HargaBeli:   Final,
		}

		if err := setup.DB.Create(&detail).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Detail Successfully Created", "data": detail})
		return
	}

	// Konversi PresentaseDiskon dari string ke int
	presentaseDiskon, err := strconv.ParseInt(diskon.PresentaseDiskon, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid diskon percentage"})
		return
	}

	// Hitung harga dengan diskon
	hargaDiskon := Menu.Harga * (1 - float64(presentaseDiskon)/100)
	Final := hargaDiskon * float64(QtyInt)

	detail := models.Detail{
		TransaksiId: TransaksiIdInt,
		MenuId:      MenuIdInt,
		Qty:         QtyInt,
		HargaBeli:   Final,
	}

	if err := setup.DB.Create(&detail).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Detail Successfully Created",
		"data":    detail,
		"diskon": gin.H{
			"id":                   diskon.Id,
			"presentase":           diskon.PresentaseDiskon,
			"harga_awal":           Menu.Harga,
			"harga_setelah_diskon": hargaDiskon,
		},
	})
}
