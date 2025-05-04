package controllers

import (
	"net/http"

	"backend_golang/setup"

	"github.com/gin-gonic/gin"
)

func UpdateRole(c *gin.Context) {
	// Ambil user_id dari parameter URL
	userId := c.Param("id")

	// Struct untuk menerima data role baru
	var requestBody struct {
		RoleId int `json:"role_id"`
	}

	// Bind JSON request ke struct
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Format data tidak valid",
		})
		return
	}

	// Query untuk update role
	query := "UPDATE users SET role_id = ? WHERE id = ?"

	// Eksekusi query
	result := setup.DB.Exec(query, requestBody.RoleId, userId)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal mengupdate role",
		})
		return
	}

	// Cek apakah ada baris yang terupdate
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User tidak ditemukan",
		})
		return
	}

	// Response sukses
	c.JSON(http.StatusOK, gin.H{
		"message": "Role berhasil diupdate",
		"role_id": requestBody.RoleId,
	})
}
