package controllers

import (
	"backend_golang/middlewares"
	"backend_golang/models"
	"backend_golang/setup"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetSiswa(c *gin.Context) {
	var siswa []models.Siswa

	// Cek role admin
	middlewares.Admin(c)
	if c.IsAborted() {
		return
	}

	if err := setup.DB.Find(&siswa).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get siswa"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": siswa})
}

func GetSiswaByUserId(c *gin.Context) {
	userId := c.Param("user_id")
	var siswa models.Siswa

	if err := setup.DB.Where("user_id = ?", userId).First(&siswa).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get siswa"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": siswa})
}

func GetSiswaById(c *gin.Context) {
	id := c.Param("id")
	var siswa models.Siswa

	if err := setup.DB.First(&siswa, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get siswa"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": siswa})
}

func CreateSiswa(c *gin.Context) {
	// Ambil data dari form
	namaSiswa := c.PostForm("nama_siswa")
	telp := c.PostForm("telp")
	alamat := c.PostForm("alamat")
	userId := c.PostForm("user_id")

	// Validasi input
	if namaSiswa == "" || telp == "" || alamat == "" || userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Semua field harus diisi"})
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
	siswa := models.Siswa{
		NamaSiswa: namaSiswa,
		Telp:      telp,
		Alamat:    alamat,
		UserId:    userIdInt,
		Foto:      filename,
	}

	if err := setup.DB.Create(&siswa).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create siswa"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Siswa created successfully"})
}

func UpdateSiswa(c *gin.Context) {
	id := c.Param("id")

	// Ambil data dari form
	namaSiswa := c.PostForm("nama_siswa")
	telp := c.PostForm("telp")
	alamat := c.PostForm("alamat")

	// Validasi input
	if namaSiswa == "" || telp == "" || alamat == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Semua field harus diisi"})
		return
	}

	// Prepare update data
	siswa := models.Siswa{
		NamaSiswa: namaSiswa,
		Telp:      telp,
		Alamat:    alamat,
	}

	// Jika ada file foto baru
	file, err := c.FormFile("foto")
	if err == nil {
		// Generate unique filename
		filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)

		// Save new file
		if err := c.SaveUploadedFile(file, "uploads/siswa/"+filename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		// Delete old foto if exists
		var oldSiswa models.Siswa
		if err := setup.DB.First(&oldSiswa, id).Error; err == nil && oldSiswa.Foto != "" {
			os.Remove("uploads/siswa/" + oldSiswa.Foto)
		}

		siswa.Foto = filename
	}

	if err := setup.DB.Model(&siswa).Where("id = ?", id).Updates(&siswa).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update siswa"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Siswa updated successfully"})
}

func DeleteSiswa(c *gin.Context) {
	id := c.Param("id")
	var siswa models.Siswa

	if err := setup.DB.Delete(&siswa, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete siswa"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Siswa deleted successfully"})
}

// GetSiswaByStand mendapatkan data siswa berdasarkan stand ID
func GetSiswaByStand(c *gin.Context) {
	// Ambil stand_id dari parameter
	standId := c.Param("stand_id")

	// Konversi stand_id ke int64
	standIdInt, err := strconv.ParseInt(standId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stand_id format"})
		return
	}

	// Cari transaksi berdasarkan stand_id
	var transaksis []models.Transaksi
	if err := setup.DB.
		Where("stan_id = ?", standIdInt).
		Preload("Siswa").
		Find(&transaksis).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get transactions"})
		return
	}

	// Buat map untuk menyimpan siswa unik
	siswaMap := make(map[int64]models.Siswa)
	for _, t := range transaksis {
		if t.Siswa.Id != 0 { // Pastikan siswa ada
			siswaMap[t.Siswa.Id] = t.Siswa
		}
	}

	// Konversi map ke slice
	var siswa []models.Siswa
	for _, s := range siswaMap {
		siswa = append(siswa, s)
	}

	c.JSON(http.StatusOK, gin.H{
		"stand_id": standIdInt,
		"data":     siswa,
	})
}
