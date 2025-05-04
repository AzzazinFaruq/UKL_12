// this page are including all of auth controllers
package controllers

import (
	"backend_golang/models"
	"backend_golang/setup"
	"backend_golang/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {

	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required,min=8"`
		RoleId   int64  `json:"role_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	user := models.User{
		Username: input.Username,
		Password: string(hashedPassword),
		RoleId:   input.RoleId,
	}

	if err := setup.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Created Succesfully"})

}

func Login(c *gin.Context) {
	var input struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		RememberMe bool   `json:"remember_me"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the user
	var user models.User
	if err := setup.DB.Preload("Role").Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password", "authenticated": false})
		return
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password", "authenticated": false})
		return
	}

	// Tentukan durasi token berdasarkan RememberMe
	var tokenDuration time.Duration
	if input.RememberMe {
		tokenDuration = 7 * 24 * time.Hour // 7 hari jika Remember Me dicentang
	} else {
		tokenDuration = 24 * time.Hour // 1 hari
	}

	// Generate token JWT menggunakan helper
	tokenString, err := utils.GenerateJWT(uint(user.Id))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to generate token, Internal Server Error", "authenticated": false})
		return
	}

	// Set cookie dengan token JWT
	c.SetCookie("Authorization", tokenString, int(tokenDuration.Seconds()), "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message":       "Login successful",
		"username":      user.Username,
		"role":          user.Role.Role,
		"authenticated": true,
	})
}

// Fungsi untuk mendapatkan data user yang sedang login
func GetCurrentUser(c *gin.Context) {
	// Ambil user ID dari context (disimpan di middleware)
	userID, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found", "status": false})
		return
	}

	// Cari user berdasarkan ID
	var user models.User
	if err := setup.DB.Preload("Role").First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found", "status": false})
		return
	}

	// Kembalikan data user (tanpa password untuk alasan keamanan)
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

// Fungsi untuk logout
func Logout(c *gin.Context) {
	// Hapus cookie Authorization dengan durasi negatif untuk menghilangkannya
	c.SetCookie("Authorization", "", -1, "/", "", false, true)

	// Kirim respon logout sukses
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
