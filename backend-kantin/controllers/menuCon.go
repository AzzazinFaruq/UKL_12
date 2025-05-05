package controllers

import (
	"backend_golang/middlewares"
	"backend_golang/models"
	"backend_golang/setup"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMenu(c *gin.Context) {
	var menu []models.Menu

	if err := setup.DB.Preload("Jenis").Find(&menu).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get Menu"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": menu})
}

func GetMenuById(c *gin.Context) {
	id := c.Param("id")
	var menu models.Menu

	if err := setup.DB.First(&menu, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get Menu"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": menu})
}

func AddMenu(c *gin.Context) {
	var menu models.Menu

	// Cek role admin
	middlewares.Admin(c)
	if c.IsAborted() {
		return
	}

	// Mendapatkan data dari form
	menu.NamaMakanan = c.PostForm("nama_makanan")
	harga, err := strconv.ParseFloat(c.PostForm("harga"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid harga"})
		return
	}
	menu.Harga = harga

	jenisId, err := strconv.ParseInt(c.PostForm("jenis_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid jenis_id"})
		return
	}
	menu.JenisId = jenisId

	menu.Deskripsi = c.PostForm("deskripsi")

	stanId, err := strconv.ParseInt(c.PostForm("stan_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stan_id"})
		return
	}
	menu.StanId = stanId

	// Mendapatkan file foto dari form
	file, err := c.FormFile("foto")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Foto is required"})
		return
	}

	// Menyimpan file foto ke direktori yang diinginkan
	filePath := filepath.Join("uploads", file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save foto"})
		return
	}
	menu.Foto = filePath

	// Menyimpan data ke database
	if err := setup.DB.Create(&menu).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add Menu"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Menu Berhasil Ditambahkan", "data": menu})
}

func UpdateMenu(c *gin.Context) {
	var menu models.Menu

	// Mendapatkan ID menu dari parameter URL
	id := c.Param("id")

	// Cek role admin
	middlewares.Admin(c)
	if c.IsAborted() {
		return
	}

	// Mencari menu berdasarkan ID
	if err := setup.DB.First(&menu, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
		return
	}

	// Memperbarui data dari form
	if namaMakanan := c.PostForm("nama_makanan"); namaMakanan != "" {
		menu.NamaMakanan = namaMakanan
	}

	if hargaStr := c.PostForm("harga"); hargaStr != "" {
		harga, err := strconv.ParseFloat(hargaStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid harga"})
			return
		}
		menu.Harga = harga
	}

	if jenisIdStr := c.PostForm("jenis_id"); jenisIdStr != "" {
		jenisId, err := strconv.ParseInt(jenisIdStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid jenis_id"})
			return
		}
		menu.JenisId = jenisId
	}

	if deskripsi := c.PostForm("deskripsi"); deskripsi != "" {
		menu.Deskripsi = deskripsi
	}

	if stanIdStr := c.PostForm("stan_id"); stanIdStr != "" {
		stanId, err := strconv.ParseInt(stanIdStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stan_id"})
			return
		}
		menu.StanId = stanId
	}

	// Memperbarui file foto jika ada
	if file, err := c.FormFile("foto"); err == nil {
		filePath := filepath.Join("uploads", file.Filename)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save foto"})
			return
		}
		menu.Foto = filePath
	}

	// Menyimpan perubahan ke database
	if err := setup.DB.Save(&menu).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Menu"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": menu})
}

func DeleteMenu(c *gin.Context) {
	id := c.Param("id")
	var menu models.Menu

	// Cek role admin
	middlewares.Admin(c)
	if c.IsAborted() {
		return
	}

	if err := setup.DB.Delete(&menu, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Menu"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Menu Berhasil Dihapus"})
}

// GetMenuByStand mendapatkan menu berdasarkan stand ID
func GetMenuByStand(c *gin.Context) {
	// Ambil stand_id dari parameter
	standId := c.Param("stand_id")

	// Konversi stand_id ke int64
	standIdInt, err := strconv.ParseInt(standId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format stand_id tidak valid"})
		return
	}

	// Cari menu berdasarkan stand_id dengan preload Jenis
	var menu []models.Menu
	if err := setup.DB.
		Preload("Jenis").
		Preload("Stan").
		Where("stan_id = ?", standIdInt).
		Find(&menu).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data menu"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"stand_id": standIdInt,
		"data":     menu,
	})
}
