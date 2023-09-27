package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/RachidMoysePolania/MotoGuard-API/helpers"
	"github.com/RachidMoysePolania/MotoGuard-API/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Register(c *gin.Context) {
	var user models.Userdata
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	encryptedpassword, err := helpers.EncryptPassword(user.Password)
	if err != nil {
		log.Println("Error encrypting the password!")
	}

	user.Password = encryptedpassword
	Db.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": "Usuario creado correctamente!"})
}

func Login(c *gin.Context) {
	var user models.Userdata
	var dbuser models.Userdata
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error Login"})
		return
	}

	Db.Where("correo = ?", user.Correo).First(&dbuser)
	if !helpers.ComparePasswords(user.Password, dbuser.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong Credentials"})
		return
	} else if user.Correo == dbuser.Correo {
		//JWT Handling
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user": dbuser.Correo,
			"exp":  time.Now().Add(time.Hour * 24).Unix(),
		})
		tokenstring, err := token.SignedString([]byte("password-for-jwt"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error Getting JWT Token"})
		}
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("Authorization", tokenstring, 3600*24*30, "", "", false, true)
		c.JSON(http.StatusOK, gin.H{"token": tokenstring})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong Credentials"})
		return
	}

}

func GetUserById(c *gin.Context) {
	userid := c.Param("id")
	var dbuser models.Userdata
	Db.Model(models.Userdata{}).Where("id = ?", userid).Preload("Road_logs").Find(&dbuser)
	c.JSON(http.StatusOK, gin.H{"data": dbuser})
}

func UpdateUserById(c *gin.Context) {
	userid := c.Param("id")
	var user models.Userdata
	var dbuser models.Userdata
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	encryptedpassword, err := helpers.EncryptPassword(user.Password)
	if err != nil {
		log.Println("Error encrypting the password!")
	}
	user.Password = encryptedpassword
	Db.Model(&dbuser).Where("id = ?", userid).Updates(user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUserById(c *gin.Context) {
	userid := c.Param("id")
	var dbuser models.Userdata
	Db.First(&dbuser, userid)
	Db.Delete(&dbuser)

	c.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("User with ID %v deleted!", userid)})
}

func GetallUsers(c *gin.Context) {
	var users []models.Userdata
	//Db.Find(&users)
	Db.Model(&models.Userdata{}).Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}
