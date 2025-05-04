package controllers

import (
	"backend_golang/middlewares"
	"backend_golang/models"
	"backend_golang/setup"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetStan(c *gin.Context) {
	var stan []models.Stan
	if err := setup.DB.Find(&stan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get stan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": stan})
}

func GetStanById(c *gin.Context) {
	id := c.Param("id")
	// Cek role admin
	middlewares.Admin(c)
	if c.IsAborted() {
		return
	}
	var stan models.Stan
	if err := setup.DB.First(&stan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stan not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": stan})
}

func CreateStan(c *gin.Context) {
	// Ambil data dari form
	namaStan := c.PostForm("nama_stan")
	telp := c.PostForm("telp")
	namaPemilik := c.PostForm("nama_pemilik")
	userId := c.PostForm("user_id")

	// Validasi input
	if namaStan == "" || telp == "" || namaPemilik == "" || userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Semua field harus diisi"})
		return
	}

	// Cek role admin
	middlewares.Admin(c)
	if c.IsAborted() {
		return
	}

	// Convert userId ke int64
	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id format"})
		return
	}

	// Handle file upload
	file, err := c.FormFile("foto")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	// Generate unique filename
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)

	// Save file
	if err := c.SaveUploadedFile(file, "uploads/siswa/"+filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Create siswa with foto
	stan := models.Stan{
		NamaStan:    namaStan,
		Telp:        telp,
		NamaPemilik: namaPemilik,
		UserId:      userIdInt,
		Foto:        filename,
	}

	if err := setup.DB.Create(&stan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create siswa"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stan created successfully"})
}

func UpdateStan(c *gin.Context) {
	id := c.Param("id")
	var Stan models.Stan

	// Cek role admin
	middlewares.Admin(c)
	if c.IsAborted() {
		return
	}

	// Cek apakah stan ada
	if err := setup.DB.First(&Stan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stan not found"})
		return
	}

	// Ambil data dari form
	namaStan := c.PostForm("nama_stan")
	namaPemilik := c.PostForm("nama_pemilik")
	telp := c.PostForm("telp")

	// Validasi input
	if namaStan == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama stan tidak boleh kosong"})
		return
	}

	// Handle file upload jika ada
	file, err := c.FormFile("foto")
	if err == nil {
		// Generate unique filename
		filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)

		// Save file
		if err := c.SaveUploadedFile(file, "uploads/siswa/"+filename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan file"})
			return
		}

		// Update foto
		Stan.Foto = filename
	}

	// Update data stan
	Stan.NamaStan = namaStan
	Stan.NamaPemilik = namaPemilik
	Stan.Telp = telp

	if err := setup.DB.Save(&Stan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate stan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stan berhasil diupdate"})
}

func DeleteStan(c *gin.Context) {
	id := c.Param("id")
	// Cek role admin
	middlewares.Admin(c)
	if c.IsAborted() {
		return
	}

	var stan models.Stan
	if err := setup.DB.Delete(&stan, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete stan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Stan deleted successfully"})
}
